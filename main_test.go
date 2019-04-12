package main

import (
	"testing"
)

type TestCase struct {
	Case   string
	Expect bool
}

func verify(t *testing.T, testCases []TestCase) {
	for _, testCase := range testCases {
		result := VerifyEmailAddress(testCase.Case)
		if result != testCase.Expect {
			t.Log(testCase)
			t.Error("failed test")
		}
	}
}

func TestVerifyEmailAddress(t *testing.T) {
	verify(t, []TestCase{
		{`@example.com`, false},
		{`abc@`, false},
		{`@`, false},
		{`abcexample.com`, false},
		{`abc@example.com`, true},
		{`a.b.c@example.com`, true},
		{`.abc@example.com`, false},
		{`abc.@example.com`, false},
		{`.abc.@example.com`, false},
		{`a..bc@example.com`, false},
		{`@@example.com`, false},
		{`"@example.com`, false},
		{`"@"@example.com`, true},
		{`(def)@example.com`, false},
		{`d(e)f@example.com`, false},
		{`{}@example.com`, true},
		{`a"b"@example.com`, false},
		{`"a"b@example.cpm`, false},
		{`😇@example.com`, false},
		{`"😇"@example.com`, false},
		{`"><script>alert('or/**/2=1#')</script>"@example.com`, true},
		{`""@example.com`, true},
		{`"a\"b"@example.com`, true},
		{`a@example`, true},
		{`a@example..com`, false},
		{`a@"example.com"`, false},
		{`a@.example.com`, false},
		{`a@example.`, false},
		{`"""@example.com`, false},
		{`"\\""@example.com`, false},
		{`"\@"@example.com`, false},
	})
}
