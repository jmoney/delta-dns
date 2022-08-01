package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoney/delta-dns/pkg/ddns"
	"github.com/robfig/cron"
)

var (
	// Info Logger
	Info *log.Logger
	// Warning Logger
	Warning *log.Logger
	// Error Logger
	Error *log.Logger
)

func init() {

	Info = log.New(os.Stdout,
		"[INFO]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stderr,
		"[WARNING]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"[ERROR]: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	c := cron.New()

	cronErr := c.AddFunc("* * * * *", ddns.Delta)
	if cronErr != nil {
		Error.Printf("cron errored out: %v", cronErr)
		return
	}

	go c.Start()
	sig := make(chan os.Signal, 5)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}
