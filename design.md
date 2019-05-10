# entrytask设计文档

## 接口设计

### pb接口

#### Login

- request:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
password|string|密码
authorized|bool|免密

- reply:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
nickname|string|昵称
profile|profile|图片地址

#### UploadProfile

- request:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
profile|string|修改后的图片地址

- reply:

字段名 | 类型 | 说明
------|-----|----
result|string|结果信息

#### ChangeNickname

- request:

字段名 | 类型 | 说明
------|-----|----
username|string|账号
nickname|string|修改后的昵称

- reply:

字段名 | 类型 | 说明
------|-----|----
result|string|结果信息

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
password|username, nickname|-|修改昵称
