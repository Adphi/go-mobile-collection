package templates

import "text/template"

var (
	_ = template.Must(Generator.New("embed").Parse(`{{ define "embed" }}
{{- /*gotype: github.com/partitio/go-mobile-collection.GeneratedType*/ -}}
type {{.Name}}Collection interface {
	Clear()
	Index(rhs {{.Ptr}}{{.Name}}) (int, error)
	Insert(i int, n {{.Ptr}}{{.Name}}) error
	Append(n {{.Ptr}}{{.Name}})
	Remove(i int) error
	Count() int
	At(i int) ({{.Ptr}}{{.Name}}, error)
	MustAt(i int) {{.Ptr}}{{.Name}}
	Iterator() {{.Name}}Iterator
}
type {{.Name}}Iterator interface {
	HasNext() bool
	Next() ({{.Ptr}}{{.Name}}, error)
}
type _{{.Name}}Collection struct {
	s []{{.Ptr}}{{.Name}}
}
// compile-time assurance that the struct matches the interface
var (
	_ {{.Name}}Collection = &_{{.Name}}Collection{}
	_ json.Marshaler = &_{{.Name}}Collection{}
	_ json.Unmarshaler = &_{{.Name}}Collection{}
)
func New{{.Name}}Collection() {{.Name}}Collection {
	return &_{{.Name}}Collection{}
}
func New{{.Name}}CollectionFrom(ss ...{{.Ptr}}{{.Name}}) {{.Name}}Collection {
	return &_{{.Name}}Collection{ss}
}
func (v *_{{.Name}}Collection) Clear() {
	v.s = v.s[:0]
}
func (v *_{{.Name}}Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}
func {{.Name}}Collection_MarshalJSONWith(this {{.Name}}Collection, marshal func({{.Ptr}}{{.Name}}) ([]byte, error)) ([]byte, error) {
	col := make([]custom{{.Name}}Marshaler, 0, this.Count())
	next := this.Iterator().Next
	for x, err := next(); err == nil; x, err = next() {
		col = append(col, custom{{.Name}}Marshaler{x, marshal})
	}
	return json.Marshal(col)
}
type custom{{.Name}}Marshaler struct {
	v       {{.Ptr}}{{.Name}}
	marshal func({{.Ptr}}{{.Name}}) ([]byte, error)
}
func (v custom{{.Name}}Marshaler) MarshalJSON() ([]byte, error) {
	return v.marshal(v.v)
}
func (v *_{{.Name}}Collection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}
func (v *_{{.Name}}Collection) Index(rhs {{.Ptr}}{{.Name}}) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("{{.Name}} not found in _{{.Name}}Collection")
}
func (v *_{{.Name}}Collection) Insert(i int, n {{.Ptr}}{{.Name}}) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_{{.Name}}Collection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, nil)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}
func (v *_{{.Name}}Collection) Append(n {{.Ptr}}{{.Name}}) {
	v.s = append(v.s, n)
}
func (v *_{{.Name}}Collection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_{{.Name}}Collection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = nil
	v.s = v.s[:len(v.s)-1]
	return nil
}
func (v *_{{.Name}}Collection) Count() int {
	return len(v.s)
}
func (v *_{{.Name}}Collection) At(i int) ({{.Ptr}}{{.Name}}, error) {
	if i < 0 || i >= len(v.s) {
		return nil, errors.Errorf("_{{.Name}}Collection invalid index %d\n", i)
	}
	return v.s[i], nil
}
func (v *_{{.Name}}Collection) MustAt(i int) {{.Ptr}}{{.Name}} {
	if x, err := v.At(i); err != nil {
		panic(err)
	} else {
		return x
	}
}
func (v *_{{.Name}}Collection) Iterator() {{.Name}}Iterator {
	return New{{.Name}}Iterator(v)
}
type _{{.Name}}Iterator struct {
	next int
	s		[]{{.Ptr}}{{.Name}}
}
func New{{.Name}}Iterator(col *_{{.Name}}Collection) {{.Name}}Iterator {
	return &_{{.Name}}Iterator{next: 0, s: col.s}
}
func (it *_{{.Name}}Iterator) HasNext() bool {
	return it.next < len(it.s)
}
func (it *_{{.Name}}Iterator) Next() ({{.Ptr}}{{.Name}}, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}
	return nil, errors.Errorf("_{{.Name}}Iterator has no more items")
}
{{ end }}`))
)
