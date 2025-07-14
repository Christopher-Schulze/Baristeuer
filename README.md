> ⚠️ **IMPORTANT NOTICE**
> 
> This project is currently **under active development** and **not yet functional**. 
> We're working hard to implement all features. Check back soon for updates!
> 
> ---

<div align="center">
  <img src="logo.png" alt="Bari$teuer Logo" width="300">
  <h1>Bari$teuer</h1>
  <p>Tax Management for Non-Profit Organizations - Simple, Secure, and Efficient</p>
</div>

Bari$teuer is a cross-platform desktop application designed to assist non-profit organizations in Germany with their tax reporting. Built with Next.js, TypeScript, and Tauri for a native desktop experience.

## 🚀 Tech Stack

- **Frontend**: Next.js 15.3.5 with TypeScript 5.3.3
- **Runtime**: Bun (instead of Node.js)
- **Build Tool**: Turbopack with SWC
- **UI Framework**: HeroUI 2.7.11 with Tailwind CSS 4.1.11
- **Desktop**: Tauri 2.0.0-rc for cross-platform native apps
- **API**: tRPC 11.4.3 for type-safe APIs
- **Database**: Prisma 6.11.1 with SQLite
- **Authentication**: NextAuth.js 4.24.11
- **State Management**: Jotai with Zod validation
- **Testing**: Vitest for unit and integration tests

## ✨ Features

- **Multi-step Tax Interview**: Guided process for tax declaration
- **Real-time Calculations**: Automatic tax computations (income, trade, solidarity, church tax)
- **PDF Generation**: Professional tax reports and exports
- **Offline Capability**: Works without internet connection
- **Type Safety**: Full TypeScript coverage with runtime validation
- **Modern UI**: Responsive design with HeroUI components
- **Cross-Platform**: Native desktop app for macOS, Windows, and Linux

## 🛠️ Development

### 📋 Voraussetzungen

- [Bun](https://bun.sh/) (neueste Version)
- [Rust](https://rustup.rs/) (für Tauri)
- [Node.js 20+](https://nodejs.org/) (Fallback)

### 🚀 Schnellstart

```bash
# Repository klonen
git clone https://github.com/Christopher-Schulze/Baristeuer.git
cd Baristeuer

# Abhängigkeiten installieren
bun install

# Datenbank einrichten
bun run prisma:generate
bun run prisma:migrate

# Entwicklungsserver starten
bun run dev
```

### Building

```bash
# Build web application
bun run build

# Build desktop application
bun run tauri build
```

### Testing

```bash
# Run tests
bun run test

# Run tests with coverage
bun run test:coverage
```

## 📖 Documentation

For detailed documentation, architecture overview, and usage instructions, see:
- [docs/DOCUMENTATION.md](docs/DOCUMENTATION.md) - Complete project documentation
- [docs/CHANGELOG.md](docs/CHANGELOG.md) - Version history and changes

## 🏗️ Project Structure

```
├── src/                    # Next.js application source
│   ├── app/               # App router pages
│   ├── components/        # React components
│   ├── lib/              # Utilities and configurations
│   └── server/           # tRPC API routes
├── src-tauri/            # Tauri desktop application
├── prisma/               # Database schema and migrations
├── docs/                 # Project documentation
└── public/               # Static assets
```

## 🔒 Security

- Type-safe API with runtime validation
- Local SQLite database for data privacy
- No external data transmission

---

_This project is for internal use and is not open for contributions._
