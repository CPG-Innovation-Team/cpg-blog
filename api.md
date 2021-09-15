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
```

# 查询所有权限
```
	uri:/auth/query/permissions
    query:
        {}
    response:
        {
            "code":"code",
            "message": "message",
            "data":
                {
                "权限1": "/admin/auth/query/roles",
                "权限2": "/permission/2",
                "权限4": "/permission/4"
                }
        }
```

# 查询所有角色
```
	uri:/auth/query/roles
    query:
        {}
    response:
        {
            "code":"code",
            "message": "message",
            "data":
                {
                    "role1": ["/permission/2", "/admin/auth/query/roles"]
                }
        }
```

# 添加单个权限
```
	uri:/auth/add/permission
    query:
        {    
        "pName": "权限4",
        "uri":"/permission/4"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 删除单个权限
```
	uri:/auth/delete/permission
    query:
        {    
        "pName": "权限4"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 新增角色
```
	uri:/auth/add/role
    query:
        {    
        "rName": "角色名"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 角色添加权限
```
	uri:/auth/role/add/permission
    query:
        {
        "rName": "角色名",
        "pName": "权限名"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 用户添加角色
```
	uri:/auth/role/add/user
    query:
        {
        "rName": "角色名",
        "uid": 111
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 用户删除角色
```
	uri:/auth/role/remove/user
    query:
        {
        "rName": "角色名",
        "uid": 111
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 删除角色
```
	uri:/auth/delete/role
    query:
        {
        "rName": "角色名"
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

