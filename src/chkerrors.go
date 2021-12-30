package src

// Check for errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
