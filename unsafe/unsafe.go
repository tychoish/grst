package unsafe

import (
	"fmt"
	"strings"
)

type RstUnsafeBuilder struct {
	indent  int
	content []string
}

func NewRstBuilder() *RstUnsafeBuilder {
	return &RstUnsafeBuilder{}
}

func (self *RstUnsafeBuilder) Len() int {
	return len(self.content)
}

func (self *RstUnsafeBuilder) AddLine(line string) (err error) {
	self.content = append(self.content, strings.TrimRight(self.IndentPadding()+line, " \t\n\r"))

	return
}

func (self *RstUnsafeBuilder) AddLines(lines []string) (err error) {
	leftPadding := self.IndentPadding()

	for i := 0; i < len(lines); i++ {
		lines[i] = leftPadding + lines[i]
	}

	self.content = append(self.content, lines...)

	return
}

func (self *RstUnsafeBuilder) GetLines() (lines []string, err error) {
	copy(lines, self.content)

	return
}

func (self *RstUnsafeBuilder) IndentPadding() string {
	return strings.Repeat(" ", self.Indent())
}

func (self *RstUnsafeBuilder) Indent() int {
	return self.indent

}

func (self *RstUnsafeBuilder) SetIndent(n int) (err error) {
	if n < 0 {
		err = fmt.Errorf("Indent must be larger than 0, %d is not.", n)
	}

	self.indent = n

	return
}
