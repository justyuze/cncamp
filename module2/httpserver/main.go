package main

import (
	"fmt"
	"net/http"
)

func main() {

	chain := InterceptorChain{HeaderInterceptor(), VersionInterceptor(), LogInterceptor()}

	http.Handle("/error", chain.Handle(errorz))

	http.Handle("/healthz", chain.Handle(healthz))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// 4.当访问 localhost/healthz 时，应返回 200
func healthz(w *StatusResponseWriter, r *http.Request) {
	w.Status(200)
	w.Write([]byte("ok"))

}

func errorz(w *StatusResponseWriter, r *http.Request) {
	// 状态码必须在 输出数据之前设置，否则Write会直接设置为200，且以第一次设置为准
	w.Status(500)
	w.Write([]byte("error"))

}
