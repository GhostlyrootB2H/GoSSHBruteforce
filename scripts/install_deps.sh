#!/bin/bash

echo "[*] Installing dependencies for GoSSHBruteforce..."

# Install masscan
if ! command -v masscan &> /dev/null; then
    echo "[*] Installing masscan..."
    sudo apt update
    sudo apt install -y masscan git build-essential
else
    echo "[✓] masscan already installed"
fi

# Install Go if not present
if ! command -v go &> /dev/null; then
    echo "[*] Installing Go..."
    wget https://go.dev/dl/go1.21.13.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.21.13.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    source ~/.bashrc
else
    echo "[✓] Go already installed"
fi

# Create directory structure
mkdir -p results data/country_ranges configs

echo "[✓] Installation complete"
echo "[*] Next steps:"
echo "    1. Download IP ranges for countries from APNIC/ARIN/RIPE"
echo "    2. Run: go mod tidy"
echo "    3. Build: go build -o gosshtool cmd/gosshtool/main.go"
