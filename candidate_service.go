package workable

import (
	"fmt"
	"net/http"
)

var _ candidateService = &candidateServiceImpl{}

type candidateService interface {
	Create(jobShortCode string, input CandidateInput) (result CandidateOutput, err error)
	List(input ListCandidatesInput, next string) (result CandidateOutput, err error)
}

type candidateServiceImpl struct {
	client    *Client
	subdomain string
}

func (s *candidateServiceImpl) Create(jobShortCode string, input CandidateInput) (result CandidateOutput, err error) {
	req, err := s.client.newRequest(s.subdomain, "POST", fmt.Sprintf("jobs/%s/candidates", jobShortCode), nil, input)
	if err != nil {
		return
	}
	err = s.client.do(req, &result)
	return
}

func (s *candidateServiceImpl) List(input ListCandidatesInput, next string) (result CandidateOutput, err error) {
	params := Params{}
	if input.JobShortCode != "" {
		params["shortcode"] = input.JobShortCode
	}
	if input.State != "" {
		params["state"] = input.State
	}
	if input.Limit != 0 {
		params["limit"] = input.Limit
	}
	if input.SinceID != "" {
		params["since_id"] = input.SinceID
	}
	if input.MaxID != "" {
		params["max_id"] = input.MaxID
	}
	if input.CreatedAfter != 0 {
		params["created_after"] = input.CreatedAfter
	}
	if input.UpdatedAfter != 0 {
		params["updated_after"] = input.UpdatedAfter
	}

	var req *http.Request
	if next != "" { // use next url
		req, err = s.client.newRequestFromURL(next, "GET", nil)
	} else {
		req, err = s.client.newRequest(s.subdomain, "GET", "jobs", params, nil)
	}
	if err != nil {
		return
	}
	err = s.client.do(req, &result)
	return
}
