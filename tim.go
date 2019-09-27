package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	// modify one single Go file
	// and print to stdout

	fn := os.Args[1]
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("%s: %s", fn, err)
	}

	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, fn, content, parser.ParseComments)
	if err != nil {
		log.Fatalf("%s: %s", fn, err)
	}

	f := &finder{
		content: content,
		fset:    fset,
		buf:     NewBuffer(content),
		pf:      parsedFile,
	}

BigLoop:
	for _, decl := range parsedFile.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.Name == "main" {
				offset := f.find(d.Type.End(), "{")
				f.buf.Insert(offset+1, "\n\tfmt.Println(\"hello world\")\n")
				break BigLoop
			}
		}
	}

	f.ensureImport("fmt")

	fmt.Fprintln(os.Stdout, string(f.buf.Bytes()))
}

type finder struct {
	content []byte
	fset    *token.FileSet
	buf     *Buffer
	pf      *ast.File
}

func (f *finder) offset(pos token.Pos) int {
	return f.fset.Position(pos).Offset
}

func (f *finder) find(pos token.Pos, text string) int {
	b := []byte(text)
	start := f.offset(pos)
	i := start
	s := f.content
	for i < len(s) {
		if bytes.HasPrefix(s[i:], b) {
			return i
		}
		if i+2 <= len(s) && s[i] == '/' && s[i+1] == '/' {
			for i < len(s) && s[i] != '\n' {
				i++
			}
			continue
		}
		if i+2 <= len(s) && s[i] == '/' && s[i+1] == '*' {
			for i += 2; ; i++ {
				if i+2 > len(s) {
					return 0
				}
				if s[i] == '*' && s[i+1] == '/' {
					i += 2
					break
				}
			}
			continue
		}
		i++
	}
	return -1
}

func (f *finder) ensureImport(pkg string) {
	hasFMT := false
	for _, pkg := range f.pf.Imports {
		if strings.HasPrefix(pkg.Path.Value, "\"fmt") {
			hasFMT = true
			break
		}
	}

	if !hasFMT {
		var offset int
		if len(f.pf.Imports) > 0 {
			offset = f.offset(f.pf.Imports[0].Pos())
			f.buf.Insert(offset, "\"fmt\"\n\t")
		} else {
			offset = f.find(f.pf.Package, "\n")
			f.buf.Insert(offset+1, "\nimport \"fmt\"\n")
		}
	}
}
