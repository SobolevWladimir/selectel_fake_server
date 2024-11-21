package repository

type RepositoryInterface interface {
	SaveFile(data []byte, filePath string) error
	GetFile(filePath string) ([]byte, error)
	DeleteFile(filepath string) error
}
