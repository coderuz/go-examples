package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// load postgres driver
	_ "github.com/lib/pq"
)

func main() {
	var db, _ = sql.Open("postgres", "user=postgres password=password dbname=workly_test sslmode=disable")
	defer db.Close()

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		// Write Header
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			emp_id, emp_id_err := strconv.Atoi(r.FormValue("EMP_ID"))
			event_code, event_code_err := strconv.Atoi(r.FormValue("EVENT_CODE"))
			dt := r.FormValue("DT")
			device_sn := r.FormValue("DEVICE_SN")

			if emp_id_err == nil && event_code_err == nil && dt != "" && device_sn != "" {
				_, dberr := db.Query("INSERT INTO inouts (EMP_ID, EVENT_CODE, DT, DEVICE_SN) VALUES ($1, $2, $3, $4)", emp_id, event_code, dt, device_sn)
				if dberr != nil {
					log.Fatal(dberr)
					println("Error")
				}
				fmt.Fprintf(w, `{"success": 1, "success_text": "Registered"}`)
				return
			}
		}

		fmt.Fprintf(w, `{"error": 1, "error_text": "Bad Request","docs": "Do 'POST /event' with {EMP_ID, EVENT_CODE, DT, DEVICE_SN} data"}`)
	})

	http.ListenAndServe(":8080", nil)
}

/*
 CREATE TABLE inouts (
	ID  SERIAL PRIMARY KEY,
    EMP_ID INT NOT NULL,
    EVENT_CODE SMALLINT NOT NULL,
    DT TIMESTAMP,
    DEVICE_SN VARCHAR(50)
 );

 INSERT INTO inouts (EMP_ID, EVENT_CODE, DT, DEVICE_SN) VALUES (1, 0, '2018-08-14 04:05:06', '4CE0460D0G')

CREATE USER workly_test WITH PASSWORD 'password';

ALTER ROLE workly_test SET client_encoding TO 'utf8';
ALTER ROLE workly_test SET default_transaction_isolation TO 'read committed';
ALTER ROLE workly_test SET timezone TO 'UTC';

GRANT ALL PRIVILEGES ON DATABASE workly_test TO workly_test;
*/
