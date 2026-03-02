package main

import "testing"

func Example() {
	main()
	// Output:
	// Hello World!
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var testcases = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello World!",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde!",
		},
		"hindi": {
			lang: "hi",
			want: "नमस्कार, संसार!",
		},
		"Greek": {
			lang: "el",
			want: `Unsupported Language: "el"`,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			want := tc.want
			got := greet(tc.lang)
			if want != got {
				t.Errorf("Test Failure : wanted: %q | Got: %q", want, got)
			}
		})
	}
}
