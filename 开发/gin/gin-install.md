## 安装
1. 安装go
2. 安装gin
        
        ```shell
        ]# go get -u github.com/gin-gonic/gin
        ]# cd /home/go/src/testpro
        ]# vim hi.go
        package main

        import (
                "net/http"
                "github.com/gin-gonic/gin"
        )

        func main() {
                r := gin.Default()
                r.GET("/", func(c *gin.Context) {
                        c.String(http.StatusOK, "hi gin!")
                })
                r.Run(":9000")
        }
        
        ]# go run hi.go #若有提示找不到"cannot find module providing package github.com/gin-gonic/gin: working directory is not part of a module"，执行以下两个步骤（需在hi.go文件目录下执行）
        ]# go mod init
        ]# go mod edit -require github.com/gin-gonic/gin@latest
        ]# go run hi.go
        [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

        [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
         - using env:   export GIN_MODE=release
         - using code:  gin.SetMode(gin.ReleaseMode)

        [GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
        [GIN-debug] Listening and serving HTTP on :9000
        [GIN] 2021/02/23 - 12:01:20 | 200 |      68.983μs |    192.168.xx.x | GET      "/"
        [GIN] 2021/02/23 - 12:01:20 | 404 |       2.389μs |    192.168.xx.x | GET      "/favicon.ico"
        [GIN] 2021/02/23 - 12:01:36 | 200 |       4.986μs |    192.168.xx.x | GET      "/"
        ```
