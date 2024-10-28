package mathutils

import "errors"

func Add(num1, num2 int) int {
	return num1 + num2
}

/*
It’s important to test edge cases to ensure your functions handle unexpected or extreme inputs gracefully.

Example: Testing Edge Cases
Suppose we modify Add2 to return an error if the sum exceeds 100
*/

func Add2(num1, num2 int) (int, error) {
	result := num1 + num2
	if result > 100 {
		return 0, errors.New("result exceeds 100")
	}
	return result, nil
}
