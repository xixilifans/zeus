package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"zeus/pkg/api/log"

	"github.com/beego/i18n"
	"github.com/gin-gonic/gin"
)

// Ignored permissions
var ignoredPerms = map[string]bool{}

// PermCheck - check permission automatically
func PermCheck(c *gin.Context) {
	route := strings.Split(c.Request.URL.RequestURI(), "?")[0]
	for _, p := range c.Params {
		route = strings.Replace(route, "/"+p.Value, "/:"+p.Key, 1)
	}
	route = strings.ToLower(c.Request.Method) + "@" + route
	uid := fmt.Sprintf("%#v", c.Value("userId"))
	if _, ok := ignoredPerms[route]; ok {
		c.Next()
		return
	}
	if accountService.CheckPermission(uid, "root", route) {
		log.Info(fmt.Sprintf("Pass permission check for %s", route))
	} else if accountService.CheckPermission(uid, "gm", route) {
		log.Info(fmt.Sprintf("Pass permission check for %s", route))
	} else {
		log.Warn(fmt.Sprintf("No permission for %s", route))
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  i18n.Tr(GetLang(), "err.Err403"),
		})
		c.Abort()
		return
	}
	c.Next()
}
