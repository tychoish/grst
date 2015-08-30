package grst

import (
	"strconv"

	. "gopkg.in/check.v1"
)

type BuilderSuite struct {
	*RstBuilder
}

var _ = Suite(&BuilderSuite{NewUnsafeBuilder()})

var _ = Suite(&BuilderSuite{NewBasicBuilder()})

var _ = Suite(&BuilderSuite{NewHashedBuilder()})

func (s *BuilderSuite) SetUpTest(c *C) {
	s.Reset()
}

func (s *BuilderSuite) TestImplements(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	var b RstGenerator

	c.Assert(s, Implements, &b)
}

func (s *BuilderSuite) TestResetAndLength(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	for i := 0; i < 10; i++ {
		c.Assert(s.Len(), Equals, i)
		s.NewLine()

	}
	c.Assert(s.Len(), Equals, 10)

	s.Reset()

	c.Assert(s.Len(), Equals, 0)
}

func (s *BuilderSuite) TestLineAndLinesAdders(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	for i := 1; i <= 5; i++ {
		s.AddLine(strconv.Itoa(i))
	}

	c.Assert(s.Len(), Equals, 5)

	one, _ := s.GetLines()

	s.Reset()
	s.AddLines([]string{"1", "2", "3", "4", "5"})
	c.Assert(s.Len(), DeepEquals, 5)

	two, _ := s.GetLines()

	for i := 0; i < s.Len(); i++ {
		c.Assert(one[i], Equals, two[i])
	}
}

func (s *BuilderSuite) TestIndenters(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Indent(), Equals, 0)

	c.Assert(s.SetIndent(5), IsNil)
	c.Assert(s.Indent(), Equals, 5)

	s.SetIndent(0)
	c.Assert(s.Indent(), Equals, 0)

	c.Assert(s.SetIndent(-1), ErrorMatches, "Indent must be larger than 0, -1 is not.")

	s.SetIndent(3)

	c.Assert(s.IndentPadding(), Equals, "   ")

	c.Assert(s.Len(), Equals, 0)
	s.AddLine("indent check")

	lns, _ := s.GetLines()

	c.Assert(lns[0], Equals, "   indent check")
}
