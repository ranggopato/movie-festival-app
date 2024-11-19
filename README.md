
# Movie Festival API

API ini berisi API untuk mengelola sistem Movie Festival. Admin bisa mengatur film, melihat statistik tontonan, dan genre terpopuler. Pengguna bisa cari, nonton, dan vote film favorit mereka. API juga dilengkapi autentikasi dan pelacakan aktivitas menonton.

## Fitur Utama:

### Admin APIs (harus login):
- **Tambah, edit, dan upload film.**
- **Lihat film dan genre terpopuler.**
- **Track viewership durasi tontonan.**

### Movie APIs - (tidak harus login):
- **Lihat dan cari film (paginasi).**
- **Cari film berdasarkan title, description, artists, dan genres.**
- **Movie viewership.**

### Sistem Voting - (harus login):
- **Vote/unvote film.**
- **Lihat daftar film yang di-vote.**

### Autentikasi:
- **Registrasi, login.**
- **Akses endpoint protected dengan token.**

## Cara Pakai:

1. **Set Variabel:**
   Gunakan `http://localhost:8080` untuk URL server dan token login otomatis diterapkan saat Anda login.

2. **Autentikasi:**
   - Login untuk mendapatkan token autentikasi. Gunakan token untuk mengakses endpoint yang memerlukan autentikasi.

3. **Admin vs. User:**
   - Beberapa fitur hanya dapat diakses oleh admin. Pastikan Anda login dengan akun admin untuk menggunakan fitur admin.

## Cara Menjalankan Aplikasi:

1. **Pastikan Go sudah terpasang** di komputer Anda. Jika belum, ikuti petunjuk di [situs resmi Go](https://golang.org/doc/install) untuk instalasi.

2. **Clone repository:**

   ```bash
   git clone https://github.com/ranggopato/movie-festival-app.git
   cd movie-festival-api
   ```

3. **Instal dependensi:**
   Di dalam folder proyek, jalankan perintah:

   ```bash
   go mod tidy
   ```

4. **Menjalankan aplikasi:**
   Untuk menjalankan aplikasi, cukup ketik perintah berikut:

   ```bash
   go run main.go
   ```

5. **Aplikasi berjalan pada port 8080.** Anda dapat mengakses API melalui `http://localhost:8080`.

6. **Login untuk mendapatkan token:**
   Gunakan endpoint `/login` untuk mendapatkan token autentikasi, dan gunakan token ini untuk mengakses endpoint lainnya yang memerlukan autentikasi.

---

### Struktur Proyek:
- `config/`: Koneksi ke database dan pengaturan lainnya.
- `controllers/`: Menangani logika untuk request dari pengguna.
- `middlewares/`: Middleware untuk autentikasi dan otorisasi.
- `models/`: Struktur data untuk entitas (misalnya, film, pengguna, dll).
- `repositories/`: Akses ke database.
- `services/`: Logika bisnis dan pengelolaan entitas.
- `utils/`: Fungsi tambahan seperti hashing password, pembuatan JWT, dll.

## Teknologi yang Digunakan:
- **Go**: Bahasa pemrograman utama.
- **GORM**: ORM untuk interaksi dengan database SQLite.
- **SQLite**: Database yang digunakan untuk menyimpan data.
- **JWT**: Untuk autentikasi dan otorisasi pengguna.

---

