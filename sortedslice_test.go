package sortedslice

import (
    "testing"
)

func Test001(t *testing.T) {
    expected := "Hello sortedslice v1.0.0\n"
    actual := Hello()
    if actual != expected {
        t.Errorf("expected %q, got %q", expected, actual)
    }
}
