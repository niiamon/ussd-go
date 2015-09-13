package ussd

// RequestAdapter adapts service provider's USSD request to conform
// to Request
type RequestAdapter interface {
	GetRequest() *Request
}

// ResponseAdapter adapts Response to service provider's
// USSD response.
type ResponseAdapter interface {
	SetResponse(Response)
}
