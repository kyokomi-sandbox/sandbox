package dao

// QuestDao is generated quest table.
type QuestDao struct {
	*Dao
}

// QuestEntity is generated quest table.
type QuestEntity struct {
    ID int
    Name string
    Detail string
}

const questSelectByID = `
select
  *
FROM
  quest
WHERE
  id = /* id */1
and
  name = /* name */"hoge"
`

func (d *QuestDao) SelectByID(args QueryArgs) (*QuestEntity, error) {

	// TODO: sqlファイルを読み込む（Dao生成時にmethod名、引数をKeyにmapで保持する
	queryString := d.queryArgs(questSelectByID, args)

	var entity QuestEntity
	err := d.QueryRow(queryString).Scan(&entity.ID, &entity.Name, &entity.Detail)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
