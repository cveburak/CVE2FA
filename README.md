<div align="center">

# CVE2fa — Local TOTP / 2FA Authenticator Manager

**A free, open-source, local-first two-factor authentication (2FA) manager and TOTP authenticator app that runs on your own computer.**
Built with Go and Vue 3 — a privacy-friendly alternative to Authy, Google Authenticator, and other cloud 2FA apps.

[![Go Report Card](https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go)](https://go.dev/dl/)
[![Vue 3](https://img.shields.io/badge/Vue-3-42b883?logo=vuedotjs)](https://vuejs.org/)
[![SQLite](https://img.shields.io/badge/Database-SQLite-003B57?logo=sqlite)](https://www.sqlite.org/)
[![Self Hosted](https://img.shields.io/badge/Local--First-Yes-success)](#)
[![License](https://img.shields.io/badge/License-Unspecified-lightgrey)](#license)

</div>

## What is CVE2fa?

CVE2fa is a **local, self-hosted two-factor authentication (2FA) manager** for generating **TOTP (Time-based One-Time Password)** codes — the same standard used by Google Authenticator, Authy, and Microsoft Authenticator. Instead of trusting a third-party cloud app with your 2FA secrets, CVE2fa runs entirely on your own computer, with encrypted local storage and encrypted backups.

It's a single self-contained binary with an embedded Vue 3 web UI, backed by a local SQLite file — no external database, no cloud account, no telemetry.

## Why keep your 2FA codes local?

- **Full control** — your TOTP secrets never leave your own machine.
- **No vendor lock-in** — encrypted export/import means your data is portable.
- **Open source** — inspect exactly how your secrets are stored and encrypted.
- **Lightweight** — one Go binary, one local SQLite file, no external dependencies at runtime.

## ✨ Features

- 🔑 Generate TOTP codes for unlimited accounts (Google, GitHub, AWS, etc.)
- 🔒 AES-256-GCM encrypted backup & restore, keys derived with Argon2id
- 📦 Encrypted import/export in a portable `.2fadata` format
- 🌙 Dark-mode UI built with DaisyUI + Tailwind CSS
- 🪪 JWT-based session authentication
- 🗄️ SQLite database — deploys as a single executable file
- 🖥️ Runs anywhere: Linux, macOS, Windows

## 📋 Requirements

- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+ and npm

## 🚀 Quick Start / Installation

```bash
# Clone the repository
git clone https://github.com/cveburak/CVE2FA.git
cd CVE2FA

# Build the frontend
cd web && npm install && npm run build && cd ..

# Build the server
go build -o CVE2fa .

# Run it
./CVE2fa
```

Default login: `admin` / `admin` (you'll be asked to change it on first login)
Runs locally at: `http://localhost:2026`

## 📖 API Reference

| Method | Path                    | Description         |
| ------ | ----------------------- | -------------------- |
| POST   | /api/auth/login         | Authenticate          |
| POST   | /api/auth/logout        | End session            |
| PUT    | /api/auth/password      | Change password        |
| GET    | /api/auth/me            | Current user            |
| GET    | /api/accounts           | List accounts             |
| POST   | /api/accounts           | Add account                |
| DELETE | /api/accounts/{id}      | Remove account               |
| GET    | /api/accounts/{id}/code | Generate TOTP code             |
| POST   | /api/export             | Export encrypted backup          |
| POST   | /api/import             | Import encrypted backup            |

## 🔐 Security

- All sensitive account data is encrypted with **AES-256-GCM** before it's written to the database
- Passwords are hashed with **Argon2id**
- Session tokens are signed with **HMAC-SHA256** (JWT)
- Backup files are encrypted with a key derived from your master password

## 🗺️ Tech Stack

| Layer     | Technology              |
| --------- | ------------------------ |
| Backend   | Go                          |
| Frontend  | Vue 3, Tailwind CSS, DaisyUI  |
| Database  | SQLite                         |
| Auth      | JWT (HMAC-SHA256)                |
| Encryption| AES-256-GCM, Argon2id              |

## 🤝 Contributing

Issues and pull requests are welcome. If you run into a bug or have a feature request, please open an issue on GitHub.

## 📄 License

No license file has been added to this repository yet — until one is added, all rights are reserved by default. Add a `LICENSE` file (e.g. MIT, Apache-2.0) to make reuse terms explicit.

---

<div align="center">

**Keywords:** local 2FA manager, self-hosted TOTP manager, open source authenticator app, Google Authenticator alternative, Authy alternative, Go 2FA app, Vue 3 authenticator, offline authenticator app

</div>
