package controller

import (
	"log"
	"net/http"

	"github.com/Anand-S23/Snippet/internal/models"
)

func (c *Controller) UploadFile(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseMultipartForm(50 << 20)
	if err != nil {
        errMsg := map[string]string {
            "error": "Unable to parse form",
        }
        log.Print(err)
        return WriteJSON(w, http.StatusBadRequest, errMsg)
	}

    file, _, err := r.FormFile("file")
	if err != nil {
        errMsg := map[string]string {
            "error": "Unable to retrieve file",
        }
        log.Print(err)
        return WriteJSON(w, http.StatusBadRequest, errMsg)
	}
	defer file.Close()

    fileID := models.NewUUID()
	err = c.store.UploadFileToS3(file, fileID)
	if err != nil {
        errMsg := map[string]string {
            "error": "Unable to upload file",
        }
        log.Printf("Unable to upload file to s3 bucket: %s\n", err)
        return WriteJSON(w, http.StatusBadRequest, errMsg)
	}

    successMsg := map[string]string {
        "message": "User created successfully",
        "fileID": fileID,
    }

    log.Println("File uploaded successfully to blob storage")
    return WriteJSON(w, http.StatusOK, successMsg)
}

