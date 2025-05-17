package controllers

import (
	"day1-go-mvc/models"
	"day1-go-mvc/views"
	"errors"
	"net/http"
	"strconv"
)

// UserController 处理用户请求
type UserController struct{}

// ListUsers 处理 GET 请求
func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// 获取查询参数
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	activeStr := r.URL.Query().Get("active")
	sortBy := r.URL.Query().Get("sortBy")

	// 获取角色统计 map
	roleStats := models.GetRoleStats()

	if idStr != "" { // if: 处理单个用户
		id, err := strconv.Atoi(idStr)
		if err != nil { // if: 验证 ID 格式
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(views.RenderError(err)))
			return
		}
		user, err := models.GetUserByID(id)
		if err != nil { // if: 检查用户是否存在
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(views.RenderError(err)))
			return
		}
		w.Write([]byte(views.RenderUser(user)))
	} else if name != "" { // else if: 按名字过滤
		users := models.FilterUsersByName(name)
		if len(users) == 0 { // if: 检查是否有匹配
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(views.RenderError(errors.New("未找到名字为 " + name + " 的用户"))))
			return
		}
		w.Write([]byte(views.RenderUsers(users, roleStats)))
	} else if activeStr != "" { // else if: 按活跃状态过滤
		active, err := strconv.ParseBool(activeStr)
		if err != nil { // if: 验证活跃状态格式
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(views.RenderError(errors.New("无效的 active 参数"))))
			return
		}
		users := models.FilterUsersByActive(active)
		if len(users) == 0 { // if: 检查是否有匹配
			status := "活跃"
			if !active {
				status = "非活跃"
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(views.RenderError(errors.New("未找到" + status + "用户"))))
			return
		}
		w.Write([]byte(views.RenderUsers(users, roleStats)))
	} else { // else: 返回所有用户
		users := models.GetUsers(sortBy)
		w.Write([]byte(views.RenderUsers(users, roleStats)))
	}
}
