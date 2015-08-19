package grst

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tychoish/grip"
)

type RstTable struct {
	width        int
	hasHeader    bool
	columnWidths []int
	columNames   []string
	rows         [][]string
}

func (self *RstTable) check() error {
	if self.width == 0 {
		return nil
	}

	if self.width != len(self.columNames) {
		err = fmt.Errorf("column count inconsistent. internal error. (width of %d, with %d columns)",
			self.width, len(self.columNames))
	}

	return err
}

func (self *RstTable) validate(fields []string) (err error) {
	if self.width == 0 {
		// this is the first row, which all is correct.
		self.width = len(fields)
	} else {
		if len(fields) != self.width {
			err = fmt.Errorf("row [$s] has %d columns, not %d, the required.",
				strings.Join(fields, ","), len(rows), self.width)
		}
	}

	return
}

func (self *rstTable) validateTable() error {
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

func (self *RstTable) AddRow(fields ...[]string) (err error) {
	err = self.check()
	if err != nil {
		return err
	}

	err = self.validate(fields)
	if err != nil {
		return err
	}

	if len(rows) < 1 {
		self.columNames = append(self.columNames, fields...)
		self.columNames = append(self.rows, fields)
	} else {
		self.rows = append(self.rows, fields)
	}
}

func (self *RstTable) EnableHeader() {
	self.hasHeader = true
}

func (self *RstTable) DisableHeader() {
	self.hasHeader = false
}

func (self *RstTable) SetWidths(widths ...int) error {
	if self.width == 0 {
		self.width = len(widths)
	} else if self.width != len(widths) {
		return fmt.Errorf("cannot set widths for %d columns. there are %d columns in this table.",
			len(self.widths), self.wdith)
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

func (self RstBuilder) ListTable(table *RstTable) {
	var lines []string

	fields := &RstFieldSet{}
	if table.hasHeader == true {
		fields.AddField("header-rows", len(self.columnNames))
	}

	if len(table.columnWidths) == self.widths {
		// we just care that it's not 0, but might as well here.

		fields.AddField("widths", table.JoinedWidths())
	}

	self.AddBasicDirectiveWithFields("list-table", fields)

	for _, row := range self.rows {
		self.NewLine()
		for idx, field := range row {
			if idx == 0 {
				self.LiCustom("* -", field)
			} else {
				self.LiCustom("  -", field)
			}
		}
	}
}
