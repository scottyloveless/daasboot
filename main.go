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
	PIMS         string
}

func main() {
	var cfg config
	var err error

	godotenv.Load()

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")
	cfg.CCID = os.Getenv("CCID")
	cfg.SiteID = os.Getenv("SITE_ID")

	flag.StringVar(&cfg.Clinic, "clinic", "", "MPH six digit shortcode")
	flag.StringVar(&cfg.PIMS, "pims", "", "Avimark or Cornerstone")
	flag.Parse()

	if cfg.Clinic == "" || cfg.PIMS == "" {
		fmt.Println("Usage: --clinic CLINIC --pims Avimark")
		os.Exit(1)
	}
	if cfg.PIMS != "Avimark" && cfg.PIMS != "Cstone" {
		fmt.Println("only two valid PIMS are Avimark and Cstone")
		os.Exit(1)
	}

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

	catalogExists, err := cfg.checkMachineCatalogName()
	if err != nil {
		panic(err)
	}

	if catalogExists {
		fmt.Println("Catalog already exists, continuing to Delivery Group...")
	} else {
		fmt.Println("Catalog name is available")
		fmt.Println("Creating new Catalog...")
		_, err = cfg.createMachineCatalog()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Catalog created successfully: MC_%s\n", cfg.Clinic)

	}

	// TODO: Create VM in vCenter here

	// TODO: Add Machine to Machine Catalog

	fmt.Println("Seeing if Delivery Group already exists...")
	dgExists, err := cfg.checkDeliveryGroupExists()
	if err != nil {
		panic(err)
	}
	if dgExists {
		// TODO: add delivery group check here
		fmt.Println("Delivery group already exists, continuing...")
	} else {
		fmt.Println("Delivery group does not exist, creating new one...")
		err = cfg.createDeliveryGroup()
		if err != nil {
			panic(err)
		}
		fmt.Println("Delivery group and PIMS apps created successfully")
	}

	os.Exit(0)
}
