package upload

import (
	"blog/conf"
	"blog/helpers"
	"blog/util/file"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

//ImageUpload 图片上传配置参数
type ImageUpload struct {
	ImagePrefixURL  string
	ImageSavePath   string
	RuntimeRootPath string
	ImageAllowExts  []string
	ImageMaxSize    int
}

//NewImageUplaod New图片上传实例
func NewImageUplaod(a *conf.App) *ImageUpload {
	return &ImageUpload{
		ImagePrefixURL:  a.ImagePrefixURL,
		ImageSavePath:   a.ImageSavePath,
		ImageAllowExts:  a.ImageAllowExts,
		ImageMaxSize:    a.ImageMaxSize,
		RuntimeRootPath: a.RuntimeRootPath,
	}
}

//GenImageSavePath 获取图片保存路径
func (i *ImageUpload) GenImageSavePath(filename string) (fullSavePath string, fullURLPath string, URLPath string, err error) {
	saveDir := i.genImageSaveDir()
	fullSaveDir := i.RuntimeRootPath + saveDir
	err = i.CheckImage(fullSaveDir)
	if err != nil {
		return "", "", "", err
	}
	filename = i.genImageName(filename)
	fullSavePath = fullSaveDir + filename
	fullURLPath = i.ImagePrefixURL + "/" + saveDir + filename
	URLPath = "/" + saveDir + filename
	return fullSavePath, fullURLPath, URLPath, nil
}

//GenImageSaveDir 生成图片保存文件夹
func (i *ImageUpload) genImageSaveDir() string {
	saveDir := i.ImageSavePath + time.Now().Format("20060102") + "/"
	return saveDir
}

//GenImageName 生成图片名称
func (i *ImageUpload) genImageName(filename string) string {
	ext := path.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	name = helpers.MD5(name + time.Now().String())
	return name + ext
}

//CheckImageExt 检查图片格式
func (i *ImageUpload) CheckImageExt(f multipart.File) bool {
	fSrc, _ := ioutil.ReadAll(f)
	ext := file.GetFileType(fSrc[:10])

	for _, allowExt := range i.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

//CheckImageSize 检查图片大小
func (i *ImageUpload) CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		return false
	}

	return size <= i.ImageMaxSize
}

//CheckImage 检查图片
func (i *ImageUpload) CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
