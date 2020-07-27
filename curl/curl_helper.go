package curl

func Request(method, pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return NewCurl().
		Request(method, pageurl, pts...)
}
func Get(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return Request("GET", pageurl, pts...)
}
func Post(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return Request("POST", pageurl, pts...)
}
func Put(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return Request("PUT", pageurl, pts...)
}
func Delete(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return Request("DELETE", pageurl, pts...)
}
func Patch(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return Request("PATCH", pageurl, pts...)
}
func Options(pageurl string, pts ...ParamHandleFunc) (result *Response, err error) {
	return Request("OPTIONS", pageurl, pts...)
}
