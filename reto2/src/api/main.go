package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	filesGRPCHandler "github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/delivery/grpc"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/delivery/http"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/delivery/rabbit"
	proto "github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/proto/files"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/application/entities"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
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
		serverFiles := grpc.NewServer()

		service := filesGRPCHandler.NewFilesGRPCHandler(nil)

		proto.RegisterFilesServer(serverFiles, service)

		err = serverFiles.Serve(lis)
		if err != nil {
			panic(fmt.Errorf("impossible to serve: %s", err))
		}
		break
	case entities.MOMServerApp:
		container := dependencies.StartDependencies()
		r := rabbit.NewRabbitHandler(container)
		chDelivery, _ := container.RabbitMQ.Consume(
			"readAll",
			"",
			true,
			false,
			false,
			false, nil)

		for ans := range chDelivery {
			fmt.Println(time.Now().String(), ans.MessageId)
			ans, _ := r.ReadAll()
			err := container.RabbitMQ.Publish("", "readAllAns", false, false,
				amqp.Publishing{
					Headers:     nil,
					ContentType: "text/plain",
					Body:        []byte(strings.Join(ans, ", ")),
				})
			if err != nil {
				fmt.Println("Error" + err.Error())
			}
		}
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
