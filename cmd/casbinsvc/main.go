package main

import (
	"casbinsvc/pkg/endpoints"
	"casbinsvc/pkg/services"
	"casbinsvc/pkg/transports"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.SetDefault("address.http", ":8081")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			panic(err)
		}
	}
	httpAddr := viper.GetString("address.http")
	var (
		service     = services.NewServer()
		endpoints   = endpoints.New(service)
		httpHandler = transports.NewHTTPHandler(endpoints)
	)
	var g group.Group
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			os.Exit(1)
		}
		g.Add(func() error {
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	g.Run()
}
