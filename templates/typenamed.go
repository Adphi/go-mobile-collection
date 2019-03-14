package templates

import "text/template"

var (
	_ = template.Must(Generator.New("typenamed").Parse(`{{ define "typeNamed" }}
{{- /*gotype: github.com/partitio/go-mobile-collection.GeneratedType*/ -}}
type {{.Name}}Collection interface {
	Clear()
	Index(rhs {{.Ptr}}{{.Name}}) (int, error)
	Insert(i int, n {{.Ptr}}{{.Name}}) error
	Append(n {{.Ptr}}{{.Name}})
	Remove(i int) error
	Count() int
	Get(i int) ({{.Ptr}}{{.Name}}, error)
	MustGet(i int) {{.Ptr}}{{.Name}}
	Iterator() {{.Name}}Iterator
}

type {{.Name}}Iterator interface {
	HasNext() bool
	Next() ({{.Ptr}}{{.Name}}, error)
}

type _{{.Name}}Collection []{{.Ptr}}{{.Name}}

// compile-time assurance that the struct matches the interface
var (
	_ {{.Name}}Collection = &_{{.Name}}Collection{}
	_ json.Marshaler = &_{{.Name}}Collection{}
	_ json.Unmarshaler = &_{{.Name}}Collection{}
)

func New{{.Name}}Collection() {{.Name}}Collection {
	var ss []{{.Ptr}}{{.Name}}
	c := _{{.Name}}Collection(ss)
	return &c
}

func New{{.Name}}CollectionFrom(ss ...{{.Ptr}}{{.Name}}) {{.Name}}Collection {
	if ss == nil {
		ss = []{{.Ptr}}{{.Name}}{}
	}
	c := _{{.Name}}Collection(ss)
	return &c
}

func (v *_{{.Name}}Collection) Clear() {
	s := *v
	*v = s[:0]
}

func (v *_{{.Name}}Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v)
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
	return json.Unmarshal(data, &v)
}

func (v *_{{.Name}}Collection) Index(rhs {{.Ptr}}{{.Name}}) (int, error) {
	for i, lhs := range *v {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("{{.Name}} not found in _{{.Name}}Collection")
}

func (v *_{{.Name}}Collection) Insert(i int, n {{.Ptr}}{{.Name}}) error {
	if i < 0 || i > len(*v) {
		return errors.Errorf("_{{.Name}}Collection error trying to insert at invalid index %d\n", i)
	}
	s := *v
	s = append(s, nil)
	copy(s[i+1:], s[i:])
	s[i] = n
	*v = s
	return nil
}

func (v *_{{.Name}}Collection) Append(n {{.Ptr}}{{.Name}}) {
	*v = append(*v, n)
}

func (v *_{{.Name}}Collection) Remove(i int) error {
	if i < 0 || i >= len(*v) {
		return errors.Errorf("_{{.Name}}Collection error trying to remove invalid index %d\n", i)
	}
	s := *v
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1]
	*v = s
	return nil
}

func (v *_{{.Name}}Collection) Count() int {
	return len(*v)
}

func (v *_{{.Name}}Collection) Get(i int) ({{.Ptr}}{{.Name}}, error) {
	if i < 0 || i >= len(*v) {
		return nil, errors.Errorf("_{{.Name}}Collection invalid index %d\n", i)
	}
	s := *v
	return s[i], nil
}

func (v *_{{.Name}}Collection) MustGet(i int) {{.Ptr}}{{.Name}} {
	if x, err := v.Get(i); err != nil {
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
	return &_{{.Name}}Iterator{next: 0, s: *col}
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

{{ end }}
`))
)