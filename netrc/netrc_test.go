// Copyright © 2010 Fazlul Shahriar <fshahriar@gmail.com>.
// See LICENSE file for license details.

package netrc_test

import (
	. "go-netrc.googlecode.com/hg/netrc"
	"testing"
)

func eqMach(a *Machine, b *Machine) bool {
	return a.Name == b.Name &&
		a.Login == b.Login &&
		a.Password == b.Password &&
		a.Account == b.Account
}

func TestParse(t *testing.T) {
	mach, mac, err := ParseFile("example.netrc")
	if err != nil {
		t.Fatal(err)
	}

	expectedMach := []*Machine{
		&Machine{"mail.google.com", "joe@gmail.com", "somethingSecret", "gmail"},
		&Machine{"ray", "demo", "mypassword", ""},
		&Machine{"", "anonymous", "joe@example.com", ""},
	}
	for i, e := range expectedMach {
		if !eqMach(e, mach[i]) {
			t.Errorf("bad machine; expected %v, got %v\n", e, mach[i])
		}
	}

	expectedMac := Macros{
		"allput": "put src/*",
	}
	for k, v := range expectedMac {
		if v != mac[k] {
			t.Errorf("bad macro for %s; expected %s, got %s\n", k, v, mac[k])
		}
	}
}
