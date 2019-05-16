package api

import "lib/server/http"

type Index struct {
	Name string
}

func (this *Index) Init() {

}

func (this *Index) Index(ctx *http.Context) {
	ctx.OkJSON(map[string]interface{}{
		"status":  0,
		"message": "success",
		"place":   "index",
	})
}

func (this *Index) Demo(ctx *http.Context) {

}
