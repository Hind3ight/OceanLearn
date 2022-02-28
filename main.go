package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":8082")
}
