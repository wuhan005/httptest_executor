package main

import (
	"github.com/thanhpk/randstr"
	"github.com/wuhan005/httptest/exec"
	"io/ioutil"
	"strings"
	"testing"
)

func allMatch(pat, str string) (bool, error) {
	return true, nil
}

func allTests(t *testing.T) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		t.Fatal("GetFile failed:", err)
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".qtf") {
			b, err := ioutil.ReadFile(file.Name())
			if err != nil {
				t.Fatal("ReadFile failed:", err)
				return
			}
			randomStr := randstr.String(16)
			b = []byte(strings.Replace(string(b), "[RANDOM_STR]", "\""+randomStr+"\"", -1))
			err = exec.ExecCases(t, string(b))
			if err != nil {
				t.Fatal("ExecCases failed:", err)
			}
		}
	}
}

func main() {
	tests := []testing.InternalTest{
		{"main", allTests},
	}
	testing.Main(allMatch, tests, nil, nil)
}
