package ddns

import (
	"log"
	"net"
	"os"
	"strings"
)

var (
	// Info Logger
	Info *log.Logger
	// Warning Logger
	Warning *log.Logger
	// Error Logger
	Error *log.Logger

	ARecords []string
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

	ARecords = []string{"0.0.0.0"}
}

func Delta() {

	domain := os.Getenv("DELTA_DOMAIN")
	iprecords, _ := net.LookupIP(domain)

	for _, ip := range iprecords {
		if !contains(ARecords, ip.String()) {
			Info.Printf("%s: New Ip discovered: %s", domain, ip.String())
		}
	}

	arecords_tmp := make([]string, len(iprecords))
	for i, ip := range iprecords {
		arecords_tmp[i] = ip.String()
	}
	ARecords = arecords_tmp
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}
