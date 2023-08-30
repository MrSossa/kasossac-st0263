package repository

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/repository/grpc"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/repository/pc"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/repository/rabbitmq"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
)

type FilesRepository interface {
	ReadAllMicro() ([]string, error)
	ReadAllGRPC() ([]string, error)
	ReadAllRabbit() ([]string, error)
}

type filesRepository struct {
	pc.PCRepository
	grpc.GRPCRepository
	rabbitmq.RabbitRQRepository
}

func NewFilesRepository(container *dependencies.Container) FilesRepository {
	return &filesRepository{
		pc.NewPCRepository(container),
		grpc.NewGRPCRepository(container),
		rabbitmq.NewRabbitRQRepository(container),
	}
}
