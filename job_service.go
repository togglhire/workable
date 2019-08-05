package workable

import "net/http"

var _ jobService = &jobServiceImpl{}

type jobService interface {
	List(input ListJobsInput, next string) (result Jobs, err error)
}

type jobServiceImpl struct {
	client    *Client
	subdomain string
}

func (s *jobServiceImpl) List(input ListJobsInput, next string) (result Jobs, err error) {
	params := Params{}
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
