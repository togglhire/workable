package workable

var _ jobService = &jobServiceImpl{}

type jobService interface {
	Get() (Jobs, error)
}

type jobServiceImpl struct {
	client *Client
}

func (j *jobServiceImpl) Get() (result Jobs, err error) {
	req, err := j.client.newRequest("GET", "partner/jobs", nil, nil)
	if err != nil {
		return
	}
	err = j.client.do(req, &result)
	return
}
