package main

import "testing"

var (
	forTheLoveOfGo                                     = Book{Author: "John Arundel", Title: "The Handmaid's Tale"}
	thePowerOfGoTools                                  = Book{Author: "John Arundel", Title: "The power of Go tools"}
	makingThingsHappen                                 = Book{Author: "JScott Berkun", Title: "Making things happen"}
	tjHarvardBusinessReviewProjectManagementHandbook   = Book{Author: "JAntonio Nieto-Rodriguez", Title: "TJHarvard Business Review Project Management Handbook"}
	TProjectManagementAbsoluteBeginnersGuide5thEdition = Book{Author: "Gregory M. Horine", Title: "Project Management Absolute Beginner's Guide (5th Edition)"}
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
				{Name: "Cami", Books: []Book{forTheLoveOfGo, thePowerOfGoTools, makingThingsHappen, tjHarvardBusinessReviewProjectManagementHandbook, TProjectManagementAbsoluteBeginnersGuide5thEdition}},
				{Name: "Oreo", Books: []Book{}},
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

func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}

func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
		return false
	}
	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}
	return true
}

func equalBooksCount(got, want map[Book]uint) bool {
	if len(got) != len(want) {
		return false
	}
	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}
	return true
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Alex", Books: []Book{forTheLoveOfGo, thePowerOfGoTools}},
				{Name: "Cami", Books: []Book{forTheLoveOfGo, thePowerOfGoTools, makingThingsHappen, tjHarvardBusinessReviewProjectManagementHandbook, TProjectManagementAbsoluteBeginnersGuide5thEdition}},
			},
			want: map[Book]uint{forTheLoveOfGo: 2, thePowerOfGoTools: 2, makingThingsHappen: 1, tjHarvardBusinessReviewProjectManagementHandbook: 1, TProjectManagementAbsoluteBeginnersGuide5thEdition: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		// "bookworm without books":            {},
		// "bookworm with twice the same book": {},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}
