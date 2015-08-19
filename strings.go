package grst

import "fmt"

// These Functions return strings that are used

func Role(name, value string) string {
	return fmt.Sprintf(":%s:`%s`", name, value)
}

func RoleWithText(name, value, text string) string {
	return fmt.Sprintf(":%s:`%s <%s>`", name, text, value)
}

func Bold(text string) string {
	return "**" + text + "**"
}

func Emph(text string) string {
	return "*" + text + "*"
}

func Pre(text string) string {
	return "``" + text + "``"
}

func InlineLink(text, link string) string {
	return fmt.Sprintf("`%s <%s>`_", text, link)
}

func FootnoteRef(name string) string {
	return "[#" + name + "]"
}
