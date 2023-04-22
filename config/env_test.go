package config

import (
	"testing"

)

func TestGetString(t *testing.T) {
	if GetString(DB_DRIVER) != "mysql" {
		t.Error("Getstring Error:" + GetString(DB_DRIVER))
	}
}
