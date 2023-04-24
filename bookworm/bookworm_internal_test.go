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
				{Name: "Alex", Books: []Book{forTheLoveOfGo, thePowerOfGoTools}},
				{Name: "Cami", Books: []Book{makingThingsHappen, tjHarvardBusinessReviewProjectManagementHandbook}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}
			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}
			if !equalBookworms(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}
