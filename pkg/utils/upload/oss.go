package upload

import (
	"context"
	"fmt"
	conf "github.com/CocaineCong/gin-mall/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

// UploadToCos 封装上传图片到腾讯云COS，然后返回状态和图片的url，单张
func UploadToCos(file multipart.File, fileName string) (path string, err error) {
	cosConfig := conf.Config.Oss
	var SecretId = cosConfig.SecretId
	var SecretKey = cosConfig.SecretKey
	var BucketName = cosConfig.BucketName
	var Region = cosConfig.Region

	var serverUrl = fmt.Sprintf("https://%s.cos.%s.myqcloud.com", BucketName, Region)
	u, _ := url.Parse(serverUrl)

	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretId,
			SecretKey: SecretKey,
		},
	})

	ext := filepath.Ext(fileName)             // 获取文件后缀（包括 ".xxx"）
	name := fileName[:len(fileName)-len(ext)] // 提取文件名（不含后缀）
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("%s_%d%s", name, timestamp, ext) // 拼接文件名、时间戳和后缀

	_, err = client.Object.Put(
		context.Background(),
		key,
		file,
		&cos.ObjectPutOptions{}, // 如果需要自定义选项，可以在这里传递
	)

	if err != nil {
		return "", err
	}

	return serverUrl + "/" + key, nil
}
