package middleware

import (
	"blog/common/helper"
	"blog/common/response"
	"blog/common/response/errx"
	"blog/common/static/cKey"
	"blog/database/model"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type PermissionMiddleware struct {
	DB    *gorm.DB
	Cache *redis.Redis
}

func NewPermissionMiddleware(db *gorm.DB, rds *redis.Redis) *PermissionMiddleware {
	return &PermissionMiddleware{
		DB:    db,
		Cache: rds,
	}
}

func (m *PermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := r.RequestURI
		roleIdStr := r.Context().Value("roleId").(string)
		roleId, _ := strconv.Atoi(roleIdStr)
		//不是超级管理员
		if roleId != 1 {
			permission := m.getPermission(roleIdStr)
			isPass := false
			for _, per := range permission {
				pList := strings.Split(per.URI, "\n")
				for _, p := range pList {
					//直接匹配
					if p == route {
						isPass = true
					}
					index := strings.Index(p, "*")
					if index > -1 {
						//去掉* 然后进行查找
						//示例 当前接口r1:/abc/ad/ee  限制：r2:/abc/ad*，如果r1中能找到/abc/ad，说明r1适用于*通配
						p := strings.Replace(p, "*", "", -1)
						if strings.Index(route, p) > -1 {
							//说明命中
							isPass = true
						}
					}
				}
			}
			if isPass == false {
				resp := response.Response{
					Code:    errx.PermissionError,
					Message: errx.GetCnMessage(errx.PermissionError),
				}
				str, _ := json.Marshal(resp)
				w.Write(str)
				return
			}
		}

		next(w, r)
	}
}

func (m *PermissionMiddleware) getPermission(roleStr string) []model.AdminPermission {
	var permission []model.AdminPermission
	cacheKey := cKey.AdminPermissionKey + roleStr
	permissionJson, _ := m.Cache.Get(cacheKey)
	if len(permissionJson) < 1 {
		var menu []model.AdminRolePermission
		m.DB.Where("role_id = ?", roleStr).Find(&menu)
		var menuId []int32
		for _, v := range menu {
			menuId = append(menuId, v.MenuID)
		}
		m.DB.Debug().Where("menu_id in (?)", menuId).Find(&permission)
		jData, _ := json.Marshal(permission)
		m.Cache.Setex(cacheKey, string(jData), cKey.AdminPermissionTtl)
	} else {
		helper.Swap(permissionJson, &permission)
	}
	return permission
}
