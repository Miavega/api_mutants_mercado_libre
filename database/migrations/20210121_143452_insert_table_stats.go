package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTableStats_20210121_143452 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableStats_20210121_143452{}
	m.Created = "20210121_143452"

	migration.Register("InsertTableStats_20210121_143452", m)
}

// Run the migrations
func (m *InsertTableStats_20210121_143452) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO schema_xmen.stats(id, count_mutant_dna, count_human_dna, ratio) VALUES (1, 0, 0, 0.0)")
}

// Reverse the migrations
func (m *InsertTableStats_20210121_143452) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE schema_xmen.stats;")
}
