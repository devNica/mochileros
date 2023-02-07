package executors

import (
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
)

type fileServiceExecutor struct {
	repositories.FileRepository
}

func NewFileServiceExecutor(repo *repositories.FileRepository) services.FileService {
	return &fileServiceExecutor{FileRepository: *repo}
}
