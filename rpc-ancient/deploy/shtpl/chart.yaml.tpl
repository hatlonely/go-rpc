namespace: ${Namespace}
name: ${Name}
replicaCount: ${ReplicaCount}

image:
  repository: ${RegistryServer}/${ImageRepository}
  tag: ${ImageTag}
  pullPolicy: Always

imagePullSecrets:
  name: ${PullSecrets}

ingress:
  host: ${IngressHost}
  secretName: ${IngressSecret}

config: |
  {
    "http": {
      "port": 80
    },
    "grpc": {
      "port": 6080
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
    "elasticsearch": {
      "uri": "http://${ElasticsearchServer}"
    },
    "service": {
      "elasticsearchIndex": "shici"
    },
    "logger": {
      "grpc": {
        "level": "Info",
        "flatMap": true,
        "writers": [{
          "type": "RotateFile",
          "rotateFileWriter": {
            "filename": "log/${Name}.rpc",
            "maxAge": "24h",
            "formatter": {
              "type": "Json"
            }
          }
        }, {
          "type": "ElasticSearch",
          "elasticSearchWriter": {
            "index": "grpc",
            "idField": "requestID",
            "timeout": "200ms",
            "msgChanLen": 200,
            "workerNum": 2,
            "elasticSearch": {
              "uri": "http://${ElasticsearchServer}"
            }
          }
        }]
      },
      "info": {
        "level": "Info",
        "writers": [{
          "type": "RotateFile",
          "rotateFileWriter": {
            "filename": "log/${Name}.rpc",
            "maxAge": "24h",
            "formatter": {
              "type": "Json"
            }
          }
        }]
      }
    }
  }