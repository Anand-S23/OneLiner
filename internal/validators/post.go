package validators

import (
	"errors"
	"fmt"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/storage"
)

func PostValidator(postData models.PostDto, store *storage.SnippetStore) map[string]string {
    errs := make(map[string]string, 3)

    err := validateName(postData.Name)
    if err != nil {
        errs["name"] = err.Error()
    }

    err = validateDescription(postData.Description)
    if err != nil {
        errs["description"] = err.Error()
    }

    err = validateFiles(postData.Files)
    if err != nil {
        errs["files"] = err.Error()
    }

    return errs
}

func validateName(name string) error {
    // TODO: Validate the name is unique in the repo

    if len(name) < 1 || len(name) > 50 {
        return errors.New("Name of repo should be at least 1 character long, and not exceed 50")
    }
    return nil
}

func validateDescription(desc string) error {
    if len(desc) > 100 {
        return errors.New("Description must be 100 charaters or less")
    }
    return nil
}

func validateFiles(files map[string]string) error {
    if len(files) == 0 {
        return errors.New("Must have at least one file in the repo")
    }

    if len(files) > MaxFilesPerRepo {
        errMsg := fmt.Sprintf("The maximum number of pages is %d", MaxFilesPerRepo) 
        return errors.New(errMsg)
    }

    return nil
}
