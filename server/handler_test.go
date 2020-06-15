package server
import (
    "testing"
)


func TestAdd(t *testing.T) {
    var (
        in = 7
        in2 = 2
        out = 10
    )
    actual := Add(in, in2)
    if actual != out {
    t.Errorf("Add(%d, %d) = %d;expected %d", in, in2, actual, out)
    }
}
