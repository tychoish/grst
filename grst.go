package grst

import (
	"fmt"
	"strings"

	"./basic"
	"./hashed"
	"./headings"
	"./unsafe"
)

type RstGenerator interface {
	AddLine(string) error
	AddLines([]string) error
	GetLines() ([]string, error)
	Len() int
	SetIndent(int) error
	Indent() int
	IndentPadding() string
}

type RstBuilder struct {
	RstGenerator
}

func NewBasicBuilder() *RstBuilder {
	return &RstBuilder{basic.NewRstBuilder()}
}

func NewUnsafeBuilder() *RstBuilder {
	return &RstBuilder{unsafe.NewRstBuilder()}
}

func NewHashedBuilder() *RstBuilder {
	return &RstBuilder{hashed.NewRstBuilder()}
}

func (self *RstBuilder) NewLine() error {
	return self.AddLine("")
}

func (self *RstBuilder) NewLines(number int) error {
	return self.AddLine(strings.Repeat("\n", number-1))
}

func (self *RstBuilder) AddDirective(name, value string, fields RstFieldSet, content RstGenerator) error {
	var lines []string

	lines = append(lines, fmt.Sprintf(".. %s:: %s", name, value))

	if fields.Len() > 0 {
		lines = append(lines, fields.resolve(3)...)

	}

	if content.Len() > 0 {
		lines = append(lines, "")
		contentLines, _ := content.GetLines()
		lines = append(lines, contentLines...)
	}

	return self.AddLines(lines)
}

func (self *RstBuilder) Append(content RstBuilder) (err error) {
	lines, err := content.GetLines()
	if err != nil {
		return err
	}

	return self.AddLines(lines)
}

func (self *RstBuilder) LiCustom(char, text string) error {
	return self.AddLine(fmt.Sprintln(char, text))
}

func (self *RstBuilder) Li(text string) error {
	return self.LiCustom("-", text)
}

func (self *RstBuilder) Field(fields RstFieldSet) error {
	return self.AddLines(fields.resolve(0))
}

func (self *RstBuilder) AddBasicDirective(name string) error {
	return self.AddDirective(name, "", RstFieldSet{}, NewBasicBuilder())
}

func (self *RstBuilder) AddBasicDirectiveWithArgument(name, value string) error {
	return self.AddDirective(name, value, RstFieldSet{}, NewBasicBuilder())
}

func (self *RstBuilder) AddBasicDirectiveWithFields(name string, fields RstFieldSet) error {
	return self.AddDirective(name, "", fields, NewBasicBuilder())
}

func (self *RstBuilder) AddBasicDirectiveWithArgumentAndFields(name, value string, fields RstFieldSet) error {
	return self.AddDirective(name, value, fields, NewBasicBuilder())
}

func (self *RstBuilder) AddBasicDirectiveWithArgumentAndContent(name, value string, content RstBuilder) error {
	return self.AddDirective(name, value, RstFieldSet{}, NewBasicBuilder())
}

func (self *RstBuilder) AddBasicDirectiveWithFieldsAndContent(name string, fields RstFieldSet, content RstBuilder) error {
	return self.AddDirective(name, "", fields, NewBasicBuilder())
}

func (self *RstBuilder) Replacement(name, value string) error {
	return self.AddLine(fmt.Sprintf(".. |%s| replace:: %s", name, value))
}

func (self *RstBuilder) Target(name string) error {
	return self.AddLine(fmt.Sprintf(".. _%s:", name))
}

func (self *RstBuilder) Title(text string) error {
	rLine := headingLine(text, headings.One)

	return self.AddLines([]string{rLine, text, rLine})
}

func (self *RstBuilder) H1(text string) error {
	return self.AddLines([]string{text, headingLine(text, headings.One)})
}

func (self *RstBuilder) H2(text string) error {
	return self.AddLines([]string{text, headingLine(text, headings.Two)})
}

func (self *RstBuilder) H3(text string) error {
	return self.AddLines([]string{text, headingLine(text, headings.Three)})
}

func (self *RstBuilder) H4(text string) error {
	return self.AddLines([]string{text, headingLine(text, headings.Four)})
}

func (self *RstBuilder) H5(text string) error {
	return self.AddLines([]string{text, headingLine(text, headings.Five)})
}

func (self *RstBuilder) H6(text string) error {
	return self.AddLines([]string{text, headingLine(text, headings.Six)})
}
