package ussd

type ValidatorDoesNotExistError struct {
	Key string
}

func (s ValidatorDoesNotExistError) Error() string {
	return s.Key + " validator does not exist"
}
