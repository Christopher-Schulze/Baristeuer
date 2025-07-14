# Baristeuer Documentation

## Project Overview

Baristeuer is a professional tax declaration software designed for clubs and non-profit organizations in Germany. It simplifies the complex process of creating and managing tax declarations by providing an intuitive step-by-step interview process that guides users through all required information for different tax areas.

## Architecture

The application is built using the modern T3-Stack with additional enhancements:

### Core Technologies
- **Next.js 15.3.5**: React framework with App Router, Turbopack, and static export capabilities
- **TypeScript 5.3.3**: Full type safety across the entire application
- **Bun**: Modern JavaScript runtime and package manager for improved performance
- **tRPC 11.4.3**: End-to-end typesafe API layer with React Query integration
- **Prisma 6.11.1**: Type-safe database ORM with SQLite backend
- **NextAuth.js 4.24.11**: Comprehensive authentication solution
- **Tailwind CSS 4.1.11**: Utility-first CSS framework with JIT compilation
- **HeroUI 2.7.11**: Modern React component library for consistent UI
- **Tauri 2.0.0-rc**: Cross-platform desktop application framework

### State Management & Data Flow
- **Jotai 2.12.5**: Atomic state management for form data
- **Zod 3.25.76**: Runtime type validation and schema definition
- **React Query**: Server state management and caching
- **Superjson**: Enhanced JSON serialization for tRPC

### Development & Testing
- **Vitest 3.2.4**: Fast unit testing framework
- **ESLint & TypeScript ESLint**: Code quality and consistency
- **PostCSS**: CSS processing and optimization

## Key Features

### Tax Declaration Management
- **Multi-Step Interview Process**: Five comprehensive steps covering all tax areas:
  1. **Grunddaten (Basic Data)**: Organization details, address, tax number
  2. **Ideeller Bereich (Ideal Sector)**: Non-profit activities and related finances
  3. **Vermögensverwaltung (Asset Management)**: Investment income and expenses
  4. **Zweckbetrieb (Purpose-Related Business)**: Mission-related commercial activities
  5. **Wirtschaftlicher Geschäftsbetrieb (Commercial Business)**: Taxable commercial activities

### Advanced Functionality
- **Dynamic Form Validation**: Real-time validation with detailed error messages
- **Automatic Tax Calculations**: Built-in tax engine with German tax law compliance
- **PDF Generation**: Professional tax declaration documents using pdf-lib
- **Data Persistence**: Automatic saving with optimistic updates
- **Responsive Design**: Mobile-first approach with dark mode support

### Technical Features
- **Type-Safe API**: Full TypeScript coverage from database to UI
- **Real-time Updates**: Optimistic UI updates with error handling
- **Cross-Platform Desktop App**: Native performance with web technologies
- **Offline Capability**: Local SQLite database for offline functionality
- **Security**: Secure authentication with session management

## Project Structure

```
src/
├── app/                          # Next.js App Router
│   ├── globals.css              # Global styles
│   ├── layout.tsx               # Root layout with providers
│   ├── page.tsx                 # Landing page
│   └── steuererklaerung/
│       └── neu/                 # New tax declaration page
├── components/
│   ├── providers/               # React context providers
│   ├── steuer/                  # Tax-related components
│   │   ├── SteuerFormular.tsx   # Main form component
│   │   └── interview-steps/     # Step-by-step interview components
│   └── ui/                      # Reusable UI components
├── lib/
│   ├── auth/                    # Authentication configuration
│   ├── pdf/                     # PDF generation utilities
│   ├── steuer/                  # Tax calculation engine
│   │   ├── engine.ts            # Core tax calculation logic
│   │   ├── formState.ts         # Jotai atoms for form state
│   │   └── steuerService.ts     # Tax service layer
│   ├── trpc/                    # tRPC client and server setup
│   ├── prisma.ts                # Database client
│   └── utils.ts                 # Utility functions
├── server/
│   ├── context.ts               # tRPC context creation
│   ├── routers/                 # API route handlers
│   └── trpc.ts                  # tRPC server configuration
├── types/                       # TypeScript type definitions
├── validations/                 # Zod schemas
└── __tests__/                   # Test files

src-tauri/                       # Tauri desktop app configuration
├── src/                         # Rust source code
├── icons/                       # Application icons
├── Cargo.toml                   # Rust dependencies
└── tauri.conf.json             # Tauri configuration

prisma/
├── schema.prisma               # Database schema
└── migrations/                 # Database migrations

docs/                           # Project documentation
├── DOCUMENTATION.md            # This file
├── CHANGELOG.md               # Version history
└── LICENSE                    # MIT License
```

## Database Schema

The application uses SQLite with Prisma ORM. Key models include:

### User Management
- **User**: User accounts with authentication data
- **Account**: OAuth provider accounts
- **Session**: User sessions
- **VerificationToken**: Email verification tokens

### Tax Data
- **Verein**: Organization/club information
  - Basic details (name, address, tax number)
  - Tax settings (Hebesatz, Kleinunternehmerregelung)
  - Financial office address
- **steuererklaerung**: Tax declarations
  - Year-specific tax data
  - Four tax areas as JSON fields:
    - `ideellerBereich`: Non-profit activities
    - `vermoegensverwaltung`: Asset management
    - `zweckbetrieb`: Purpose-related business
    - `wirtschaftlicherGeschaeftsbetrieb`: Commercial business

## Tax Calculation Engine

The `SteuerEngine` class implements German tax law calculations:

### Core Calculations
- **Income Tax (Einkommensteuer)**: Progressive tax rates
- **Trade Tax (Gewerbesteuer)**: Municipal business tax with exemptions
- **Solidarity Surcharge (Solidaritätszuschlag)**: Additional federal tax
- **Church Tax (Kirchensteuer)**: Optional religious tax

### Tax Constants (2024)
- Income tax exemption: €11,604
- Trade tax exemption: €5,000
- Solidarity surcharge rate: 5.5%
- Standard municipal multiplier: 400%

## Development Guidelines

### Code Quality
- **TypeScript**: Strict mode enabled, no `any` types
- **ESLint**: Enforced code style and best practices
- **Testing**: Unit tests for business logic with Vitest
- **Validation**: Zod schemas for all data structures

### State Management
- **Form State**: Jotai atoms for reactive form data
- **Server State**: React Query for API data caching
- **Validation**: Real-time validation with error states

### Performance
- **Turbopack**: Fast development builds
- **Static Export**: Optimized production builds
- **Code Splitting**: Automatic route-based splitting
- **Optimistic Updates**: Immediate UI feedback

## Build and Deployment

### Development
```bash
bun install                    # Install dependencies
bun run dev                   # Start development server
bun run prisma:migrate        # Run database migrations
bun run test                  # Run unit tests
```

### Production
```bash
bun run build                 # Build Next.js application
bun run tauri:build          # Build desktop application
```

### Desktop Application
The Tauri build creates native installers for:
- **macOS**: `.dmg` installer
- **Windows**: `.msi` installer
- **Linux**: `.deb` and `.rpm` packages

## Security Considerations

- **Authentication**: Secure session management with NextAuth.js
- **Data Validation**: Server-side validation with Zod schemas
- **SQL Injection**: Protected by Prisma ORM
- **XSS Protection**: React's built-in XSS protection
- **CSRF Protection**: Built into NextAuth.js
- **Local Data**: SQLite database for sensitive tax data
