package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {
	fmt.Println(message, rails)

	messageArray := []rune(message)

	matrix := make([][]rune, rails)
	for rail := range matrix {
		matrix[rail] = make([]rune, len(messageArray))
	}

	for i, rail, direction := 0, 0, 1; i < len(message); i, rail = i+1, rail+direction {
		matrix[rail][i] = messageArray[i]

		if (rail + direction) == rails {
			direction = -1
		} else if (rail + direction) == -1 {
			direction = 1
		}
	}

	fmt.Println(matrix)

	retStr := ""

	for i := 0; i < rails; i++ {
		retStr += string(matrix[i])
	}

	return strings.Replace(retStr, string("\x00"), "", -1)

}

func Decode(message string, rails int) string {

	messageArray := []rune(message)
	retArray := make([]rune, len(message))

	cipher_char := 1

	c_len := len(message)

	for i := 1; i <= rails; i++ {
		flag := true

		plain_char := i
		for plain_char <= c_len {
			retArray[plain_char-1] = messageArray[cipher_char-1]
			cipher_char++

			if i == 1 || i == rails {
				plain_char = plain_char + (rails-1)*2
			} else {
				if flag {
					plain_char = plain_char + (rails-i)*2
					flag = false
				} else {
					plain_char = plain_char + ((i - 1) * 2)
					flag = true
				}
			}
		}
	}
	return string(retArray)
}
