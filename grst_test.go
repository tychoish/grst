package grst

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type GeneratorSuite struct {
	*RstBuilder
}

var _ = Suite(&GeneratorSuite{NewUnsafeBuilder()})

var _ = Suite(&GeneratorSuite{NewBasicBuilder()})

var _ = Suite(&GeneratorSuite{NewHashedBuilder()})

func (s *GeneratorSuite) SetUpTest(c *C) {
	s.Reset()
}

func (s *GeneratorSuite) TestConstructor(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	var b RstGenerator

	c.Assert(s, Implements, &b)
}

func (s *GeneratorSuite) TestNewLine(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	var result string

	c.Assert(s.Len(), Equals, 0)
	s.NewLine()
	c.Assert(s.Len(), Equals, 1)
	s.Reset()

	for i := 0; i < 10; i++ {
		s.NewLine()
	}
	c.Assert(s.Len(), Equals, 10)
	result, _ = s.Resolve()
	c.Assert(strings.Count(result, "\n"), Equals, 10)

	s.Reset()
	s.NewLines(10)
	c.Assert(s.Len(), Equals, 10)

	result, _ = s.Resolve()
	c.Assert(strings.Count(result, "\n"), Equals, 10)

	// s.Reset()
}
func (s *GeneratorSuite) TestHeadingTitle(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)
	t := "this is the title"
	s.Title(t)
	c.Logf(s.Resolve())
	c.Assert(s.Len(), Equals, 3)

	lines, _ := s.GetLines()
	for _, ln := range lines {
		c.Assert(len(ln), Equals, len(t))
	}

}

func (s *GeneratorSuite) TestHeadings(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	text := "heading text"

	type headingMethod func(string) error
	for idx, fn := range []headingMethod{s.H1, s.H2, s.H3, s.H4, s.H5, s.H6} {
		c.Logf("testing heading level: %d", idx)
		err := fn(text)
		c.Assert(err, IsNil)
		c.Assert(s.Len(), Equals, 2)
		lines, err := s.GetLines()
		for _, ln := range lines {
			c.Assert(len(ln), Equals, len(text))
		}

		s.Reset()
	}

}

func (s *GeneratorSuite) TestTarget(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	err := s.Target("target-tag")
	c.Assert(err, IsNil)
	c.Assert(s.Len(), Equals, 1)

	lines, err := s.GetLines()
	c.Assert(err, IsNil)
	c.Assert(len(lines), Equals, 1)
	c.Assert(len(lines[0]), Equals, len("target-tag")+5)

	c.Assert(s.Len(), Equals, 1)
}

func (s *GeneratorSuite) TestReplacement(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	err := s.Replacement("replace", "value")
	c.Assert(err, IsNil)
	c.Assert(s.Len(), Equals, 1)

	lines, err := s.GetLines()
	c.Assert(err, IsNil)
	c.Assert(len(lines), Equals, 1)
	c.Assert(len(lines[0]), Equals, len("replace")+len("value")+16)

	c.Assert(s.Len(), Equals, 1)
}

func (s *GeneratorSuite) TestRstField(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	f := &RstFieldSet{}
	fieldTester(c, f)

	c.Assert(s.Len(), Equals, 0)
	s.Field(*f)

	c.Assert(s.Len(), Equals, 10)

	lines, err := s.GetLines()
	c.Assert(err, IsNil)
	c.Assert(len(lines), Equals, 10)
	for _, ln := range lines {
		c.Assert(len(ln), Equals, 12)
		c.Assert(ln, Equals, ":name: value")

	}

}
func (s *GeneratorSuite) TestLi(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	s.Li("this is a bullet")

	c.Assert(s.Len(), Equals, 1)

	result, err := s.Resolve()
	c.Assert(result, Equals, "- this is a bullet\n")
	c.Assert(err, IsNil)

	s.Reset()
	for i := 1; i > 11; i++ {
		s.Li("bullet")
		c.Assert(s.Len(), Equals, i)
	}

	lines, _ := s.GetLines()
	for _, content := range lines {
		c.Assert(content, Equals, "- bullet")
	}
}

func (s *GeneratorSuite) TestLiCustomChar(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	chars := []string{"*", "-", "~"}
	for idx, char := range chars {
		s.LiCustom(char, "bullet")
		c.Assert(s.Len(), Equals, idx+1)
		lines, _ := s.GetLines()
		c.Assert(lines[idx], Equals, char+" bullet")
	}

	lines, _ := s.GetLines()

	for idx, ln := range lines {
		c.Assert(string(ln[0]), Equals, chars[idx])

	}
}

func (s *GeneratorSuite) TestAddBasicDirective(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	s.AddBasicDirective("directive")
	c.Assert(s.Len(), Equals, 1)

	result, _ := s.GetLines()
	c.Assert(result[0], Equals, ".. directive::")
}

func (s *GeneratorSuite) TestAddBasicDirectiveWithArgument(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	c.Assert(s.Len(), Equals, 0)

	s.AddBasicDirectiveWithArgument("directive", "argument")
	c.Assert(s.Len(), Equals, 1)

	result, _ := s.GetLines()
	c.Assert(result[0], Equals, ".. directive:: argument")
	text, _ := s.Resolve()
	c.Assert(text, Equals, ".. directive:: argument\n")
}

func (s *GeneratorSuite) TestAddBasicDirectiveWIthFields(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	f := RstFieldSet{}

	c.Assert(s.Len(), Equals, 0)
	c.Assert(f.Len(), Equals, 0)

	f.AddField("name", "value")
	c.Assert(f.Len(), Equals, 1)

	s.AddBasicDirectiveWithFields("directive", f)
	c.Assert(s.Len(), Equals, 2)
	lines, _ := s.Resolve()

	c.Assert(lines, Equals, ".. directive::\n   :name: value\n")
}

func (s *GeneratorSuite) TestAddBasicDirectiveWithFieldsAndArgument(c *C) {
	c.Logf("running test with %s backend", s.Builder())

	f := RstFieldSet{}

	c.Assert(s.Len(), Equals, 0)
	c.Assert(f.Len(), Equals, 0)

	f.AddField("name", "value")
	c.Assert(f.Len(), Equals, 1)

	s.AddBasicDirectiveWithArgumentAndFields("directive", "argument", f)
	c.Assert(s.Len(), Equals, 2)
	lines, _ := s.Resolve()

	c.Assert(lines, Equals, ".. directive:: argument\n   :name: value\n")
}
