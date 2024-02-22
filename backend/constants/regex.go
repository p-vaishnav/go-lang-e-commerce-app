package constants

var Regex = struct {
	Email  string
	Mobile string
}{
	Email:  `[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}`,
	Mobile: "^91[0-9]{10}$", // at this point only indian numbers are valid
}
