package main

import (
  "os"
  "fmt"
  "net/http"
  "database/sql"
  _ "github.com/mattn/go-oci8"
)

func main(){

  http.HandleFunc("/",hello)

  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }

  err := http.ListenAndServe(":"+port,nil)
  if err!=nil {
    panic(err)
  }

}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res,"Hello World")
    
    db, err := sql.Open("oci8", getDSN())
    if err != nil {
      fmt.Fprintln(res,err)
      return
    }
    defer db.Close()

    if err = testSelect(db); err != nil {
      fmt.Fprintln(res,err)
      return
    }

    rows, err := db.Query("select tname from tab where rownum < 5")
    if err != nil {
      fmt.Fprintln(res,err)
      return 
    }
    defer rows.Close()

    for rows.Next() {
      var tname string
      rows.Scan(&tname)
      fmt.Fprintln(res,tname)
    }
    
}

func getDSN() string {
  return "scott/tiger@dbip:1523/tk101t"
}

func testSelect(db *sql.DB) error {
    rows, err := db.Query("select 3.14, 'foo' from dual")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var f1 float64
        var f2 string
        rows.Scan(&f1, &f2)
        println(f1, f2) // 3.14 foo
    }
    return nil
}

