package validators

import (
	"errors"
	"fmt"
	"mime/multipart"
)

const MaxFilesPerRepo = 5

func UploadValidator(files []*multipart.FileHeader) error {
    if len(files) > MaxFilesPerRepo {
        errMsg := fmt.Sprintf("The maximum number of pages is %d", MaxFilesPerRepo) 
        return errors.New(errMsg)
    }

    filenameSet := make(map[string]struct{})

    for _, fileHandler := range files {
        if _, exists := filenameSet[fileHandler.Filename]; exists {
            return errors.New("Multiple files have the same name")
        }

        filenameSet[fileHandler.Filename] = struct{}{}
    }

    return nil
}

