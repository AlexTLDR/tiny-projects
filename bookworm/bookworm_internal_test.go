

type testCase struct {
	bookwormsFile string
	want          []bookworm
	wantErr       bool
}
tests := map[string]testCase{}

var (
	forTheLoveOfGo    = Book{Author: "John Arundel", Title: "The Handmaid's Tale"}
	thePowerOfGoTools = Book{Author: "John Arundel", Title: "The power of Go tools"}
)