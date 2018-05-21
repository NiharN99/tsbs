package main

import (
	"time"

	"bitbucket.org/440-labs/influxdb-comparisons/query"
)

// TimescaleDBDevopsLastPointPerHost produces TimescaleDB-specific queries for the devops lastpoint case
type TimescaleDBDevopsLastPointPerHost struct {
	TimescaleDBDevops
}

// NewTimescaleDBDevopsLastPointPerHost returns a new TimescaleDBDevopsLastPointPerHost for given paremeters
func NewTimescaleDBDevopsLastPointPerHost(start, end time.Time, scale int) QueryGenerator {
	underlying := newTimescaleDBDevopsCommon(start, end, scale)
	return &TimescaleDBDevopsLastPointPerHost{
		TimescaleDBDevops: *underlying,
	}

}

// Dispatch fills in the query.Query
func (d *TimescaleDBDevopsLastPointPerHost) Dispatch() query.Query {
	q := query.NewTimescaleDB() // from pool
	d.LastPointPerHost(q)
	return q
}