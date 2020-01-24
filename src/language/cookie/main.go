package main

import "fmt"

func main() {
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			// 设置cookie
			c.SetCookie(
				"key_cookie",   // 设置cookie的key
				"value_cookie", // 设置cookie的值
				60,             // 过期时间
				"/",            // 所在目录
				"127.0.0.1",    //域名
				false,          // 是否只能通过https访问
				true)           // 是否允许别人通过js获取自己的cookie
		}
		fmt.Println("cookie的值是:", cookie)
	})
	r.Run(":8000")
}
