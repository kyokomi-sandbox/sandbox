package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	yaml "gopkg.in/yaml.v2"
)

type Sample struct {
	ID   int
	Name string
}
type ConfigYaml struct {
	Database map[string]ConfigDatabase
}

type ConfigSetting struct {
	Main map[string]ConfigDatabase
}

type ConfigDatabase struct {
	User     string
	Password string
	Host     string
	DB       string
}

func main() {

	var configPath string
	flag.StringVar(&configPath, "path", "./", "config file path")
	flag.Parse()

	log.SetFlags(log.Llongfile)

	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	var conf ConfigYaml
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalln(err)
	}

	pp.Println(conf)

	dbConf := conf.Database["main"]
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.DB,
	)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var sample Sample
	err = db.QueryRow("SELECT * FROM sample LIMIT ?", 1).Scan(&sample.ID, &sample.Name)
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(sample)
}
