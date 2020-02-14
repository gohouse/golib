package curl

import (
	"net/http"
	"net/url"
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
	val interface{}

	h   H
	cli http.Client
}
type ParamHandleFunc func(*Param)

func ParamString(arg string) ParamHandleFunc {
	return func(p *Param) {
		p.pt = PT_String
		p.val = arg
	}
}
func ParamJson(arg []byte) ParamHandleFunc {
	return func(p *Param) {
		p.pt = PT_Json
		p.val = arg
	}
}
func ParamForm(arg url.Values) ParamHandleFunc {
	return func(p *Param) {
		p.pt = PT_Form
		p.val = arg
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
