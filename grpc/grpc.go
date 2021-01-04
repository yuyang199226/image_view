package grpc

type Godview interface {
    Localize() string

}

type GodviewImpl struct {
    Godview

}

func (c GodviewImpl) Localize() string {
    return "hello world"
}
