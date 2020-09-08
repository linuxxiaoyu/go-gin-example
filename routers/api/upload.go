package api

import (
	"net/http"
	"path/filepath"

	"github.com/linuxxiaoyu/go-gin-example/pkg/app"
	"github.com/linuxxiaoyu/go-gin-example/pkg/upload"

	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/go-gin-example/pkg/e"
	"github.com/linuxxiaoyu/go-gin-example/pkg/logging"
)

// @Summary Upload an image
// @Produce  json
// @Param image formData file true "image"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /upload [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		appG.Response(http.StatusOK, code, data)
		return
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := filepath.Join(fullPath, imageName)
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = filepath.Join(savePath, imageName)
			}
		}
	}

	appG.Response(http.StatusOK, code, data)
}
