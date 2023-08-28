package http

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/usecase"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
)

type FilesHTTPHandler struct {
	usecase usecase.FileUseCases
}

func NewFilesHTTPHandler(container *dependencies.Container) *FilesHTTPHandler {
	filesHTTPHandler := FilesHTTPHandler{
		usecase: usecase.NewFileUseCases(container),
	}
	return &filesHTTPHandler
}
