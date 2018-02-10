package main

import (
	"goweb/controller"
	"goweb/core"
	"goweb/tool/FyLogger"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

// regList , 注册所有的控制器
var regList = make(map[string]interface{})

// init , 所有的控制器必须手动注册进来，这点和PHP完全不同，PHP可以动态引入，而GO不行
// go 只可反射已经存在的struct，而这个存在，则必须是预先加载好的
func init() {
	FyLogger.DebugLog("server init....")
	regList["index"] = &controller.IndexController{}
	regList["psych"] = &controller.PsychController{}
}

// main func
func main() {
	Start(":8080")
}

// Start func
func Start(port string) {
	server := &http.Server{
		Handler:        newMux(),
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	FyLogger.DebugLog("server start at " + port)
	log.Fatalln(server.ListenAndServe())
}

func newMux() *FyMux {
	m := &FyMux{}
	m.p.New = func() interface{} {
		return &core.Context{}
	}

	return m
}

// FyMux , 自定义的路由
// 每次连接过来，使用异步的连接池来处理请求
type FyMux struct {
	p sync.Pool
}

// 根据go的语法，实现了某接口的方法，就是继承了某个接口
// 所以，在此处，只要实现了ServeHTTP方法，就是继承了handler接口
// 注：handler接口是go专门用来处理http请求的
func (m *FyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	startTime := time.Now()
	// 连接池里面保存的接口为： core.Context
	ctx := m.p.Get().(*core.Context)
	// 不记录fav
	defer FyLogger.AccessLog(ctx, startTime)
	defer m.p.Put(ctx)
	// 设置w和r
	ctx.Set(w, r)

	// 往response的头里面写数据
	w.Header().Set("Server", "FyServer V1.0")

	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	pathInfo := strings.Split(path, "/")
	defaultController := "index"
	defaultAction := "Index"

	if len(pathInfo) > 1 {
		defaultController = pathInfo[1]
	}
	if len(pathInfo) > 2 {
		defaultAction = strings.Title(strings.ToLower(pathInfo[2]))
	}
	ctl, ok := regList[defaultController]
	if !ok {
		//http.NotFound(w, r)
		FyLogger.DebugLog("controller not found", defaultController)
		ctx.NotFound()
		return
	}

	//defaultController += "Controller"
	defaultAction += "Action"

	ref := reflect.ValueOf(ctl)
	act := ref.MethodByName(defaultAction)
	if !act.IsValid() {
		FyLogger.DebugLog("action not found", defaultController+":"+defaultAction)
		ctx.NotFound()
		return
	}
	/*
		refR := reflect.ValueOf(r)
		refW := reflect.ValueOf(w)
		act.Call([]reflect.Value{refW, refR})
	*/
	controller := ref.Interface().(core.IApp)
	controller.Init(ctx)
	act.Call(nil)
}
