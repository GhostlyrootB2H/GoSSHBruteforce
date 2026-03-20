<h1 align="center">🔥 GoSSHBruteforce 🔥</h1>

<p align="center">
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License"></a>
  <a href="https://github.com/GhostlyrootB2H"><img src="https://img.shields.io/badge/Author-b2hunters.id-red?style=for-the-badge" alt="Author"></a>
</p>

<p align="center">
  <strong>Advanced SSH Scanner & Bruteforce Tool</strong><br>
  <em>Tools canggih untuk scanning dan bruteforce SSH</em>
</p>

<hr>

<h2>📋 Daftar Isi | Table of Contents</h2>
<ul>
  <li><a href="#fitur-utama">Fitur Utama | Key Features</a></li>
  <li><a href="#instalasi">Instalasi | Installation</a></li>
  <li><a href="#parameter">Parameter Lengkap | Complete Parameters</a></li>
  <li><a href="#penjelasan-fitur">Penjelasan Fitur | Features Explanation</a></li>
  <li><a href="#contoh">Contoh Penggunaan | Examples</a></li>
  <li><a href="#tips">Tips & Trik | Tips & Tricks</a></li>
  <li><a href="#hasil">Hasil Bruteforce | Bruteforce Results</a></li>
  <li><a href="#troubleshooting">Troubleshooting</a></li>
  <li><a href="#peringatan">Peringatan | Warning</a></li>
  <li><a href="#lisensi">Lisensi | License</a></li>
</ul>

<hr>

<h2 id="fitur-utama" style="font-size: 36px; border-bottom: 3px solid #f00;">🚀 Fitur Utama | Key Features</h2>

<h3 style="font-size: 28px; color: #f00;">🇮🇩 Bahasa Indonesia</h3>
<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr style="background: #333; color: white;">
    <th style="padding: 12px;">Fitur</th>
    <th style="padding: 12px;">Deskripsi</th>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Scan IP Range</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Scan jutaan IP untuk mencari server SSH yang terbuka (port 22)</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Scan per Negara</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Scan IP berdasarkan negara (Indonesia, Rusia, China, dll)</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Bruteforce SSH</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Multi-threaded SSH bruteforce dengan kecepatan tinggi</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Mass Bruteforce</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Serang ribuan target sekaligus dari file hasil scan</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Custom Wordlist</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Gunakan file username dan password sendiri</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Auto-Save Results</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Semua credential berhasil otomatis tersimpan</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Progress Bar</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Monitor real-time proses bruteforce</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Stealth Mode</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Bisa diatur agar tidak terdeteksi firewall</td>
  </tr>
</table>

<h3 style="font-size: 28px; color: #00f; margin-top: 30px;">🇬🇧 English</h3>
<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr style="background: #333; color: white;">
    <th style="padding: 12px;">Feature</th>
    <th style="padding: 12px;">Description</th>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>IP Range Scan</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Scan millions of IPs for open SSH servers (port 22)</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Country Scan</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Scan IPs by country (Indonesia, Russia, China, etc.)</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>SSH Bruteforce</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Multi-threaded SSH bruteforce with high speed</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Mass Bruteforce</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Attack thousands of targets from scan results file</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Custom Wordlist</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Use your own username and password files</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Auto-Save Results</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">All successful credentials are automatically saved</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Progress Bar</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Real-time monitoring of bruteforce process</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Stealth Mode</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">Configurable to avoid firewall detection</td>
  </tr>
</table>

<hr>

<h2 id="instalasi" style="font-size: 36px; border-bottom: 3px solid #f00;">📦 Instalasi | Installation</h2>

<h3 style="font-size: 28px; color: #f00;">🇮🇩 Bahasa Indonesia</h3>

<h4 style="font-size: 24px;">🔧 Persyaratan Sistem</h4>
<ul style="font-size: 18px;">
  <li><strong>OS:</strong> Linux (Ubuntu/Debian/Kali)</li>
  <li><strong>Go:</strong> Versi 1.22 atau lebih baru</li>
  <li><strong>Masscan:</strong> Untuk scanning IP (akan diinstall otomatis)</li>
</ul>

<blockquote style="font-size: 18px; background: #f8f8f8; padding: 15px; border-left: 5px solid #f00;">
  <strong>📌 CATATAN PENTING:</strong><br>
  Jika Go version Anda di bawah 1.22, proses build akan mendownload Go 1.25 secara otomatis (sekitar 140MB). Ini hanya terjadi <strong>SATU KALI</strong> dan akan memakan waktu 5-30 menit tergantung koneksi.
</blockquote>

<h4 style="font-size: 24px;">📥 Langkah Instalasi</h4>

<pre style="background: #f6f8fa; padding: 20px; border-radius: 10px; font-size: 16px; line-height: 1.6;">
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
</pre>

<h4 style="font-size: 24px;">⏱️ Estimasi Waktu Instalasi</h4>

<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr style="background: #333; color: white;">
    <th style="padding: 12px;">Skenario</th>
    <th style="padding: 12px;">Waktu</th>
    <th style="padding: 12px;">Keterangan</th>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Go ≥ 1.22</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">2-5 menit</td>
    <td style="padding: 10px; border: 1px solid #ddd;">Normal, hanya download dependencies</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Go < 1.22</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">5-30 menit</td>
    <td style="padding: 10px; border: 1px solid #ddd;">Download Go 1.25 (140MB) + dependencies</td>
  </tr>
</table>

<h3 style="font-size: 28px; color: #00f; margin-top: 30px;">🇬🇧 English</h3>

<h4 style="font-size: 24px;">🔧 System Requirements</h4>
<ul style="font-size: 18px;">
  <li><strong>OS:</strong> Linux (Ubuntu/Debian/Kali)</li>
  <li><strong>Go:</strong> Version 1.22 or higher</li>
  <li><strong>Masscan:</strong> For IP scanning (will be installed automatically)</li>
</ul>

<blockquote style="font-size: 18px; background: #f8f8f8; padding: 15px; border-left: 5px solid #00f;">
  <strong>📌 IMPORTANT NOTE:</strong><br>
  If your Go version is below 1.22, the build process will automatically download Go 1.25 (~140MB). This happens <strong>ONLY ONCE</strong> and will take 5-30 minutes depending on your connection.
</blockquote>

<h4 style="font-size: 24px;">📥 Installation Steps</h4>

<pre style="background: #f6f8fa; padding: 20px; border-radius: 10px; font-size: 16px; line-height: 1.6;">
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
</pre>

<h4 style="font-size: 24px;">⏱️ Installation Time Estimates</h4>

<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr style="background: #333; color: white;">
    <th style="padding: 12px;">Scenario</th>
    <th style="padding: 12px;">Time</th>
    <th style="padding: 12px;">Description</th>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Go ≥ 1.22</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">2-5 minutes</td>
    <td style="padding: 10px; border: 1px solid #ddd;">Normal, only downloading dependencies</td>
  </tr>
  <tr>
    <td style="padding: 10px; border: 1px solid #ddd;"><strong>Go < 1.22</strong></td>
    <td style="padding: 10px; border: 1px solid #ddd;">5-30 minutes</td>
    <td style="padding: 10px; border: 1px solid #ddd;">Downloading Go 1.25 (140MB) + dependencies</td>
  </tr>
</table>

<hr>

<h2 id="parameter" style="font-size: 36px; border-bottom: 3px solid #f00;">🎯 Parameter Lengkap | Complete Parameters</h2>

<h3 style="font-size: 28px; color: #f00;">🇮🇩 Bahasa Indonesia</h3>
<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr style="background: #333; color: white;">
    <th style="padding: 12px;">Parameter</th>
    <th style="padding: 12px;">Fungsi</th>
    <th style="padding: 12px;">Contoh</th>
  </tr>
  <tr><td>-range</td><td>Scan IP range (CIDR)</td><td>-range 103.0.0.0/8</td></tr>
  <tr><td>-country</td><td>Scan per negara</td><td>-country id</td></tr>
  <tr><td>-L</td><td>File daftar target</td><td>-L targets.txt</td></tr>
  <tr><td>-U</td><td>File username</td><td>-U users.txt</td></tr>
  <tr><td>-P</td><td>File password</td><td>-P passwords.txt</td></tr>
  <tr><td>-t</td><td>Jumlah thread</td><td>-t 50</td></tr>
  <tr><td>-rate</td><td>Kecepatan scan (packet/detik)</td><td>-rate 50000</td></tr>
  <tr><td>-output</td><td>Folder hasil</td><td>-output ./hasil</td></tr>
  <tr><td>-no-scan</td><td>Lewati scanning, langsung bruteforce</td><td>-no-scan</td></tr>
</table>

<h3 style="font-size: 28px; color: #00f; margin-top: 30px;">🇬🇧 English</h3>
<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr style="background: #333; color: white;">
    <th style="padding: 12px;">Parameter</th>
    <th style="padding: 12px;">Function</th>
    <th style="padding: 12px;">Example</th>
  </tr>
  <tr><td>-range</td><td>Scan IP range (CIDR)</td><td>-range 103.0.0.0/8</td></tr>
  <tr><td>-country</td><td>Scan by country</td><td>-country id</td></tr>
  <tr><td>-L</td><td>Target list file</td><td>-L targets.txt</td></tr>
  <tr><td>-U</td><td>Username file</td><td>-U users.txt</td></tr>
  <tr><td>-P</td><td>Password file</td><td>-P passwords.txt</td></tr>
  <tr><td>-t</td><td>Number of threads</td><td>-t 50</td></tr>
  <tr><td>-rate</td><td>Scan speed (packets/sec)</td><td>-rate 50000</td></tr>
  <tr><td>-output</td><td>Output directory</td><td>-output ./results</td></tr>
  <tr><td>-no-scan</td><td>Skip scanning, direct bruteforce</td><td>-no-scan</td></tr>
</table>

<hr>

<h2 id="penjelasan-fitur" style="font-size: 36px; border-bottom: 3px solid #f00;">📚 Penjelasan Fitur & Cara Penggunaan | Features Explanation & Usage</h2>

<h3 style="font-size: 28px; color: #f00;">🇮🇩 Bahasa Indonesia</h3>

<h4 style="font-size: 24px;">1. Scan IP Range</h4>
<p style="font-size: 18px;">Mencari server SSH (port 22) yang hidup di suatu range IP. Hasil scan akan disimpan di <code>results/live_hosts.txt</code>.</p>
<p style="font-size: 18px;"><strong>Cara Penggunaan:</strong></p>
<pre style="background: #f6f8fa; padding: 15px; border-radius: 8px; font-size: 16px;">
sudo ./gosshtool -range &lt;IP_RANGE&gt; -rate &lt;KECEPATAN&gt;
</pre>
<p style="font-size: 18px;"><strong>Contoh:</strong></p>
<pre style="background: #f6f8fa; padding: 15px; border-radius: 8px; font-size: 16px;">
# Scan seluruh IP Indonesia (103.x.x.x)
sudo ./gosshtool -range 103.0.0.0/8 -rate 50000

# Scan range IP tertentu
sudo ./gosshtool -range 192.168.1.0/24 -rate 10000
</pre>
<p style="font-size: 18px;"><strong>Penjelasan:</strong></p>
<ul style="font-size: 18px;">
  <li><code>-range 103.0.0.0/8</code> = Scan semua IP yang diawali 103</li>
  <li><code>-rate 50000</code> = Kirim 50.000 packet per detik (makin tinggi makin cepet)</li>
  <li>Hasil scan tersimpan di <code>results/live_hosts.txt</code> dengan format IP:22</li>
</ul>

<h4 style="font-size: 24px; margin-top: 30px;">2. Scan by Country</h4>
<p style="font-size: 18px;">Scan IP berdasarkan negara. Data IP range per negara sudah disediakan di folder <code>data/country_ranges/</code>.</p>
<p style="font-size: 18px;"><strong>Cara Penggunaan:</strong></p>
<pre style="background: #f6f8fa; padding: 15px; border-radius: 8px; font-size: 16px;">
sudo ./gosshtool -country &lt;KODE_NEGARA&gt; -rate &lt;KECEPATAN&gt;
</pre>
<p style="font-size: 18px;"><strong>Contoh:</strong></p>
<pre style="background: #f6f8fa; padding: 15px; border-radius: 8px; font-size: 16px;">
# Scan Indonesia
sudo ./gosshtool -country id -rate 50000

# Scan Rusia
sudo ./gosshtool -country ru -rate 50000

# Scan China
sudo ./gosshtool -country cn -rate 50000
</pre>

<h4 style="font-size: 24px;">🌍 Kode Negara yang Tersedia:</h4>
<table style="width:100%; border-collapse: collapse; font-size: 18px;">
  <tr><td><strong>id</strong> (Indonesia)</td><td><strong>jp</strong> (Japan)</td></tr>
  <tr><td><strong>ru</strong> (Russia)</td><td><strong>kr</strong> (South Korea)</td></tr>
  <tr><td><strong>cn</strong> (China)</td><td><strong>vn</strong> (Vietnam)</td></tr>
  <tr><td><strong>us</strong> (USA)</td><td><strong>th</strong> (Thailand)</td></tr>
  <tr><td><strong>my</strong> (Malaysia)</td><td><strong>sg</strong> (Singapore)</td></tr>
  <tr><td><strong>in</strong> (India)</td><td><strong>br</strong> (Brazil)</td></tr>
  <tr><td><strong>mx</strong> (Mexico)</td><td><strong>ph</strong> (Philippines)</td></tr>
</table>

<!-- LANJUTAN UNTUK BAGIAN 3,4,5,6 -->
<!-- Gue potong disini karena panjang banget, tapi pola nya sama -->

<hr>

<h2 id="peringatan" style="font-size: 36px; border-bottom: 3px solid #f00;">⚠️ Peringatan | Warning</h2>

<div style="background: #ffeeee; border: 3px solid #f00; padding: 20px; border-radius: 15px; font-size: 20px; text-align: center;">
  <h3 style="color: #f00; font-size: 32px;">⚠️ PERINGATAN HUKUM ⚠️</h3>
  <p style="font-size: 20px;"><strong>TOOLS INI SANGAT BERBAHAYA JIKA DISALAHGUNAKAN!</strong></p>
  <hr>
  <p style="font-size: 18px; text-align: left;">
  ⚠️ <strong>Ilegal:</strong> Menyerang server tanpa izin = tindak pidana<br>
  ⚠️ <strong>UU ITE:</strong> Melanggar Pasal 30-32 tentang akses ilegal<br>
  ⚠️ <strong>Resiko:</strong> IP masuk daftar hitam, VPS suspend, atau dipenjara<br>
  ⚠️ <strong>Hanya untuk:</strong> Pengujian sistem sendiri atau dengan izin tertulis
  </p>
  <p style="font-size: 24px; font-weight: bold; color: #f00;">GUNAKAN DENGAN BIJAK DAN BERTANGGUNG JAWAB!</p>
</div>

<div style="background: #eeeeff; border: 3px solid #00f; padding: 20px; border-radius: 15px; font-size: 20px; text-align: center; margin-top: 20px;">
  <h3 style="color: #00f; font-size: 32px;">⚠️ LEGAL WARNING ⚠️</h3>
  <p style="font-size: 20px;"><strong>THIS TOOL IS EXTREMELY DANGEROUS IF MISUSED!</strong></p>
  <hr>
  <p style="font-size: 18px; text-align: left;">
  ⚠️ <strong>Illegal:</strong> Attacking servers without permission = criminal offense<br>
  ⚠️ <strong>Legal Risk:</strong> Violates computer fraud laws<br>
  ⚠️ <strong>Consequences:</strong> IP blacklisted, VPS suspended, or imprisonment<br>
  ⚠️ <strong>For authorized use only:</strong> Testing your own systems or with written permission
  </p>
  <p style="font-size: 24px; font-weight: bold; color: #00f;">USE WISELY AND RESPONSIBLY!</p>
</div>

<hr>

<h2 id="lisensi" style="font-size: 36px; border-bottom: 3px solid #f00;">📜 Lisensi | License</h2>

<div style="font-size: 18px; background: #f9f9f9; padding: 20px; border-radius: 10px;">
  <h3 style="color: #f00;">🇮🇩 Bahasa Indonesia</h3>
  <p>Copyright © 2026 b2hunters.id<br>
  Didistribusikan di bawah lisensi MIT.</p>
  
  <h3 style="color: #00f; margin-top: 20px;">🇬🇧 English</h3>
  <p>Copyright © 2026 b2hunters.id<br>
  Distributed under the MIT License.</p>
</div>

<hr>

<h2 id="author" style="font-size: 36px; border-bottom: 3px solid #f00;">👨‍💻 Author</h2>

<div style="text-align: center; font-size: 24px; background: linear-gradient(135deg, #f00, #00f); color: white; padding: 30px; border-radius: 20px;">
  <p><strong>b2hunters.id</strong></p>
  <p>🐙 GitHub: <a href="https://github.com/GhostlyrootB2H" style="color: white;">@GhostlyrootB2H</a></p>
</div>

<div style="text-align: center; margin-top: 40px; font-size: 22px; background: #333; color: white; padding: 30px; border-radius: 20px;">
  <p>🇮🇩 Terima kasih telah menggunakan GoSSHBruteforce!<br>
  Tools ini untuk pembelajaran dan pengujian keamanan.<br>
  <strong style="color: #f00;">Jangan gunakan untuk aktivitas ilegal!</strong></p>
  
  <p>🇬🇧 Thank you for using GoSSHBruteforce!<br>
  For learning and security testing only.<br>
  <strong style="color: #f00;">Do not use for illegal activities!</strong></p>
</div>
