package main

import "flag"

var (
	a *string
	d *string
	r *string
)

func init() {
	a = flag.String("a", "", "RUN_ADDRESS")
	d = flag.String("d", "", "DATABASE_URI")
	r = flag.String("r", "", "ACCRUAL_SYSTEM_ADDRESS")
}
