#!/bin/bash

dirs=(providers/config providers/casdoor providers/cert providers/captcha providers/database/mysql providers/database/postgres providers/database/redis providers/database/sqlite providers/grpcs providers/hashids providers/http providers/httpclient providers/jwt providers/k8s providers/log providers/swagger providers/uuid providers/single_flight providers/micro_service services/http services/grpc)
for dir in ${dirs[@]}; do
    cd $dir
    go mod tidy
    cd -
done