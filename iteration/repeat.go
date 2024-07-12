package iteration

func Repeat(character string, repititionCount int) string {
	var repeated string

	for i := 0; i < repititionCount; i++ {
		repeated += character
	}

	return repeated
}
