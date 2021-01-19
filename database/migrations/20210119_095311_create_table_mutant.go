package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableMutant_20210119_095311 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableMutant_20210119_095311{}
	m.Created = "20210119_095311"

	migration.Register("CreateTableMutant_20210119_095311", m)
}

// Run the migrations
func (m *CreateTableMutant_20210119_095311) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA schema_xmen;")
	m.SQL("ALTER SCHEMA schema_xmen OWNER TO postgres;")
	m.SQL("SET search_path TO pg_catalog,public,schema_xmen;")
	m.SQL("CREATE TABLE schema_xmen.mutant (id serial NOT NULL,dna varchar NOT NULL,is_mutant boolean NOT NULL,CONSTRAINT mutant_pk PRIMARY KEY (id));")
	m.SQL("ALTER TABLE schema_xmen.mutant OWNER TO postgres;")
}

// Reverse the migrations
func (m *CreateTableMutant_20210119_095311) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE schema_xmen.mutant;")
}
