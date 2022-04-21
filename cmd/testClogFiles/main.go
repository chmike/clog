package main

import "github.com/chmike/clog"

func main() {
	clog.SetFileOutput("/tmp/testClog/test", 1024, 3)

	for i := 0; i < 1000; i++ {
		if i%20 == 0 {
			clog.Info("test file clog")
		} else {
			clog.Print("test file clog")
		}
	}
}
