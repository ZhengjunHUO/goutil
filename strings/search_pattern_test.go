package strings

import (
	"testing"
	"reflect"
)

func TestPatternFinder(t *testing.T) {
	pats := []string{"AABA", "ABABCABAB", "AAAAB", "ABABAC", "AAAA", "小猫咪"}
	txts := []string{"AABAACAADAABAABA", "ABABDABACDABABCABAB", "AAAAAAAAAAAAAAAAAB", "ABABABCABABABCABABABC", "AAAAABAAABA", "福福是一只小猫咪"}
	expect := [][]int{
		[]int{0,9,12},
		[]int{10},
		[]int{13},
		[]int{},
		[]int{0,1},
		[]int{5},
	}

	for i := range pats {
		rslt := NewPatternFinder(pats[i]).FindIn(txts[i])
		if !reflect.DeepEqual(rslt, expect[i]) {
			t.Errorf("FindIn returns %v, expect %v", rslt, expect[i])
		}
	}
}
