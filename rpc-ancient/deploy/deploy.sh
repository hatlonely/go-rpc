#!/usr/bin/env bash

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
    docker tag "${ImageRepository}:${ImageTag}" "${RegistryServer}/${ImageRepository}:${ImageTag}"
    docker push "${RegistryServer}/${ImageRepository}:${ImageTag}"
}

function SQLTpl() {
    environment=$1
    kubectl run -n prod -it --rm sql --image=mysql:5.7.30 --restart=Never -- \
      mysql -uroot -h${MysqlServer} -p${MysqlRootPassword} -e "$(cat "tmp/${environment}/create_table.sql")"
}

function CreateNamespaceIfNotExists() {
    kubectl get namespaces "${Namespace}" 2>/dev/null 1>&2 && return 0
    kubectl create namespace "${Namespace}" &&
    Info "create namespace ${Namespace} success" ||
    Warn "create namespace ${Namespace} failed"
}

function CreatePullSecretsIfNotExists() {
    kubectl get secret "${PullSecrets}" -n "${Namespace}" 2>/dev/null 1>&2 && return 0
    kubectl create secret docker-registry "${PullSecrets}" \
        --docker-server="${RegistryServer}" \
        --docker-username="${RegistryUsername}" \
        --docker-password="${RegistryPassword}" \
        --namespace="prod" &&
    Info "[kubectl create secret docker-registry ${PullSecrets}] success" ||
    Warn "[kubectl create secret docker-registry ${PullSecrets}] failed"
}

function Render() {
    environment=$1
    variable=$2
    sh tpl.sh render "${environment}" "${variable}"
}

function Test() {
    kubectl run -n "${Namespace}" -it --rm "${Name}" --image="${RegistryServer}/${ImageRepository}:${ImageTag}" --restart=Never -- /bin/bash
}

function Install() {
    environment=$1
    helm install "${Name}" -n "${Namespace}" "./chart/${Name}" -f "tmp/${environment}/chart.yaml"
}

function Upgrade() {
    environment=$1
    helm upgrade "${Name}" -n "${Namespace}" "./chart/${Name}" -f "tmp/${environment}/chart.yaml"
}

function Diff() {
    environment=$1
    helm diff upgrade "${Name}" -n "${Namespace}" "./chart/${Name}" -f "tmp/${environment}/chart.yaml"
}

function Delete() {
    helm delete "${Name}" -n "${Namespace}"
}

function Restart() {
    kubectl get pods -n "${Namespace}" | grep "${Name}" | awk '{print $1}' | xargs kubectl delete pods -n "${Namespace}"
}

function Help() {
    echo "sh deploy.sh <environment> <action>"
    echo "example"
    echo "  sh deploy.sh prod build"
    echo "  sh deploy.sh prod sql"
    echo "  sh deploy.sh prod secret"
    echo "  sh deploy.sh prod render"
    echo "  sh deploy.sh prod install"
    echo "  sh deploy.sh prod upgrade"
    echo "  sh deploy.sh prod delete"
    echo "  sh deploy.sh prod diff"
    echo "  sh deploy.sh prod test"
    echo "  sh deploy.sh prod restart"
}

function main() {
    if [ -z "$2" ]; then
        Help
        return 0
    fi

    environment=$1
    action=$2

    # shellcheck source=tmp/$1/environment.sh
    source "tmp/$1/environment.sh"

    if [ "${K8sContext}" != "$(kubectl config current-context)" ]; then
        Warn "context [${WebOffice_K8S_Context}] not match [$(kubectl config current-context)]"
        return 1
    fi

    case "${action}" in
        "build") Build;;
        "sql") SQLTpl "${environment}";;
        "secret") CreatePullSecretsIfNotExists;;
        "render") Render "${environment}" "$3";;
        "install") Install "${environment}";;
        "upgrade") Upgrade "${environment}";;
        "diff") Diff "${environment}";;
        "delete") Delete;;
        "test") Test;;
        "restart") Restart;;
        *) Help;;
    esac
}

main "$@"
