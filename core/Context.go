package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context ， 上下文对象
type Context struct {
	RseW       http.ResponseWriter
	Req        *http.Request
	outMessage *httpMessage
}

// Set 方法，初始化上下文对象
func (c *Context) Set(w http.ResponseWriter, r *http.Request) {
	c.RseW = w
	c.Req = r
	c.outMessage = new(httpMessage)
	c.outMessage.statusCode = http.StatusOK
	c.outMessage.jsonData = new(jsonMessage)
	c.outMessage.jsonData.ReSet()
}

// Redirect 302 response
func (c *Context) Redirect(url string) {
	http.Redirect(c.RseW, c.Req, url, 302)
}

// Echo , 写入ResponseWrite
func (c *Context) Echo(args interface{}) {
	fmt.Fprintln(c.RseW, args)
}

// Show , 输出JSON
func (c *Context) Show() {
	output := c.outMessage.jsonData.ToMap()
	msg, err := json.Marshal(output)
	if err != nil {
		c.outMessage.statusCode = 500
		output = c.outMessage.jsonData.ReSet().ToMap()
		msg, _ = json.Marshal(output)
	}
	c.RseW.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.RseW.Header().Set("X-Content-Type-Options", "nosniff")
	c.RseW.WriteHeader(c.outMessage.statusCode)

	c.Echo(string(msg))
}

// NotFound , 自定义来处理404函数
func (c *Context) NotFound() {
	c.SetStatusCode(http.StatusNotFound).SetDataMessage("Query Not Found").Show()
}

// HTTPError ，返回502
func (c *Context) HTTPError() {
	c.SetStatusCode(http.StatusBadGateway).SetDataMessage("Service Error").Show()
}

// SetData , 需要输出的json数据
func (c *Context) SetData(data interface{}) *Context {
	c.outMessage.jsonData.data = data
	return c
}

// SetDataMessage , 需要输出json数据的message
func (c *Context) SetDataMessage(msg string) *Context {
	c.outMessage.jsonData.message = msg
	return c
}

// SetDataCode , 数据对应的code(业务定义的code)
func (c *Context) SetDataCode(code int) *Context {
	c.outMessage.jsonData.code = code
	return c
}

// GetStatusCode 返回本次请求的http状态码
func (c *Context) GetStatusCode() int {
	return c.outMessage.statusCode
}

// SetStatusCode , 返回的http请求的状态码
func (c *Context) SetStatusCode(code int) *Context {
	c.outMessage.statusCode = code
	return c
}
