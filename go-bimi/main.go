package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// parseBIMIRecord extracts the 'l=' logo URL from a BIMI record string
func parseBIMIRecord(txt string) string {
	// Simple regex to find 'l=' followed by a URL
	re := regexp.MustCompile(`l=([^;]+)`)
	match := re.FindStringSubmatch(txt)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// getBIMIRecord queries the DNS for the BIMI TXT record
func getBIMIRecord(domain string) (string, error) {
	lookupName := "default._bimi." + domain
	txts, err := net.LookupTXT(lookupName)
	if err != nil {
		return "", err
	}
	for _, txt := range txts {
		if strings.HasPrefix(txt, "v=BIMI1;") {
			return txt, nil
		}
	}
	return "", fmt.Errorf("No valid BIMI record found")
}

// downloadLogo fetches the SVG from the provided URL
func downloadLogo(url string, domain string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	// Print the SVG content to stdout
	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: bimi <domain>")
		os.Exit(1)
	}
	domain := os.Args[1]
	txtRecord, err := getBIMIRecord(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving BIMI record: %v\n", err)
		os.Exit(1)
	}

	logoURL := parseBIMIRecord(txtRecord)
	if logoURL == "" {
		fmt.Fprintln(os.Stderr, "No logo URL found in BIMI record.")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Found BIMI logo URL: %s\n", logoURL)
	err = downloadLogo(logoURL, domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading logo: %v\n", err)
		os.Exit(1)
	}
}
