package curl

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ParamType int

const (
	PT_String ParamType = iota + 1 // string
	PT_Json                        // json
	PT_Form                        // map[string][]string
	//PT_Header
)

type H map[string]string
type Param struct {
	pt  ParamType
	val io.Reader

	h   H
	cli http.Client
}
type ParamHandleFunc func(*Param)

func ParamString(arg string) ParamHandleFunc {
	return func(p *Param) {
		p.pt = PT_String
		p.val = strings.NewReader(arg)
	}
}
func ParamJson(arg *[]byte) ParamHandleFunc {
	return func(p *Param) {
		p.pt = PT_Json
		p.val = bytes.NewBuffer(*arg)
		// 设置header头
		var newMap = H{"Content-Type": "application/json"}
		for k, v := range p.h {
			newMap[k] = v
		}
		p.h = newMap
	}
}
func ParamForm(arg url.Values) ParamHandleFunc {
				//rd = strings.NewReader(c.param.val.(url.Values).Encode())
	return func(p *Param) {
		p.pt = PT_Form
		p.val = strings.NewReader(arg.Encode())
		// 设置header头
		var newMap = H{"Content-Type": "application/x-www-form-urlencoded"}
		for k, v := range p.h {
			newMap[k] = v
		}
		p.h = newMap
	}
}
func ParamHeader(h H) ParamHandleFunc {
	return func(p *Param) {
		//p.pt = PT_Header
		p.h = h
	}
}

// ParamClient 主要设置 http.Client 信息的
func ParamClient(cli http.Client) ParamHandleFunc {
	return func(p *Param) {
		p.cli = cli
	}
}
