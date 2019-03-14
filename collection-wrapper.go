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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func process(inputPath []string) error {
	types := map[string][]GeneratedType{}
	packagePaths := map[string]string{}
	for _, p := range inputPath {
		packageName, ts := loadFile(p)
		packagePaths[packageName] = filepath.Dir(p)
		if pt, ok := types[packageName]; ok {
			types[packageName] = append(pt, ts...)
			continue
		}
		types[packageName] = ts
	}

	for p, t := range types {
		if t == nil {
			continue
		}
		outputPath := fmt.Sprintf("%s/%s_collection.go", packagePaths[p], p)

		output, err := CreateOrReplace(outputPath)
		if err != nil {
			return err
		}
		if err := render(output, p, t); err != nil {
			output.Close()
			return fmt.Errorf("Could not generate go code: %s", err)
		}
		output.Close()
	}
	return nil
}

func CreateOrReplace(outputPath string) (*os.File, error) {
	if _, err := os.Stat(outputPath); err == nil {
		if err := os.Remove(outputPath); err != nil {
			return nil, fmt.Errorf("Could not remote output file %s", outputPath)
		}
	}
	output, err := os.Create(outputPath)
	if err != nil {
		return nil, fmt.Errorf("Could not open output file: %s", err)
	}
	return output, nil
}

func main() {
	cmd := &cobra.Command{
		Use:  "go-mobile-collection file...",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var files []string
			for _, path := range args {
				// Ignore vendor
				if path == "vendor" {
					continue
				}
				// Check files
				p, err := os.Stat(path)
				if err != nil {
					log.Fatalf("Invalid file or directory: %s", path)
				}
				if !p.IsDir() {
					if !strings.HasSuffix(path, ".go") {
						return fmt.Errorf("Invalid file: %s", path)
					}
					files = append(files, path)
					continue
				}
				path = strings.TrimSuffix(path, "/")
				fs, err := ioutil.ReadDir(path)
				if err != nil {
					return err
				}
				for _, f := range fs {
					if strings.HasSuffix(f.Name(), ".go") {
						files = append(files, fmt.Sprintf("%s/%s", path, f.Name()))
					}
				}

			}
			if err := process(files); err != nil {
				return err
			}

			natives, err := cmd.Flags().GetString("natives")
			if err != nil {
				return err
			}
			if natives == "" {
				return nil
			}
			var gts []GeneratedType
			for _, n := range nativesTypes {
				gts = append(gts, NewGeneratedType(n, typeInterface, false))
			}
			out, err := CreateOrReplace("native_collection.go")
			if err != nil {
				return err
			}
			defer out.Close()
			return render(out, natives, gts)
		},
	}
	cmd.Flags().StringP("natives", "n", "", "Generate collection for native types (string, ints, floats) inside given package")
	cmd.Execute()

}
