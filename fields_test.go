package grst

import . "gopkg.in/check.v1"

type FieldSuite struct{}

var _ = Suite(&FieldSuite{})

func (s *FieldSuite) TestRstField(c *C) {
	f := &RstFieldSet{}
	fieldTester(c, f)
}
func fieldTester(c *C, f *RstFieldSet) {
	// add a bunch of elements to the field set
	for i := 0; i < 10; i++ {
		c.Assert(f.Len(), Equals, i)
		f.AddField("name", "value")
	}

	// check that we have all the fields
	c.Assert(f.Len(), Equals, 10)

	// check that the number of elements you put into the fieldSet
	// are the number you get out.
	c.Assert(f.resolve(0), HasLen, 10)

	// check resolution
	c.Assert(f.resolve(0), DeepEquals, []string{":name: value",
		":name: value", ":name: value", ":name: value", ":name: value",
		":name: value", ":name: value", ":name: value", ":name: value",
		":name: value"})

	// check resolution with indent
	c.Assert(f.resolve(2), DeepEquals, []string{"  :name: value",
		"  :name: value", "  :name: value", "  :name: value", "  :name: value",
		"  :name: value", "  :name: value", "  :name: value", "  :name: value",
		"  :name: value"})

	// check the iterator and it's output.
	for fp := range f.Iter() {
		c.Assert(fp, DeepEquals, &RstFieldPair{"name", "value"})
		c.Assert(fp.Name, Equals, "name")
		c.Assert(fp.Value, Equals, "value")
	}
}
