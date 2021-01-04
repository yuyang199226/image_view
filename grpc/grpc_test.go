package grpc

import (
    "fmt"
    "testing"
    "imageview/grpc/mocks"
)

func TestGodviewMock {
    mockgod := &mocks.Godview{}
    mockResFn := func() string {
        return "mock hello"
    }
    mockgod.On("Localize", )

}
