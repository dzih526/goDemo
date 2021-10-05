package standardProject

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"standardProject/conf"
	"standardProject/dao"
	"standardProject/router"
	"syscall"
)
var signals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

func main() {
	wire.Build(router.Routers,conf.InitConfig,dao.InitDB)
	rctx, cancel := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(rctx)
	sigs := make(chan os.Signal,1)
	signal.Notify(sigs, signals...)
	g.Go(func() error {
		for {
			sig := <-sigs
			fmt.Println(sig)
		}
		return nil
	})

	g.Go(func() error {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "done",
			})
		})
		return  r.Run()
	})

	if err:=g.Wait();err!=nil{
		fmt.Println("Something Wrong")
		cancel()
	}
}