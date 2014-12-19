package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/gchaincl/dotsql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/lann/squirrel"

	"testing"
)

type Test struct {
	db         *sql.DB
	dbx        *sqlx.DB
	dbgorm     gorm.DB
	dbbeegoorm orm.Ormer
}

var test Test = Test{}

type User struct {
	A, B string
}

// Model Struct
type BeeUser struct {
	Id int `orm:"auto"`
	A  string
	B  string
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", ":memory:")
	orm.RegisterModel(new(BeeUser))
}

func TestMain(m *testing.M) {
	var err error

	test.db, err = sql.Open("sqlite3", ":memory:")
	panicIfErr(err)

	_, err = test.db.Exec("CREATE TABLE users(a, b UNIQUE)")
	panicIfErr(err)

	_, err = test.db.Exec("INSERT INTO users VALUES(1, 2)")
	panicIfErr(err)

	test.dbx = sqlx.NewDb(test.db, "sqlite3")

	test.dbgorm, err = gorm.Open("sqlite3", ":memory:")
	panicIfErr(err)

	user := &User{"1", "2"}
	test.dbgorm.CreateTable(user)
	test.dbgorm.Create(user)

	err = orm.RunSyncdb("default", true, false)
	panicIfErr(err)
	test.dbbeegoorm = orm.NewOrm()
	beeuser := &BeeUser{A: "1", B: "2"}
	test.dbbeegoorm.Insert(beeuser)

	os.Exit(m.Run())
}

func BenchmarkNative(b *testing.B) {
	db := test.db

	var t1, t2 string

	for i := 0; i < b.N; i++ {
		rows, err := db.Query("SELECT * from users")
		panicIfErr(err)

		for rows.Next() {
			err := rows.Scan(&t1, &t2)
			panicIfErr(err)
		}
	}
}

func BenchmarkBeego(b *testing.B) {
	db := test.dbbeegoorm
	db.Using("default") // Using default, you can use other database

	for i := 0; i < b.N; i++ {
		var users []*BeeUser
		qs := test.dbbeegoorm.QueryTable("bee_user")
		_, err := qs.All(&users)
		panicIfErr(err)
	}
}

func BenchmarkSqlX(b *testing.B) {
	db := test.dbx
	var user User

	for i := 0; i < b.N; i++ {
		rows, err := db.Queryx("SELECT * from users")
		panicIfErr(err)

		for rows.Next() {
			err := rows.StructScan(&user)
			panicIfErr(err)
		}
	}
}

func BenchmarkDotSQL(b *testing.B) {
	db := test.db

	dot, err := dotsql.LoadFromString(`
	-- name: select
	SELECT * from users`)
	panicIfErr(err)

	var t1, t2 string

	for i := 0; i < b.N; i++ {
		rows, err := dot.Query(db, "select")
		panicIfErr(err)

		for rows.Next() {
			err := rows.Scan(&t1, &t2)
			panicIfErr(err)
		}
	}
}

func BenchmarkSqrl(b *testing.B) {
	db := test.db

	var t1, t2 string

	for i := 0; i < b.N; i++ {
		rows, err := squirrel.Select("*").From("users").RunWith(db).Query()
		panicIfErr(err)

		for rows.Next() {
			err := rows.Scan(&t1, &t2)
			panicIfErr(err)
		}
	}
}

func BenchmarkGorm(b *testing.B) {
	db := test.dbgorm

	users := []User{}

	for i := 0; i < b.N; i++ {
		db.Find(&users)
	}
}
