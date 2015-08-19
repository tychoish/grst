package basic

import (
	"strings"
	"sync"
)

type RstBasicBuilder struct {
	indent  int
	content []string
	lock    *sync.RWMutex
}

func NewRstBasicBuilder() *RstBasicBuilder {
	return &RstBasicBuilder{
		lock: &sync.RWMutex{},
	}
}

func (self *RstBasicBuilder) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return len(self.content)
}

func (self *RstBasicBuilder) add(line string) (err error) {
	self.content = append(self.content, strings.TrimRight(self.IndentPadding()+line, " \t\n\r"))

	return
}

func (self *RstBasicBuilder) addMultiple(lines []string) (err error) {
	leftPadding := self.IndentPadding()

	for i := 0; i < len(lines); i++ {
		lines[i] = leftPadding + lines[i]
	}

	self.content = append(self.content, lines...)

	return
}

func (self *RstBasicBuilder) AddLine(line string) (err error) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.add(line)

	return
}

func (self *RstBasicBuilder) AddLines(lines []string) (err error) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.addMultiple(lines)

	return
}

func (self *RstBasicBuilder) GetLines() (lines []string, err error) {
	self.lock.RLock()
	defer self.lock.RUnlock()

	var output []string
	copy(output, self.content)

	leftPadding := self.IndentPadding()

	for i := 0; i < len(lines); i++ {
		lines[i] = leftPadding + lines[i]
	}

	return output, nil
}
