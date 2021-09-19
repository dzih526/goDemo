package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

var signals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

func main() {

		g := new(errgroup.Group)

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
		}
}