package main

import "github.com/chmike/clog"

func main() {
	clog.SetLevel(clog.Debug2Level)
	clog.Infoln("starting...", 10)
	clog.Println("normal progress")
	clog.Debug("this is a debug message")
	clog.Debug1("this is a debug1 message")
	clog.Debug2("this is a debug2 message")
	clog.Warning("this is a warning")
	clog.Error("this is an error")
	clog.Println("multiline\nmessage")

	c := clog.New("clog_test", clog.Debug2Level)
	c.Infoln("starting...", 10)
	c.Println("normal progress")
	c.Debug("this is a debug message")
	c.Debug1("this is a debug1 message")
	c.Debug2("this is a debug2 message")
	c.Println("multiline\nmessage")
	c.Warning("this is a warning")
	c.Error("this is an error")
	c.Fatal("this is a fatal error")

	c.Info("unprinted message")
}
