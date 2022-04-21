package main

import "github.com/chmike/clog"

func main() {
	clog.SetFileOutput(clog.Options{
		Prefix:   "/tmp/testClog/test",
		MaxLen:   1024,
		MaxFiles: 3,
		Level:    clog.Debug2Level,
	})

	for i := 0; i < 1000; i++ {
		if i%20 == 0 {
			clog.Info("test file clog")
		} else {
			clog.Print("test file clog")
		}
	}
}
