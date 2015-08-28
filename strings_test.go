package grst

import (
	"strings"

	. "gopkg.in/check.v1"

	h "./headings"
)

type StringSuite struct{}

var _ = Suite(&StringSuite{})

func (s *FieldSuite) TestHeadingLineGenerator(c *C) {
	for _, char := range []string{h.One, h.Two, h.Three, h.Four, h.Five, h.Six} {
		for i := 1; i < 11; i++ {
			heading := strings.Repeat("x", i)
			hl := headingLine(heading, char)
			c.Assert(len(hl), Equals, i)
			c.Assert(len(heading), Equals, len(hl))

			hl = headingLine(heading, "++")
			c.Assert(len(hl), Equals, i)
			c.Assert(len(heading), Equals, len(hl))
		}
	}

}

func (s *FieldSuite) TestRoles(c *C) {
	c.Assert(Role("name", "value"), Equals, ":name:`value`")
	c.Assert(RoleWithText("name", "value", "text"), Equals, ":name:`text <value>`")
}

func (s *FieldSuite) TestFormatting(c *C) {
	c.Assert(Bold("text"), Equals, "**text**")
	c.Assert(Bold("complex text"), Equals, "**complex text**")

	c.Assert(Emph("text"), Equals, "*text*")
	c.Assert(Emph("complex text"), Equals, "*complex text*")

	c.Assert(Pre("text"), Equals, "``text``")
	c.Assert(Pre("complex text"), Equals, "``complex text``")
}

func (s *FieldSuite) TestRefs(c *C) {
	c.Assert(InlineLink("text", "link"), Equals, "`text <link>`_")
	c.Assert(FootnoteRef("note"), Equals, "[#note]_")
}
