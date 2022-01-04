package src

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "go-template-test" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
