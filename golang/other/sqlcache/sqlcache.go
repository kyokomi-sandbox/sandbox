package sqlcache

import (
	"database/sql"
	"log"

	"time"

	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/k0kubun/pp"
	"gopkg.in/redis.v3"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

type Quest struct {
	ID        int
	Name      string
	Detail    string
	CreatedAt time.Time
}

type sqlCache struct {
	RDB    *sql.DB
	KVS    *redis.Client
	Memory map[string]interface{}
}

func (s *sqlCache) Close() {
	s.RDB.Close()
	s.KVS.Close()
}

func New() sqlCache {
	s := sqlCache{}

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	s.RDB = db
	s.RDB.SetMaxOpenConns(80)

	opt := redis.Options{}
	opt.Addr = "localhost:6379"
	opt.Network = "tcp"
	s.KVS = redis.NewClient(&opt)

	s.Memory = map[string]interface{}{}

	return s
}

func (s *sqlCache) QueryMemoryCache(query string, args ...interface{}) Quest {
	q := Quest{}

	key := fmt.Sprintf("%s:%v", query, args)

	data, ok := s.Memory[key]
	if !ok {
		q = s.Query(query, args...)

		// cache
		s.Memory[key] = interface{}(q)
	} else {
		q = data.(Quest)
	}

	return q
}

func (s *sqlCache) QueryKVSCache(query string, args ...interface{}) Quest {
	q := Quest{}

	key := fmt.Sprintf("%s:%v", query, args)

	res, err := s.KVS.HGetAllMap(key).Result()
	if len(res) == 0 || err != nil {
		q = s.Query(query, args...)

		// cache
		err = s.KVS.HMSet(key,
			"id", strconv.Itoa(q.ID),
			"name", q.Name,
			"detail", q.Detail,
			"created_at", q.CreatedAt.Format("2006-01-02 15:04:05"),
		).Err()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		q.ID, _ = strconv.Atoi(res["id"])
		q.Name = res["name"]
		q.Detail = res["detail"]
		q.CreatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", res["created_at"], time.Local)
	}

	return q
}

func (s *sqlCache) Query(query string, args ...interface{}) Quest {
	rows, err := s.RDB.Query(query, args...)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var q Quest
		rows.Scan(&q.ID, &q.Name, &q.Detail, &q.CreatedAt)
		return q
	}

	return Quest{}
}
