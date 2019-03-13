// //go:generate gomobile bind -target=android -v -o ../build/collections.aar

package natives

// @collection-wrapper
type Str string

func NewStr(s Str) Str {
	ss := Str(s)
	return ss
}

func (s *Str)ToStd() Str {
	return Str(*s)
}

// @collection-wrapper
type Int int

func NewInt(i int) Int {
	ii := Int(i)
	return ii
}

func (i *Int)ToStd() int {
	return int(*i)
}

// @collection-wrapper
type Int32 int32

func NewInt32(i int32) Int32 {
	ii := Int32(i)
	return ii
}

func (i *Int32)ToStd() int32 {
	return int32(*i)
}

// @collection-wrapper
type Int64 int64

func NewInt64(i int64) Int64 {
	ii := Int64(i)
	return ii
}

func (i *Int64)ToStd() int64 {
	return int64(*i)
}

// @collection-wrapper
type Float32 float32

func NewFloat32(i float64) Float32 {
	ii := Float32(i)
	return ii
}

func (i *Float32)ToStd() float32 {
	return float32(*i)
}

// @collection-wrapper
type Float64 float64

func NewFloat64(i float64) Float64 {
	ii := Float64(i)
	return ii
}

func (i *Float64)ToStd() float64 {
	return float64(*i)
}
