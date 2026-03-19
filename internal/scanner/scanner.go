package scanner

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
    "strings"

    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/models"
    "github.com/fatih/color"
)

type Scanner struct {
    MasscanPath string
    OutputDir   string
    ExcludeFile string
    Rate        int
}

func NewScanner(outputDir string) *Scanner {
    return &Scanner{
        MasscanPath: "masscan",
        OutputDir:   outputDir,
        ExcludeFile: "configs/exclude.conf",
        Rate:        10000,
    }
}

func (s *Scanner) CheckMasscan() error {
    _, err := exec.LookPath(s.MasscanPath)
    if err != nil {
        return fmt.Errorf("masscan not found. Install with: sudo apt install masscan -y")
    }
    return nil
}

func (s *Scanner) ScanRange(targetRange string, rate int, excludeInternal bool) ([]models.Target, error) {
    color.Cyan("[*] Starting masscan scan on %s at %d packets/sec", targetRange, rate)

    // Prepare output file
    rawOutput := filepath.Join(s.OutputDir, "masscan_raw.txt")
    filteredOutput := filepath.Join(s.OutputDir, "live_hosts.txt")

    // Build masscan command
    args := []string{
        "-p22",
        targetRange,
        "--rate", fmt.Sprintf("%d", rate),
        "-oJ", rawOutput,
    }

    // Add exclude file if exists
    if excludeInternal {
        if _, err := os.Stat(s.ExcludeFile); err == nil {
            args = append(args, "--excludefile", s.ExcludeFile)
            color.Yellow("[*] Using exclude file: %s", s.ExcludeFile)
        } else {
            // Default excludes
            args = append(args, "--exclude", "255.255.255.255")
            args = append(args, "--exclude", "10.0.0.0/8")
            args = append(args, "--exclude", "172.16.0.0/12")
            args = append(args, "--exclude", "192.168.0.0/16")
            args = append(args, "--exclude", "127.0.0.0/8")
        }
    }

    color.Blue("[>] Running: masscan %s", strings.Join(args, " "))

    // Execute masscan
    cmd := exec.Command(s.MasscanPath, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        return nil, fmt.Errorf("masscan failed: %v", err)
    }

    color.Green("[✓] Scan completed. Parsing results...")

    // Parse results
    targets, err := s.parseMasscanJSON(rawOutput)
    if err != nil {
        return nil, err
    }

    // Save filtered targets
    f, err := os.Create(filteredOutput)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    for _, t := range targets {
        f.WriteString(fmt.Sprintf("%s:%d\n", t.IP, t.Port))
    }

    color.Green("[✓] Found %d live SSH hosts", len(targets))
    color.Green("[✓] Saved to: %s", filteredOutput)

    return targets, nil
}

func (s *Scanner) parseMasscanJSON(filename string) ([]models.Target, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var targets []models.Target
    scanner := bufio.NewScanner(file)

    // Simple regex to extract IP from masscan JSON
    // Masscan JSON format: {"ip":"1.2.3.4", "ports": [{"port":22,"proto":"tcp","status":"open"}]}
    ipRegex := regexp.MustCompile(`"ip":"([0-9.]+)"`)
    portRegex := regexp.MustCompile(`"port":(\d+)`)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "open") && strings.Contains(line, "22") {
            ipMatch := ipRegex.FindStringSubmatch(line)
            portMatch := portRegex.FindStringSubmatch(line)

            if len(ipMatch) > 1 && len(portMatch) > 1 {
                // Filter out private IPs again for safety
                target := models.Target{
                    IP:   ipMatch[1],
                    Port: 22, // We only scan port 22
                }
                targets = append(targets, target)
            }
        }
    }

    return targets, scanner.Err()
}

// ScanByCountry - Scan IP ranges for specific country
func (s *Scanner) ScanByCountry(country string, rate int) ([]models.Target, error) {
    color.Cyan("[*] Loading IP ranges for country: %s", country)

    rangeFile := fmt.Sprintf("data/country_ranges/%s.txt", strings.ToLower(country))
    if _, err := os.Stat(rangeFile); err != nil {
        return nil, fmt.Errorf("country range file not found: %s", rangeFile)
   }

    file, err := os.Open(rangeFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var allTargets []models.Target
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        cidr := strings.TrimSpace(scanner.Text())
        if cidr == "" || strings.HasPrefix(cidr, "#") {
            continue
        }

        color.Yellow("[*] Scanning CIDR: %s", cidr)
        targets, err := s.ScanRange(cidr, rate, true)
        if err != nil {
            color.Red("[!] Error scanning %s: %v", cidr, err)
            continue
        }
        allTargets = append(allTargets, targets...)
    }

    return allTargets, nil
}
