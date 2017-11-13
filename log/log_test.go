package log

import (
	"testing"
)

//const checkMark = "\u2713"
//const ballotX = "\u2717"

func TestLog(t *testing.T) {
	t.Log("Test InitLog start...")
	InitLog("testLog", 1, "debug", true)

	Log.Info("Info yes!")
	Log.Error("Error yes!")
	Log.Debug("Debug yes!")
	Log.Warning("Warning yes!")
	Log.Critical("Critical yes!")

}
