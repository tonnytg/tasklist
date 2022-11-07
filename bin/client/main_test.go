package main_test

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestMainFlags(t *testing.T) {

	expected := []byte(`Client Side
pass help as argument to get help.
./taskclient help`)

	cmd := exec.Command("go", "run", "main.go")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	for i, v := range expected {
		if v != out.Bytes()[i] {
			t.Errorf("expected: %q, got: %q", string(expected), out.String())
			break
		}
	}
}
