package hashed

import (
	"fmt"
	"strings"
	"sync"
)

type RstHashedBuilder struct {
	content map[int]string
	indent  int
	line    int

	sync.RWMutex
}

func NewRstBuilder() *RstHashedBuilder {
	return &RstHashedBuilder{
		content: make(map[int]string),
	}
}

func (self *RstHashedBuilder) Builder() string {
	return "hashed"
}

func (self *RstHashedBuilder) Reset() {
	self.Lock()
	defer self.Unlock()

	self.content = make(map[int]string)
	self.line = 0
	self.indent = 0
}

func (self *RstHashedBuilder) AddLine(line string) (err error) {
	self.Lock()
	defer self.Unlock()

	self.line += 1
	self.content[self.line] = strings.TrimRight(self.IndentPadding()+line, " \t\n\r")

	return
}

func (self *RstHashedBuilder) AddLines(lines []string) (err error) {
	self.Lock()
	defer self.Unlock()

	leftPadding := self.IndentPadding()

	for i := 0; i < len(lines); i++ {
		self.line += 1
		self.content[self.line] = strings.TrimRight(leftPadding+lines[i], " \t\n\r")

	}

	return
}

func (self *RstHashedBuilder) GetLines() (lines []string, err error) {
	self.RLock()
	defer self.RUnlock()

	for i := 0; i <= self.line+1; i++ {
		ln, ok := self.content[i]
		if ok {
			lines = append(lines, ln)
		}
	}

	return
}

func (self *RstHashedBuilder) Len() int {
	return self.line
}

func (self *RstHashedBuilder) IndentPadding() string {
	return strings.Repeat(" ", self.Indent())
}

func (self *RstHashedBuilder) Indent() int {
	return self.indent

}
func (self *RstHashedBuilder) SetIndent(n int) (err error) {
	if n < 0 {
		err = fmt.Errorf("Indent must be larger than 0, %d is not.", n)
	}

	self.Lock()
	defer self.Unlock()

	self.indent = n

	return
}
