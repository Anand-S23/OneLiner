package controller

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/validators"
	"github.com/gorilla/mux"
)

func (c *Controller) UploadFile(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseMultipartForm(50 << 20)
	if err != nil {
        log.Print(err)
        return BadRequestError(w, "Unable to parse form")
	}

    file, _, err := r.FormFile("file")
	if err != nil {
        log.Print(err)
        return BadRequestError(w, "Unable to retrieve file")
	}
	defer file.Close()

    fileID := models.NewUUID()
	err = c.store.UploadFileToS3(file, fileID)
	if err != nil {
        log.Printf("Unable to upload file to s3 bucket: %s\n", err)
        errMsg := ErrorMessage("Unable to upload file")
        return WriteJSON(w, http.StatusInternalServerError, errMsg)
	}

    successMsg := map[string]string {
        "message": "User created successfully",
        "fileID": fileID,
    }

    log.Println("File uploaded successfully to blob storage")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) UploadFiles(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseMultipartForm(50 << 20)
	if err != nil {
        log.Print(err)
        return BadRequestError(w, "Unable to parse form")
	}

    files := r.MultipartForm.File["files"]
    err = validators.UploadValidator(files)
    if err != nil {
        return BadRequestError(w, err.Error())
    }

    var wg sync.WaitGroup
    var mu sync.Mutex
    uploadedFiles := make(map[string]string)

    for _, fileHeader := range files {
        wg.Add(1)

        go func(fileHeader *multipart.FileHeader) {
            defer wg.Done()

            file, err := fileHeader.Open()
            if err != nil {
                log.Println("Error opening file:", err)
                return
            }
            defer file.Close()

            fileID := models.NewUUID()
            err = c.store.UploadFileToS3(file, fileID)
            if err != nil {
                log.Printf("Unable to upload file to s3 bucket: %s\n", err)
                return
            }

            mu.Lock()
            uploadedFiles[fileHeader.Filename] = fileID
            mu.Unlock()
        }(fileHeader)
    }

    wg.Wait()

    if len(uploadedFiles) != len(files) {
        // TODO: Delete rest of files from s3?
        log.Printf("%d of %d files uploaded to S3", len(uploadedFiles), len(files))
        return InternalServerError(w)
    }

    log.Println("Files uploaded successfully to blob storage")
    return WriteJSON(w, http.StatusOK, uploadedFiles)
}

func (c *Controller) GetPostsForCurrentUser(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)
    posts, err := c.store.GetPostsByUser(models.NewPostRecordPK(currentUserID))
    if err != nil {
        log.Printf("Could not get posts for %s user: %s", currentUserID, err)
        return InternalServerError(w)
    }

    log.Printf("Retrived %d posts for %s user successfully", len(posts), currentUserID)
    return WriteJSON(w, http.StatusOK, posts)
}

func (c *Controller) CreatePost(w http.ResponseWriter, r *http.Request) error {
    var postData models.PostDto
    err := json.NewDecoder(r.Body).Decode(&postData)
    if err != nil {
        return BadRequestError(w, "Could not parse post data")
    }

    postErrs := validators.PostValidator(postData, c.store)
    if len(postErrs) != 0 {
        log.Println("Failed to create new post, invalid data")
        return WriteJSON(w, http.StatusBadRequest, postErrs)
    }

    currentUserID := r.Context().Value("user_id").(string)
    post := models.NewPost(postData, currentUserID)
    err = c.store.PutPost(post)
    if err != nil {
        return InternalServerError(w)
    }

    successMsg := map[string]string {
        "message": "Post created successfully",
    }
    log.Println("Post created successfully")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) ReadPost(w http.ResponseWriter, r *http.Request) error {
    errMsg := ErrorMessage("Invalid Post ID")

    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    post := c.store.GetPost(models.NewPostRecordSK(postID))
    if post.ID == "" {
        log.Printf("Could not get post with sk %s\n", models.NewPostRecordSK(postID))
        return PageNotFoundError(w)
    }

    log.Printf("Returning infromation about post with id %s\n", post.ID)
    return WriteJSON(w, http.StatusOK, post)
}

func (c *Controller) UpdatePost(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        return PageNotFoundError(w)
    }

    post := c.store.GetPost(models.NewPostRecordSK(postID))
    if post.ID == "" {
        log.Printf("Could not get post with sk %s\n", models.NewPostRecordSK(postID))
        return PageNotFoundError(w)
    }

    if post.UserID != currentUserID {
        log.Printf("%s is trying to delete post by %s", post.ID, currentUserID)
        return UnauthorizedError(w)
    }

    var postData models.PostDto
    err := json.NewDecoder(r.Body).Decode(&postData)
    if err != nil {
        return BadRequestError(w, "Could not parse post data")
    }

    postErrs := validators.PostValidator(postData, c.store)
    if len(postErrs) != 0 {
        log.Println("Failed to create new post, invalid data")
        return WriteJSON(w, http.StatusBadRequest, postErrs)
    }

    post.Name = postData.Name
    post.Description = postData.Description
    post.Files = postData.Files

    err = c.store.PutPost(post)
    if err != nil {
        return InternalServerError(w)
    }

    successMsg := map[string]string {
        "message": "Post updated successfully",
    }
    log.Println("Post updated successfully")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) DeletePost(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        return PageNotFoundError(w)
    }

    post := c.store.GetPost(models.NewPostRecordSK(postID))
    if post.ID == "" {
        log.Printf("Could not get post with sk %s\n", models.NewPostRecordSK(postID))
        return PageNotFoundError(w)
    }

    if post.UserID != currentUserID {
        log.Printf("%s is trying to delete post by %s", post.ID, currentUserID)
        return UnauthorizedError(w)
    }

    err := c.store.DeletePost(models.NewPostRecordPK(post.UserID), models.NewPostRecordSK(post.ID))
    if err != nil {
        return InternalServerError(w)
    }

    return nil
}

