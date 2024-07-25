package api

import (
	"encoding/json"
	"leonardoramosantos/tfc_plan_notifier/config"
	"log"
)

type Organization struct {
	Id         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type orgData struct {
	Data []Organization `json:"data"`
	Meta Meta           `json:"meta"`
}

func GetOrganizations(tfcAPIConfig *config.TFCApi) []Organization {
	var result []Organization

	var curr_page = 0
	var total_pages = 0

	for should_continue := true; should_continue; should_continue = (curr_page > total_pages) {
		curr_page += 1
		var response_body = tfcAPIConfig.CallAPIListObjects("organizations", curr_page)
		var request_result orgData
		if err := json.Unmarshal(response_body, &request_result); err != nil {
			log.Fatalf("Error: ", err)
		}
		result = append(result, request_result.Data...)
		total_pages = request_result.Meta.Pagination.TotalPages
	}

	return result
}
