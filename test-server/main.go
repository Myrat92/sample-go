package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.GET("/fib", func(c *gin.Context) {
		num := c.Query("num")
		n, err := strconv.Atoi(num)
		if err != nil {
			c.String(http.StatusBadRequest, "Request not a number format")
		}

		res := fib(n)
		c.String(http.StatusOK, "Fibonacci is %s", strconv.Itoa(res))
	})

	r.Run()
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
