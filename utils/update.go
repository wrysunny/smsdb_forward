package utils

import (
	"log"
)

func update(ID int) {
	updatesql, _ := db.Exec("update `inbox` set `Processed` = 'true' WHERE `id` = ?", ID)
	affected, err := updatesql.RowsAffected()
	if err != nil {
		log.SetPrefix("[Error]")
		log.Println("exec update affected error:", err.Error())
	}

	log.SetPrefix("[Info]")
	log.Println("update affected ", affected, "line")
}
