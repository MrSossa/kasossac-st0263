package pc

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
	"io/ioutil"
	"log"
)

type PCRepository interface {
	ReadAllMicro() ([]string, error)
}

type pcRepository struct {
}

func NewPCRepository(container *dependencies.Container) PCRepository {
	return &pcRepository{}
}

func (repository pcRepository) ReadAllMicro() ([]string, error) {
	files, err := ioutil.ReadDir("reto2/files")
	if err != nil {
		return nil, err
	}
	fileList := make([]string, 0)
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	mock := []string{"file1.txt,file2.txt"}
	log.Println(mock)
	return mock, nil
}
