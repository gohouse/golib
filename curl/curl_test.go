package curl

import "testing"

func TestNewCurl(t *testing.T) {
	c := NewCurl(ParamHeader(H{"content-type": "application/json"}))
	res := c.Get("http://10.10.35.201:8081/api/v1/appdownloadurl")
	if res.Error != nil {
		t.Error(res.Error.Error())
	}
	t.Log(res.Response.Status)
	t.Log(res.String())

	var res2 = map[string]interface{}{}
	res.BindJson(&res2)
	t.Log(res2)
}
