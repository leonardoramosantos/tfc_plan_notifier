package api

import (
	"encoding/json"
	"time"
)

type StatusTimestamps struct {
	PlanPlannedAt time.Time `json:"planned-at"`
	PlanErroredAt time.Time `json:"errored-at"`
}

type RunAttributes struct {
	Status     string           `json:"status"`
	Timestamps StatusTimestamps `json:"status-timestamps"`
}

type Run struct {
	Id      string        `json:"id"`
	RunAttr RunAttributes `json:"attributes"`
}

type runData struct {
	Data []Run `json:"data"`
	Meta Meta  `json:"meta"`
}

// Specialization of TFC api to get plan runs
func (x *TFCApi) GetRuns(workspace_id string) []Run {
	var result []Run

	var curr_page = 0
	var total_pages = 0

	// Loop to load all pages of data
	for should_continue := true; should_continue; should_continue = (curr_page > total_pages) {
		curr_page += 1
		var response_body = x.CallAPIListObjects("workspaces/"+workspace_id+"/runs", curr_page)
		var request_result runData
		if err := json.Unmarshal(response_body, &request_result); err != nil {
			log.Fatalf("Error: ", err)
		}
		result = append(result, request_result.Data...)
		total_pages = request_result.Meta.Pagination.TotalPages
	}

	return result
}
