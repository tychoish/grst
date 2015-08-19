package grst

import (
	"fmt"
	"strings"
)

type RstFieldPair struct {
	Name  string
	Value string
}

type RstFieldSet struct {
	fieldPairs []RstFieldPair
}

func (self *RstFieldSet) Len() int {
	return len(self.fieldPairs)
}

func (self *RstFieldSet) Iter() <-chan *RstFieldPair {
	ch := make(chan *RstFieldPair)

	if self.Len() != 0 {
		go func() {
			for i := 0; i < self.Len(); i++ {
				ch <- &self.fieldPairs[i]
			}
		}()

	}

	return ch
}

func (self *RstFieldSet) AddField(name, value string) {
	f := RstFieldPair{
		Name:  name,
		Value: value,
	}
	self.fieldPairs = append(self.fieldPairs, f)
}

func (self *RstFieldSet) resolve(indent int) (lines []string) {
	for field := range self.Iter() {
		lines = append(lines, fmt.Sprintf("%s:%s: %s",
			strings.Repeat(" ", indent),
			field.Name, field.Value))
	}

	return
}
