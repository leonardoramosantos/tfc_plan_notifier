package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/op/go-logging"
)

// Structure to encapsulate all Terraform Cloud API Endpoints
type TFCApi struct {
	token               string
	default_path_prefix string
}

var log = logging.MustGetLogger("tfc_plan_notifier")

func (x *TFCApi) CallAPIListObjectsOnlyLastOne(endpointPath string) []byte {
	var full_url = x.default_path_prefix + endpointPath + "?page[size]=2"
	log.Debugf("URL: %s", full_url)

	var client = resty.New()
	var resp, err = client.R().
		SetHeader("Authorization", "Bearer "+x.token).
		Get(full_url)

	if err != nil {
		log.Errorf("TFC API Call Error. Err: %v", err)
	}

	return resp.Body()
}

// Generalization of all TFC requests
func (x *TFCApi) CallAPIListObjects(endpointPath string, page ...int) []byte {
	var full_url = x.default_path_prefix + endpointPath + "?"
	if len(page) > 0 {
		full_url = full_url + fmt.Sprintf("page[number]=%d", page[0])
	}

	var client = resty.New()
	var resp, err = client.R().
		SetHeader("Authorization", "Bearer "+x.token).
		Get(full_url)

	if err != nil {
		log.Errorf("TFC API Call Error. Err: %v", err)
	}

	return resp.Body()
}

func GetTFCApi(token string) *TFCApi {
	var newObj = TFCApi{}
	newObj.default_path_prefix = "https://app.terraform.io/api/v2/"
	newObj.token = token
	return &newObj
}
