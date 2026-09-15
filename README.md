# CVE2fa

Kendi sunucunuzda barındırabileceğiniz (self-hosted), web tabanlı bir TOTP (Zaman Tabanlı Tek Seferlik Şifre) yöneticisi. Go ve Vue 3 ile geliştirilmiştir.

## Özellikler

- Birden fazla hesap için TOTP kodları oluşturma
- Argon2id anahtar türetme işlevi ile AES-256-GCM şifreli yedekleme/geri yükleme
- Şifreli içe/dışa aktarma (.2fadata)
- Koyu tema destekli arayüz (DaisyUI + Tailwind CSS)
- JWT oturum yönetimi
- SQLite veritabanı (tek bir çalıştırılabilir dosya olarak konuşlandırma)

## Gereksinimler

- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+ and npm

## Kurulum

```bash
# Depoyu klonlayın
git clone https://github.com/cveburak/CVE2FA.git
cd CVE2FA

# Ön yüzü (frontend) derleyin
cd web && npm install && npm run build && cd ..

# Sunucuyu derleyin
go build -o CVE2fa .

# Çalıştırın
./CVE2fa
```

Varsayılan giriş bilgileri: `admin` / `admin`  
Sunucu: `http://localhost:2026`

## API

| Method | Path                    | Description        |
| ------ | ----------------------- | ------------------ |
| POST   | /api/auth/login         | Authenticate       |
| POST   | /api/auth/logout        | End session        |
| PUT    | /api/auth/password      | Change password    |
| GET    | /api/auth/me            | Current user       |
| GET    | /api/accounts           | List accounts      |
| POST   | /api/accounts           | Add account        |
| DELETE | /api/accounts/{id}      | Remove account     |
| GET    | /api/accounts/{id}/code | Generate TOTP code |
| POST   | /api/export             | Export encrypted   |
| POST   | /api/import             | Import encrypted   |

## Security

- Tüm hassas veriler veri tabanına kaydedilmeden önce AES-256-GCM ile şifrelenir
- Şifreler Argon2id ile özetlenir (hash)
- Oturum jetonları (token) HMAC-SHA256 ile imzalanır
- Yedekleme dosyaları, ana şifrenizden türetilen bir anahtar ile şifrelenir
