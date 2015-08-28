package grst

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type GeneratorSuite struct {
	b *RstBuilder
}

var _ = Suite(&GeneratorSuite{b: NewUnsafeBuilder()})

var _ = Suite(&GeneratorSuite{b: NewBasicBuilder()})

var _ = Suite(&GeneratorSuite{b: NewHashedBuilder()})

func (s *GeneratorSuite) SetUpTest(c *C) {
	s.b.Reset()
}

func (s *GeneratorSuite) TestConstructor(c *C) {
	var b RstGenerator

	c.Assert(s.b, Implements, &b)
}

func (s *GeneratorSuite) TestNewLine(c *C) {
	var result string

	c.Assert(s.b.Len(), Equals, 0)
	s.b.NewLine()
	c.Assert(s.b.Len(), Equals, 1)
	s.b.Reset()

	for i := 0; i < 10; i++ {
		s.b.NewLine()
	}
	c.Assert(s.b.Len(), Equals, 10)
	result, _ = s.b.Resolve()
	c.Assert(strings.Count(result, "\n"), Equals, 10)

	s.b.Reset()
	s.b.NewLines(10)
	c.Assert(s.b.Len(), Equals, 10)

	result, _ = s.b.Resolve()
	c.Assert(strings.Count(result, "\n"), Equals, 10)

	// s.b.Reset()
}
