package pc

import "github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"

type PCRepository interface {
	ReadAllMicro() ([]string, error)
}

type pcRepository struct {
}

func NewPCRepository(container *dependencies.Container) PCRepository {
	return &pcRepository{}
}

func (repository pcRepository) ReadAllMicro() ([]string, error) {
	return []string{"hola", "pepe"}, nil
}
