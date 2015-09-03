package ussd

import (
	"testing"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/suite"
	"github.com/samora/ussd-go/sessionstores"
)

type DataBagSuite struct {
	suite.Suite
	store      sessionstores.Store
	databag    *DataBag
	request    *Request
	key, value string
}

func (d *DataBagSuite) SetupSuite() {
	d.store = sessionstores.NewRedis("localhost:6379")
	err := d.store.Connect()
	d.Nil(err)
	d.request = &Request{}
	d.request.Mobile = "233246662003"
	d.request.Operator = "vodafone"
	d.request.ServiceCode = "*123#"
	d.request.SessionId = StrRandom(10)
	d.request.Type = strInitiation
	d.request.Message = d.request.ServiceCode
	d.key = "name"
	d.value = "Samora"
	d.databag = newDataBag(d.store, d.request)
}

func (d *DataBagSuite) TearDownSuite() {
	err := d.store.Close()
	d.Nil(err)
}

func (d *DataBagSuite) TestDataBag() {
	name := d.request.Mobile + "DataBag"

	err := d.databag.Set(d.key, d.value)
	d.Nil(err)
	val, err := d.store.HashGetValue(name, d.key)
	d.Nil(err)
	d.Equal(d.value, val)

	val, err = d.databag.Get(d.key)
	d.Nil(err)
	d.Equal(d.value, val)

	exists, err := d.databag.Exists(d.key)
	d.Nil(err)
	d.True(exists)

	err = d.databag.Delete(d.key)
	d.Nil(err)
	exists, err = d.databag.Exists(d.key)
	d.Nil(err)
	d.False(exists)

	err = d.databag.Clear()
	d.Nil(err)
	exists, err = d.store.HashExists(name)
	d.Nil(err)
	d.False(exists)
}

func TestDataBagSuite(t *testing.T) {
	suite.Run(t, new(DataBagSuite))
}
