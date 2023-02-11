package utils

import (
	"fmt"
	"log"
)

func query() {
	rowslen, _ := db.Query("select count(*) from `inbox` where `Processed` = 'false'")
	var number int
	for rowslen.Next() {
		rowslen.Scan(&number)
	}
	rowslen.Close()
	if number > 0 {
		rows, _ := db.Query("select `id` from `inbox` where `Processed` = 'false'")
		var idlist []int
		for rows.Next() {
			var ID int
			rows.Scan(&ID)
			idlist = append(idlist, ID)
		}
		rows.Close()
		for _, i := range idlist {
			querydetail(i)
		}
	}
}

func querydetail(ID int) {
	rows, _ := db.Query("select `ReceivingDateTime`,`SenderNumber`,`TextDecoded` from `inbox` where `id` = ?", ID)
	var (
		ReceivingDateTime string
		SenderNumber      string
		TextDecoded       string
	)
	for rows.Next() {
		rows.Scan(&ReceivingDateTime, &SenderNumber, &TextDecoded)
	}
	rows.Close()
	log.Println(fmt.Sprintf("query unread message detail :\nReceivingDateTime:%s\nSenderNumber:%s\nTextDecoded:%s\n", ReceivingDateTime, SenderNumber, TextDecoded))
	if TextDecoded != "" {
		if send(ReceivingDateTime, SenderNumber, TextDecoded) {
			update(ID)
		} else {
			update(ID)
		}
	} else {
		update(ID)
	}

}
