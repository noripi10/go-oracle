package libs

import (
	"database/sql"
	"fmt"
)

func GetShain(db *sql.DB, cd string) (string, string, error) {
	rows, err := db.Query("SELECT shain_code, shain_name FROM M_SHAIN WHERE shain_code =:CD", cd)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return "", "", err
	}

	defer rows.Close()

	var _shainCode, _shainName string
	for rows.Next() {
		rows.Scan(&_shainCode, &_shainName)
		break
	}

	return _shainCode, _shainName, nil
}
