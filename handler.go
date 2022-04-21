package clog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// handler manages msg output.
type handler interface {
	output(msg string)
	close()
}

// Handler outputting to stderr.
type stdOutHandler struct {
}

func (s stdOutHandler) output(msg string) {
	os.Stdout.Write([]byte(msg))
}

func (s stdOutHandler) close() {
}

// activeHandler managing output of messages.
var activeHandler handler = stdOutHandler{}

type fileHandler struct {
	mtx      sync.Mutex
	prefix   string   // file name prefix
	maxFiles int      // maximum number of files to keep
	maxLen   int      // maximum byte length of file
	bytes    int      // number of bytes written in current file
	file     *os.File // current logging file
}

func (f *fileHandler) close() {
	if f.file != nil {
		f.file.Close()
		f.file = nil
		f.bytes = 0
	}
}

func (f *fileHandler) output(msg string) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	if f.file != nil && f.bytes+len(msg) > f.maxLen {
		f.close()
	}
	if f.file == nil {
		fileDir := filepath.Dir(f.prefix)
		fileBase := filepath.Base(f.prefix)
		symlinkName := f.prefix + "_latest"
		symlinkBase := filepath.Base(symlinkName)
		// make sure directory exist
		if err := os.MkdirAll(fileDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "clog error: %s\n", err)
			return
		}
		// purge old files if required
		files, err := os.ReadDir(fileDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "clog error: %s\n", err)
			return
		}
		clogFiles := make([]string, 0, len(files))
		for _, e := range files {
			if strings.HasPrefix(e.Name(), fileBase) && e.Name() != symlinkBase {
				clogFiles = append(clogFiles, e.Name())
			}
		}
		for len(clogFiles) >= f.maxFiles {
			os.Remove(filepath.Join(fileDir, clogFiles[0]))
			clogFiles = clogFiles[1:]
		}
		// create new file
		fileName := fmt.Sprint(f.prefix, "_", time.Now().Format("20060102-150405.000000"))
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "clog error: %s\n", err)
			return
		}
		f.file = file
		f.bytes = 0
		// create or update symlink
		os.Remove(symlinkName)
		os.Symlink(filepath.Base(file.Name()), symlinkName)
	}
	f.file.Write([]byte(msg))
	f.bytes += len(msg)
}

// SetFileOutput sets the output to files with the given prefix and maximum byte length.
// Erase oldest file when maxFiles is reached.
func SetFileOutput(prefix string, maxLen, maxFiles int) {
	if maxLen < 1024 {
		maxLen = 1024
	}
	activeHandler.close()
	activeHandler = &fileHandler{
		prefix:   prefix,
		maxLen:   maxLen,
		maxFiles: maxFiles,
	}
}
