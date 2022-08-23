package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grahms/fronteira/oidc"
	"net/http"
	"net/url"
)

type Request struct {
	engine     *gin.Engine
	targetHost string
}

func NewRequest(targetHost string) *Request {
	r := gin.Default()

	req := Request{
		engine:     r,
		targetHost: targetHost,
	}
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

func (r *Request) MakeHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		endpoint := r.targetHost + c.Request.RequestURI
		c.Redirect(http.StatusMultipleChoices, endpoint)
	}
}

func (r *Request) AuthMiddleware(op Operation) gin.HandlerFunc {

	oidcParams := oidc.InitParams{
		Router:       r.engine,
		ClientId:     "xx-xxx-xxx",
		ClientSecret: "xx-xxx-xxx",
		Issuer:       url.URL{Host: "https://accounts.google.com/"}, //add '.well-known/openid-configuration' to see it's a good link
		ClientUrl:    url.URL{Host: r.targetHost},                   //your website's url
		Scopes:       op.Scopes,
		ErrorHandler: func(c *gin.Context) {
			//gin_oidc pushes a new error before any "ErrorHandler" invocation
			message := c.Errors.Last().Error()
			//redirect to ErrorEndpoint with error message
			//redirectToErrorPage(c, "http://example2.domain/error", message)
			//when "ErrorHandler" ends "c.Abort()" is invoked - no further handlers will be invoked
			c.JSON(500, message)
			c.Abort()
		},
		PostLogoutUrl: url.URL{Host: r.targetHost},
	}

	return oidc.Init(oidcParams)
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
