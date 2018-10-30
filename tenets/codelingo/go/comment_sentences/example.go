package main

func main() {
	doThisOrThat(false)
}

// doThis is good.
func doThis() {
	return
}

//doThis2 is good.
func doThis2() {
	return
}

// doThis is bad.
func doThat() {
	return
}

//doThis is bad.
func doWhat() {
	return
}

// Request represents a request to run a command. This is good.
type Request struct{}

// Writes the JSON encoding of req to w. This is Badd!
func Encode() {}

// doWhat is bad
func doWhat() {
	return
}
