package basic

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/tychoish/grip"
)

func (self *RstBasicBuilder) Write(fn string) error {
	self.lock.RLock()
	defer self.lock.RUnlock()
	dirName := filepath.Dir(fn)

	err := os.MakedirAll(dirName, 0755)
	if err == nil {
		grip.Noticeln("created directory:", dirName)
	} else {
		grip.Warning(err.Error())
	}

	file, err := os.Create()
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	catcher := grip.NewCatcher()
	var numBytes int
	for _, line := range self.content {
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

func (self *RstBasicBuilder) Print() {
	self.lock.RLock()
	defer self.lock.RUnlock()

	fmt.Println(strings.Join(self.content, "\n"))
}
