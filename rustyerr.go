package rustyerr

type okVariant[T any] struct {
    value T
}

type errVariant struct {
    errval error
}

type variant interface{}

type Result[T any] struct {
    variant
}

func (res Result[T]) Unwrap() T {
    switch res.variant.(type) {
    case okVariant[T]:
        return res.variant.(okVariant[T]).value
    case errVariant:
        panic(res.variant.(errVariant).errval)
    default:
        panic("unkown result variant")
}
}

func Ok[T any](obj T) Result[T] {
    return Result[T]{
        variant: okVariant[T]{
            value: obj,
        },
    }
}

func Err(err error) Result[error] {
    return Result[error]{
        variant: errVariant{
            errval: err,
        },
    }
}

