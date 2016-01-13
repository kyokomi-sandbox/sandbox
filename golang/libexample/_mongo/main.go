package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type Person struct {
	ID      string
	Balance int
	Valid   bool
}

func main() {
	log.SetFlags(log.Llongfile)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	tc := session.DB("txn_test").C("tc")

	runner := txn.NewRunner(tc)
	accounts := session.DB("txn_test").C("accounts")

	if err := accounts.Insert(bson.M{"_id": 0, "balance": 300}); err != nil {
		log.Fatalln(err)
	}
	if err := accounts.Insert(bson.M{"_id": 1, "balance": 100}); err != nil {
		log.Fatalln(err)
	}
	if err := accounts.Insert(bson.M{"_id": 2, "balance": 200}); err != nil {
		log.Fatalln(err)
	}
	if err := accounts.Insert(bson.M{"_id": 3, "balance": 400}); err != nil {
		log.Fatalln(err)
	}

	ops := []txn.Op{
		{
			C:      "accounts",
			Id:     0,
			Assert: bson.D{{"$or", []bson.D{{{"balance", 100}}, {{"balance", 300}}}}},
			Update: bson.D{{"$inc", bson.D{{"balance", 100}}}},
		},
		{
			C:      "accounts",
			Id:     1,
			Assert: bson.D{{"balance", 100}},
			Update: bson.D{{"$inc", bson.D{{"balance", 150}}}},
		},
	}

	if err := runner.Run(ops, "", nil); err != nil {
		log.Fatalln(err)
	}
}
