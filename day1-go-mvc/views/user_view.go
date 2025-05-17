package views

import (
	"day1-go-mvc/models"
	"fmt"
	"strings"
)

// RenderUsers 渲染用户列表
func RenderUsers(users []models.User, roleStats map[string]int) string {
	var builder strings.Builder
	builder.WriteString("用户列表:\n")
	for _, user := range users { // for 循环遍历 slice
		status := "非活跃"
		if user.Active { // if 语句检查状态
			status = "活跃"
		}
		builder.WriteString(fmt.Sprintf("ID: %d, 名字: %s, 状态: %s, 角色: %s\n", user.ID, user.Name, status, user.Role))
	}
	// 添加角色统计
	builder.WriteString("\n角色统计:\n")
	for role, count := range roleStats { // for 循环遍历 map
		builder.WriteString(fmt.Sprintf("%s: %d\n", role, count))
	}
	return builder.String()
}

// RenderUser 渲染单个用户
func RenderUser(user models.User) string {
	status := "非活跃"
	if user.Active { // if 语句检查状态
		status = "活跃"
	}
	return fmt.Sprintf("用户: ID: %d, 名字: %s, 状态: %s, 角色: %s\n", user.ID, user.Name, status, user.Role)
}

// RenderError 渲染错误信息
func RenderError(err error) string {
	return fmt.Sprintf("错误: %s\n", err.Error())
}
