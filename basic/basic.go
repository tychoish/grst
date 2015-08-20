package basic

import (
	"fmt"
	"strings"
	"sync"
)

type RstBasicBuilder struct {
	indent  int
	content []string
	lock    *sync.RWMutex
}

func NewRstBuilder() *RstBasicBuilder {
	return &RstBasicBuilder{
		lock: &sync.RWMutex{},
	}
}

func (self *RstBasicBuilder) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return len(self.content)
}

func (self *RstBasicBuilder) AddLine(line string) (err error) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.content = append(self.content, strings.TrimRight(self.IndentPadding()+line, " \t\n\r"))

	return
}

func (self *RstBasicBuilder) AddLines(lines []string) (err error) {
	self.lock.Lock()
	defer self.lock.Unlock()

	leftPadding := self.IndentPadding()

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimRight(leftPadding+lines[i], " \t\n\r")
	}

	self.content = append(self.content, lines...)

	return
}

func (self *RstBasicBuilder) GetLines() (output []string, err error) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	copy(output, self.content)

	return
}

func (self *RstBasicBuilder) IndentPadding() string {
	return strings.Repeat(" ", self.Indent())
}

func (self *RstBasicBuilder) Indent() int {
	return self.indent

}
func (self *RstBasicBuilder) SetIndent(n int) (err error) {
	if n < 0 {
		err = fmt.Errorf("Indent must be larger than 0, %d is not.", n)
	}

	self.lock.Lock()
	defer self.lock.Unlock()

	self.indent = n

	return
}
