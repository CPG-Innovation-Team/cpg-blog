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

# 登出
```
	uri:/admin/logout
    query:
        {}
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
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

# 查询用户列表
```
	uri:/admin/user/query/list
    query:{"state":0}//只能为1\2\3 用户状态 1-正常;2-禁发文;3-冻结
    response:
        {
            "code":"code",
            "message": "message",
            "data": [
            {
            "UID": 111,
            "UserName": "chenxi",
            "Email": "chenxi@qq.com",
            "Nickname": "",
            "Avatar": "",
            "Gender": 0,
            "Introduce": "",
            "State": 0,
            "IsRoot": 1
            }
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

# 删除文章
```
	uri:/admin/article/delete
    query:
        {"sn": 12334}
    response:
        {
            "code":"code",
            "message": "message",
            "data":""
        }
```

# 更新文章
```
	uri:/admin/article/update
    query:
        {
            "sn":"1", //必传
            "title": "update22",
            "cover": "update22",
            "content": "update222",
            "tags": "update222",
            "state":"" //必传，且不能为空
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":""
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
                    "sn":"12323"
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
                    "state":1,  //必传参数，精确查询，默认查询未上线文章 文章状态 0-未审核;1-已上线;2-下线;3-用户删除'
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
                }],
            "PageNum": 1,
            "PageSize": 10,
            "Total": 15,    //总条数
            "TotalPage": 2  //总页数
                }             
        }
```

# 查询所有权限
```
	uri:/admin/auth/query/permissions
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
	uri:/admin/auth/query/roles
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
	uri:/admin/auth/add/permission
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
	uri:/admin/auth/delete/permission
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
	uri:/admin/auth/add/role
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
	uri:/admin/auth/role/add/permission
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
	uri:/admin/auth/role/add/user
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
	uri:/admin/auth/role/remove/user
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
	uri:/admin/auth/delete/role
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

# 点赞文章/评论
```
	uri:/like
    query: //点赞,参数不允许同时为空/存在，二选一
        {
            "sn":"1967400744308736",
            "comment_id":0
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 取消点赞文章/评论
```
	uri:/like/cancel
    query: //取消点赞,参数不允许同时为空/存在，二选一
        {
            "sn":"1967400744308736",
            "comment_id":0
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 新增评论
```
	uri:/admin/comment/add
    query: //必填
        {
            "sn":1967400744308736,
            "content":""
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
            "commentId":122
            }
        }
```

# 删除评论
```
	uri:/admin/comment/delete
    query: //必填
        {
            "commentId":1967400744308736
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 新增评论回复
```
	uri:/comment/reply
    query: //必填
        {
            "commentId":111,
            "content":""
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
            "commentReplyId":122
            }
        }
```

# 删除评论回复
```
	uri:/comment/reply/delete
    query: //必填
        {
            "id":111 //回复id
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 查询文章所有评论及回复
```
	uri:/comment/list
    query: //必填
        {
            "sn":111
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
                1:{ //楼层
	                cid: 111, //自增ID
	                sn: 111, //文章sn号
		            uid: 111, //评论用户uid
		            content: "hello", //评论内容
		            zanNum： 111, //点赞数
		            floor： 1, //第几楼
		            state： 1, //状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
		            createdAt: 2020.0189
		            replyList:[{
		                	id: 111, //自增ID
	                        cid: 111, //评论cid
	                        uid: 111, //回复用户uid
	                        content: "hello", //回复内容
	                        state： 1 //状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	                        createdAt: 2020.0189
		                    },
		                    {
		                	id: 111, //自增ID
	                        cid: 111, //评论cid
	                        uid: 111, //回复用户uid
	                        content: "hello", //回复内容
	                        state： 1 //状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	                        createdAt: 2020.0189
		                    }]
		            }
            }
        }
```

# 查询所有需要审核的文章
```
	uri:/admin/review/query/article/list
    query:{}
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
              "ArticleMap": {
                "21828941648556032": {
                     "Aid": 1,
                     "Sn": 21828941648556032,
                     "Title": "test",
                     "Uid": 1,
                     "Cover": "www.baidu.com",
                     "Content": "this is test",
                     "Tags": "test",
                     "State": 0,
                     "CreatedAt": "2021-11-17T13:40:25.252+08:00",
                     "UpdatedAt": "2021-11-17T17:28:21.634+08:00"
                }}}
        }
```

# 查询所有审核失败的文章
```
	uri:/admin/review/query/article/failed/list
    query:{}
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
              "ArticleMap": {
                "21828941648556032": {
                     "Aid": 1,
                     "Sn": 21828941648556032,
                     "Title": "test",
                     "Uid": 1,
                     "Cover": "www.baidu.com",
                     "Content": "this is test",
                     "Tags": "test",
                     "State": 0,
                     "CreatedAt": "2021-11-17T13:40:25.252+08:00",
                     "UpdatedAt": "2021-11-17T17:28:21.634+08:00"
                }}}
        }
```

# 审核文章
```
	uri:/admin/review/article
    query:
        {
          "sn":"21828941648556032",
          "state": false
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 查询所有需要审核的评论
```
	uri:/admin/review/query/comment/list
    query:{}
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
                "CommentMap": {
                    "1": {
                    "Cid": 1,
                    "Sn": 21828941648556032,
                    "UID": 1,
                    "Content": "add content 01",
                    "ZanNum": 0,
                    "Floor": 1,
                    "State": 0,
                    "CreatedAt": "2021-11-17T17:12:35+08:00"
                        }
                    }
                }
        }
```

# 查询所有审核失败的评论
```
	uri:/admin/review/query/comment/failed/list
    query:{}
        response:
        {
            "code":"code",
            "message": "message",
            "data":{
                "CommentMap": {
                    "1": {
                    "Cid": 1,
                    "Sn": 21828941648556032,
                    "UID": 1,
                    "Content": "add content 01",
                    "ZanNum": 0,
                    "Floor": 1,
                    "State": 0,
                    "CreatedAt": "2021-11-17T17:12:35+08:00"
                        }
                    }
                }
        }
```

# 审核评论
```
	uri:/admin/review/comment
    query:
        {
            "commentId": 1,
            "state": false
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```

# 查询所有需要审核的回复
```
	uri:/admin/review/query/reply/list
    query:{}
    response:
        {
            "code":"code",
            "message": "message",
            "data":{
                "ReplyMap": {
                    "1": {
                    "Id": 1,
                    "Cid": 1,
                    "UID": 1,
                    "Content": "add reply",
                    "State": 0,
                    "CreatedAt": "2021-11-17T17:26:26+08:00"
                    }
                }    
            }
        }
```

# 查询所有审核失败的回复
```
	uri:/admin/review/query/reply/failed/list
    query:{}
        response:
        {
            "code":"code",
            "message": "message",
            "data":{
                "ReplyMap": {
                    "1": {
                    "Id": 1,
                    "Cid": 1,
                    "UID": 1,
                    "Content": "add reply",
                    "State": 0,
                    "CreatedAt": "2021-11-17T17:26:26+08:00"
                    }
                }    
            }
        }
```

# 审核回复
```
	uri:/admin/review/reply
    query:
        {
            "commentId": 1,
            "state": false
        }
    response:
        {
            "code":"code",
            "message": "message",
            "data":{}
        }
```