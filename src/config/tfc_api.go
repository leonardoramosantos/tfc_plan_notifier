package config

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/op/go-logging"
)

type TFCApi struct {
	is_authenticated    bool
	token               string
	default_path_prefix string
}

var log = logging.MustGetLogger("tfc_plan_notifier")

func (x *TFCApi) CallAPIListObjects(endpointPath string, page ...int) []byte {
	if !x.is_authenticated {
		x.authenticate()
	}

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

func (x *TFCApi) authenticate() {
	var resp *http.Response
	var err error
	var body []byte

	resp, err = http.Get("test")
	if err != nil {
		log.Fatalf("An Error Occurred! %v", err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error Reading Body %v", err)
	}
	x.is_authenticated = true
	fmt.Println(body)
}

func GetConfig() *TFCApi {
	var newObj = TFCApi{}
	newObj.is_authenticated = true
	newObj.default_path_prefix = "https://app.terraform.io/api/v2/"
	//TODO - Get token from ENV
	newObj.token = ""
	return &newObj
}
