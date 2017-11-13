package conf

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestInitConfig(t *testing.T) {

	t.Log("Test InitConfig start")

	err := InitConfig("conf.toml")
	if err != nil {
		t.Fatal("Should be able to get configure file parsed.", ballotX, err)
	}

	t.Log("Should be able to get configure file parsed.", checkMark)

}
