//go:generate gomobile bind -target=android -v -o ../build/collections.aar
// WARNING - These collections are not thread-safe

package natives

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Float32Collection interface {
	Clear()
	Index(rhs float32) (int, error)
	Insert(i int, n float32) error
	Append(n float32)
	Remove(i int) error
	Count() int
	At(i int) (float32, error)
	MustAt(i int) float32
	Iterator() Float32Iterator
}
type Float32Iterator interface {
	HasNext() bool
	Next() (float32, error)
}
type _Float32Collection struct {
	s []float32
}

// compile-time assurance that the struct matches the interface
var (
	_ Float32Collection = &_Float32Collection{}
	_ json.Marshaler    = &_Float32Collection{}
	_ json.Unmarshaler  = &_Float32Collection{}
)

func NewFloat32Collection() Float32Collection {
	return &_Float32Collection{}
}
func NewFloat32CollectionFrom(ss ...float32) Float32Collection {
	return &_Float32Collection{ss}
}
func (v *_Float32Collection) Clear() {
	v.s = v.s[:0]
}
func (v *_Float32Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func Float32Collection_MarshalJSONWith(this Float32Collection, marshal func(float32) ([]byte, error)) ([]byte, error) {
	col := make([]customFloat32Marshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, customFloat32Marshaler{x, marshal})
	}
	return json.Marshal(col)
}

type customFloat32Marshaler struct {
	v       float32
	marshal func(float32) ([]byte, error)
}

func (v customFloat32Marshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_Float32Collection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_Float32Collection) Index(rhs float32) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("Float32 not found in _Float32Collection")
}
func (v *_Float32Collection) Insert(i int, n float32) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_Float32Collection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, 0)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_Float32Collection) Append(n float32) {
	v.s = append(v.s, n)
}
func (v *_Float32Collection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_Float32Collection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = 0
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_Float32Collection) Count() int {
	return len(v.s)
}
func (v *_Float32Collection) At(i int) (float32, error) {
	if i < 0 || i >= len(v.s) {
		return 0, errors.Errorf("_Float32Collection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_Float32Collection) MustAt(i int) float32 {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_Float32Collection) Iterator() Float32Iterator {
	return NewFloat32Iterator(v)
}

type _Float32Iterator struct {
	next int
	s    []float32
}

func NewFloat32Iterator(col *_Float32Collection) Float32Iterator {
	return &_Float32Iterator{next: 0, s: col.s}
}
func (it *_Float32Iterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_Float32Iterator) Next() (float32, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return 0, errors.Errorf("_Float32Iterator has no more items")
}

type Float64Collection interface {
	Clear()
	Index(rhs float64) (int, error)
	Insert(i int, n float64) error
	Append(n float64)
	Remove(i int) error
	Count() int
	At(i int) (float64, error)
	MustAt(i int) float64
	Iterator() Float64Iterator
}
type Float64Iterator interface {
	HasNext() bool
	Next() (float64, error)
}
type _Float64Collection struct {
	s []float64
}

// compile-time assurance that the struct matches the interface
var (
	_ Float64Collection = &_Float64Collection{}
	_ json.Marshaler    = &_Float64Collection{}
	_ json.Unmarshaler  = &_Float64Collection{}
)

func NewFloat64Collection() Float64Collection {
	return &_Float64Collection{}
}
func NewFloat64CollectionFrom(ss ...float64) Float64Collection {
	return &_Float64Collection{ss}
}
func (v *_Float64Collection) Clear() {
	v.s = v.s[:0]
}
func (v *_Float64Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func Float64Collection_MarshalJSONWith(this Float64Collection, marshal func(float64) ([]byte, error)) ([]byte, error) {
	col := make([]customFloat64Marshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, customFloat64Marshaler{x, marshal})
	}
	return json.Marshal(col)
}

type customFloat64Marshaler struct {
	v       float64
	marshal func(float64) ([]byte, error)
}

func (v customFloat64Marshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_Float64Collection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_Float64Collection) Index(rhs float64) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("Float64 not found in _Float64Collection")
}
func (v *_Float64Collection) Insert(i int, n float64) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_Float64Collection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, 0)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_Float64Collection) Append(n float64) {
	v.s = append(v.s, n)
}
func (v *_Float64Collection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_Float64Collection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = 0
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_Float64Collection) Count() int {
	return len(v.s)
}
func (v *_Float64Collection) At(i int) (float64, error) {
	if i < 0 || i >= len(v.s) {
		return 0, errors.Errorf("_Float64Collection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_Float64Collection) MustAt(i int) float64 {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_Float64Collection) Iterator() Float64Iterator {
	return NewFloat64Iterator(v)
}

type _Float64Iterator struct {
	next int
	s    []float64
}

func NewFloat64Iterator(col *_Float64Collection) Float64Iterator {
	return &_Float64Iterator{next: 0, s: col.s}
}
func (it *_Float64Iterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_Float64Iterator) Next() (float64, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return 0, errors.Errorf("_Float64Iterator has no more items")
}

type IntCollection interface {
	Clear()
	Index(rhs int) (int, error)
	Insert(i int, n int) error
	Append(n int)
	Remove(i int) error
	Count() int
	At(i int) (int, error)
	MustAt(i int) int
	Iterator() IntIterator
}
type IntIterator interface {
	HasNext() bool
	Next() (int, error)
}
type _IntCollection struct {
	s []int
}

// compile-time assurance that the struct matches the interface
var (
	_ IntCollection    = &_IntCollection{}
	_ json.Marshaler   = &_IntCollection{}
	_ json.Unmarshaler = &_IntCollection{}
)

func NewIntCollection() IntCollection {
	return &_IntCollection{}
}
func NewIntCollectionFrom(ss ...int) IntCollection {
	return &_IntCollection{ss}
}
func (v *_IntCollection) Clear() {
	v.s = v.s[:0]
}
func (v *_IntCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func IntCollection_MarshalJSONWith(this IntCollection, marshal func(int) ([]byte, error)) ([]byte, error) {
	col := make([]customIntMarshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, customIntMarshaler{x, marshal})
	}
	return json.Marshal(col)
}

type customIntMarshaler struct {
	v       int
	marshal func(int) ([]byte, error)
}

func (v customIntMarshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_IntCollection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_IntCollection) Index(rhs int) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("Int not found in _IntCollection")
}
func (v *_IntCollection) Insert(i int, n int) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_IntCollection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, 0)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_IntCollection) Append(n int) {
	v.s = append(v.s, n)
}
func (v *_IntCollection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_IntCollection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = 0
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_IntCollection) Count() int {
	return len(v.s)
}
func (v *_IntCollection) At(i int) (int, error) {
	if i < 0 || i >= len(v.s) {
		return 0, errors.Errorf("_IntCollection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_IntCollection) MustAt(i int) int {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_IntCollection) Iterator() IntIterator {
	return NewIntIterator(v)
}

type _IntIterator struct {
	next int
	s    []int
}

func NewIntIterator(col *_IntCollection) IntIterator {
	return &_IntIterator{next: 0, s: col.s}
}
func (it *_IntIterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_IntIterator) Next() (int, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return 0, errors.Errorf("_IntIterator has no more items")
}

type Int32Collection interface {
	Clear()
	Index(rhs int32) (int, error)
	Insert(i int, n int32) error
	Append(n int32)
	Remove(i int) error
	Count() int
	At(i int) (int32, error)
	MustAt(i int) int32
	Iterator() Int32Iterator
}
type Int32Iterator interface {
	HasNext() bool
	Next() (int32, error)
}
type _Int32Collection struct {
	s []int32
}

// compile-time assurance that the struct matches the interface
var (
	_ Int32Collection  = &_Int32Collection{}
	_ json.Marshaler   = &_Int32Collection{}
	_ json.Unmarshaler = &_Int32Collection{}
)

func NewInt32Collection() Int32Collection {
	return &_Int32Collection{}
}
func NewInt32CollectionFrom(ss ...int32) Int32Collection {
	return &_Int32Collection{ss}
}
func (v *_Int32Collection) Clear() {
	v.s = v.s[:0]
}
func (v *_Int32Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func Int32Collection_MarshalJSONWith(this Int32Collection, marshal func(int32) ([]byte, error)) ([]byte, error) {
	col := make([]customInt32Marshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, customInt32Marshaler{x, marshal})
	}
	return json.Marshal(col)
}

type customInt32Marshaler struct {
	v       int32
	marshal func(int32) ([]byte, error)
}

func (v customInt32Marshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_Int32Collection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_Int32Collection) Index(rhs int32) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("Int32 not found in _Int32Collection")
}
func (v *_Int32Collection) Insert(i int, n int32) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_Int32Collection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, 0)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_Int32Collection) Append(n int32) {
	v.s = append(v.s, n)
}
func (v *_Int32Collection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_Int32Collection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = 0
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_Int32Collection) Count() int {
	return len(v.s)
}
func (v *_Int32Collection) At(i int) (int32, error) {
	if i < 0 || i >= len(v.s) {
		return 0, errors.Errorf("_Int32Collection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_Int32Collection) MustAt(i int) int32 {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_Int32Collection) Iterator() Int32Iterator {
	return NewInt32Iterator(v)
}

type _Int32Iterator struct {
	next int
	s    []int32
}

func NewInt32Iterator(col *_Int32Collection) Int32Iterator {
	return &_Int32Iterator{next: 0, s: col.s}
}
func (it *_Int32Iterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_Int32Iterator) Next() (int32, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return 0, errors.Errorf("_Int32Iterator has no more items")
}

type Int64Collection interface {
	Clear()
	Index(rhs int64) (int, error)
	Insert(i int, n int64) error
	Append(n int64)
	Remove(i int) error
	Count() int
	At(i int) (int64, error)
	MustAt(i int) int64
	Iterator() Int64Iterator
}
type Int64Iterator interface {
	HasNext() bool
	Next() (int64, error)
}
type _Int64Collection struct {
	s []int64
}

// compile-time assurance that the struct matches the interface
var (
	_ Int64Collection  = &_Int64Collection{}
	_ json.Marshaler   = &_Int64Collection{}
	_ json.Unmarshaler = &_Int64Collection{}
)

func NewInt64Collection() Int64Collection {
	return &_Int64Collection{}
}
func NewInt64CollectionFrom(ss ...int64) Int64Collection {
	return &_Int64Collection{ss}
}
func (v *_Int64Collection) Clear() {
	v.s = v.s[:0]
}
func (v *_Int64Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func Int64Collection_MarshalJSONWith(this Int64Collection, marshal func(int64) ([]byte, error)) ([]byte, error) {
	col := make([]customInt64Marshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, customInt64Marshaler{x, marshal})
	}
	return json.Marshal(col)
}

type customInt64Marshaler struct {
	v       int64
	marshal func(int64) ([]byte, error)
}

func (v customInt64Marshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_Int64Collection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_Int64Collection) Index(rhs int64) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("Int64 not found in _Int64Collection")
}
func (v *_Int64Collection) Insert(i int, n int64) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_Int64Collection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, 0)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_Int64Collection) Append(n int64) {
	v.s = append(v.s, n)
}
func (v *_Int64Collection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_Int64Collection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = 0
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_Int64Collection) Count() int {
	return len(v.s)
}
func (v *_Int64Collection) At(i int) (int64, error) {
	if i < 0 || i >= len(v.s) {
		return 0, errors.Errorf("_Int64Collection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_Int64Collection) MustAt(i int) int64 {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_Int64Collection) Iterator() Int64Iterator {
	return NewInt64Iterator(v)
}

type _Int64Iterator struct {
	next int
	s    []int64
}

func NewInt64Iterator(col *_Int64Collection) Int64Iterator {
	return &_Int64Iterator{next: 0, s: col.s}
}
func (it *_Int64Iterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_Int64Iterator) Next() (int64, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return 0, errors.Errorf("_Int64Iterator has no more items")
}

type StringCollection interface {
	Clear()
	Index(rhs string) (int, error)
	Insert(i int, n string) error
	Append(n string)
	Remove(i int) error
	Count() int
	At(i int) (string, error)
	MustAt(i int) string
	Iterator() StringIterator
}
type StringIterator interface {
	HasNext() bool
	Next() (string, error)
}
type _StringCollection struct {
	s []string
}

// compile-time assurance that the struct matches the interface
var (
	_ StringCollection = &_StringCollection{}
	_ json.Marshaler   = &_StringCollection{}
	_ json.Unmarshaler = &_StringCollection{}
)

func NewStringCollection() StringCollection {
	return &_StringCollection{}
}
func NewStringCollectionFrom(ss ...string) StringCollection {
	return &_StringCollection{ss}
}
func (v *_StringCollection) Clear() {
	v.s = v.s[:0]
}
func (v *_StringCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func StringCollection_MarshalJSONWith(this StringCollection, marshal func(string) ([]byte, error)) ([]byte, error) {
	col := make([]customStringMarshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, customStringMarshaler{x, marshal})
	}
	return json.Marshal(col)
}

type customStringMarshaler struct {
	v       string
	marshal func(string) ([]byte, error)
}

func (v customStringMarshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_StringCollection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_StringCollection) Index(rhs string) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("String not found in _StringCollection")
}
func (v *_StringCollection) Insert(i int, n string) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_StringCollection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, "")
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_StringCollection) Append(n string) {
	v.s = append(v.s, n)
}
func (v *_StringCollection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_StringCollection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = ""
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_StringCollection) Count() int {
	return len(v.s)
}
func (v *_StringCollection) At(i int) (string, error) {
	if i < 0 || i >= len(v.s) {
		return "", errors.Errorf("_StringCollection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_StringCollection) MustAt(i int) string {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_StringCollection) Iterator() StringIterator {
	return NewStringIterator(v)
}

type _StringIterator struct {
	next int
	s    []string
}

func NewStringIterator(col *_StringCollection) StringIterator {
	return &_StringIterator{next: 0, s: col.s}
}
func (it *_StringIterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_StringIterator) Next() (string, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return "", errors.Errorf("_StringIterator has no more items")
}
