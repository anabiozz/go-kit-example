package main

import (
	"context"
	"fmt"
	"go-kit-test/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	var svc server.Service
	svc = server.LoremService{}
	endpoint := server.Endpoints{
		LoremEndpoint: server.MakeLoremEndpoint(svc),
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := server.MakeHTTPHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Starting server at port 8085")
		handler := r
		errChan <- http.ListenAndServe(":8085", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)

}
