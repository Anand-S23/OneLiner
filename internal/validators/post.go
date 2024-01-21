package validators

import (
	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/storage"
)

func PostValidator(postData models.PostDto, store *storage.SnippetStore) map[string]string {
    errs := make(map[string]string, 2)

    return errs
}
