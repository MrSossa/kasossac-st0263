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
	return usecase.repository.ReadAllMicro()
}

func (usecase fileUseCases) ReadAll() ([]string, error) {
	return usecase.repository.ReadAllGRPC()
}
