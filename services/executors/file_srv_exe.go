package executors

import (
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
)

type fileServiceExecutor struct {
	repositories.FileRepo
}

func NewFileServiceExecutor(repo *repositories.FileRepo) services.FileService {
	return &fileServiceExecutor{FileRepo: *repo}
}
