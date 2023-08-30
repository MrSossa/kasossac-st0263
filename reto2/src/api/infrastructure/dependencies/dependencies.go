package dependencies

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/proto/files"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type Container struct {
	GRPCClient files.FilesClient
	RabbitMQ   *amqp.Channel
}

func StartDependencies() *Container {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	serviceClient := files.NewFilesClient(conn)

	//connRabbit, err := amqp.Dial("amqp://guest:guest@18.207.23.165:5672/")
	//if err != nil {
	//	panic(err)
	//}
	//ch, err := connRabbit.Channel()
	//
	//_, err = ch.QueueDeclare("readAll", false, false, false, false, nil)
	//_, err = ch.QueueDeclare("readAllAns", false, false, false, false, nil)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	container := Container{
		GRPCClient: serviceClient,
		//RabbitMQ:   ch,
	}
	return &container
}
