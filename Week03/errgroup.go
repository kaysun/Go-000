package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BarHandler struct {

}

func (handerl BarHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	//server := &http.Server{
	//	Addr: "8080",
	//}

	mux := http.NewServeMux()
	mux.Handle("/bar", &BarHandler{})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	g.Go(func() error {
		fmt.Println("http server start")
		go func() {
			fmt.Println("block here line40")
			<- ctx.Done()
			fmt.Println("ctx Done line42")
			// new context
			httpCtx, cancel := context.WithTimeout(context.Background(), 5000 * time.Millisecond)
			defer cancel()
			// http server shutdown
			if err := server.Shutdown(httpCtx); err != nil {
				fmt.Println(fmt.Sprintf("http server shutdown occurred error, err=%+v", err))
			}
			fmt.Println("http server shutdown finished")
		}()
		return server.ListenAndServe()
		//return http.ListenAndServe(":8080", &BarHandler{})
	})

	signalChan := make(chan os.Signal)
	g.Go(func() error {
		// 监听信号
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		for {
			fmt.Println("block here line61")
			select {
			case <- ctx.Done():
				fmt.Println("context done")
				return ctx.Err()
			case <- signalChan:
				fmt.Println("receive signal")
				return errors.New("system exit signal")
			}
		}
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(fmt.Sprintf("err=%+v", err))
	}

	fmt.Println("the end")
}