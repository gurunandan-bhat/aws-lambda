package main

import (
	"aws-lambda/config"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	signer "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

func main() {

	cfg, err := config.Configuration()
	if err != nil {
		log.Fatalf("error parsing configuration: %s", err)
	}

	params := url.Values{}
	params.Add("vCategoryUrlName", "boxers")
	url := url.URL{
		Scheme:   "https",
		Host:     cfg.APIGateway,
		Path:     "categoryProducts",
		RawQuery: params.Encode(),
	}

	creds := aws.Credentials{
		AccessKeyID:     cfg.AccessKey,
		SecretAccessKey: cfg.SecretAccessKey,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		log.Fatalf("error constructing request: %s", err)
	}
	req.Header.Add("Accept", "*/*")

	payloadHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	s := signer.NewSigner()
	if err := s.SignHTTP(
		context.Background(),
		creds,
		req,
		payloadHash,
		"execute-api",
		cfg.AWSRegion,
		time.Now(),
	); err != nil {
		log.Fatalf("error signing request: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error requesting sub categories: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading resing response body: %s", err)
	}
	fmt.Println(string(respBody))
}
