package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		input  http.Header
		outStr string
		outErr error
	}{
		// headers set in next loop
		{input: nil, outStr: "pickles", outErr: nil},
		{input: nil, outStr: "trend", outErr: nil},
		{input: nil, outStr: "", outErr: errors.New("malformed authorization header")},
		{input: nil, outStr: "", outErr: ErrNoAuthHeaderIncluded},
		{input: nil, outStr: "mamma-2lda9fi-dwa", outErr: nil},
	}

	for i := range tests {
		if i == 2 {
			tests[i].input = http.Header{}
			tests[i].input.Add("Authorization", "malf0rm3d-key")
		} else if i == 3 {
			tests[i].input = http.Header{}
		} else {
			tests[i].input = http.Header{}
			tests[i].input.Add("Authorization", fmt.Sprintf("ApiKey %v", tests[i].outStr))
		}
	}

	for _, testCase := range tests {
		str, err := GetAPIKey(testCase.input)
		if err != nil && testCase.outErr != nil { // both err not nil
			if str != testCase.outStr || err.Error() != testCase.outErr.Error() {
				t.Fatalf("expected:\n%v | %v\ngot: %v | %v",
					testCase.outStr, testCase.outErr,
					str, err)
			}
		} else if err == nil && testCase.outErr == nil { // both err nil
			if str != testCase.outStr {
				t.Fatalf("expected:\n%v | %v\ngot: %v | %v",
					testCase.outStr, testCase.outErr,
					str, err)
			}
		} else { // err different
			t.Fatalf("expected:\n%v | %v\ngot: %v | %v",
				testCase.outStr, testCase.outErr,
				str, err)
		}
	}
}
