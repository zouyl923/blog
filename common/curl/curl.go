package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type Curl struct {
	NewRequest func() *Curl

	client  *http.Client      //请求客户端
	request *http.Request     //请求对象
	TimeOut time.Duration     //请求超时
	Body    io.Reader         //请求体
	Headers map[string]string //设置header
	Cookies []*http.Cookie    //cookies
	Method  string            //请求方式
	Url     string            //请求地址

	Params map[string]interface{} //get参数
	Data   map[string]interface{} //post请求参数 content-type application/json
	Form   map[string]interface{} //表单请求参数 content-type form-data
}

func NewRequest() *Curl {
	return &Curl{
		client:  &http.Client{Timeout: 3 * time.Second},
		request: &http.Request{},
		Headers: map[string]string{},
		Params:  map[string]interface{}{},
		Data:    map[string]interface{}{},
		Form:    map[string]interface{}{},
	}
}

func (c *Curl) SetTimeout(timeout time.Duration) *Curl {
	c.TimeOut = timeout
	c.client.Timeout = timeout
	return c
}

func (c *Curl) SetBody(body io.Reader) *Curl {
	c.Body = body
	return c
}

func (c *Curl) SetHeader(k string, v string) *Curl {
	c.Headers[k] = v
	return c
}

func (c *Curl) SetCookie(cookie *http.Cookie) *Curl {
	c.Cookies = append(c.Cookies, cookie)
	return c
}

func (c *Curl) SetMethod(method string) *Curl {
	c.Method = method
	return c
}

func (c *Curl) SetUrl(url string) *Curl {
	c.Url = url
	return c
}
func (c *Curl) SetParam(k string, v interface{}) *Curl {
	fmt.Println(k, v)
	c.Params[k] = v
	return c
}
func (c *Curl) SetData(k string, v interface{}) *Curl {
	c.Params[k] = v
	return c
}

func (c *Curl) SetForm(k string, v interface{}) *Curl {
	c.Params[k] = v
	return c
}

func (c *Curl) Send() ([]byte, error) {
	//获取request
	c.request, _ = http.NewRequest(c.Method, c.Url, c.Body)
	c.request.Header.Set("Content-Type", "application/json")
	//设置自定义header
	for k, v := range c.Headers {
		c.request.Header.Set(k, v)
	}
	//设置自定义cookie
	for _, v := range c.Cookies {
		c.request.AddCookie(v)
	}
	//发送请求
	resp, err := c.client.Do(c.request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Curl) Get() ([]byte, error) {
	data, _ := json.Marshal(c.Params)
	body := bytes.NewReader(data)
	c.Body = body
	c.Method = "GET"
	return c.Send()
}

func (c *Curl) Post() ([]byte, error) {
	data, _ := json.Marshal(c.Data)
	body := bytes.NewReader(data)
	c.Body = body
	c.Method = "POST"
	return c.Send()
}

func (c *Curl) Put() ([]byte, error) {
	data, _ := json.Marshal(c.Data)
	body := bytes.NewReader(data)
	c.Body = body
	c.Method = "PUT"
	return c.Send()
}

func (c *Curl) Delete() ([]byte, error) {
	data, _ := json.Marshal(c.Data)
	body := bytes.NewReader(data)
	c.Body = body
	c.Method = "DELETE"
	return c.Send()
}

// filePath 本地文档地址
func (c *Curl) FormPost(filePath string, fileField string) ([]byte, error) {
	//文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//创建一个multipart类型的写文件
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 额外参数
	for key, val := range c.Data {
		_ = writer.WriteField(key, val.(string))
	}
	//创建一个新的form-data头
	part, err := writer.CreateFormFile(fileField, file.Name())
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		fmt.Println(" post err=", err)
	}
	c.Body = body
	c.Method = "POST"
	c.SetHeader("Content-Type", writer.FormDataContentType())
	return c.Send()
}
