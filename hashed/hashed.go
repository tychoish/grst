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

	*sync.RWMutex
}

func NewRstBuilder() *RstHashedBuilder {
	return &RstHashedBuilder{
		content: make(map[int]string),
	}
}

func (self *RstHashedBuilder) AddLine(line string) (err error) {
	self.Lock()
	defer self.Unlock()

	self.content[self.line] = strings.TrimRight(self.IndentPadding()+line, " \t\n\r")
	self.line += 1

	return
}

func (self *RstHashedBuilder) AddLines(lines []string) (err error) {
	self.Lock()
	defer self.Unlock()

	leftPadding := self.IndentPadding()

	for i := 0; i < len(lines); i++ {
		self.content[self.line] = strings.TrimRight(leftPadding+lines[i], " \t\n\r")
		self.line += 1
	}

	return
}

func (self *RstHashedBuilder) GetLines() (lines []string, err error) {
	self.RLock()
	defer self.RUnlock()

	for i := 0; i <= self.line; i++ {
		lines = append(lines, self.content[i])
	}

	return
}

func (self *RstHashedBuilder) Len() int {
	return self.line + 1
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
