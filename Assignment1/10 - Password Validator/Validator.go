package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	pass := []string{"root123root", // The hardest one!
		"RootRoot123", "rootRoot_123", "Root_1_root", "123", "rootrootroot",
		"rootroorrrrr1", "RootRoot**123", "RootRoot,123"}

	for i := range pass {
		passToTest := pass[i]
		_, err := validatePassword(passToTest)
		if err != nil {
			fmt.Println(passToTest, " :-( ", err)
		} else {
			fmt.Println(passToTest, ":-)")
		}
	}

}

func validatePassword(password string) (bool, error) {
	// Rule 1.
	if len(password) < 9 {
		return false, errors.New("Too short, should contains at least 9 characters")
	}
	byteStr := []byte(password)
	// Rule 2. at least 1 number
	re := regexp.MustCompile(`.*\d.*`)
	matches := re.Find(byteStr)
	if len(matches) < 1 {
		return false, errors.New("Password should contain at least 1 digit")
	}
	// Rule 3. at least 1 upper case
	re = regexp.MustCompile(`.*[A-Z]+.*`)
	matches = re.Find(byteStr)
	if len(matches) < 1 {
		return false, errors.New("Password should contain at least one uppercase letter")
	}
	// Rule 4. at least 1 special character
	re = regexp.MustCompile(`.*[_?*&!#<>@.,]+.*`)
	matches = re.Find(byteStr)
	if len(matches) < 1 {
		return false, errors.New("Password should contain at least one punctuation char")
	}

	return true, nil
}
