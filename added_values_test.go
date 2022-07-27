package goja

import (
	"testing"
)

func TestNumberEquality(t *testing.T) {
	vm := New()

	res, err := vm.RunString(`var a = new Number('5') 
	a`)
	if err != nil {
		t.Fatal(err)
	}
	if !valueInt(5).Equals(res) {
		t.Fatal("values are not equal")
	}
}

func TestInt64SameAsFloat(t *testing.T) {
	if !valueInt64(5).SameAs(valueFloat(5.0)) {
		t.Fatal("values are not equal")
	}
}
