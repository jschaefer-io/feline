package main

import "fmt"

// token.Value.(float64)

func main() {
	file, fileErr := prepareFile("./testfile.feline")
	if fileErr != nil {
		panic(fileErr)
	}
	test(file)
}

func test(file File) {
	for _, line := range file.lines {
		fmt.Printf("Line %d:", line.number)
		for _, token := range line.tokens {
			fmt.Println("\t" + token.ToString())
		}
	}
}

/*
func tokenlist(...list int)[]int{
	return list
}

func startswith tokenlist  endswith tokenlist
	-> erst startswith
	-> dann sub-parse process
	-> endswith


# Expression besteht aus:
- Term
	Token (Value)
	group/Expression
	functionCall
- Operator (Type)

# variable set besteht aus:
let (keyword)
xyz (keyword)
type (keyword)
= Operator
group/Expression
STOP (token)

# function define besteht aus:
keyword
group
	keyword
	type keyword
	komma delimiter
scope

# function call besteht aus:
keyword
group
	keyword
	type keyword
	komma delimiter
*/

func prepareFile(path string) (File, error) {
	file := NewFile(path)
	err := file.Prepare()
	if err != nil {
		return file, err
	}
	return file, nil
}
