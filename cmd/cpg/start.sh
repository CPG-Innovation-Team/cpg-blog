#!/usr/bin/env bash
# shellcheck disable=SC2009
# shellcheck disable=SC2126

num=$(ps -elf | grep main | wc -l)
if [ "${num}" -gt 1 ]; then
killall main
else
echo "start deploy"
fi

cd /data/cpg-blog/cmd/cpg/ && go build main.go
nohup /data/cpg-blog/cmd/cpg/main > /data/cpg-blog/cmd/cpg/gin.log 2>&1 &
ps -elf | grep main