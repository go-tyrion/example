package api

import (
	"encoding/json"
	"lib/config"
	"lib/core/log"
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

	json.Marshal(msg)

	log.ShowFile()
	log.Info("controller")

	logger := log.NewLogger()
	logger.Info("log2")

	ctx.OkJSON(map[string]interface{}{
		"status":  0,
		"message": result,
		"data": map[string]interface{}{
			"id":      1,
			"name":    "eden",
			"brokers": config.Strings("kafka.brokers", "kafka.ini", ","),
		},
	})
}

func (this *Index) Demo(ctx *http.Context) {

}
