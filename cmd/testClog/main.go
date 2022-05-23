package main

import (
	"encoding/json"

	"github.com/chmike/clog"
)

type test struct {
	LogLevel clog.LevelType `json:"logLevel"`
}

var v = test{
	LogLevel: clog.Debug2Level,
}

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

	clog.Info("test json encoding")
	b, _ := json.MarshalIndent(v, "", "  ")
	clog.Println(string(b))

	var v2 test
	json.Unmarshal(b, &v2)
	if v != v2 {
		clog.Fatal("mismatch")
	}

	c := clog.New("clog_test", clog.Debug2Level)
	c.Infoln("starting...", 10)
	c.Println("normal progress")
	c.Debug("this is a debug message")
	c.Debug1("this is a debug1 message")
	c.Debug2("this is a debug2 message")
	c.Infoln("multiline\nmessage")
	c.Warning("this is a warning")
	c.Error("this is an error")
	c.New("test").Println("create a sub tagged logger")
	c.Fatal("this is a fatal error")

	c.Info("unprinted message")
}
