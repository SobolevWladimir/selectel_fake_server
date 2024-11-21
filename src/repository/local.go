package repository

import (
	"fmt"
);

type RepositoryLocal struct {
	RootDir string
}

func (rp *RepositoryLocal) SaveFile(file []byte, filePath string) error{
	fmt.Println("filePath", filePath)
	return nil;
}
