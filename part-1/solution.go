package main

import (
	"fmt"
	"os"
)

// state is a state of the deterministic finite automaton.
type state uint8

const (
	stateS state = iota
	stateA
	stateP
	stateO
	stateJ
	stateT
	stateD1
	stateD2
	stateK
	stateD3
	stateD4
	stateD5
	stateD6
	stateF
	stateERR
)

// symbolClass is an ASCII input class of the automaton.
type symbolClass uint8

const (
	symbolP symbolClass = iota
	symbolO
	symbolD
	symbolH
	symbolHash
	symbolDash
	symbolN0
	symbolOther
)

// deltaS returns the transition from S.
func deltaS(a symbolClass) state {
	if a == symbolP {
		return stateP
	}

	if a == symbolH || a == symbolHash || a == symbolDash || a == symbolN0 {
		return stateA
	}

	return stateERR
}

// deltaA returns the transition from A.
func deltaA(a symbolClass) state {
	if a == symbolP {
		return stateP
	}

	return stateERR
}

// deltaP returns the transition from P.
func deltaP(a symbolClass) state {
	if a == symbolO {
		return stateO
	}

	return stateERR
}

// deltaO returns the transition from O.
func deltaO(a symbolClass) state {
	if a == symbolHash || a == symbolDash {
		return stateJ
	}

	if a == symbolH {
		return stateT
	}

	if a == symbolD {
		return stateD1
	}

	return stateERR
}

// deltaJ returns the transition from J.
func deltaJ(a symbolClass) state {
	if a == symbolH {
		return stateT
	}

	if a == symbolD {
		return stateD1
	}

	return stateERR
}

// deltaT returns the transition from T.
func deltaT(a symbolClass) state {
	if a == symbolD {
		return stateD1
	}

	return stateERR
}

// deltaD1 returns the transition from D1.
func deltaD1(a symbolClass) state {
	if a == symbolD {
		return stateD2
	}

	return stateERR
}

// deltaD2 returns the transition from D2.
func deltaD2(a symbolClass) state {
	if a == symbolH || a == symbolDash {
		return stateK
	}

	if a == symbolD {
		return stateD3
	}

	return stateERR
}

// deltaK returns the transition from K.
func deltaK(a symbolClass) state {
	if a == symbolD {
		return stateD3
	}

	return stateERR
}

// deltaD3 returns the transition from D3.
func deltaD3(a symbolClass) state {
	if a == symbolD {
		return stateD4
	}

	return stateERR
}

// deltaD4 returns the transition from D4.
func deltaD4(a symbolClass) state {
	if a == symbolD {
		return stateD5
	}

	return stateERR
}

// deltaD5 returns the transition from D5.
func deltaD5(a symbolClass) state {
	if a == symbolD {
		return stateD6
	}

	return stateERR
}

// deltaD6 returns the transition from D6.
func deltaD6(a symbolClass) state {
	if a == symbolH || a == symbolHash || a == symbolDash || a == symbolN0 {
		return stateF
	}

	return stateERR
}

// deltaF returns the transition from F.
func deltaF(_ symbolClass) state {
	return stateERR
}

// deltaERR returns the transition from ERR.
func deltaERR(_ symbolClass) state {
	return stateERR
}

// delta applies the transition function for the current state.
func delta(q state, a symbolClass) state {
	functions := [...]func(symbolClass) state{
		deltaS, deltaA, deltaP, deltaO, deltaJ, deltaT, deltaD1, deltaD2,
		deltaK, deltaD3, deltaD4, deltaD5, deltaD6, deltaF, deltaERR,
	}

	return functions[q](a)
}

// symbol classifies one input byte using the same ASCII rules as the C++ code.
func symbol(c byte) symbolClass {
	if c >= '0' && c <= '9' {
		return symbolD
	}

	switch c {
	case 'p':
		return symbolP
	case 'o':
		return symbolO
	case ' ', '\t', '\n', '\r', '\v', '\f':
		return symbolH
	case '#':
		return symbolHash
	case '-':
		return symbolDash
	}

	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' {
		return symbolOther
	}

	return symbolN0
}

// accepts reports whether the automaton accepts the whole word.
func accepts(word string) bool {
	q := stateS
	for i := 0; i < len(word); i++ {
		q = delta(q, symbol(word[i]))
	}

	return q == stateD6 || q == stateF
}

// main prints one numeric acceptance result for each command-line argument.
func main() {
	for _, word := range os.Args[1:] {
		if accepts(word) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
