package api

import (
	"blog/http/response"
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

//Verify 验证
func captchaVerify(captchaID string, captchaValue string) bool {
	return captcha.VerifyString(captchaID, captchaValue)
}

//GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	r := response.NewR(c)
	data := make(map[string]string)

	length := captcha.DefaultLen
	captchaID := captcha.NewLen(length)

	ImageURL := srv.C.App.ImagePrefixURL + "/captcha/" + captchaID + ".png"
	data["captcha_id"] = captchaID
	data["image_url"] = ImageURL
	r.SuccessResponse(data, "验证码")
}

//GetCaptchaImg 获取验证码图片
func GetCaptchaImg(c *gin.Context) {
	// captchaID := c.Param("captcha_id")
	// fmt.Println("GetCaptchaPng : " + captchaID)
	ServeHTTP(c.Writer, c.Request)
}

//ServeHTTP 生成图片
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	// fmt.Println("file : " + file)
	// fmt.Println("ext : " + ext)
	// fmt.Println("id : " + id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	// fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

//Serve 生成
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
