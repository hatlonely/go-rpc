#!/usr/bin/env bash

source tmp/environment.sh

function Trac() {
    echo "[TRAC] [$(date +"%Y-%m-%d %H:%M:%S")] $1"
}

function Info() {
    echo "\033[1;32m[INFO] [$(date +"%Y-%m-%d %H:%M:%S")] $1\033[0m"
}

function Warn() {
    echo "\033[1;31m[WARN] [$(date +"%Y-%m-%d %H:%M:%S")] $1\033[0m"
    return 1
}

function Build() {
    cd .. && make image && cd -
    docker login --username="${RegistryUsername}" --password="${RegistryPassword}" "${RegistryServer}"
    docker tag ${Image}:${Version} ${RegistryServer}/${Image}:${Version}
    docker push ${RegistryServer}/${Image}:${Version}
}

function SQLTpl() {
    cat > tmp/create_table.sql <<EOF
CREATE DATABASE IF NOT EXISTS ${MysqlDatabase};
CREATE USER IF NOT EXISTS '${MysqlUsername}'@'%' IDENTIFIED BY '${MysqlPassword}';
GRANT ALL PRIVILEGES ON ${MysqlDatabase}.* TO '${MysqlUsername}'@'%';

USE ${MysqlDatabase};
CREATE TABLE \`accounts\` (
  \`id\` bigint NOT NULL AUTO_INCREMENT,
  \`email\` varchar(64) DEFAULT NULL,
  \`phone\` varchar(64) DEFAULT NULL,
  \`name\` varchar(32) DEFAULT NULL,
  \`password\` varchar(32) DEFAULT NULL,
  \`birthday\` timestamp NULL DEFAULT '1970-01-02 00:00:00',
  \`gender\` int DEFAULT NULL,
  \`avatar\` varchar(512) DEFAULT NULL,
  PRIMARY KEY (\`id\`),
  UNIQUE KEY \`email_idx\` (\`email\`),
  UNIQUE KEY \`phone_idx\` (\`phone\`),
  UNIQUE KEY \`name_idx\` (\`name\`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8
EOF
    kubectl run -n prod -it --rm sql --image=mysql:5.7.30 --restart=Never -- mysql -uroot -hmysql -p${MysqlRootPassword} -e "$(cat tmp/create_table.sql)"
}

function CreateNamespaceIfNotExists() {
    kubectl get namespaces "${Namespace}" 2>/dev/null 1>&2 && return 0
    kubectl create namespace "${Namespace}" &&
    Info "create namespace ${Namespace} success" ||
    Warn "create namespace ${Namespace} failed"
}

function CreatePullSecretsIfNotExists() {
    kubectl get secret "${PullSecrets}" -n "${Namespace}" 2>/dev/null 1>&2 && return 0
    kubectl create secret docker-registry ${PullSecrets} \
        --docker-server="${RegistryServer}" \
        --docker-username="${RegistryUsername}" \
        --docker-password="${RegistryPassword}" \
        --namespace="prod" &&
    Info "[kubectl create secret docker-registry ${PullSecrets}] success" ||
    Warn "[kubectl create secret docker-registry ${PullSecrets}] failed"
}

function CreateConfigMap() {
    CreateNamespaceIfNotExists || return 1

cat > tmp/${ConfigmapFile} <<EOF
{
  "http": {
    "port": 80
  },
  "grpc": {
    "port": 6080
  },
  "account": {
    "accountExpiration": "5m",
    "captchaExpiration": "30m"
  },
  "redis": {
    "addr": "${RedisAddr}",
    "password": "${RedisPassword}",
    "dialTimeout": "200ms",
    "readTimeout": "200ms",
    "writeTimeout": "200ms",
    "maxRetries": 3,
    "poolSize": 20,
    "db": 0
  },
  "mysql": {
    "username": "${MysqlUsername}",
    "password": "${MysqlPassword}",
    "database": "${MysqlDatabase}",
    "host": "${MysqlServer}",
    "port": 3306,
    "connMaxLifeTime": "60s",
    "maxIdleConns": 10,
    "maxOpenConns": 20
  },
  "email": {
    "from": "${EmailFrom}",
    "password": "${EmailPassword}",
    "server": "${EmailServer}",
    "port": ${EmailPort}
  },
  "logger": {
    "grpc": {
      "level": "Info",
      "writers": [{
        "type": "RotateFile",
        "filename": "log/account.grpc",
        "maxAge": "24h"
      }]
    },
    "warn": {
      "level": "Warn",
      "writers": [{
        "type": "RotateFile",
        "filename": "log/account.err",
        "maxAge": "24h"
      }]
    },
    "info": {
      "level": "Info",
      "writers": [{
        "type": "RotateFile",
        "filename": "log/account.log",
        "maxAge": "24h"
      }]
    }
  }
}
EOF

    kubectl get configmap "${Configmap}" -n "${Namespace}" 2>/dev/null 1>&2 && return 0

    kubectl create configmap "${Configmap}" -n "${Namespace}" --from-file=${ConfigmapFile}=tmp/${ConfigmapFile} &&
    Info "[kubectl create configmap "${Configmap}" -n "${Namespace}" --from-file=${ConfigmapFile}=tmp/${ConfigmapFile}] success" ||
    Warn "[kubectl create configmap "${Configmap}" -n "${Namespace}" --from-file=${ConfigmapFile}=tmp/${ConfigmapFile}] fail"
}

function CreateDeployment() {
    cat > tmp/job.yaml <<EOF
apiVersion: batch/v1
kind: Job
metadata:
  name: ${Name}
  namespace: ${Namespace}
spec:
  parallelism: 1
  completions: 1
  activeDeadlineSeconds: 1800
  backoffLimit: 1
  template:
    metadata:
      name: ${Name}
    spec:
      imagePullSecrets:
      - name: ${PullSecrets}
      containers:
      - name: ${Name}
        imagePullPolicy: Always
        image: ${Image}:${Version}
        command: [ "bin/databus", "-c", "config/shici.json" ]
        volumeMounts:
        - name: ${Name}-data
          mountPath: /var/docker/${Name}/data
        - name: ${Name}-config
          mountPath: /var/docker/${Name}/config
      volumes:
      - name: ${Name}-data
        persistentVolumeClaim:
          claimName: ${PVCName}
      - name: ${Name}-config
        projected:
          sources:
          - configMap:
              name: ${Configmap}
              items:
                - key: ${ConfigmapFile}
                  path: shici.json
      restartPolicy: OnFailure
EOF

    kubectl get job -n "${Namespace}" "${Name}" && kubectl delete job -n "${Namespace}" "${Name}"
    kubectl apply -f tmp/job.yaml &&
    Info "[kubectl apply -f tmp/job.yaml] success" ||
    Warn "[kubectl apply -f tmp/job.yaml] failed"
}

function Help() {
    echo "sh deploy.sh <action>"
    echo "example"
    echo "  sh deploy.sh build"
    echo "  sh deploy.sh sql"
    echo "  sh deploy.sh configmap"
    echo "  sh deploy.sh secret"
    echo "  sh deploy.sh job"
}

function main() {
    if [ -z "$1" ]; then
        Help
        return 0
    fi

    case "$1" in
        "build") Build;;
        "sql") SQLTpl;;
        "secret") CreatePullSecretsIfNotExists;;
        "configmap") CreateConfigMap;;
        "job") CreateJob;;
    esac
}

main "$@"
