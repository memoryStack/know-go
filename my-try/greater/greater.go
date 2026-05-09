package greater

// func (m MyInt) Greater(v MyInt) bool {

func IsGreater[T interface {Greater(T) bool}](x, y T) bool {
	return x.Greater(y)
}
