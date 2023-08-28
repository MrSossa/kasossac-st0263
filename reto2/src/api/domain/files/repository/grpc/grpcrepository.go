package grpc

import (
	"context"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/proto/files"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
	"strings"
)

type GRPCRepository interface {
	ReadAllGRPC() ([]string, error)
}

type grpcRepository struct {
	GRPCClient files.FilesClient
}

func NewGRPCRepository(container *dependencies.Container) GRPCRepository {
	return &grpcRepository{
		GRPCClient: container.GRPCClient,
	}
}

func (repository *grpcRepository) ReadAllGRPC() ([]string, error) {
	res, err := repository.GRPCClient.ReadAll(context.Background(), &files.ReadAllRequest{})
	if err != nil {
		return nil, err
	}
	return strings.Split(res.Files, ", "), nil
}
