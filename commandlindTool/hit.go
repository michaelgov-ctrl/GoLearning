package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"
)

const (
	bannerText = `
 __  __    __     _____
/\ \_\ \  /\ \   /\__ _\
\ \  __ \ \ \ \ \/_/\ \/
 \ \_\ \_\ \ \ \   \ \_\
  \/_/\/_/  \/_/    \/_/
`
	usageText = `
Usage:
	hit [options] url
Options
	-n
		Number of requests to make
	-c
		Concurrency level
`
)

func banner() string {
	return bannerText[1:]
}
func usage() string {
	return usageText[1:]
}

func main() {
	if err := run(flag.CommandLine, os.Args[1:], os.Stdout); err != nil {
		os.Exit(1)
	}

}

func run(s *flag.FlagSet, args []string, out io.Writer) error {
	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}
	if err := f.parse(s, args); err != nil {
		return err
	}
	fmt.Println(banner())
	fmt.Fprintf(out, "Making %d requests to %s with a concurrency level of %d.\n", f.n, f.url, f.c)

	// hit pkg integration here
	var sum Result
	sum.Merge(&Result{
		Bytes:    1000,
		Status:   http.StatusOK,
		Duration: time.Second,
	})
	sum.Merge(&Result{
		Bytes:    1000,
		Status:   http.StatusTeapot,
		Duration: 2 * time.Second,
	})
	sum.Finalize(2 * time.Second)
	sum.Fprint(out)
	return nil
}
