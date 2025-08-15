# 接口说明

> 所有接口返回文本格式：`Content-Type: text/plain; charset=utf-8`，包含用户信息和角色统计数据。

| 接口路径 | HTTP 方法 | 查询参数         | 功能描述                           | 返回示例                     |
|----------|------------|------------------|------------------------------------|------------------------------|
| /users   | GET        | 无参数           | 获取所有用户列表（默认按 ID 排序） | 用户列表（默认按 ID 排序）   |
| /users   | GET        | sortBy=id        | 获取按 ID 排序的用户列表           | 按 ID 排序的用户列表         |
| /users   | GET        | sortBy=name      | 获取按名字排序的用户列表           | 按名字排序的用户列表         |
| /users   | GET        | id=数字          | 获取指定 ID 的单个用户             | 单个用户信息                 |
| /users   | GET        | name=字符串      | 获取指定名字的用户列表             | 过滤后的用户列表             |
| /users   | GET        | active=true/false| 获取指定活跃状态的用户列表         | 过滤后的用户列表             |



###  所有用户（按 ID 排序）：curl http://localhost:8080/users?sortBy=id

```txt
用户列表:
ID: 1, 名字: Alice, 状态: 活跃, 角色: Admin
ID: 2, 名字: Bob, 状态: 非活跃, 角色: User
ID: 3, 名字: Charlie, 状态: 活跃, 角色: User
ID: 4, 名字: Alice Smith, 状态: 活跃, 角色: Admin

角色统计:
Admin: 2
User: 2
```

### 所有用户（按名字排序）：curl http://localhost:8080/users?sortBy=name

```txt
用户列表:
ID: 1, 名字: Alice, 状态: 活跃, 角色: Admin
ID: 4, 名字: Alice Smith, 状态: 活跃, 角色: Admin
ID: 2, 名字: Bob, 状态: 非活跃, 角色: User
ID: 3, 名字: Charlie, 状态: 活跃, 角色: User

角色统计:
Admin: 2
User: 2
```

### 单个用户：curl http://localhost:8080/users?id=1

```txt
用户: ID: 1, 名字: Alice, 状态: 活跃, 角色: Admin
按名字过滤：curl http://localhost:8080/users?name=Alice
用户列表:
ID: 1, 名字: Alice, 状态: 活跃, 角色: Admin

角色统计:
Admin: 2
User: 2
```

### 按活跃状态过滤：curl http://localhost:8080/users?active=true

```txt
用户列表:
ID: 1, 名字: Alice, 状态: 活跃, 角色: Admin
ID: 3, 名字: Charlie, 状态: 活跃, 角色: User
ID: 4, 名字: Alice Smith, 状态: 活跃, 角色: Admin

角色统计:
Admin: 2
User: 2
```




