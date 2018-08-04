package main

import (
	"fmt"

	"log"

	"github.com/kyokomi-sandbox/go-sandbox/libexample/_doma/dao"
    "github.com/k0kubun/pp")

func main() {
    log.SetFlags(log.Llongfile)
	fmt.Println("Hello doma!")
    
    opts := dao.Options{
        Debug: true,
    }
	d, err := dao.NewDao(opts)
	if err != nil {
		log.Fatalln(err)
	}
    defer d.Close()

	q, err := d.Quest.SelectByID(dao.QueryArgs{
        "id": 1,
        "name": "quest1",
    })
    if err != nil {
        log.Fatalln(err)
    }
    
    pp.Println(q)
}
