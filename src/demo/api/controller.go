package api

import "lib/server/http"

type Index struct {
	Name string
}

func (this *Index) Init() {

}

func (this *Index) Index(ctx *http.Context) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.OkJSON(map[string]interface{}{
		"status":  0,
		"message": "success",
		"data": map[string]interface{}{
			"id":   1,
			"name": "eden",
		},
	})
}

func (this *Index) Demo(ctx *http.Context) {

}
