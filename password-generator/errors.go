package gopasswordgenerator

type PasswordSizeError struct{}

func (e *PasswordSizeError) Error() string { return "Password length must be at least 8 characters." }
