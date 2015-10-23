package owl

import (
	"testing"
)

func TestRegToPost(t *testing.T) {
	r, err := RegToPost("a|b")
	if err != nil {
		t.Errorf("RegToPost(`a|b`) failed. error:%s", err.Error())
	}
	if r != "ab|" {
		t.Errorf("RegToPost(`a|b`) failed. Got %s", r)
	}
	r, err = RegToPost("(((ab)c)d)e|fg")
	if err != nil {
		t.Errorf("RegToPost(`(((ab)c)d)e|fg`) failed. error:%s", err.Error())
	}
	if r != "ab.c.d.e.fg.|" {
		t.Errorf("RegToPost(`(((ab)c)d)e|fg`) failed. Got %s", r)
	}
	r, err = RegToPost("((abc*)(a+))c?")
	if err != nil {
		t.Errorf("RegToPost(`((abc*)(a+))c?`) failed. error:%s", err.Error())
	}
	if r != "ab.c*.a+.c?." {
		t.Errorf("RegToPost(`((abc*)(a+))c?`) failed. Got %s", r)
	}
}
