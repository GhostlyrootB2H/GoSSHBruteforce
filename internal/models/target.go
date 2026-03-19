package models

import (
    "fmt"
    "net"
    "strconv"
    "strings"
)

type Target struct {
    IP   string
    Port int
}

func (t Target) String() string {
    return fmt.Sprintf("%s:%d", t.IP, t.Port)
}

func ParseTarget(line string) (*Target, error) {
    line = strings.TrimSpace(line)
    if line == "" {
        return nil, fmt.Errorf("empty line")
    }

    var ip string
    var port int = 22 // default SSH port

    if strings.Contains(line, ":") {
        parts := strings.Split(line, ":")
        if len(parts) != 2 {
            return nil, fmt.Errorf("invalid format: %s", line)
        }
        ip = parts[0]
        p, err := strconv.Atoi(parts[1])
        if err != nil {
            return nil, fmt.Errorf("invalid port: %s", parts[1])
        }
        port = p
    } else {
        ip = line
    }

    // Validate IP
    if net.ParseIP(ip) == nil {
        return nil, fmt.Errorf("invalid IP: %s", ip)
    }

    return &Target{IP: ip, Port: port}, nil
}

type Credential struct {
    Target   Target
    Username string
    Password string
}

func (c Credential) String() string {
    return fmt.Sprintf("%s - %s:%s", c.Target.String(), c.Username, c.Password)
}
