package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

// クラスによるAdapterパターン（継承を使ったもの）
// インスタンスによるAdapterパターン（委譲を使ったもの）

// Print is Printer
type Print interface {
	printWeak()
	printStrong()
}

// Banner is banner
type Banner struct {
	name string
}

func (b Banner) showWithParam() {
	fmt.Println("(" + b.name + ")")
}

func (b Banner) showWithAster() {
	fmt.Println("*" + b.name + "*")
}

// PrintBannerAdapter is Print interface Banner adapter
type PrintBannerAdapter struct {
	banner *Banner
}

// NewPrintBannerAdapter is PrintBannerAdapter Construct.
func NewPrintBannerAdapter(name string) Print {
	p := PrintBannerAdapter{}
	p.banner = &Banner{}
	p.banner.name = name

	return &p
}

func (p PrintBannerAdapter) printWeak() {
	p.banner.showWithParam()
}

func (p PrintBannerAdapter) printStrong() {
	p.banner.showWithAster()
}

var _ Print = (*PrintBannerAdapter)(nil)

// -------------------

type FileIO interface {
	ReadFromFile() error
	WriteToFile() error
	SetValue(key string, value string)
	GetValue(key string) string
}

type FileJson struct {
	filename string
	values map[string]string
}

func NewFileJson(filename string) FileIO {
	f := FileJson{}
	f.filename = filename
	f.values = make(map[string]string)
	return &f
}

func (f *FileJson) ReadFromFile() error {

	data, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &f.values); err != nil {
		return err
	}

	return nil
}

func (f *FileJson) WriteToFile() error {
	data, err := json.Marshal(f.values)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(f.filename, data, 644); err != nil {
		return err
	}
	return nil
}

func (f *FileJson) SetValue(key string, value string) {
	f.values[key] = value
}

func (f FileJson) GetValue(key string) string {
	return f.values[key]
}

var _ FileIO = (*FileJson)(nil)

func main() {
	p := NewPrintBannerAdapter("hello")
	p.printWeak()
	p.printStrong()

	f := NewFileJson("./file.txt")
	if err := f.ReadFromFile(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(f.GetValue("year"))
	f.SetValue("hoge", "fuga")

	if err := f.WriteToFile(); err != nil {
		log.Fatalln(err)
	}
}
