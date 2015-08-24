package grst

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/tychoish/grip"
)

type RstTable struct {
	width          int
	hasHeader      bool
	columnWidths   []int
	maxFieldWidths []int
	columnNames    []string
	rows           [][]string

	*sync.RWMutex
}

func (self *RstTable) check() (err error) {
	if self.width == 0 {
		return
	}

	if self.width != len(self.columnNames) && self.width != len(self.maxFieldWidths) {
		err = fmt.Errorf("column count inconsistent. internal error. (width of %d, with (%d, %d) columns, )",
			self.width, len(self.columnNames), len(self.maxFieldWidths))
	}

	return
}

func (self *RstTable) validate(fields []string) (err error) {
	if self.width == 0 {
		// this is the first row, which all is correct.
		self.width = len(fields)
		self.maxFieldWidths = make([]int, self.width)

		return nil
	} else {
		catcher := grip.NewCatcher()

		if len(fields) != self.width {
			catcher.Add(fmt.Errorf("row [$s] has %d columns, not %d, the required.",
				strings.Join(fields, ","), len(self.rows), self.width))
		}
		if len(fields) != len(self.maxFieldWidths) {
			catcher.Add(fmt.Errorf("row [$s] has %d columns, not %d, the required.",
				strings.Join(fields, ","), len(self.rows), len(self.maxFieldWidths)))
		}

		return catcher.Resolve()
	}

}

func (self *RstTable) validateTable() error {
	catcher := grip.NewCatcher()

	catcher.Add(self.check()) // to validate width and column names

	for idx, row := range self.rows {
		if len(row) != self.width {
			catcher.Add(fmt.Errorf("row number %d has %d columns, not the same width as the table (%d)",
				idx, len(row), self.width))
		}
	}

	return catcher.Resolve()
}

func (self *RstTable) AddRow(fields ...string) (err error) {
	self.Lock()
	defer self.Unlock()

	err = self.check()
	if err != nil {
		return
	}

	err = self.validate(fields)
	if err != nil {
		return
	}

	for idx, field := range fields {
		if len(field) > self.maxFieldWidths[idx] {
			self.maxFieldWidths[idx] = len(field)
		}
	}

	if len(self.rows) < 1 {
		self.columnNames = append(self.columnNames, fields...)
		self.rows = append(self.rows, fields)
	} else {
		self.rows = append(self.rows, fields)
	}

	return
}

func (self *RstTable) EnableHeader() {
	self.hasHeader = true
}

func (self *RstTable) DisableHeader() {
	self.hasHeader = false
}

func (self *RstTable) SetWidths(widths ...int) error {
	self.Lock()
	defer self.Unlock()

	if self.width == 0 {
		self.width = len(widths)
	} else if self.width != len(widths) {
		return fmt.Errorf("cannot set widths for %d columns. there are %d columns in this table.",
			len(widths), self.width)
	}

	var total int
	for _, w := range widths {
		total += w
	}

	if total != 100 {
		return fmt.Errorf("cannot set column widths that add up to more than 100 (%d)", total)
	} else {
		self.columnWidths = widths

		return nil
	}
}

func (self *RstTable) JoinedWidths() string {
	var parts []string

	for _, w := range self.columnWidths {
		parts = append(parts, strconv.Itoa(w))
	}

	return strings.Join(parts, " ")
}

func (self *RstBuilder) ListTable(table *RstTable) error {
	table.RLock()
	defer table.RUnlock()

	lines := NewUnsafeBuilder()

	fields := RstFieldSet{}
	if table.hasHeader == true {
		fields.AddField("header-rows", strconv.Itoa(len(table.columnNames)))
	}

	if len(table.columnWidths) == table.width {
		// we just care that it's not 0, but might as well here.

		fields.AddField("widths", table.JoinedWidths())
	}

	lines.AddBasicDirectiveWithFields("list-table", fields)

	for _, row := range table.rows {
		lines.NewLine()
		for idx, field := range row {
			if idx == 0 {
				lines.LiCustom("* -", field)
			} else {
				lines.LiCustom("  -", field)
			}
		}
	}

	outputLines, err := lines.GetLines()
	if err != nil {
		return err
	} else {
		return self.AddLines(outputLines)
	}
}

func (self *RstBuilder) StandardTable(table *RstTable) error {
	table.RLock()
	defer table.RUnlock()

	lines := NewUnsafeBuilder()
	catcher := grip.NewCatcher()

	columnLines := make([]string, len(table.maxFieldWidths))
	for idx, col := range table.maxFieldWidths {
		columnLines[idx] = strings.Repeat("-", col)
	}
	rowSeperator := "+" + strings.Join(columnLines, "+") + "+"

	catcher.Add(lines.AddLine(rowSeperator))
	for _, row := range table.rows {
		paddedFields := make([]string, len(table.maxFieldWidths))

		for idx, field := range row {
			if len(field) < table.maxFieldWidths[idx] {
				paddedFields[idx] = field + strings.Repeat(" ", table.maxFieldWidths[idx]-len(field))
			} else {
				paddedFields[idx] = field
			}
		}

		catcher.Add(lines.AddLine("|" + strings.Join(paddedFields, "|") + "|"))
		catcher.Add(lines.AddLine(rowSeperator))
	}

	return catcher.Resolve()

}
