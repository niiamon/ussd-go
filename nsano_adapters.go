package ussd

type NsanoRequest struct {
	MSISDN  string `json:"msisdn"`
	Network string `json:"network"`
	Message string `json:"msg"`
}

func (n *NsanoRequest) GetRequest() *Request {
	return &Request{
		Mobile:  n.MSISDN,
		Message: n.Message,
		Network: n.Network,
	}
}

type NsanoResponse struct {
	USSDResp ussdResp
}

type ussdResp struct {
	Action string `json:"action"`
	Menus  string `json:"menus"`
	Title  string `json:"title"`
}

func (n *NsanoResponse) SetResponse(response *Response) {
	if response.Release {
		n.USSDResp.Menus = response.Message
		n.USSDResp.Action = "prompt"
	} else {
		n.USSDResp.Action = "input"
	}
}
