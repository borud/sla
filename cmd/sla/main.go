package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hako/durafmt"
)

const day = 24 * time.Hour
const week = 7 * day
const month = 30 * day
const year = 365 * day

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("\nUsage:\n\t%s <sla in %%>\n\n", os.Args[0])
		os.Exit(-1)
	}

	sla, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("not a float")
	}

	uptime := sla / 100.0
	downtime := 1.0 - uptime
	fmt.Println()
	fmt.Printf("  Downtime : %f\n", 100.0-sla)
	fmt.Printf("  Uptime   : %f\n", sla)
	fmt.Println()
	fmt.Printf("  Daily    : %v\n", durafmt.Parse(time.Duration(float64(day)*downtime)))
	fmt.Printf("  Weekly   : %v\n", durafmt.Parse(time.Duration(float64(week)*downtime)))
	fmt.Printf("  Month    : %v\n", durafmt.Parse(time.Duration(float64(month)*downtime)))
	fmt.Printf("  Year     : %v\n", durafmt.Parse(time.Duration(float64(year)*downtime)))
	fmt.Println()
}
