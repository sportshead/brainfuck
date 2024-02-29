package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

const PageSize = 30000

type (
	Page  []byte
	Table map[int]Page
)

func getInput() []byte {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<filename>")
		os.Exit(1)
	}
	filename := os.Args[1]

	data, _ := os.ReadFile(filename)
	return data
}

func main() {
	input := getInput()

	memory := make(Table)
	memory[0] = make(Page, PageSize)

	ip := 0
	page := 0
	idx := 0

	for ip < len(input) {
		// fmt.Println(ip, page, idx, string(input[ip]))
		switch input[ip] {
		case '+':
			pg := memory[page]
			pg[idx]++
		case '-':
			pg := memory[page]
			pg[idx]--
		case '.':
			fmt.Print(string(memory[page][idx]))
		case ',':
			handleInput(memory, page, idx)
		case '>':
			idx++
			page, idx = wrapPointer(memory, page, idx)
		case '<':
			idx--
			page, idx = wrapPointer(memory, page, idx)
		case '[':
			if memory[page][idx] == 0 {
				pairs := 1
				for pairs > 0 && ip < len(input) {
					ip++
					switch input[ip] {
					case '[':
						pairs++
					case ']':
						pairs--
					}
				}
			}
		case ']':
			if memory[page][idx] != 0 {
				pairs := 1
				for pairs > 0 && ip > 0 {
					ip--
					switch input[ip] {
					case '[':
						pairs--
					case ']':
						pairs++
					}
				}
			}
		}
		ip++
	}
	// fmt.Println()
	// fmt.Println(memory)
}

func handleInput(memory Table, page, idx int) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	byte := make([]byte, 1)
	_, err = os.Stdin.Read(byte)
	if byte[0] == 0x04 || err == io.EOF {
		byte[0] = 0x00
	} else if err != nil {
		fmt.Println("Error reading from stdin:", err)
		byte[0] = 0x00
	}
	if byte[0] == 0x0D { // convert CR to LF
		byte[0] = 0x0A
	}
	pg := memory[page]
	pg[idx] = byte[0]
}

func wrapPointer(memory Table, page, idx int) (int, int) {
	if idx >= PageSize {
		page++
		idx = 0
	} else if idx < 0 {
		page--
		idx = PageSize - 1
	}
	if _, ok := memory[page]; !ok {
		memory[page] = make(Page, PageSize)
	}
	return page, idx
}
