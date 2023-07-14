package usecase

import (
	"database/sql"
)

type KeywordUsecase struct {
	db *sql.DB
}

func NewKeywordUsecase(db *sql.DB) KeywordUsecase {
	return KeywordUsecase{db: db}
}

func (k *KeywordUsecase) Add(keyword string) error {
	query := `INSERT INTO "keywords" ("keyword")
			   VALUES ($1);`

	_, err := k.db.Exec(query, keyword)

	if err != nil {
		return err
	}

	return nil
}

func (k *KeywordUsecase) GetAll() ([]string, error) {
	rows, err := k.db.Query("SELECT keyword FROM keywords")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keywords []string

	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			return nil, err
		}

		keywords = append(keywords, keyword)
	}

	return keywords, nil
}

func (k *KeywordUsecase) Delete(keyword string) error {
	//sql, args, err := squirrel.Delete("keyword").From("keywords").Where(squirrel.Eq{"keyword": keyword}).PlaceholderFormat(squirrel.Dollar).ToSql()
	query := `DELETE FROM keywords WHERE keyword=$1`

	_, err := k.db.Exec(query, keyword)
	if err != nil {
		return err
	}

	return nil
}
