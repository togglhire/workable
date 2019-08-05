package workable

var _ accountService = &accountServiceImpl{}

type accountService interface {
	List() (Accounts, error)
}

type accountServiceImpl struct {
	client *Client
}

func (s *accountServiceImpl) List() (result Accounts, err error) {
	req, err := s.client.newRequest("", "GET", "accounts", nil, nil)
	if err != nil {
		return
	}
	err = s.client.do(req, &result)
	return
}
