package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thinkeridea/go-extend/exnet"
)

type Interceptor func(*StatusResponseWriter, *http.Request, InterceptorHandlerFunc)

type InterceptorChain []Interceptor

type InterceptorHandlerFunc func(w *StatusResponseWriter, r *http.Request)

func (cont InterceptorHandlerFunc) Intercept(interceptor Interceptor) InterceptorHandlerFunc {
	return func(w *StatusResponseWriter, request *http.Request) {
		interceptor(w, request, cont)
	}
}

func (f InterceptorHandlerFunc) ServeHTTP(w *StatusResponseWriter, r *http.Request) {
	f(w, r)
}

// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
func HeaderInterceptor() Interceptor {
	return func(w *StatusResponseWriter, r *http.Request, next InterceptorHandlerFunc) {

		header := r.Header
		for key, value := range header {
			for _, vv := range value {
				w.Header().Add(key, vv)
			}

		}
		next(w, r)
	}
}

//2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func VersionInterceptor() Interceptor {
	return func(w *StatusResponseWriter, r *http.Request, next InterceptorHandlerFunc) {
		os.Setenv("VERSION", "1.1")
		version := os.Getenv("VERSION")
		w.Header().Set("VERSION", version)
		next(w, r)
	}
}

//3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func LogInterceptor() Interceptor {
	return func(w *StatusResponseWriter, r *http.Request, next InterceptorHandlerFunc) {

		ip := exnet.ClientPublicIP(r)
		if ip == "" {
			ip = exnet.ClientIP(r)
		}
		next(w, r)
		log.Printf("客户端IP为：%s \n", ip)
		log.Printf("HTTP 返回码：%d \n", w.statusCode)

	}
}

func (chain InterceptorChain) Handle(handler InterceptorHandlerFunc) http.Handler {

	for _, interceptor := range chain {
		handler = handler.Intercept(interceptor)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &StatusResponseWriter{w, 0}

		handler.ServeHTTP(rw, r)
	})

}
