package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	engine     *gin.Engine
	targetHost string
}

func NewRequest(targetHost string) *Request {
	r := gin.Default()
	req := Request{engine: r, targetHost: targetHost}
	req.NoRoute()
	return &req
}

func (r *Request) Handler(op Operation) {
	method := op.Method
	path := op.Path

	switch method {
	case "POST":
		r.post(path, r.AuthMiddleware(op), r.MakeHandler())
	case "GET":
		r.get(path, r.AuthMiddleware(op), r.MakeHandler())
	}

}

func (r *Request) NoRoute() {
	r.engine.NoRoute(func(c *gin.Context) {
		endpoint := r.targetHost + c.Request.RequestURI
		c.Redirect(http.StatusMultipleChoices, endpoint)
	})
}

func (r *Request) MakeHandler() func(*gin.Context) {

	return func(c *gin.Context) {
		endpoint := r.targetHost + c.Request.RequestURI
		c.Redirect(http.StatusMultipleChoices, endpoint)
	}
}

func (r *Request) AuthMiddleware(op Operation) func(*gin.Context) {

	return func(c *gin.Context) {
		c.Next()
		fmt.Printf("allowed roles %v", op.Roles)
		println()
	}
}
func (r *Request) post(endpoint string, middleware gin.HandlerFunc, handler gin.HandlerFunc) {

	r.engine.POST(endpoint, middleware, handler)
}

func (r *Request) get(endpoint string, middleware gin.HandlerFunc, handler gin.HandlerFunc) {
	r.engine.GET(endpoint, middleware, handler)
}

func (r *Request) Run(port string) {
	err := r.engine.Run(port)
	if err != nil {
		return
	}
}
