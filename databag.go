package ussd

import (
	"encoding/json"
	"github.com/samora/ussd-go/sessionstores"
)

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

// Set value in databag
func (d DataBag) Set(key, value string) error {
	return d.store.HashSetValue(d.name, key, value)
}

// Get value from databag
func (d DataBag) Get(key string) (string, error) {
	return d.store.HashGetValue(d.name, key)
}

// Exists verifies if value is in databag
func (d DataBag) Exists(key string) (bool, error) {
	return d.store.HashValueExists(d.name, key)
}

// Delete value from databag
func (d DataBag) Delete(key string) error {
	return d.store.HashDeleteValue(d.name, key)
}

// Clear databag
func (d DataBag) Clear() error {
	return d.store.HashDelete(d.name)
}

// SetMarshal marshals v into json and puts in databag
func (d DataBag) SetMarshal(key string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return d.Set(key, string(b))
}

// GetUnmarshal retrieves json from databag and unmarshals into v
func (d DataBag) GetUnmarshal(key string, v interface{}) error {
	str, err := d.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), v)
}
