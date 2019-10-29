package api

import (
	"blog/http/response"
	"errors"

	"github.com/gin-gonic/gin"
)

//UploadImage 图片上传接口
func UploadImage(c *gin.Context) {
	r := response.NewR(c)
	data := make(map[string]string)

	f, image, err := c.Request.FormFile("image")
	if err != nil {
		r.FailedResponse(err)
		return
	}

	if image == nil {
		r.FailedResponse(errors.New("参数有误"))
		return
	}
	//生成保存路径 生成完整url地址
	fullSavePath, fullURLPath, URLPath, err := srv.ImageUpload.GenImageSavePath(image.Filename)
	if err != nil {
		r.FailedResponse(err)
		return
	}
	//检查图片格式
	if ok := srv.ImageUpload.CheckImageExt(f); !ok {
		r.FailedResponse(errors.New("图片格式不支持"))
		return
	}
	//检查图片尺寸
	if ok := srv.ImageUpload.CheckImageSize(f); !ok {
		r.FailedResponse(errors.New("图片尺寸过大"))
		return
	}
	//保存图片
	if err := c.SaveUploadedFile(image, fullSavePath); err != nil {
		err = errors.New("图片保存失败")
		r.FailedResponse(err)
		return
	}

	data["image_url"] = fullURLPath
	data["image_save_url"] = URLPath
	r.SuccessResponse(data, "图片上传成功")
}
