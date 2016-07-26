package controllers

import (
	json "github.com/bitly/go-simplejson"
	"github.com/deepglint/mock-json/app"
	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

func (c Api) GetMock(id string) revel.Result {
	println(id)
	s := app.GetJson()
	if id != "" {

		j, _ := json.NewJson([]byte(s))
		res, _ := j.Get(id).MarshalJSON()
		return c.RenderText(string(res))
	}
	//println(string(s))
	return c.RenderText(s)

}
