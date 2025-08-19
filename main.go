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
}

func main() {
	var cfg config

	godotenv.Load()

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")
	cfg.CCID = os.Getenv("CCID")

	flag.StringVar(&cfg.Clinic, "clinic", "CLINIC", "MPH six digit shortcode")
	flag.Parse()

	var err error
	cfg.BearerToken, cfg.TokenExpiry, err = cfg.getBearerToken()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Clinic: %s\nClientID: %s\nClientSecret: %s\nCCID: %s\nToken: %s\nTokenExpiry: %v\n", cfg.Clinic, os.Getenv("CLIENT_ID"), cfg.ClientSecret, cfg.CCID, cfg.BearerToken, cfg.TokenExpiry)

	machineCatalogs, err := cfg.getMachineCatalogs()
	if err != nil {
		log.Fatalf("error retrieving machine catalog list: %v", err)
	}

	for _, mc := range machineCatalogs.Items {
		fmt.Println(mc.Name)
	}
}
