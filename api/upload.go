package api

import (
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/zaggash/led-matrix-ui/utils"
)

type imageUploadForm struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func uploadFile(ctx *gin.Context) {
	var form imageUploadForm
	pixelsFolder := utils.GetPixelsFolder()
	err := ctx.ShouldBind(&form)

	if err != nil {
		ctx.Data(http.StatusUnprocessableEntity, "text/html", createUploadFeedback("No file is sent", "danger"))
		log.Println(err)
		return
	}
	if err := validateFile(ctx, form.File); err != nil {
		ctx.Data(http.StatusUnprocessableEntity, "text/html", createUploadFeedback("Unable to validate file", "danger"))
		log.Println(err)
		return
	}

	newFileName := uuid.New().String()
	newFileName += "_" + form.File.Filename

	if err := ctx.SaveUploadedFile(form.File, pixelsFolder+"/upload/"+newFileName); err != nil {
		ctx.Data(http.StatusUnsupportedMediaType, "text/html", createUploadFeedback("Unable to save the file", "danger"))
		log.Println(err)
		return
	}

	ctx.Data(http.StatusOK, "text/html", createUploadFeedback("Your file has been successfully uploaded", "success"))
}

// For the color check Bootstrap documentation
// https://getbootstrap.com/docs/5.3/helpers/color-background/#overview
func createUploadFeedback(message string, statusColor string) []byte {
	answer := `<div id="form-answer" class="text-center text-bg-` + statusColor + ` rounded-2 mt-2">` + message + `</div>`
	return []byte(answer)
}

func validateFile(ctx *gin.Context, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return errors.Errorf("unable to open file: %w", err)
	}
	defer src.Close()

	err = validateContentType(src, "image/")
	if err != nil {
		return err
	}

	return nil
}

func validateContentType(src multipart.File, expectedMime string) error {
	buffer := make([]byte, 512)
	_, err := src.Read(buffer)
	if err != nil {
		return errors.Errorf("unable to read file header: %w", err)
	}

	// Function to check if a string has a certain prefix
	checkType := func(s string) error {
		if !strings.HasPrefix(s, expectedMime) {
			return errors.New("invalid file type")
		}
		return nil
	}

	// Check the content type
	contentType := http.DetectContentType(buffer)
	err = checkType(contentType)
	if err != nil {
		return err
	}

	// Check the mime type
	mime := mimetype.Detect(buffer)
	err = checkType(mime.String())
	if err != nil {
		return err
	}

	return nil
}
