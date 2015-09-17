package ussd

type validatorDoesNotExistError struct {
	Key string
}

func (s validatorDoesNotExistError) Error() string {
	return s.Key + " validator does not exist"
}
