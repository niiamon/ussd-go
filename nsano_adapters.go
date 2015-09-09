package ussd

import (
	"net/url"
)

type NsanoRequest struct {
	MSISDN  string `json:"msisdn"`
	Network string `json:"network"`
	Message string `json:"msg"`
}

func (n *NsanoRequest) GetRequest() *Request {
	message, err := url.QueryUnescape(n.Message)
	if err != nil {
		message = n.Message
	}
	network, err := url.QueryUnescape(n.Network)
	if err != nil {
		network = n.Network
	}
	return &Request{
		Mobile:  n.MSISDN,
		Message: message,
		Network: network,
	}
}

type NsanoResponse struct {
	USSDResp *ussdResp
}

type ussdResp struct {
	Action string `json:"action"`
	Menus  string `json:"menus"`
	Title  string `json:"title"`
}

func (n *NsanoResponse) SetResponse(response *Response) {
	n.USSDResp = new(ussdResp)
	n.USSDResp.Title = response.Message
	if response.Release {
		n.USSDResp.Action = "prompt"
	} else {
		n.USSDResp.Action = "input"
	}
}
