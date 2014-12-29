package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ttacon/chalk"
	"strings"
	"path/filepath"
	"bufio"
	"io"
	"github.com/philhofer/msgp/parse"
	"github.com/philhofer/msgp/gen"
	"text/template"
	"bytes"
	"go/format"
	"github.com/k0kubun/pp"
)

var (
	// command line flags
	out     string // output file
	file    string // input file (or directory)
	pkg     string // output package name

)

func init() {
	flag.StringVar(&out, "o", "", "output file")
	flag.StringVar(&file, "file", "", "input file")
	flag.StringVar(&pkg, "pkg", "", "output package")
}

func main() {
	flag.Parse()

	// GOFILE and GOPACKAGE are
	// set by `go generate`
	if file == "" {
		file = os.Getenv("GOFILE")
	}
	if pkg == "" {
		pkg = os.Getenv("GOPACKAGE")
	}

	if file == "" {
		fmt.Println(chalk.Red.Color("No file to parse."))
		os.Exit(1)
	}

	err := DoAll(pkg, file)
	if err != nil {
		fmt.Println(chalk.Red.Color(err.Error()))
		os.Exit(1)
	}
}

var (
	baseTemplateText  = `
// DecodeMsg implements the msgp.Decodable interface
func ({{.Varname}} *{{.Value.Struct.Name}}) DecodeMsg(dc *msgp.Reader) (err error) {
	{{if not .Value.Struct.AsTuple}}var field []byte; _ = field{{end}}
	{{if .Value.Struct.AsTuple}}
	{ {{/* tuples get their own blocks so that we don't clobber 'ssz'*/}}
		var ssz uint32
		ssz, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if ssz != {{len .Value.Struct.Fields}} {
			err = msgp.ArrayError{Wanted: {{len .Value.Struct.Fields}}, Got: ssz}
			return
		}
		{{range .Value.Struct.Fields}}{{template "ElemTempl" .Value.Struct.FieldElem}}{{end}}
	}
	{{else}}
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz:=uint32(0); xplz<isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}

		// TODO:
	}
	{{end}}

	return
}
`
	sampleTemplate = template.Must(template.New("base").Parse(baseTemplateText))
)

// DoAll writes all methods using the associated file and package.
// (The package is only relevant for writing the new file's package declaration.)
func DoAll(gopkg string, gofile string) error {
	var (
		outwr  *bufio.Writer // location to write methods
	)

	fmt.Println("gopkg: ", gopkg, " gofile: ", gofile)

	newfile := createNewFileName(gopkg, out, gofile)
	fmt.Println("fileName: ", newfile)

	file, err := os.Create(newfile)
	if err != nil {
		return err
	}
	defer file.Close()

	outwr = bufio.NewWriter(file)

	err = writePkgHeader(outwr, gopkg)
	if err != nil {
		return err
	}

	err = writeImportHeader(outwr, "github.com/philhofer/msgp/msgp")
	if err != nil {
		return err
	}

	elems, _, err := parse.GetElems(gofile)
	if err != nil {
		return err
	}

	for _, el := range elems {

		p, ok := el.(*gen.Ptr)
		pp.Println(p, ok, !ok || p.Value.Type() != gen.StructType, p.Value.Type(), gen.StructType)
		if !ok || p.Value.Type() != gen.StructType {
			continue
		}

		pp.Println(p.Value.TypeName())

		if err := execAndFormat(sampleTemplate, outwr, p); err != nil {
			fmt.Println(err)
			continue
		}
	}

	err = outwr.Flush()
	if err != nil {
		return err
	}
	fmt.Print(chalk.Green.Color("\u2713\n"))

	return nil
}

///////////////////////////
// あとでまとめる

// createNewFileName generateするfile名をいい感じに作成する
// output先を指定した場合は_genをつけない
func createNewFileName(pkgName string, out string, gofile string) string {
	var isDir bool
	if fInfo, err := os.Stat(gofile); err == nil && fInfo.IsDir() {
		isDir = true
	}

	var newfile string // new file name
	if out != "" {
		newfile = out
		if pre := strings.TrimPrefix(out, gofile); len(pre) > 0 &&
				!strings.HasSuffix(out, ".go") {
			newfile = filepath.Join(gofile, out)
		}
	} else {
		// small sanity check if gofile == . or dir
		// let's just stat it again, not too costly
		if isDir {
			gofile = filepath.Join(gofile, pkgName)
		}
		// new file name is old file name + _gen.go
		newfile = strings.TrimSuffix(gofile, ".go") + "_gen.go"
	}

	return newfile
}

// fileの中身作成

func writePkgHeader(w io.Writer, name string) error {
	_, err := io.WriteString(w, fmt.Sprintf("package %s\n\n", name))
	if err != nil {
		return err
	}

	// TODO: option NOTE
//	_, err = io.WriteString(w, "// NOTE: THIS FILE WAS PRODUCED BY THE\n// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)\n// DO NOT EDIT\n\n")
	return err
}

func writeImportHeader(w io.Writer, imports ...string) error {
	_, err := io.WriteString(w, "import (\n")
	if err != nil {
		return err
	}
	for _, im := range imports {
		_, err = io.WriteString(w, fmt.Sprintf("\t%q\n", im))
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(w, ")\n\n")
	return err
}

// execAndFormat executes a template and formats the output, using buf as temporary storage
func execAndFormat(t *template.Template, w io.Writer, i interface{}) error {
	buf := bytes.NewBuffer(nil)

	err := t.Execute(buf, i)
	if err != nil {
		return fmt.Errorf("template: %s", err)
	}
	bts, err := format.Source(buf.Bytes())
	if err != nil {
		w.Write(buf.Bytes())
		return fmt.Errorf("gofmt: %s", err)
	}
	_, err = w.Write(bts)
	return err
}
