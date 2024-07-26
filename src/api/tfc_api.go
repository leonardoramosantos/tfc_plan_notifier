package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/op/go-logging"
)

type TFCApi struct {
	token               string
	default_path_prefix string
}

var log = logging.MustGetLogger("tfc_plan_notifier")

func (x *TFCApi) CallAPIListObjects(endpointPath string, page ...int) []byte {
	var full_url = x.default_path_prefix + endpointPath + "?"
	if len(page) > 0 {
		full_url = full_url + fmt.Sprintf("page[number]=%d", page[0])
	}

	var client = resty.New()

	log.Debugf("Calling %s", full_url, x.token)
	var resp, err = client.R().
		SetHeader("Authorization", "Bearer "+x.token).
		Get(full_url)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	return resp.Body()
}

func GetTFCApi(token string) *TFCApi {
	var newObj = TFCApi{}
	newObj.default_path_prefix = "https://app.terraform.io/api/v2/"
	newObj.token = token
	return &newObj
}
