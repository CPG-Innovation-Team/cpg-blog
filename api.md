# 登录
```
	uri:/login
    query:
        {
            "username": "chenxi666",
            "email": "chxi@163.com",
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":
                {
                    "token":""
                }
        }
```

# 注册
```
	uri:/register
    query:
        {
            "username": "chenxi666",    //必填
            "email": "chxi@163.com",    //必填
            "passCode": "dddfdfdfsdd",  
            "passwd": "chenXi951026.",  //必填
            "nickname": "cx",   
            "avatar": "www.baidu.com",
            "gender": 1,
            "introduce": "hello"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":   
                {
                    "token":""
                }
        }             
```

# 查询用户信息
```
	uri:/admin/user/query/info
    query:
        {
            "username": "chenxi666",    //两者填一则可
            "email": "chxi@163.com",
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":   
                {
                    "username": "chenxi666",
                    "email": "chxi@163.com",
                    "passCode": "dddfdfdfsdd",
                    "passwd": "chenXi951026.",
                    "nickname": "cx",
                    "avatar": "www.baidu.com",
                    "gender": 1,
                    "introduce": "hello",
                    "state":0, 
                    "isRoot": "true",
                    "createdAt":1232312,
                    "updatedAt":232323         
                }
        }             
```

# 修改用户信息
```
	uri:/admin/user/update/info
    query:
        {
            "username": "chenxi666",    //必填
            "email": "chxi@163.com",    //必填
            "passCode": "dddfdfdfsdd",
            "passwd": "chenXi951026.",  //必填
            "nickname": "cx",
            "avatar": "www.baidu.com",
            "gender": 1,
            "introduce": "hello",
            "state":0, 
            "isRoot": "true"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":   
                {}
        }             
```

# 新增文章
```
	uri:/admin/article/add
    query:
        {
            "title": "hello",   //必填
            "cover": "ddddd",
            "content": "this is content",   //必填
            "tags": "ffff"  //必填
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":   
                {
                    "sn":12323
                }
        }
```        
        
# 查询文章详情
```
	uri:/admin/article/info
    query:
        {
            "sn":123
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":   
                {
                    "sn":12323
                    "title": "hello",
                    "author":"chen",
                    "cover": "ddddd",
                    "content": "this is content",
                    "tags": "ffff",
                    "state":0,
                    "createdAt":1232312,
                    "updatedAt":232323 
                }
        }
```

# 查询文章列表
```
	uri:/admin/article/list
    query:
        {
            "isAllMyselfArticles":true,   //默认false,根据条件搜索所有的文章，否则查询自身所有文章
            "article":                //以下条件根据页面提供的搜索条件进行组合查询
            {
                    "aid":123,  //精确查询
                    "sn":12323  //精确查询
                    "title": "hello",
                    "uid":122,  //精确查询
                    "tags": "ffff", //tag之间使用逗号隔开
                    "state":0,  //精确查询
                    "viewNum":true, //默认false，根据浏览量倒序查询
                    "cmtNum":true,  //默认false，根据评论量倒序查询
                    "zanNum":true   //默认false，根据点赞数倒序查询
            }
            "page":
            {
                    "pageNum":1,
                    "pageSize":10
            }
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":
                {
                "articleDetailList":
                    [{
                    "sn":12323
                    "title": "hello",
                    "uid":122,
                    "author":"chen",
                    "cover": "ddddd",
                    "content": "this is content",
                    "tags": "ffff",
                    "state":0,
                    "viewNum":22,
                    "cmtNum":1,
                    "zanNum":22 
                },{
                    "sn":12323
                    "title": "hello",
                    "uid":122,
                    "author":"chen",
                    "cover": "ddddd",
                    "content": "this is content",
                    "tags": "ffff",
                    "state":0,
                    "viewNum":1,
                    "cmtNum":1,
                    "zanNum":2 
                }]
                }             
        }



