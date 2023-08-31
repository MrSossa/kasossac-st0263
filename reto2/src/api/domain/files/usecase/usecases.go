package usecase

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/repository"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
)

type FileUseCases interface {
	ReadAllMicro() ([]string, error)
	ReadAll() ([]string, error)
}

type fileUseCases struct {
	repository repository.FilesRepository
}

func NewFileUseCases(container *dependencies.Container) FileUseCases {
	return fileUseCases{
		repository: repository.NewFilesRepository(container),
	}
}

func (usecase fileUseCases) ReadAllMicro() ([]string, error) {
	ans, err := usecase.repository.ReadAllMicro()
	return ans, err
}

func (usecase fileUseCases) ReadAll() ([]string, error) {
	ans, err := usecase.repository.ReadAllGRPC()
	if err == nil {
		return ans, err
	}
	return usecase.repository.ReadAllRabbit()
}
