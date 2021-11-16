package nestedtext_test

import (
	"testing"

	"github.com/knadh/koanf/parsers/nestedtext"
)

func TestParser(t *testing.T) {
	// test input in NestedText format
	ntsource := `ports:
  - 6483
  - 8020
  - 9332
timeout: 20
`
	// Test decoder
	nt := nestedtext.Parser()
	c, err := nt.Unmarshal([]byte(ntsource))
	if err != nil {
		t.Fatal("Unmarshal of NestedText input failed")
	}
	t.Logf("config-tree: %#v", c)
	timeout := c["timeout"]
	if timeout != "20" {
		t.Errorf("expected timeout-parameter to be 20, is %q", timeout)
	}

	// test encoder
	out, err := nt.Marshal(c)
	if err != nil {
		t.Fatal("Marshal of config to NestedText failed")
	}
	t.Logf("config-text: %q", string(out))
	if string(out) != ntsource {
		t.Errorf("expected output of Marshal(…) to equal input to Unmarshal(…); didn't")
	}
}
