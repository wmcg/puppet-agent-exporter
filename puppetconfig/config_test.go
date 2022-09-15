package puppetconfig

import (
	"testing"

	"gopkg.in/ini.v1"
)

func TestLoadReport(t *testing.T) {
	config, _ := ini.Load("puppet.conf")

	expectedServer := "puppet-server.example.com"
	actualServer := GetSetting(config, "server")
	if actualServer != expectedServer {
		t.Fatalf("'%+v' != '%+v'", actualServer, expectedServer)
	}

	// Test that 'agent' overides 'main'
	expectedEnv := "overider"
	actualEnv := GetSetting(config, "environment")
	if actualEnv != expectedEnv {
		t.Fatalf("'%+v' != '%+v'", actualEnv, expectedEnv)
	}
}
