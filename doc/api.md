## /user

#### 公共Header参数

| 参数名 | 示例值 | 参数描述      |
| --- | --- |-----------|
| Authorization |  | jwt-token |

#### 公共Query参数

| 参数名 | 示例值 | 参数描述 |
| --- | --- | ---- |
| 暂无参数 |

#### 公共Body参数

| 参数名 | 示例值 | 参数描述 |
| --- | --- | ---- |
| 暂无参数 |

## /user/发送验证码

#### 接口URL

> 192.168.61.128:8888/captcha

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |
| Author | - | String | 是 | - |

#### 请求Body参数

```json
{"email":"3239089828@qq.com"}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| email | 3239089828@qq.com | String | 是 | 登录邮箱 |

## /user/注册并登录

> 192.168.61.128:8888/register

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |
| Authorization | - | String | 是 | - |

#### 请求Body参数

```json
{"email":"3239089828@qq.com","captcha":"950478","name":"muskonu","password":"12345678"}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| email | 3239089828@qq.com | String | 是 | 登录邮箱 |
| captcha | - | String | 是 | 邮箱验证码 |
| name | muskonu | String | 是 | 用户名 |
| password | 12345678 | String | 是 | 登陆密码 |

#### 成功响应示例

```json
{
    "200": {
        "accessToken": "",
        "accessExpire": "",
        "refreshAfter": "",
        "name": ""
    }
}
```

## /user/邮箱密码登录

#### 接口URL

> 192.168.61.128:8888/login/pass

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{"email":"3239089828@qq.com","password":"12345678"}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| email | 3239089828@qq.com | String | 是 | 登录邮箱 |
| password | 12345678 | String | 是 | - |

#### 成功响应示例

```json
{"200":{"accessToken":"","accessExpire":"","refreshAfter":"","name":""}}
```

## /user/修改用户密码

#### 接口URL

> 192.168.61.128:8888/password

#### 请求方式

> PUT

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{"email":"3239089828@qq.com","password":"12345678"}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| email | 3239089828@qq.com | String | 是 | 登录邮箱 |
| password | 12345678 | String | 是 | 登陆密码 |

## /user/修改用户名

#### 接口URL

> 192.168.61.128:8888/user/name

#### 请求方式

> PUT

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{"name":"muskonu"}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| name | muskonu | String | 是 | 用户名 |


## /user/邮箱验证码登录

#### 接口URL

> 192.168.61.128:8888/login/captcha

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{
    "email": "3239089828@qq.com",
    "captcha": ""
}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| email | 3239089828@qq.com | String | 是 | 登录邮箱 |
| captcha | - | String | 是 | 邮箱验证码 |

#### 成功响应示例

```json
{"200":{"accessToken":"","accessExpire":"","refreshAfter":"","name":""}}
```

## /todo

```text
暂无描述
```

#### 公共Header参数

| 参数名 | 示例值 | 参数描述      |
| --- | --- |-----------|
| Authorization |  | jwt token |

#### 公共Query参数

| 参数名 | 示例值 | 参数描述 |
| --- | --- | ---- |
| 暂无参数 |

#### 公共Body参数

| 参数名 | 示例值 | 参数描述 |
| --- | --- | ---- |
| 暂无参数 |

## /todo/删除待办事项

#### 接口URL

> 192.168.61.128:8888/todo

#### 请求方式

> DELETE

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{
    "id": 1
}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| id | 1 | String | 是 | - |


## /todo/改变完成情况

#### 接口URL

> 192.168.61.128:8888/todo

#### 请求方式

> PATCH

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{
    "id": 1
}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| id | 1 | String | 是 | - |

## /todo/添加待办事项

#### 接口URL

> 192.168.61.128:8888/todo

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{
    "content": "起床",
    "dueDate": 1695889860,
    "recurrence": 1
}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| content | 起床 | String | 是 | 待办事项内容 |
| dueDate | 1695889259 | Integer | 是 | 任务到期时间 |
| recurrence | 1 | String | 是 | 执行周期 |

## /todo/修改待办事项

#### 接口URL

> 192.168.61.128:8888/todo

#### 请求方式

> PUT

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Body参数

```json
{
    "id": 1,
    "content": "睡觉",
    "dueDate": 1695889259,
    "recurrence": 1
}
```

| 参数名        | 示例值 | 参数类型 | 是否必填 | 参数描述 |
|------------| --- | ---- | ---- | ---- |
| id         | 1 | String | 是 | - |
| content    | 睡觉 | String | 是 | 待办事项内容 |
| dueDate    | 1695889259 | String | 是 | 任务到期时间 |
| recurrence | 1 | String | 是 | - |

## /todo/查询待办事项

#### 接口URL

> 192.168.61.128:8888/todo?cursor=0&pageSize=5&isCompleted=true

#### 请求方式

> GET

#### Content-Type

> json

#### 请求Header参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Accept | application/json | Text | 是 | - |
| Content-Type | application/json | Text | 是 | - |

#### 请求Query参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| cursor | 0 | Integer | 是 | 游标 |
| pageSize | 5 | Integer | 是 | - |
| isCompleted | true | Boolean | 是 | - |

#### 成功响应示例

```json
{
    "200": {
        "todos": [
            {
                "id": "",
                "isCompleted": "",
                "content": "",
                "dueDate": "",
                "repeat": ""
            }
        ]
    }
}
```
