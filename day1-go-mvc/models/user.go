package models

import (
    "errors"
    "sort"
)

// User 结构体
type User struct {
    ID     int
    Name   string
    Active bool
    Role   string
}

// GetUsers 返回用户 slice，支持排序
func GetUsers(sortBy string) []User {
    users := []User{
        {ID: 1, Name: "Alice", Active: true, Role: "Admin"},
        {ID: 2, Name: "Bob", Active: false, Role: "User"},
        {ID: 3, Name: "Charlie", Active: true, Role: "User"},
        {ID: 4, Name: "Alice Smith", Active: true, Role: "Admin"},
    }

    // if 语句处理排序
    if sortBy == "name" {
        sort.Slice(users, func(i, j int) bool {
            return users[i].Name < users[j].Name
        })
    } else if sortBy == "id" {
        sort.Slice(users, func(i, j int) bool {
            return users[i].ID < users[j].ID
        })
    }
    return users
}

// GetUserByID 根据 ID 获取用户
func GetUserByID(id int) (User, error) {
    users := GetUsers("")
    for _, user := range users { // for 循环遍历 slice
        if user.ID == id {
            return user, nil
        }
    }
    return User{}, errors.New("用户未找到")
}

// FilterUsersByName 根据名字过滤
func FilterUsersByName(name string) []User {
    users := GetUsers("")
    var filtered []User
    for _, user := range users { // for 循环实现过滤
        if user.Name == name {
            filtered = append(filtered, user) // 动态扩展 slice
        }
    }
    return filtered
}

// FilterUsersByActive 根据活跃状态过滤
func FilterUsersByActive(active bool) []User {
    users := GetUsers("")
    var filtered []User
    for _, user := range users { // for 循环实现过滤
        if user.Active == active {
            filtered = append(filtered, user)
        }
    }
    return filtered
}

// GetRoleStats 返回角色统计 map
func GetRoleStats() map[string]int {
    users := GetUsers("")
    roleStats := make(map[string]int) // 初始化 map
    for _, user := range users { // for 循环统计角色
        roleStats[user.Role]++ // map 键值操作
    }
    return roleStats
}
