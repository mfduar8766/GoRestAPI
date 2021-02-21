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
