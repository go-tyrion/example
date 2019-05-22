package api

import (
	"demo/kafka"
	"encoding/json"
	"lib/config"
	"lib/server/http"
)

type Index struct {
	Name string
}

func (this *Index) Init() {

}

func (this *Index) Index(ctx *http.Context) {
	var result string
	ctx.SetHeader("Content-Type", "application/json")

	msg := map[string]interface{}{
		"id":     1,
		"name":   "Tyrion",
		"mobile": "13207182679",
	}
	data, _ := json.Marshal(msg)

	err := kafka.ProducerInstance.Product("test-topic", data)
	if err != nil {
		result = err.Error()
	} else {
		result = "ok"
	}

	ctx.OkJSON(map[string]interface{}{
		"status":  0,
		"message": result,
		"data": map[string]interface{}{
			"id":      1,
			"name":    "eden",
			"brokers": config.Strings("brokers", "kafka.ini", ","),
		},
	})
}

func (this *Index) Demo(ctx *http.Context) {

}
