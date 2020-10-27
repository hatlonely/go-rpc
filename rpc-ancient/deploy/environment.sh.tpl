#!/usr/bin/env bash

Namespace="prod"
Name="go-rpc-ancient"
RegistryServer="{{.registry.server}}"
RegistryUsername="{{.registry.username}}"
RegistryPassword="{{.registry.password}}"
MysqlServer="{{.mysql.server}}"
MysqlRootPassword="{{.mysql.rootPassword}}"
MysqlUsername="{{.mysql.username}}"
MysqlPassword="{{.mysql.password}}"
MysqlDatabase="ancient"
RedisAddr="{{.redis.addr}}"
RedisPassword="{{.redis.password}}"
PullSecrets="hatlonely-pull-secrets"
Image="hatlonely/go-rpc-ancient"
Version="1.0.0"
ReplicaCount=2
IngressHost="k8s.ancient.hatlonely.com"
IngressSecret="k8s-secret"