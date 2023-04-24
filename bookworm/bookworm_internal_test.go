package main

import "testing"

var (
	forTheLoveOfGo                                   = Book{Author: "John Arundel", Title: "The Handmaid's Tale"}
	thePowerOfGoTools                                = Book{Author: "John Arundel", Title: "The power of Go tools"}
	makingThingsHappen                               = Book{Author: "JScott Berkun", Title: "Making things happen"}
	tjHarvardBusinessReviewProjectManagementHandbook = Book{Author: "JAntonio Nieto-Rodriguez", Title: "TJHarvard Business Review Project Management Handbook"}
)

// type testCase struct {
// 	bookwormsFile string
// 	want          []Bookworm
// 	wantErr       bool
// }

func TestLoadBookworms_success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
	want          []Bookworm
	wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Alex", Books: []Book{forTheLoveOfGo,thePowerOfGoTools }},
				{Name: "Cami", Books: []Book{makingThingsHappen,tjHarvardBusinessReviewProjectManagementHandbook }},
			},
		}
	}
}
