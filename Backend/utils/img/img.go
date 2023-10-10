package img

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

func ProcessAvatar(userID uint, file *multipart.File) (*bytes.Buffer, string, string, error) {
	// 检查文件类型
	buffer := make([]byte, 512) // 512 bytes should be enough for the DetectContentType function
	_, err := (*file).Read(buffer)
	if err != nil && err != io.EOF {
		log.Println(err)
		return nil, "", "", err
	}
	// 检查内容类型
	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		log.Printf("unsupported file format: %s\n", contentType)
		return nil, "", "", fmt.Errorf("unsupported file format: %s", contentType)
	}
	log.Println(contentType)
	// 重置文件读取位置
	_, err = (*file).Seek(0, io.SeekStart)
	if err != nil {
		log.Println(err)
		return nil, "", "", err
	}
	// 将文件解码为图像
	img, _, err := image.Decode((*file))
	if err != nil {
		log.Println(err)
		return nil, "", "", err
	}
	// 压缩图片至200x200
	newImg := resize.Resize(200, 200, img, resize.Lanczos3)
	// 转换图片为PNG格式
	var b bytes.Buffer
	err = png.Encode(&b, newImg)
	if err != nil {
		log.Println(err)
		return nil, "", "", err
	}
	currentTime := time.Now().UnixNano() // 获取当前的Unix纳秒时间戳
	AvatarID := fmt.Sprintf("%d_%d", userID, currentTime)
	// 生成新的文件名
	filename := AvatarID + ".png"
	contentType = "image/png"
	return &b, filename, contentType, nil
}
