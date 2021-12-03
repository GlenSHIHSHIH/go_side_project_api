#!/bin/bash

#db setting
export db_host="172.21.0.2"
export db_port="3306"
export db_name="jiyoung_shopee"
export db_username="glen"
export db_password="1qaz@WSX"

#redis setting
export redis_host="172.21.0.3"
export redis_port="6379"
export redis_password="1qaz@WSX"
export redis_db="0"

#web setting
export web_host="kumkum.com"
export web_port="8010"
export web_imgUrl="https://cf.shopee.tw/file/"
export jwt_secret="f9946c78-f48a-435d-acc4-4bf469ef2680"
export jwt_token_time="480"
export jwt_ref_token_time="1440"

cd /go/golang-shopee/docker_build_go
nohup ./compiler/golang_file shopee-api &
