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
