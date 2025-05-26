package hw2

import "fmt"

func Err(a int) error {
	if a%5 == 0 {
		return fmt.Errorf("the number %d is divisible by 5", a)
	}
	return nil
}
