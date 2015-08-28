package grst

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tychoish/grip"
)

func (self *RstBuilder) Write(fn string) error {
	lines, err := self.GetLines()
	if err != nil {
		return err
	}

	dirName := filepath.Dir(fn)
	err = os.MkdirAll(dirName, 0755)
	if err == nil {
		grip.Noticeln("created directory:", dirName)
	} else {
		grip.Warning(err.Error())
	}

	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	catcher := grip.NewCatcher()
	var numBytes int
	for _, line := range lines {
		nb, err := fmt.Fprintln(w, line)
		numBytes += nb
		catcher.Add(err)
	}

	catcher.Add(w.Flush())

	if catcher.HasErrors() == false {
		grip.Debugf("wrote %d bytes to file '%f'.", numBytes, fn)
	}

	return catcher.Resolve()
}

func (self *RstBuilder) Print() error {
	lines, err := self.GetLines()

	fmt.Println(strings.Join(lines, "\n"))

	return err
}

func (self *RstBuilder) Resolve() (string, error) {
	lines, err := self.GetLines()

	return strings.Join(lines, "\n") + "\n", err
}
