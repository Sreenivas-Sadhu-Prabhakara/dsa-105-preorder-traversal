package main

import (
	"reflect"
	"testing"
)

func TestPreorder(t *testing.T) {
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	want := []int{4, 2, 1, 3, 6, 5, 7}
	if got := Preorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("Preorder = %v, want %v", got, want)
	}
	if got := Preorder(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Preorder(nil) = %v", got)
	}
	chain := tn(1, nil, tn(2, tn(3, nil, nil), nil))
	if got := Preorder(chain); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Preorder(chain) = %v, want [1 2 3]", got)
	}
}
