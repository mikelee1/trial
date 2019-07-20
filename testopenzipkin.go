package main

import (
	"google.golang.org/grpc/metadata"
	"context"
	"github.com/op/go-logging"
)

var logger11 *logging.Logger

func init()  {
	logger11 = logging.MustGetLogger("testopenzipkin")
}

func main() {
	md := metadata.MD{"test":[]string{"a"}}
	ctx := metadata.NewOutgoingContext(context.Background(),md)
	logger11.Info(metadata.FromOutgoingContext(ctx))
}


