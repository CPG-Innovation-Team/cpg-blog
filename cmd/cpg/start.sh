#!/usr/bin/env bash
# shellcheck disable=SC2009
# shellcheck disable=SC2126
# shellcheck disable=SC2006

num=`ps -elf | grep main | wc -l`
one=1
if [ $num -gt $one ]; then
killall main
else
echo "start deploy"
fi

cd /data/cpg-blog/cmd/cpg/ && go build main.go
nohup /data/cpg-blog/cmd/cpg/main > /data/cpg-blog/cmd/cpg/gin.log 2>&1 &
ps -elf | grep main