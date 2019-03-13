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
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

type typeType int

const (
	typeStruct typeType = iota
	typeInterface
)

func loadFile(inputPath string) (string, []GeneratedType) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Could not parse file: %s", err)
	}

	packageName := identifyPackage(f)
	if packageName == "" {
		log.Fatalf("Could not determine package name of %s", inputPath)
	}

	collectionTypes := map[string]typeType{}
	var types []GeneratedType
	for _, decl := range f.Decls {
		typeName, tType, ok, typeNamed := identifyCollectionType(decl)
		if ok {
			collectionTypes[typeName] = tType
			collectionType := NewGeneratedType(typeName, tType, typeNamed)
			types = append(types, collectionType)
			continue
		}
	}

	return packageName, types
}

func identifyPackage(f *ast.File) string {
	if f.Name == nil {
		return ""
	}
	return f.Name.Name
}

func identifyCollectionType(decl ast.Decl) (typeName string, tType typeType, match bool, typeNamed bool) {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	if genDecl.Doc == nil {
		return
	}

	found := false
	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@collection-wrapper") {
			if strings.Contains(comment.Text, "typenamed") {
				typeNamed = true
			}
			found = true
			break
		}
	}
	if !found {
		return
	}

	for _, spec := range genDecl.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			if typeSpec.Name != nil {
				typeName = typeSpec.Name.Name
				if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
					tType = typeInterface
				} else {
					tType = typeStruct
				}
				break
			}
		}
	}
	if typeName == "" {
		return
	}

	match = true
	return
}
