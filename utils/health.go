package utils

import (
	"fmt"
	"log"
)

func health() {
	err := db.Ping()
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Fatalln("sqlite3 health: bad status,", err.Error())
	}
	rows, err := db.Query("select `UpdatedInDB`,`InsertIntoDB`,`TimeOut`,`Send`,`Receive`,`IMEI`,`IMSI`,`NetCode`,`NetName`,`Client`,`Battery`,`Signal`,`Sent`,`Received` from 'phones' LIMIT 0,30")
	if err != nil {
		log.SetPrefix("[Info] ")
		log.Println("sqlite3 query failed:", err.Error())
	}
	log.SetPrefix("[Info] ")
	log.Println("sqlite3 health.")

	for rows.Next() {
		var (
			UpdatedInDB  string
			InsertIntoDB string
			TimeOut      string
			Send         string
			Receive      string
			IMEI         string
			IMSI         string
			NetCode      string
			NetName      string
			Client       string
			Battery      int
			Signal       int
			Sent         int
			Received     int
		)
		err = rows.Scan(&UpdatedInDB, &InsertIntoDB, &TimeOut, &Send, &Receive, &IMEI, &IMSI, &NetCode, &NetName, &Client, &Battery, &Signal, &Sent, &Received)
		log.SetPrefix("[Info] ")
		log.Print("phone status:\n")
		log.Println(fmt.Sprintf("\nUpdatedInDB: %s\nInsertIntoDB: %s\nTimeOut: %s\nSend: %s\nReceive: %s\nIMEI: %s\nIMSI: %s\nNetCode: %s\nNetName: %s\nClient: %s\nBattery: %d\nSignal: %d\nSent: %d\nReceived: %d", UpdatedInDB, InsertIntoDB, TimeOut, Send, Receive, IMEI, IMSI, NetCode, NetName, Client, Battery, Signal, Sent, Received))
		if err != nil {
			log.SetPrefix("[Error] ")
			fmt.Println("get phone status failed.", err.Error())
		}
	}
	rows.Close()
}
