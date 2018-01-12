// Copyright 2014 Brett Slatkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	generatedTemplate = template.Must(template.New("render").Parse(`// generated by collection-wrapper -- DO NOT EDIT
// WARNING - These collections are not thread-safe

package {{.Package}}

import (
  "fmt"
  "encoding/json"
)

{{range .Types}}
type {{.Name}}Collection struct {
  s []*{{.Name}}
}

func New{{.Name}}Collection() *{{.Name}}Collection {
  return &{{.Name}}Collection{}
}

func (v *{{.Name}}Collection) Clear() {
  v.s = v.s[:0]
}

func (v *{{.Name}}Collection) MarshalJSON() ([]byte, error) {
  return json.Marshal(v.s)
}

func (v *{{.Name}}Collection) UnmarshalJSON(data []byte) error {
  return json.Unmarshal(data, &v.s)
}

func (v *{{.Name}}Collection) Index(rhs *{{.Name}}) (int, error) {
  for i, lhs := range v.s {
    if lhs == rhs {
      return i, nil
    }
  }
  return -1, fmt.Errorf("{{.Name}} not found in {{.Name}}Collection")
}

func (v *{{.Name}}Collection) Insert(i int, n *{{.Name}}) error {
  if i < 0 || i > len(v.s) {
    return fmt.Errorf("{{.Name}}Collection error trying to insert at invalid index %d\n", i)
  }
  v.s = append(v.s, nil)
  copy(v.s[i+1:], v.s[i:])
  v.s[i] = n
  return nil
}

func (v *{{.Name}}Collection) Append(n *{{.Name}}) {
  v.s = append(v.s, n)
}

func (v *{{.Name}}Collection) Remove(i int) error {
  if i < 0 || i >= len(v.s) {
    return fmt.Errorf("{{.Name}}Collection error trying to remove invalid index %d\n", i)
  }
  copy(v.s[i:], v.s[i+1:])
  v.s[len(v.s)-1] = nil
  v.s = v.s[:len(v.s)-1]
  return nil
}

func (v *{{.Name}}Collection) Count() int {
  return len(v.s)
}

func (v *{{.Name}}Collection) At(i int) (*{{.Name}}, error) {
  if i < 0 || i >= len(v.s) {
    return nil, fmt.Errorf("{{.Name}}Collection invalid index %d\n", i)
  }
  return v.s[i], nil
}

{{end}}`))
)

type GeneratedType struct {
	Name string
}

func getRenderedPath(inputPath string) (string, error) {
	if !strings.HasSuffix(inputPath, ".go") {
		return "", fmt.Errorf("Input path %s doesn't have .go extension", inputPath)
	}
	trimmed := strings.TrimSuffix(inputPath, ".go")
	dir, file := filepath.Split(trimmed)
	return filepath.Join(dir, fmt.Sprintf("%s_collection.go", file)), nil
}

type generateTemplateData struct {
	Package string
	Types   []GeneratedType
}

func render(w io.Writer, packageName string, types []GeneratedType) error {
	return generatedTemplate.Execute(w, generateTemplateData{packageName, types})
}
