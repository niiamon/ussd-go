package sessionstores

import "github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/suite"

// StoreSuite should be used to test each session store implementation
type StoreSuite struct {
	suite.Suite
	store            Store
	name, key, value string
}

// NewStoreSuite is used to create a suite of tests to validate each
// session store.
func NewStoreSuite(store Store) *StoreSuite {
	return &StoreSuite{store: store}
}

func (s *StoreSuite) SetupSuite() {
	s.name = "hash"
	s.key = "test"
	s.value = "works"
	err := s.store.Connect()
	s.Nil(err)
}

func (s *StoreSuite) TearDownSuite() {
	err := s.store.Close()
	s.Nil(err)
}

func (s *StoreSuite) TestStore() {
	// SetValue
	err := s.store.SetValue(s.key, s.value)
	s.Nil(err)

	// GetValue
	val, err := s.store.GetValue(s.key)
	s.Nil(err)
	s.Equal(s.value, val)

	// ValueExists
	exists, err := s.store.ValueExists(s.key)
	s.Nil(err)
	s.True(exists)

	// DeleteValue
	err = s.store.DeleteValue(s.key)
	s.Nil(err)
	exists, err = s.store.ValueExists(s.key)
	s.Nil(err)
	s.False(exists)

	// HashSetValue
	err = s.store.HashSetValue(s.name, s.key, s.value)
	s.Nil(err)

	// HashGetValue
	val, err = s.store.HashGetValue(s.name, s.key)
	s.Nil(err)
	s.Equal(s.value, val)

	// HashValueExists
	exists, err = s.store.HashValueExists(s.name, s.key)
	s.Nil(err)
	s.True(exists)

	// HashDeleteValue
	err = s.store.HashDeleteValue(s.name, s.key)
	s.Nil(err)
	exists, err = s.store.HashValueExists(s.name, s.key)
	s.Nil(err)
	s.False(exists)

	// HashExists
	err = s.store.HashSetValue(s.name, s.key, s.value)
	s.Nil(err)
	exists, err = s.store.HashExists(s.name)
	s.Nil(err)
	s.True(exists)

	// HashDelete
	err = s.store.HashDelete(s.name)
	s.Nil(err)
	exists, err = s.store.HashExists(s.name)
	s.Nil(err)
	s.False(exists)
}
