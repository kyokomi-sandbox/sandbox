package dao

import "database/sql"

type Options struct {
    Debug bool
}

type Dao struct {
	*sql.DB
    options Options
    
    // generate dao
    Quest QuestDao
}

func (d *Dao) setupDao() {

    // generate dao
    d.Quest = QuestDao{Dao: d}
    
}
