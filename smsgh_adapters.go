package ussd

type SmsghRequest struct {
	Mobile      string
	SessionId   string
	ServiceCode string
	Type        string
	Message     string
	Operator    string
}

func (s *SmsghRequest) GetRequest() *Request {
	return &Request{
		Mobile:  s.Mobile,
		Message: s.Message,
		Network: s.Operator,
	}
}

type SmsghResponse struct {
	Type, Message string
}

func (s *SmsghResponse) SetResponse(response Response) {
	s.Message = response.Message
	if response.Release {
		s.Type = "Release"
	} else {
		s.Type = "Response"
	}
}
