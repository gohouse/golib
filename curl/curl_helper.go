package curl

func Request(method, pageurl string, pts ...ParamHandleFunc) *Result {
	return NewCurl(pts...).Request(method,pageurl)
}
func Get(pageurl string, pts ...ParamHandleFunc) *Result {
	return Request("GET",pageurl,pts...)
}
func Post(pageurl string, pts ...ParamHandleFunc) *Result {
	return Request("POST",pageurl,pts...)
}
func Put(pageurl string, pts ...ParamHandleFunc) *Result {
	return Request("PUT",pageurl,pts...)
}
func Delete(pageurl string, pts ...ParamHandleFunc) *Result {
	return Request("DELETE",pageurl,pts...)
}
func Patch(pageurl string, pts ...ParamHandleFunc) *Result {
	return Request("PATCH",pageurl,pts...)
}
func Options(pageurl string, pts ...ParamHandleFunc) *Result {
	return Request("OPTIONS",pageurl,pts...)
}
