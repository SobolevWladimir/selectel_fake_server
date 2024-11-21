package repository

import (
	"os"
	"path"
	"path/filepath"
)

type RepositoryLocal struct {
	RootDir string
}

func (rp *RepositoryLocal) createFilePath(filePath string) string {
	return filepath.Join(rp.RootDir, filePath)
}
func (rp *RepositoryLocal) SaveFile(data []byte, filePath string) error {
	pathInDisk := rp.createFilePath(filePath)
	dir := path.Dir(pathInDisk)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(pathInDisk)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RepositoryLocal) GetFile(filePath string) ([]byte, error) {
	pathInDisk := rp.createFilePath(filePath)
	return  os.ReadFile(pathInDisk);
}

