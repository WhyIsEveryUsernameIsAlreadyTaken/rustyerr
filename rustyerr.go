package rustyerr

type okVariant struct {
    value any
}

type errVariant struct {
    errval error
}

type variant interface{}

type Result struct {
    variant
}

func (res Result) Unwrap() any {
    switch res.variant.(type) {
    case okVariant:
        return res.variant.(okVariant).value
    case errVariant:
        panic(res.variant.(errVariant).errval)
    default:
        panic("unkown result variant")
}
}

func Ok(obj any) Result {
    return Result{
        variant: okVariant{
            value: obj,
        },
    }
}

func Err(err error) Result {
    return Result{
        variant: errVariant{
            errval: err,
        },
    }
}
