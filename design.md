# entrytask设计文档

## 接口设计

### http接口


### pb接口

#### Login

request

字段名 | 类型 | 说明
------|-----|----
username|string|账号
password|string|密码

reply

字段名 | 类型 | 说明
------|-----|----
username|string|账号
nickname|string|昵称
profile|bytes|图片地址

#### UploadProfile

#### ChangeNickname

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
