package railfence

func Encode(message string, rails int) string {
	panic("Please implement the Encode function")
}

func Decode(message string, rails int) string {

	messageArray := []rune(message)
	length := len(messageArray)

	decodedArray := make([]rune, length)
	i := 0

	for rail := 0; rail < rails; rail++ {
		for index := 0 + rail; index < length; index += 2 * (rails - rail - 1) {
			decodedArray[index] = messageArray[i]
			i++
		}
	}

	return string(messageArray)

}
