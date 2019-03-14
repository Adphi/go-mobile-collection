package templates

import "text/template"

var (
	_ = template.Must(Generator.New("typenamed").Parse(`{{ define "typeNamed" }}
type {{.Name | Title}}Collection interface {
	Clear()
	Index(rhs {{.Ptr}}{{.Name}}) (int, error)
	Insert(i int, n {{.Ptr}}{{.Name}}) error
	Append(n {{.Ptr}}{{.Name}})
	Remove(i int) error
	Count() int
	Get(i int) ({{.Ptr}}{{.Name}}, error)
	Set(i int, n {{.Ptr}}{{.Name}}) error
	MustGet(i int) {{.Ptr}}{{.Name}}
	Iterator() {{.Name | Title}}Iterator
}

type {{.Name | Title}}Iterator interface {
	HasNext() bool
	Next() ({{.Ptr}}{{.Name}}, error)
}

type _{{.Name | Title}}Collection []{{.Ptr}}{{.Name}}

// compile-time assurance that the struct matches the interface
var (
	_ {{.Name | Title}}Collection = &_{{.Name | Title}}Collection{}
	_ json.Marshaler = &_{{.Name | Title}}Collection{}
	_ json.Unmarshaler = &_{{.Name | Title}}Collection{}
)

func New{{.Name | Title}}Collection() {{.Name | Title}}Collection {
	var ss []{{.Ptr}}{{.Name}}
	c := _{{.Name | Title}}Collection(ss)
	return &c
}

func New{{.Name | Title}}CollectionFrom(ss ...{{.Ptr}}{{.Name}}) {{.Name | Title}}Collection {
	if ss == nil {
		ss = []{{.Ptr}}{{.Name}}{}
	}
	c := _{{.Name | Title}}Collection(ss)
	return &c
}

func (v *_{{.Name | Title}}Collection) Clear() {
	s := *v
	*v = s[:0]
}

func (v *_{{.Name | Title}}Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v)
}

func {{.Name | Title}}Collection_MarshalJSONWith(this {{.Name | Title}}Collection, marshal func({{.Ptr}}{{.Name}}) ([]byte, error)) ([]byte, error) {
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

func (v *_{{.Name | Title}}Collection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v)
}

func (v *_{{.Name | Title}}Collection) Index(rhs {{.Ptr}}{{.Name}}) (int, error) {
	for i, lhs := range *v {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("{{.Name}} not found in _{{.Name | Title}}Collection")
}

func (v *_{{.Name | Title}}Collection) Insert(i int, n {{.Ptr}}{{.Name}}) error {
	if i < 0 || i > len(*v) {
		return errors.Errorf("_{{.Name | Title}}Collection error trying to insert at invalid index %d\n", i)
	}
	s := *v
	s = append(s, {{.DefaultValue}})
	copy(s[i+1:], s[i:])
	s[i] = n
	*v = s
	return nil
}

func (v *_{{.Name | Title}}Collection) Append(n {{.Ptr}}{{.Name}}) {
	*v = append(*v, n)
}

func (v *_{{.Name | Title}}Collection) Remove(i int) error {
	if i < 0 || i >= len(*v) {
		return errors.Errorf("_{{.Name | Title}}Collection error trying to remove invalid index %d\n", i)
	}
	s := *v
	copy(s[i:], s[i+1:])
	s[len(s)-1] = {{.DefaultValue}}
	s = s[:len(s)-1]
	*v = s
	return nil
}

func (v *_{{.Name | Title}}Collection) Count() int {
	return len(*v)
}

func (v *_{{.Name | Title}}Collection) Get(i int) ({{.Ptr}}{{.Name}}, error) {
	if i < 0 || i >= len(*v) {
		return {{.DefaultValue}}, errors.Errorf("_{{.Name | Title}}Collection invalid index %d\n", i)
	}
	s := *v
	return s[i], nil
}

func (v *_{{.Name | Title}}Collection) Set(i int, n {{.Ptr}}{{.Name}}) error {
	if i < 0 || i > len(*v) {
		return errors.Errorf("_{{.Name | Title}}Collection error trying to insert at invalid index %d\n", i)
	}
	s := *v
	s[i] = n
	*v = s
	return nil
}

func (v *_{{.Name | Title}}Collection) MustGet(i int) {{.Ptr}}{{.Name}} {
	if x, err := v.Get(i); err != nil {
		panic(err)
	} else {
		return x
	}
}

func (v *_{{.Name | Title}}Collection) Iterator() {{.Name | Title}}Iterator {
	return New{{.Name | Title}}Iterator(v)
}

type _{{.Name | Title}}Iterator struct {
	next int
	s		[]{{.Ptr}}{{.Name}}
}

func New{{.Name | Title}}Iterator(col *_{{.Name | Title}}Collection) {{.Name | Title}}Iterator {
	return &_{{.Name | Title}}Iterator{next: 0, s: *col}
}

func (it *_{{.Name | Title}}Iterator) HasNext() bool {
	return it.next < len(it.s)
}

func (it *_{{.Name | Title}}Iterator) Next() ({{.Ptr}}{{.Name}}, error) {
	if it.HasNext() {
		val := it.s[it.next]
		it.next = it.next + 1
		return val, nil
	}

	return {{.DefaultValue}}, errors.Errorf("_{{.Name | Title}}Iterator has no more items")
}
{{ end }}
`))
)
