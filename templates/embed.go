package templates

import "text/template"

var (
	_ = template.Must(Generator.New("embed").Parse(`{{ define "embed" }}
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

type _{{.Name | Title}}Collection struct {
	s []{{.Ptr}}{{.Name}}
}

// compile-time assurance that the struct matches the interface
var (
	_ {{.Name | Title}}Collection = &_{{.Name | Title}}Collection{}
	_ json.Marshaler = &_{{.Name | Title}}Collection{}
	_ json.Unmarshaler = &_{{.Name | Title}}Collection{}
)

func New{{.Name | Title}}Collection() {{.Name | Title}}Collection {
	return &_{{.Name | Title}}Collection{}
}

func New{{.Name | Title}}CollectionFrom(ss ...{{.Ptr}}{{.Name}}) {{.Name | Title}}Collection {
	return &_{{.Name | Title}}Collection{ss}
}

func (v *_{{.Name | Title}}Collection) Clear() {
	v.s = v.s[:0]
}

func (v *_{{.Name | Title}}Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
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
	return json.Unmarshal(data, &v.s)
}

func (v *_{{.Name | Title}}Collection) Index(rhs {{.Ptr}}{{.Name}}) (int, error) {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i, nil
		}
	}
	return -1, errors.Errorf("{{.Name}} not found in _{{.Name | Title}}Collection")
}

func (v *_{{.Name | Title}}Collection) Insert(i int, n {{.Ptr}}{{.Name}}) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_{{.Name | Title}}Collection error trying to insert at invalid index %d\n", i)
	}
	v.s = append(v.s, {{.DefaultValue}})
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
	return nil
}

func (v *_{{.Name | Title}}Collection) Append(n {{.Ptr}}{{.Name}}) {
	v.s = append(v.s, n)
}

func (v *_{{.Name | Title}}Collection) Remove(i int) error {
	if i < 0 || i >= len(v.s) {
		return errors.Errorf("_{{.Name | Title}}Collection error trying to remove invalid index %d\n", i)
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = {{.DefaultValue}}
	v.s = v.s[:len(v.s)-1]
	return nil
}

func (v *_{{.Name | Title}}Collection) Count() int {
	return len(v.s)
}

func (v *_{{.Name | Title}}Collection) Get(i int) ({{.Ptr}}{{.Name}}, error) {
	if i < 0 || i >= len(v.s) {
		return {{.DefaultValue}}, errors.Errorf("_{{.Name | Title}}Collection invalid index %d\n", i)
	}
	return v.s[i], nil
}

func (v *_{{.Name | Title}}Collection) Set(i int, n {{.Ptr}}{{.Name}}) error {
	if i < 0 || i > len(v.s) {
		return errors.Errorf("_{{.Name | Title}}Collection error trying to insert at invalid index %d\n", i)
	}
	v.s[i] = n
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
	return &_{{.Name | Title}}Iterator{next: 0, s: col.s}
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
{{ end }}`))
)
