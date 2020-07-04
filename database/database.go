package database

import (
    "fmt"
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

)



type Config struct {
	Addr         string          // for trace
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  time.Duration   // connect max life time.
	QueryTimeout time.Duration   // query sql timeout
	ExecTimeout  time.Duration   // execute sql timeout
	TranTimeout  time.Duration   // transaction sql timeout

}


func NewMySQL() (*sql.DB, error) {
    dsn := "root:1055979970YuYa@tcp(127.0.0.1:3306)/business_engine?charset=utf8&readTimeout=30s&writeTimeout=30s"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
    d.SetMaxOpenConns(100)
	d.SetMaxIdleConns(100)
	d.SetConnMaxLifetime(time.Duration(time.Minute))
	return d, nil
}
type EventRecord struct {
    SiteId string
    Type string
    FileName string
}

func Query(db *sql.DB) ([]EventRecord, error){
    s := "select siteId,type,filename from bi_offline_events_status where id > ?;"
    rows, err := db.Query(s, 1)
    if err != nil {
        return nil, err
    }
    var events []EventRecord 
    for rows.Next() {
    
    var r EventRecord
    if err := rows.Scan(&r.SiteId, &r.Type, &r.FileName); err != nil {
        return nil, err
    }
    fmt.Println(r)
    events = append(events, r)


    }
    return events, nil

}



