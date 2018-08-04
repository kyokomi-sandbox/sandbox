package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/pquerna/cachecontrol/cacheobject"
)

func main() {


	respDir, err := cacheobject.ParseResponseCacheControl("private, max-age=-1526432994")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(respDir.MaxAge)
}
