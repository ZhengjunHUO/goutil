package errors

import (
	"strings"
	"fmt"
)

// ErrBuf accumulates and stores the errors in an internal buffer
type ErrBuf struct {
        builder *strings.Builder
}

// NewErrBuf returns an initialized ErrBuf instance
func NewErrBuf() *ErrBuf {
        return &ErrBuf{
                builder: new(strings.Builder),
        }
}

func (b *ErrBuf) Append(e error) {
        if e != nil {
                b.builder.WriteString(e.Error() + "\n")
        }
}

func (b *ErrBuf) AppendString(s string) {
        b.builder.WriteString(s + "\n")
}

func (b *ErrBuf) ToError() error {
        if b.builder.Len() == 0 {
                return nil
        }

        return fmt.Errorf("%v", b.builder.String())
}

func (b *ErrBuf) ToString() string {
        return b.builder.String()
}
