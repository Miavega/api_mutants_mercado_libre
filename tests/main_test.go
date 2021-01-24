package test

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
)

func TestMain(m *testing.M) {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgUrls := "us-east-2.console.aws.amazon.com"
	pgDb := os.Getenv("mutants_db")
	pgSchema := os.Getenv("POSTGRES_SCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

	exitVal := m.Run()

	os.Exit(exitVal)
}
