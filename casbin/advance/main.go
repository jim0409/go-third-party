package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:100;uniqueIndex:unique_index"`
	V0    string `gorm:"size:100;uniqueIndex:unique_index"`
	V1    string `gorm:"size:100;uniqueIndex:unique_index"`
	V2    string `gorm:"size:100;uniqueIndex:unique_index"`
	V3    string `gorm:"size:100;uniqueIndex:unique_index"`
	V4    string `gorm:"size:100;uniqueIndex:unique_index"`
	V5    string `gorm:"size:100;uniqueIndex:unique_index"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:secret@tcp(127.0.0.1:3306)/mysql?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection to MySQL DB error : %v", err)
		return
	}
	if err != nil {
		log.Printf("連接數據庫錯誤: %v", err)
		return
	}

	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	if err != nil {
		log.Printf("配置數據庫錯誤: %v", err)
		return
	}

	// e, err := casbin.NewEnforcer("rbac_models.conf", a)
	e, err := casbin.NewSyncedEnforcer("rbac_models.conf", a)
	if err != nil {
		log.Printf("初始化casbin錯誤: %v", err)
		return
	}

	// 從 DB 加載策略
	e.LoadPolicy()

	//获取router路由对象
	r := gin.New()
	//增加policy
	r.POST("/api/v1/add", func(c *gin.Context) {
		// fmt.Println("增加Policy")

		if ok, _ := e.AddRoleForUser("admin", "administrator"); !ok {
			// fmt.Println("Policy已存在")
			c.JSON(204, map[string]string{
				"msg": "policy already exists",
			})
		} else {
			// fmt.Println("增加成功")
			c.JSON(200, map[string]string{
				"msg": "add policy success",
			})
		}

		if ok, _ := e.AddPolicy("administrator", "/api/v1/hello", "GET"); !ok {
			// fmt.Println("Policy已存在")
			c.JSON(204, map[string]string{
				"msg": "policy already exists",
			})
		} else {
			// fmt.Println("增加成功")
			c.JSON(200, map[string]string{
				"msg": "add policy success",
			})
		}

		if ok, _ := e.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
			// fmt.Println("Policy已存在")
			c.JSON(204, map[string]string{
				"msg": "policy already exists",
			})
		} else {
			// fmt.Println("增加成功")
			c.JSON(200, map[string]string{
				"msg": "add policy success",
			})
		}
		c.Abort()
	})
	//删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		// fmt.Println("删除Policy")
		if ok, _ := e.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
			c.JSON(404, map[string]string{
				"msg": "policy not exist",
			})
		} else {
			c.JSON(200, map[string]string{
				"msg": "success delete",
			})
		}

		c.Abort()
	})
	//获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		// fmt.Println("查看policy")
		list := e.GetPolicy()
		m := make(map[int]interface{})

		for _, vlist := range list {
			for k, v := range vlist {
				m[k] = v
			}
		}
		c.JSON(200, m)

		c.Abort()
	})
	//使用自定义拦截器中间件
	r.Use(Authorize(e))

	//创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("start listening service..")
	})

	r.Run(":9000") // 開始監聽 port 9000
}

// middleware 做權限驗證
// func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
func Authorize(e *casbin.SyncedEnforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		// 獲取請求的 URI
		obj := c.Request.URL.RequestURI()

		// 獲取請求的方法
		act := c.Request.Method

		// 獲取用戶的角色，預設是 admin
		sub := "admin"

		//判断策略中是否存在

		if ok, _ := e.Enforce(sub, obj, act); ok {
			// fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			// fmt.Println("很遗憾,权限验证没有通过")
			c.JSON(403, map[string]string{
				"msg": "forbidden",
			})
			c.Abort()
		}
	}
}
