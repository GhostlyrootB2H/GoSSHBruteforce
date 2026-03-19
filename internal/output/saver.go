package output

import (
    "fmt"
    "os"
    "sync"
    "time"
    
    "github.com/GhostlyrootB2H/GoSSHBruteforce/internal/models"
)

type ResultSaver struct {
    filename string
    mu       sync.Mutex
    file     *os.File
}

func NewResultSaver(filename string) *ResultSaver {
    // Ensure directory exists
    os.MkdirAll("results", 0755)

    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil
    }

    return &ResultSaver{
        filename: filename,
        file:     file,
    }
}

func (r *ResultSaver) Save(cred models.Credential) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    line := fmt.Sprintf("[%s] %s - %s:%s\n", timestamp, cred.Target.String(), cred.Username, cred.Password)
    
    _, err := r.file.WriteString(line)
    return err
}

func (r *ResultSaver) Close() error {
    return r.file.Close()
}
