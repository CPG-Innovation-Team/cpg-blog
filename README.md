# cpg-blog
## cpg-blog-structure(V1)

## Development Language: Golang

### 服务介绍：

### 主要技术栈：

GIN

Redis

GROM

Docker

k8s

### 数据存储：

Mysql 5.7+

MongoDB

Redis 集群

### 主要功能：
![主要功能](https://cpg-innovate.atlassian.net/4e8d9825-b899-4ce9-bc78-57ab0f893947#media-blob-url=true&id=033a884f-eb6a-4be9-a12e-6cc063432006&collection=&contextId=10148&mimeType=image%2Fpng&name=image-20210801-094252.png&size=56879&width=518&height=570)

### 数据库设计：
[数据库设计](https://cpg-innovate.atlassian.net/browse/CPGI-124)

### 调用关系：

物理调用关系
![调用关系](blob:https://cpg-innovate.atlassian.net/e806e6fe-ea8e-46d9-b67e-7e261b85778a#media-blob-url=true&id=d95c4d64-837e-4831-9f5a-25913b269b8e&collection=&contextId=10148&mimeType=image%2Fpng&name=image-20210801-101611.png&size=46965&width=499&height=538)

### 内部依赖关系：
![依赖关系](blob:https://cpg-innovate.atlassian.net/f18ce653-fa6b-4122-8ae6-8c7bfc5b0431#media-blob-url=true&id=daab0f94-1d0a-4175-a0e4-5183da035329&collection=&contextId=10148&mimeType=image%2Fpng&name=image-20210801-102459.png&size=31181&width=375&height=505)

### 目录结构：
```
.
│  go.mod[项目依赖配置文件]
│  go.sum
│      
├─api[所有接口定义]
├─cmd[存放启动文件和控制器]
│  ├─controller
│  │      controller.go
│  │      
│  └─main
│          cpcApplicationMain.go
│          
├─internal[所有服务源码]
│  └─server
│      ├─server1
│      │      handler.go
│      │      router.go[该服务所有请求地址]
│      │      server1.go
│      │      
│      └─server2
│              handler.go
│              router.go
│              server2.go
│              
├─pkg[外部应用程序使用的库代码]
└─test[测试代码]
```
