package config

import (
	"os"
	"testing"
)

func TestInitEnv(t *testing.T) {
	os.Setenv("APPLICATION_PORT", "4444")

	InitEnv()

	if Manager.ApplicationPort != 4444 {
		t.Errorf("Environment variables are not loading properly. Expected: %d , Actual: %d", 4444, Manager.ApplicationPort)
	}
}
