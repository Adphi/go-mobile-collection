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
	"github.com/partitio/go-mobile-collection/templates"
	"io"
	"sort"
	"strings"
)

type GeneratedType struct {
	Name      string
	Type      typeType
	Ptr       string // used when type is interface
	DefaultValue string
	TypeNamed bool
}


var nativesTypes = []string{"int", "int32", "int64", "float32", "float64", "string"}

func NewGeneratedType(name string, tType typeType, typeNamed bool) GeneratedType {
	ptr := "*"
	if tType == typeInterface {
		ptr = ""
	}
	var defaultValue string

	switch strings.ToLower(name) {
	case "int", "int32", "int64", "float32", "float64":
		defaultValue = "0"
	case "string":
		defaultValue = "\"\""
	default:
		defaultValue = "nil"
	}

	return GeneratedType{
		Name:      name,
		Type:      tType,
		Ptr:       ptr,
		DefaultValue: defaultValue,
		TypeNamed: typeNamed,
	}
}

type GenerateTemplateData struct {
	Package string
	Types   []GeneratedType
}

func render(w io.Writer, packageName string, types []GeneratedType) error {
	sort.Slice(types, func(i, j int) bool { return types[i].Name < types[j].Name })
	return templates.Generator.Execute(w, GenerateTemplateData{packageName, types})
}
