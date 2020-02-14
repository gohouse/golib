package curl

import (
	"bytes"
	"errors"
	"github.com/gohouse/t"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Curl struct {
	param *Param
}
type Result struct {
	t.T
	Request *http.Request
	Response *http.Response
	Error error
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

func (c *Curl) Request(method, pageurl string, pts ...ParamHandleFunc) (result *Result) {
	var b []byte
	//var err error
	result = &Result{}
	for _, o := range pts {
		o(c.param)
	}
	var rd io.Reader = c.buildParam()
	// "https://www.sex.com/gifs/?page=2"

	client := &http.Client{
		Transport:     c.param.cli.Transport,
		CheckRedirect: c.param.cli.CheckRedirect,
		Jar:           c.param.cli.Jar,
		Timeout:       c.param.cli.Timeout,
	}
	//var req *http.Request
	//var resp *http.Response
	result.Request, result.Error = http.NewRequest(strings.ToUpper(method), pageurl, rd)
	if result.Error != nil {
		return
	}
	//req.Header.Add("referer", "https://www.sex.com")
	if c.param.h != nil {
		for k, v := range c.param.h {
			result.Request.Header.Add(k, v)
		}
	}
	//处理返回结果
	result.Response, result.Error = client.Do(result.Request)
	if result.Error != nil {
		return
	}
	defer result.Response.Body.Close()
	if result.Response.StatusCode >= 300 {
		result.Error = errors.New(result.Response.Status)
		return
	}

	b, result.Error = ioutil.ReadAll(result.Response.Body)
	if result.Error != nil {
		return
	}

	result.T = t.New(b)
	return result
}
func (c *Curl) Get(pageurl string, pts ...ParamHandleFunc) *Result {
	return c.Request("GET", pageurl, pts...)
}
func (c *Curl) Post(pageurl string, pts ...ParamHandleFunc) *Result {
	return c.Request("POST", pageurl, pts...)
}
func (c *Curl) Put(pageurl string, pts ...ParamHandleFunc) *Result {
	return c.Request("PUT", pageurl, pts...)
}
func (c *Curl) Delete(pageurl string, pts ...ParamHandleFunc) *Result {
	return c.Request("DELETE", pageurl, pts...)
}
func (c *Curl) Patch(pageurl string, pts ...ParamHandleFunc) *Result {
	return c.Request("PATCH", pageurl, pts...)
}
func (c *Curl) Options(pageurl string, pts ...ParamHandleFunc) *Result {
	return c.Request("OPTIONS", pageurl, pts...)
}

func (c *Curl) buildParam() (rd io.Reader) {
	if c.param.pt > 0 {
		switch c.param.pt {
		case PT_String:
			rd = strings.NewReader(c.param.val.(string))
		case PT_Json:
			rd = bytes.NewBuffer(c.param.val.([]byte))
		case PT_Form:
			rd = strings.NewReader(c.param.val.(url.Values).Encode())
			var newMap = H{"content-type": "application/x-www-form-urlencoded"}
			for k, v := range c.param.h {
				newMap[k] = v
			}
			c.param.h = newMap
		}
	}
	return
}
