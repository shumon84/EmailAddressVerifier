package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

const (
	LowerAlphabet = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Alphabet      = LowerAlphabet + UpperAlphabet
	Numeric       = "0123456789"
	AlphaNumeric  = Alphabet + Numeric
)

const debug = false

const (
	DomainPartAvailableChars   = AlphaNumeric + "!#$%&'*+-/=?^_`{|}~"
	DotAtomAvailableChars      = AlphaNumeric + "!#$%&'*+-/=?^_`{|}~"
	QuotedStringAvailableChars = AlphaNumeric + "!#$%&'*+-/=?^_`{|}~(),.:;<>@[]"
)

func IsContains(s []byte, b byte) bool {
	for _, v := range s {
		if v == b {
			return true
		}
	}
	return false
}

type StateMachine struct {
	io.Reader
}

func NewStateMachine(reader io.Reader) *StateMachine {
	return &StateMachine{
		Reader: reader,
	}
}

func (s *StateMachine) Next() (byte, bool) {
	state := make([]byte, 1)
	_, err := s.Read(state)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if err == io.EOF {
		return 0, false
	}
	return state[0], true
}

func (s *StateMachine) IsAcceptance() bool {
	return s.Q1()
}

func StackTrace() {
	for i := 1; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		fmt.Println(fn.Name())
	}
}

func (s *StateMachine) Reject() bool {
	if debug {
		StackTrace()
	}
	return false
}

func (s *StateMachine) Acceptance() bool {
	if debug {
		StackTrace()
	}
	return true
}

func (s *StateMachine) Q1() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(DotAtomAvailableChars), current) {
		return s.Q7()
	} else if current == '"' {
		return s.Q2()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q2() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(QuotedStringAvailableChars), current) {
		return s.Q4()
	} else if current == '\\' {
		return s.Q3()
	} else if current == '"' {
		return s.Q5()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q3() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if current == '\\' || current == '"' {
		return s.Q4()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q4() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(QuotedStringAvailableChars), current) {
		return s.Q4()
	} else if current == '\\' {
		return s.Q3()
	} else if current == '"' {
		return s.Q5()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q5() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if current == '@' {
		return s.Q9()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q6() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(DotAtomAvailableChars), current) {
		return s.Q7()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q7() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(DotAtomAvailableChars), current) {
		return s.Q7()
	} else if current == '.' {
		return s.Q8()
	} else if current == '@' {
		return s.Q9()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q8() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(DotAtomAvailableChars), current) {
		return s.Q7()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q9() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(DomainPartAvailableChars), current) {
		return s.Q10()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q10() bool {
	current, ok := s.Next()
	if !ok {
		return s.Acceptance()
	}

	if IsContains([]byte(DomainPartAvailableChars), current) {
		return s.Q10()
	} else if current == '.' {
		return s.Q11()
	} else {
		return s.Reject()
	}
}

func (s *StateMachine) Q11() bool {
	current, ok := s.Next()
	if !ok {
		return s.Reject()
	}

	if IsContains([]byte(DomainPartAvailableChars), current) {
		return s.Q10()
	} else {
		return s.Reject()
	}
}

func VerifyEmailAddress(emailAddress string) bool {
	stateMachine := NewStateMachine(strings.NewReader(emailAddress))
	ok := stateMachine.IsAcceptance()
	return ok
}

func main() {
	cin := bufio.NewScanner(os.Stdin)
	for cin.Scan() {
		fmt.Println(VerifyEmailAddress(cin.Text()))
	}
}
