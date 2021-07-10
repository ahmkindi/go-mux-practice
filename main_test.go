package main_test

import(
  "os"
  "testing"
  "log"
  "github.com/ahmkindi/go-mux-practice"
)

var a main.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
  id SERIAL,
  name TEXT NOT NULL,
  price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  CONSTRAINT products_pbkey PRIMARY KEY (id)
)`

func TestMain(m *testing.M){
  a.Initialize(
    os.Getenv("APP_DB_USERNAME"),
    os.Getenv("APP_DB_PASSWORD"),
    os.Getenv("APP_DB_NAME"))

  ensureTableExists()
  code := m.Run()
  clearTable()
  os.Exit(code)
}

func ensureTableExists(){
  if _, err := a.DB.Exec(tableCreationQuery); err != nil {
    log.Fatal(err)
  }
}

func clearTable() {
  a.DB.Exec("DELETE FROM products")
  a.DB.Exec("ALTTER SEQUENCE products_id_seg RESTART WITH 1")
}
