package utils

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// GetCountryCIDRs returns list of CIDR ranges for a given country
// Data source: This should be regularly updated from IP2Location or similar
func GetCountryCIDRs(countryCode string) ([]string, error) {
    filename := fmt.Sprintf("data/country_ranges/%s.txt", strings.ToLower(countryCode))
    
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("country data for %s not found", countryCode)
    }
    defer file.Close()

    var cidrs []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(line, "#") {
            cidrs = append(cidrs, line)
        }
    }

    return cidrs, scanner.Err()
}

// Common country codes
var CountryNames = map[string]string{
    "id": "Indonesia",
    "ru": "Russia", 
    "cn": "China",
    "us": "United States",
    "jp": "Japan",
    "kr": "South Korea",
    "vn": "Vietnam",
    "th": "Thailand",
    "my": "Malaysia",
    "sg": "Singapore",
    "ph": "Philippines",
    "in": "India",
    "br": "Brazil",
    "mx": "Mexico",
}
