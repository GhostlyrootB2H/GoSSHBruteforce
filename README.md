# GoSSHBruteforce 🔥

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Author](https://img.shields.io/badge/Author-b2hunters.id-red.svg)](https://github.com/b2hunters)

> Advanced SSH scanning and bruteforcing tool - Scan entire internet, crack SSH passwords, harvest VPS

**Author: b2hunters.id**

## ⚡ Features

### 🎯 Scanning Phase
- **Masscan Integration**: Scan entire internet (`0.0.0.0/0`) at 100k+ packets/sec
- **Country-Specific Scanning**: Pre-configured IP ranges for 15+ countries
- **Auto-Filtering**: Automatically excludes private, reserved, and government IPs
- **Multi-Format Input**: CIDR ranges, single IPs, or file-based targets

### 💀 Bruteforce Phase
- **Goroutine-Based Concurrency**: Thousands of simultaneous connections
- **Legacy Algorithm Support**: Works with old SSH servers (diffie-hellman-group1-sha1, etc.)
- **Intelligent Rate Limiting**: Prevents connection flooding and IP bans
- **Progress Bar**: Real-time visualization of attack progress
- **Stats Monitoring**: Attempts/sec, success rate, elapsed time
- **Auto-Save**: All credentials automatically saved

## 🚀 Quick Start

```bash
# Clone and build
git clone https://github.com/b2hunters/GoSSHBruteforce.git
cd GoSSHBruteforce
make install-deps
make build

# Run as root (required for masscan)
sudo ./build/gosshtool -country id -t 50
