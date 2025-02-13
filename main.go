package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func main() {

	port := flag.String("port", "", "端口号")
	path := flag.String("path", "", "路径")
	username := flag.String("u", "", "用户名")
	password := flag.String("p", "", "密码")

	flag.Parse()

	if *port == "" {
		fmt.Println("缺少参数: -port，用于指定服务运行端口")
		os.Exit(1)
	} else if *path == "" {
		fmt.Println("缺少参数: -path，用于指定分享路径")
		os.Exit(1)
	} else if *username != "" && *password == "" {
		fmt.Println("缺少参数: -p，用于指定分享的密码")
		os.Exit(1)
	} else if *username == "" && *password != "" {
		fmt.Println("缺少参数: -u，用于指定分享的用户名")
		os.Exit(1)
	}

	if len(*username) != 0 && len(*password) != 0 {

		fs := &webdav.Handler{
			FileSystem: webdav.Dir(*path),
			LockSystem: webdav.NewMemLS(),
		}

		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			u, p, ok := req.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if u != *username || p != *password {
				http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
				return
			}
			fs.ServeHTTP(w, req)
		})

		http.ListenAndServe(fmt.Sprint(":", *port), nil)
	} else {
		http.ListenAndServe(fmt.Sprint(":", *port), &webdav.Handler{
			FileSystem: webdav.Dir(*path),
			LockSystem: webdav.NewMemLS(),
		})
	}
}
