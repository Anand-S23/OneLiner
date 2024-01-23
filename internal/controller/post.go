package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/validators"
	"github.com/gorilla/mux"
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

func (c *Controller) GetPostsForCurrentUser(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)
    posts, err := c.store.GetPostsByUser(currentUserID)
    if err != nil {
        log.Printf("Could not get posts for %s user: %s", currentUserID, err)
        return InternalServerError(w)
    }

    successMsg := map[string][]models.Post {
        "posts": posts,
    }
    log.Printf("Retrived posts for %s successfully", currentUserID)
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) CreatePost(w http.ResponseWriter, r *http.Request) error {
    var postData models.PostDto
    err := json.NewDecoder(r.Body).Decode(&postData)
    if err != nil {
        errMsg := map[string]string {
            "error": "Could not parse post data",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
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
    errMsg := map[string]string {"error": "Invalid Post ID"}

    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    post := c.store.GetPost(models.NewPostRecordSK(postID))
    if post.ID == "" {
        log.Printf("Could not get post with sk %s\n", models.NewPostRecordSK(postID))
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    log.Printf("Returning infromation about post with id %s\n", post.ID)
    return WriteJSON(w, http.StatusOK, post)
}

func (c *Controller) UpdatePost(w http.ResponseWriter, r *http.Request) error {
    currentUserID := r.Context().Value("user_id").(string)

    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        errMsg := map[string]string {"error": "Invalid Post ID"}
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    post := c.store.GetPost(models.NewPostRecordSK(postID))
    if post.ID == "" {
        errMsg := map[string]string {"error": "Invalid Post ID"}
        log.Printf("Could not get post with sk %s\n", models.NewPostRecordSK(postID))
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    if post.UserID != currentUserID {
        errMsg := map[string]string {"error": "Unauthorized"}
        log.Printf("%s is trying to delete post by %s", post.ID, currentUserID)
        return WriteJSON(w, http.StatusUnauthorized, errMsg)
    }

    var postData models.PostDto
    err := json.NewDecoder(r.Body).Decode(&postData)
    if err != nil {
        errMsg := map[string]string {
            "error": "Could not parse post data",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
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
        errMsg := map[string]string {"error": "Invalid Post ID"}
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    post := c.store.GetPost(models.NewPostRecordSK(postID))
    if post.ID == "" {
        errMsg := map[string]string {"error": "Invalid Post ID"}
        log.Printf("Could not get post with sk %s\n", models.NewPostRecordSK(postID))
        return WriteJSON(w, http.StatusNotFound, errMsg)
    }

    if post.UserID != currentUserID {
        errMsg := map[string]string {"error": "Unauthorized"}
        log.Printf("%s is trying to delete post by %s", post.ID, currentUserID)
        return WriteJSON(w, http.StatusUnauthorized, errMsg)
    }

    err := c.store.DeletePost(models.NewPostRecordPK(post.UserID), models.NewPostRecordSK(post.ID))
    if err != nil {
        return InternalServerError(w)
    }

    return nil
}

