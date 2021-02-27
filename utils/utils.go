package utils

import "fmt"

// MustNotError - Used to log errors
func MustNotError(err error) error {
	if err != nil {
		fmt.Printf("Error occured due to: %+v\n", err.Error())
		return err
	}
	return nil
}

// CreateMessage - Used to create a byte error
func CreateMessage(mesage string) []byte {
	return []byte(mesage)
}
