package main

import (
	"fmt"
	"net"
	"os"

	filesGRPCHandler "github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/delivery/grpc"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/delivery/http"
	proto "github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/proto/files"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/application/entities"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	compilationType := os.Args[1]
	switch entities.AppType(compilationType) {
	case entities.GRPCServerApp:
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			panic(fmt.Errorf("cannot create listener: %s", err))
		}
		container := dependencies.StartDependencies()

		serverFiles := grpc.NewServer()

		service := filesGRPCHandler.NewFilesGRPCHandler(container)

		proto.RegisterFilesServer(serverFiles, service)

		err = serverFiles.Serve(lis)
		if err != nil {
			panic(fmt.Errorf("impossible to serve: %s", err))
		}
		break
	case entities.MOMServerApp:
		break
	case entities.ApiGWApp:
		r := gin.Default()

		container := dependencies.StartDependencies()
		h := http.NewFilesHTTPHandler(container)

		r.GET("/readAll", h.ReadAll())

		r.Run()
		break
	}
}
