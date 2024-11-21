package error_handling

import (
	"errors"
	"fmt"
)

func Guess(guess int) (string, error) {

	myNumber := 50

	if guess < 1 || guess > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}

	if guess > myNumber {
		return "Daha küçük bir sayı giriniz", nil
	}

	if guess < myNumber {
		return "Daha büyük bir sayı giriniz", nil
	}

	return "Bildin", nil
}

func Demo2() {
	message, err := Guess(51)
	if err == nil {
		fmt.Println(message)
	} else {
		fmt.Println(message, err)
	}
	// fmt.Println(Guess(500))
	// fmt.Println(Guess(-50))
}
