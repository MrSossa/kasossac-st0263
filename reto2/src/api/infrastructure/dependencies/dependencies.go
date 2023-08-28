package dependencies

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/proto/files"

	"google.golang.org/grpc"
)

type Container struct {
	GRPCClient files.FilesClient
}

func StartDependencies() *Container {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	serviceClient := files.NewFilesClient(conn)

	container := Container{
		GRPCClient: serviceClient,
	}
	return &container
}
