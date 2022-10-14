# 集合列表

## 用户集合
```json5
{
    "account":"账号",
    "password":"密码",
    "nickname":"昵称",
    "sex":1,  //0-未知 1-男 2-女
    "email":"邮箱",
    "avatar":"头像",
    "created_at":1, //创建时间
    "updated_at":1,  //更新时间
}
```
## 消息集合
```json5
{
    "user_identity": "用户的唯一标识",
    "room_identity": "房间的唯一标识",
    "data": "发送的数据",
    "created_at": 1,	//创建时间
    "updated_at": 1,	//更新时间
}
```

## 房间集合
```json5
{
  "number": "房间号",
  "name": "房间名称",
  "info": "房间简介",
  "user_indentity": "房间的创建者",
  "create_at": 1,	//房间的创建时间
  "update_at": 1,	//房间的更新时间
}
```
### 核心包
https://github.com/gorilla/websocket

### 扩展安装
```shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
```

