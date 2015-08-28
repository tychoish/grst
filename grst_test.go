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
	c.Logf("running test with %s backend", s.b.Builder())

	var b RstGenerator

	c.Assert(s.b, Implements, &b)
}

func (s *GeneratorSuite) TestNewLine(c *C) {
	c.Logf("running test with %s backend", s.b.Builder())

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

func (s *GeneratorSuite) TestHeadingTitle(c *C) {
	c.Logf("running test with %s backend", s.b.Builder())

	c.Assert(s.b.Len(), Equals, 0)
	t := "this is the title"
	s.b.Title(t)
	c.Logf(s.b.Resolve())
	c.Assert(s.b.Len(), Equals, 3)

	lines, _ := s.b.GetLines()
	for _, ln := range lines {
		c.Assert(len(ln), Equals, len(t))
	}

}

func (s *GeneratorSuite) TestHeadings(c *C) {
	c.Logf("running test with %s backend", s.b.Builder())

	c.Assert(s.b.Len(), Equals, 0)

	text := "heading text"

	type headingMethod func(string) error
	for idx, fn := range []headingMethod{s.b.H1, s.b.H2, s.b.H3, s.b.H4, s.b.H5, s.b.H6} {
		c.Logf("testing heading level: %d", idx)
		err := fn(text)
		c.Assert(err, IsNil)
		c.Assert(s.b.Len(), Equals, 2)
		lines, err := s.b.GetLines()
		for _, ln := range lines {
			c.Assert(len(ln), Equals, len(text))
		}

		s.b.Reset()
	}

}

func (s *GeneratorSuite) TestTarget(c *C) {
	c.Logf("running test with %s backend", s.b.Builder())

	c.Assert(s.b.Len(), Equals, 0)

	err := s.b.Target("target-tag")
	c.Assert(err, IsNil)
	c.Assert(s.b.Len(), Equals, 1)

	lines, err := s.b.GetLines()
	c.Assert(err, IsNil)
	c.Assert(len(lines), Equals, 1)
	c.Assert(len(lines[0]), Equals, len("target-tag")+5)

	c.Assert(s.b.Len(), Equals, 1)
}

func (s *GeneratorSuite) TestReplacement(c *C) {
	c.Logf("running test with %s backend", s.b.Builder())

	c.Assert(s.b.Len(), Equals, 0)

	err := s.b.Replacement("replace", "value")
	c.Assert(err, IsNil)
	c.Assert(s.b.Len(), Equals, 1)

	lines, err := s.b.GetLines()
	c.Assert(err, IsNil)
	c.Assert(len(lines), Equals, 1)
	c.Assert(len(lines[0]), Equals, len("replace")+len("value")+16)

	c.Assert(s.b.Len(), Equals, 1)
}

func (s *GeneratorSuite) TestRstField(c *C) {
	f := &RstFieldSet{}
	fieldTester(c, f)

	c.Assert(s.b.Len(), Equals, 0)
	s.b.Field(*f)

	c.Assert(s.b.Len(), Equals, 10)

	lines, err := s.b.GetLines()
	c.Assert(err, IsNil)
	c.Assert(len(lines), Equals, 10)
	for _, ln := range lines {
		c.Assert(len(ln), Equals, 12)
		c.Assert(ln, Equals, ":name: value")

	}

}
