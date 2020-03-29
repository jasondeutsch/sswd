package main

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	tt := []struct {
		cmd string
	}{
		{cmd: "ls -la"},
		{cmd: "pwd"},
		{cmd: "mkdir test"},
		{cmd: "rm -rf test"},
	}

	for _, tc := range tt {
		b, err := run(tc.cmd)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(b))
	}
}
