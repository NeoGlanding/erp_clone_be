package controllers

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/libs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func PostFile(c *gin.Context) {

	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	file, _ := c.FormFile("file")

	id := uuid.New()

	fileSavedLocation := fmt.Sprintf("temp/%s", file.Filename) 
	
	
	c.SaveUploadedFile(file, fileSavedLocation)

	
	tempFile, err := os.Open(fileSavedLocation);

	if err != nil {
		fmt.Println(err)
		helpers.SetInternalServerError(c, "Error while opening file E-0")
	}

	defer tempFile.Close()

	bucket := libs.FirebaseStorageBucket.Object(id.String())
	writer := bucket.NewWriter(context.Background())
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer writer.Close()

	if _, err := io.Copy(writer, tempFile); err != nil {
		fmt.Println(err)
		helpers.SetInternalServerError(c, "Error while processing file E-1")
		return
	}

	os.Remove(fileSavedLocation)

	c.Set("data", user)
}