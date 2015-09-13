package ussd

import "github.com/samora/ussd-go/sessionstores"

// DataBag is a key-value store.
// Used to store data that will be available across USSD session's
// request.
type DataBag struct {
	name  string
	store sessionstores.Store
}

func newDataBag(store sessionstores.Store, request *Request) *DataBag {
	name := request.Mobile + "DataBag"
	return &DataBag{
		name:  name,
		store: store,
	}
}

func (d DataBag) Set(key, value string) error {
	return d.store.HashSetValue(d.name, key, value)
}

func (d DataBag) Get(key string) (string, error) {
	return d.store.HashGetValue(d.name, key)
}

func (d DataBag) Exists(key string) (bool, error) {
	return d.store.HashValueExists(d.name, key)
}

func (d DataBag) Delete(key string) error {
	return d.store.HashDeleteValue(d.name, key)
}

func (d DataBag) Clear() error {
	return d.store.HashDelete(d.name)
}
