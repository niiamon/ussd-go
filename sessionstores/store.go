package sessionstores

const (
	ExpiryTime = 35
)

// Store interface should be implemented by all session stores
// ussd supports.
type Store interface {
	// Connect to session store.
	Connect() error

	// Key-Value
	SetValue(key, value string) error
	GetValue(string) (string, error)
	ValueExists(string) (bool, error)
	DeleteValue(string) error

	// Hash
	HashSetValue(name, key, value string) error
	HashGetValue(name, key string) (string, error)
	HashValueExists(name, key string) (bool, error)
	HashDeleteValue(name, key string) error
	HashExists(string) (bool, error)
	HashDelete(string) error

	// Close connection to session store.
	Close() error
}
