package workable

var _ candidateService = &candidateServiceImpl{}

type candidateService interface {
	Post(input CandidateInput) (result CandidateOutput, err error)
}

type candidateServiceImpl struct {
	client    *Client
	subdomain string
}

func (s *candidateServiceImpl) Post(input CandidateInput) (result CandidateOutput, err error) {
	req, err := s.client.newRequest(s.subdomain, "POST", "candidates", nil, input)
	if err != nil {
		return
	}
	err = s.client.do(req, &result)
	return
}
