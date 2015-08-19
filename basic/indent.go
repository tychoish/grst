package basic

import (
	"fmt"
	"strings"
)

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
