# 测试用例：

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


