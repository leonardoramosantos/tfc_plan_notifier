package api

import (
	"encoding/json"
)

type Organization struct {
	Id         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type orgData struct {
	Data []Organization `json:"data"`
	Meta Meta           `json:"meta"`
}

// Specialization of TFC api to get organizations
func (x *TFCApi) GetOrganizations() []Organization {
	var result []Organization

	var curr_page = 0
	var total_pages = 0

	// Loop to load all pages of data
	for should_continue := true; should_continue; should_continue = (curr_page > total_pages) {
		curr_page += 1
		var response_body = x.CallAPIListObjects("organizations", curr_page)
		var request_result orgData
		if err := json.Unmarshal(response_body, &request_result); err != nil {
			log.Fatalf("Error: ", err)
		}
		result = append(result, request_result.Data...)
		total_pages = request_result.Meta.Pagination.TotalPages
	}

	return result
}
