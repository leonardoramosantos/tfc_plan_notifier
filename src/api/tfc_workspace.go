package api

import (
	"encoding/json"
)

type Workspace struct {
	Id         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type wksData struct {
	Data []Workspace `json:"data"`
	Meta Meta        `json:"meta"`
}

func (x *TFCApi) GetWorkspaces(organization_name string) []Workspace {
	var result []Workspace

	var curr_page = 0
	var total_pages = 0

	for should_continue := true; should_continue; should_continue = (curr_page > total_pages) {
		curr_page += 1
		var response_body = x.CallAPIListObjects("organizations/"+organization_name+"/workspaces", curr_page)
		var request_result wksData
		if err := json.Unmarshal(response_body, &request_result); err != nil {
			log.Fatalf("Error: ", err)
		}
		result = append(result, request_result.Data...)
		total_pages = request_result.Meta.Pagination.TotalPages
	}

	return result
}
