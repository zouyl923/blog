package cos

import (
	"blog/common/helper"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	SecretId  string
	SecretKey string
	Cos       struct {
		AppId  int
		Bucket string
		REGION string
		Domain string
	}
}

type Cos struct {
	client *cos.Client
	config Config
	domain string
	ctx    context.Context
}

func NewClient(config interface{}) (*Cos, error) {
	c := Config{}
	helper.ExchangeStruct(config, &c)
	var bucketUrl *url.URL
	if len(c.Cos.Domain) > 0 {
		//设置定义域名
		bucketUrl, _ = url.Parse(c.Cos.Domain)
	} else {
		//使用源域名
		bucketUrl, _ = url.Parse("https://cos" + c.Cos.Bucket + c.Cos.REGION + "myqcloud.com")
	}
	serviceUrl, _ := url.Parse("https://cos" + c.Cos.REGION + "myqcloud.com")
	client := cos.NewClient(&cos.BaseURL{
		BucketURL:  bucketUrl,
		ServiceURL: serviceUrl,
	}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.SecretId,
			SecretKey: c.SecretKey,
		},
	})
	return &Cos{
		client: client,
		config: c,
		domain: bucketUrl.String(),
		ctx:    context.Background(),
	}, nil
}

func (c *Cos) UploadLocal(fileName string, filePtah string, dir string) (string, error) {
	t := time.Now().Format("20060102150405")
	ext := filepath.Ext(fileName)
	name := t + "_" + helper.Md5(time.Now().String()+fileName)
	key := dir + "/" + name + ext
	_, _, err := c.client.Object.Upload(
		c.ctx, key, filePtah, nil,
	)
	if err != nil {
		return "", err
	}
	url := c.domain + "/" + key
	return url, nil
}

func (c *Cos) GetPresignedURL(uri string) (string, error) {
	u, _ := url.Parse(uri)
	key := strings.Trim(u.Path, "/")
	signUrl, err := c.client.Object.GetPresignedURL(c.ctx, http.MethodGet, key, c.config.SecretId, c.config.SecretKey, 24*time.Hour, nil)
	if err != nil {
		return "", err
	}
	return signUrl.String(), nil
}
