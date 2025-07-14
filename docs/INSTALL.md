# Bari$teuer Installation Guide

This guide will help you install and set up Bari$teuer on your system.

## 📋 System Requirements

- **Operating System**: Windows 10/11, macOS 10.15+, or Linux (64-bit)
- **Memory**: Minimum 4 GB RAM (8 GB recommended)
- **Disk Space**: At least 500 MB free space
- **Screen Resolution**: Minimum 1280x720 pixels

## 🚀 Quick Start (for Developers)

### 1. Clone the repository

```bash
git clone https://github.com/Christopher-Schulze/Baristeuer.git
cd Baristeuer
```

### 2. Install Dependencies

#### Using Bun (recommended):
```bash
bun install
```

#### Or using npm:
```bash
npm install
```

### 3. Set Up Environment Variables

Create a `.env` file in the root directory with the following content:

```env
DATABASE_URL="file:./dev.db"
NEXTAUTH_SECRET="your-secret-key-here"
NEXTAUTH_URL="http://localhost:3000"
```

### 4. Set Up Database

```bash
# Generate Prisma Client
bun run prisma:generate

# Run database migrations
bun run prisma:migrate

# Seed test data (optional)
bun run prisma:seed
```

### 5. Start the Application

#### Development Environment:
```bash
bun run dev
```

#### Production Build:
```bash
bun run build
bun run start
```

## 💻 Building the Desktop Version

### 1. Set Up Rust and Tauri

```bash
# Install Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Install Tauri CLI
cargo install tauri-cli
```

### 2. Build the Desktop Application

```bash
# Build the web application
bun run build

# Create the desktop version
bun run tauri:build
```

The final application will be available in the `src-tauri/target/release` directory.

## 🔧 Troubleshooting

### Common Issues

1. **Missing Dependencies**:
   - Ensure all system dependencies are installed
   - Run `bun install` again

2. **Database Issues**:
   - Delete `prisma/dev.db` and re-run migrations
   - Check database file permissions

3. **Build Errors**:
   - Delete `node_modules` and reinstall dependencies
   - Ensure you're using the latest version of Bun

## 📞 Support

For questions or issues, please contact:
- Email: support@baristeuer.de
- GitHub Issues: [https://github.com/Christopher-Schulze/Baristeuer/issues](https://github.com/Christopher-Schulze/Baristeuer/issues)
