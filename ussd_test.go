package ussd

import (
	"testing"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/suite"
	"github.com/samora/ussd-go/sessionstores"
)

type UssdSuite struct {
	suite.Suite
	ussd    *Ussd
	request *Request
	store   *sessionstores.Redis
}

func (u *UssdSuite) SetupSuite() {
	u.request = &Request{}
	u.request.Mobile = "233246662003"
	u.request.Operator = "vodafone"
	u.request.ServiceCode = "*123#"
	u.request.Type = strInitiation
	u.request.Message = u.request.ServiceCode

	u.store = sessionstores.NewRedis("localhost:6379")

	u.ussd = New(u.store, "demo", "Menu")
	u.ussd.Middleware(addData("global", "i'm here"))
	u.ussd.Ctrl(new(demo))
}

// func (u *UssdSuite) TearDownSuite() {
// 	u.ussd.end()
// }

func (u *UssdSuite) TestUssd() {

	u.Equal(1, len(u.ussd.middlewares))
	u.Equal(2, len(u.ussd.ctrls))

	response := u.ussd.process(u.request)
	u.Equal(strResponse, response.Type)
	u.Contains(response.Message, "Welcome")

	u.request.Message = "1"
	u.request.Type = strResponse
	response = u.ussd.process(u.request)
	u.Equal(strResponse, response.Type)
	u.Contains(response.Message, "Enter Name")

	u.request.Message = "Samora"
	response = u.ussd.process(u.request)
	u.Equal(strResponse, response.Type)
	u.Contains(response.Message, "Select Sex")

	u.request.Message = "1"
	response = u.ussd.process(u.request)
	u.Equal(strRelease, response.Type)
	u.Contains(response.Message, "Master Samora")

	u.request.Message = u.request.ServiceCode
	u.request.Type = strInitiation
	u.ussd.process(u.request)
	u.request.Message = "0"
	u.request.Type = strResponse
	response = u.ussd.process(u.request)
	u.Equal(strRelease, response.Type)
	u.Equal("Bye bye.", response.Message)
}

func TestUssdSuite(t *testing.T) {
	suite.Run(t, new(UssdSuite))
}
