package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type authcar_platelist struct{
	id int
	card_info_id int
	licence_no  string
	is_parking  int
	parking_time string
}

func main(){

	db, err := sql.Open("mysql", "root:a1b2c3@tcp(192.168.2.253:3306)/acdbs?parseTime=true")
	if err != nil{
		log.Printf("get err ... \n")
		log.Fatal(err)
	}

	defer db.Close()
	

	err = db.Ping()
	if err != nil{
		log.Printf("err db ping()")
		log.Fatalln(err)
	}

	rows, err := db.Query("select * from authcar_platelist")

	if err != nil{
		log.Printf("Query err ... \n")
		log.Fatalln(err)
	}

	for rows.Next() {
		// log.Printf("row next", rows)
		// var id, card_info_id, licence_no, is_parking, parking_time interface{}
		var authcardb authcar_platelist
		err = rows.Scan(
			&authcardb.id,  
			&authcardb.card_info_id, 
			&authcardb.licence_no, 
			&authcardb.is_parking, 
			&authcardb.parking_time)

		if err != nil{
			log.Fatalln(err)
		}
		log.Printf("found row containing: %d, %d, %s, %d",
			authcardb.id,  
			authcardb.card_info_id, 
			authcardb.licence_no, 
			authcardb.is_parking)

		// DefaultTimeLoc := time.Local
		parking_time, _ := time.Parse("2006-01-02 15:04:05", authcardb.parking_time)
		log.Println(parking_time)

	}
	rows.Close()

	/*
	cols, _ := rows.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	*/
}