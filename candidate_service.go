package workable

import "fmt"

var _ candidateService = &candidateServiceImpl{}

type candidateService interface {
	Create(jobShortCode string, input CandidateInput) (result CandidateOutput, err error)
	// List(jobShortCode string)
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
