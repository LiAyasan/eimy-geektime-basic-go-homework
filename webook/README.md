# 第二周作业

> 当前项目功能对标第二周第六讲的项目

* 要求
  * 补充 /users/edit 接口，允许用户补充信息：昵称、生日、个人简介
  * 补充 /users/profile 接口，展示用户信息

## V0.1

### POST /users/edit

> 更新于 2023/08/13

**Request**

| 参数名      | 类型     | 允许为空 | 备注            |
|----------|--------|------|---------------|
| nickname | string | true | 长度不超过8        |
| birthday | string | true | 格式：0000-00-00 |
| details | string | true | 长度不超过300      |

字段传空代表保持原状。

若 details 原本为空并传空值则修改成默认值：“此人很神秘，什么也没写”。

需要登录态。

**Response**

* 若 nickname 长度超过8，则提示“昵称长度不合法”；
* 若 birthday 不符合正常时间，则提示“生日日期有误”；
* 若 details 长度超过300，则提示“简介长度不合法”；
* 若出现其他错误，则提示“系统故障”；
* 若没出现错误，则提示“修改成功”。

### GET /users/profile

请求返回，返回一串字符串

![/users/profile接口请求](/image/img.png)

### user表

```sql
ALTER TABLE users ADD nickname VARCHAR(8);
ALTER TABLE users ADD birthday VARCHAR(10);
ALTER TABLE users ADD details VARCHAR(300);
```