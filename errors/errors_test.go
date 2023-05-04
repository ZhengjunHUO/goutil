package errors

import (
	"fmt"
	"testing"
)

func TestErrBuf(t *testing.T) {
	eb := NewErrBuf()
	if eb.ToError() != nil {
		t.Errorf("ToError() returns %v, expect nil", eb.ToError())
	}

	err1 := "Error checking foo"
	err2 := "Error checking bar"
	eb.Append(fmt.Errorf("%s", err1))
	if eb.ToString() != err1+"\n" {
		t.Errorf("ToString() returns %s, expect %s", eb.ToString(), err1)
	}

	eb.AppendString(err2)
	if eb.ToError().Error() != fmt.Sprintf("%v\n%v\n", err1, err2) {
		t.Errorf("ToError() returns %s, expect %s", eb.ToError(), fmt.Errorf("%v\n%v\n", err1, err2))
	}
}
