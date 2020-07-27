package curl

import (
	"errors"
	"github.com/gohouse/t"
	"io/ioutil"
	"net/http"
	"strings"
)

type Curl struct {
	param *Param
}
type Response struct {
	*http.Response
	t.T
}

// NewCurl 初始化
// pts 主要是用来设置 header
func NewCurl(pts ...ParamHandleFunc) *Curl {
	var p = &Param{}
	for _, o := range pts {
		o(p)
	}
	return &Curl{p}
}

func (c *Curl) Request(method, pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	var b []byte
	result = &Response{}
	for _, o := range pts {
		o(c.param)
	}

	client := &http.Client{
		Transport:     c.param.cli.Transport,
		CheckRedirect: c.param.cli.CheckRedirect,
		Jar:           c.param.cli.Jar,
		Timeout:       c.param.cli.Timeout,
	}
	var req *http.Request
	//var resp *http.Response
	req, err = http.NewRequest(strings.ToUpper(method), pageurl, c.param.val)
	if err != nil {
		return
	}

	if c.param.h != nil {
		for k, v := range c.param.h {
			req.Header.Add(k, v)
		}
	}
	//处理返回结果
	result.Response, err = client.Do(req)
	if err != nil {
		return
	}
	defer result.Response.Body.Close()
	if result.Response.StatusCode >= 300 {
		err = errors.New(result.Response.Status)
		return
	}

	b, err = ioutil.ReadAll(result.Response.Body)
	if err != nil {
		return
	}

	result.T = t.New(b)
	return result, nil
}
func (c *Curl) Get(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return c.Request("GET", pageurl, pts...)
}
func (c *Curl) Post(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return c.Request("POST", pageurl, pts...)
}
func (c *Curl) Put(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return c.Request("PUT", pageurl, pts...)
}
func (c *Curl) Delete(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return c.Request("DELETE", pageurl, pts...)
}
func (c *Curl) Patch(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return c.Request("PATCH", pageurl, pts...)
}
func (c *Curl) Options(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return c.Request("OPTIONS", pageurl, pts...)
}
