package main

// func init() {
// 	gin.SetMode(gin.ReleaseMode)
// }

// func main() {
// 	routes := gin.Default()
// 	routes.Use(CORSMiddleware())

// 	// test server connection
// 	routes.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	routes.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "welcome.html", gin.H{
// 			"content": "welcome to youtube analytics",
// 		})
// 	})

// 	// svc routes
// 	svc := routes.Group("/svc")
// 	{
// 		svc.GET("/get_insights", handlers.GetInsights)
// 	}

// 	// run server
// 	fmt.Println("server running on port.. :", 8080)
// 	routes.Run(fmt.Sprintf(":%v", 8080))
// }

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// c.Writer.Header().Set("Content-Type", "*; application/json; charset=utf-8; application/x-www-form-urlencoded;")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 		} else {
// 			c.Next()
// 		}

// 		c.Next()
// 	}
// }
