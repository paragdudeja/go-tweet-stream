package errors

// Function for handling errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
