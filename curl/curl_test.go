package curl

import "testing"

func TestNewCurl(t *testing.T) {
	c := NewCurl(ParamHeader(H{"content-type": "application/json"}))
	res,err := c.Get("http://10.10.35.201:8081/api/v1/appdownloadurl")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(res.Response.Status)
	t.Log(res.String())

	var res2 = map[string]interface{}{}
	res.Bind(&res2)
	t.Log(res2)
}
