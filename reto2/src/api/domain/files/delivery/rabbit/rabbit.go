package rabbit

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/usecase"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
)

type RabbitHandler struct {
	usecase usecase.FileUseCases
}

func NewRabbitHandler(container *dependencies.Container) *RabbitHandler {
	return &RabbitHandler{
		usecase: usecase.NewFileUseCases(container),
	}
}

func (handler RabbitHandler) ReadAll() ([]string, error) {
	return handler.usecase.ReadAllMicro()
}
