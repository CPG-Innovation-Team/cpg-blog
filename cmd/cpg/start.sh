cd /data/cpg-blog/cmd/cpg/ && go build main.go
nohup /data/cpg-blog/cmd/cpg/main > /data/cpg-blog/cmd/cpg/gin.log 2>&1 &
ps -elf | grep main