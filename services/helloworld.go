package services

import (
	"github.com/danielmoledo/hello-world/model"
	"github.com/danielmoledo/hello-world/repository"
	"time"
)

func HelloWorld() (*model.HelloWorld, error) {
	auditlog := model.HelloWorld{Timestamp: time.Now().In(time.UTC)}

	return repository.Create(&auditlog)
}
