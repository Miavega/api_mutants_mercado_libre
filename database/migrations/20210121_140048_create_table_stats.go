package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableStats_20210121_140048 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableStats_20210121_140048{}
	m.Created = "20210121_140048"

	migration.Register("CreateTableStats_20210121_140048", m)
}

// Run the migrations
func (m *CreateTableStats_20210121_140048) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE schema_xmen.stats (id serial NOT NULL,count_mutant_dna integer NOT NULL,count_human_dna integer NOT NULL,ratio numeric(5,1) NOT NULL,CONSTRAINT stats_pk PRIMARY KEY (id));")
	m.SQL("ALTER TABLE schema_xmen.stats OWNER TO postgres;")
}

// Reverse the migrations
func (m *CreateTableStats_20210121_140048) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE schema_xmen.stats;")
}
