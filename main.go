package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	Clinic       string
	ClientID     string
	ClientSecret string
	CCID         string
	BearerToken  string
	TokenExpiry  time.Time
	SiteID       string
}

func main() {
	var cfg config
	var err error

	godotenv.Load()

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")
	cfg.CCID = os.Getenv("CCID")
	cfg.SiteID = os.Getenv("SITE_ID")

	flag.StringVar(&cfg.Clinic, "clinic", "CLINIC", "MPH six digit shortcode")
	flag.Parse()

	cfg.BearerToken, cfg.TokenExpiry, err = cfg.getBearerToken()
	if err != nil {
		panic(err)
	}

	fmt.Println("Fetching machine catalogs...")
	machineCatalogs, err := cfg.getMachineCatalogs()
	if err != nil {
		log.Fatalf("error retrieving machine catalog list: %v", err)
	}

	for _, mc := range machineCatalogs.Items {
		fmt.Println(mc.Name)
	}

	mcNameExists, err := cfg.checkMachineCatalogName()
	if err != nil {
		panic(err)
	}

	if mcNameExists {
		fmt.Println("Catalog already exists")
	} else {
		fmt.Println("Catalog name is available")
	}
}
