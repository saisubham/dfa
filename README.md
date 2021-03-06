# DFA in Go
Implementation of a simple Deterministic Finite Automata in Go.

## Installation
```
git clone https://github.com/saisubham/dfa.git
```

## Compiling
```
cd dfa
go mod tidy
```

## Testing
```
go test -v
```

## Sample program
```
func main() {
	dfa, err := MakeDFA(
		[]int{0, 1, 2},   // all states
		[]rune{'0', '1'}, // all symbols
		0,                // initial state
		[]int{0},         // final states
	)
	if err != nil {
		t.Fail()
	}
	err = dfa.AddTransitions([]*Transitions{
		{0, '0', 0},
		{0, '1', 1},
		{1, '0', 2},
		{1, '1', 0},
		{2, '0', 1},
		{2, '1', 2},
	})
	if err != nil {
		t.Fail()
	}
	out, err := dfa.Run("1001")
	if err != nil {
		log.Printf(err.Error())
	}

	if out {
		fmt.Println("accepted")
	} else {
		fmt.Println("rejected")
	}
}
```

## Running
Assuming sample program is stored in main.go
```
go run main.go
```
