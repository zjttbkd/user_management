# entrytask设计文档

## 接口设计

### http服务

TODO

### grpc服务

#### pb接口

接口名|输入包 | 返回包
------|-----|-------
登录鉴权|LoginRequest|UsrInfoReply
查询信息|QueryRequest|UsrInfoReply
修改图片|ProfileRequest|UsrInfoReply
修改昵称|NicknameRequest|UsrInfoReply

#### 请求包字段

- LoginRequest:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
password|string|密码

- QueryRequest:

字段名 | 类型 | 说明
------|-----|----
username|string|账号

- ProfileRequest:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
profile|string|修改后的图片地址

- NicknameRequest:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
nickname|string|修改后的昵称

#### 返回包字段

- UsrInfoReply:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
nickname|string|昵称
profile|profile|图片地址


## redis设计

### 键值对

- key: username
- value: { "Profile": profile, "Nickname": nickname }

### 接口函数

函数名 | 输入字段 | 输出字段 | 功能描述
------|---------|---------|-------
cacheUserInfo|struct{ usernam,nickname, profile }|-|增加缓存
fetchUserInfo|username|struct{ usernam,nickname, profile }|获取缓存
delUserInfo|username|-|修改昵称、图片后删除缓存


## 数据库设计

### 用户表

user_mgn_db.user_info_tab_0000000x

字段名 | 类型 | 说明
------|-----|----
uid|bigint unsigned|内部id（自增主键）
username|varhcar(128)|账号（索引）
password|varchar(128)|密码
nickname|varchar(255)|昵称
profile|varchar(255)|图片地址
create_time|int unsigned|创建时间
modify_time|int unsigned|修改时间（索引）

### 接口函数

函数名 | 输入字段 | 输出字段 | 功能描述
------|---------|---------|-------
queryInfo|usernam|password, nickname, profile|获取信息
uploadProfile|username, profile|-|修改图片地址
changeNickname|username, nickname|-|修改昵称
