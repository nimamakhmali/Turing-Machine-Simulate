# Turing Machine Simulator - Three Phases

A complete Turing Machine Simulator project with three distinct phases: Core, Web, and Desktop.

## Project Overview

This project demonstrates the evolution of a Turing Machine Simulator from a simple CLI application to a full desktop application.

## Phase Structure

### 📁 Phase 1: Core (`phase1-core/`)
**Pure Go implementation with CLI interface**
- Core Turing machine logic
- Command line interface
- Unit tests
- Example configurations
- No external dependencies

### 🌐 Phase 2: Web (`phase2-web/`)
**Web application with modern UI**
- Go backend with Gin framework
- HTML/CSS/JavaScript frontend
- RESTful API
- Interactive web interface
- Real-time simulation

### 🖥️ Phase 3: Desktop (`phase3-desktop/`)
**Desktop application with Electron**
- Cross-platform desktop app
- Native menu system
- Portable executable
- Integrated backend
- Professional UI

## Quick Start

### Phase 1: Core
```bash
cd phase1-core
go run main.go -config input/example_tm.json -tape "101"
```

### Phase 2: Web
```bash
cd phase2-web/backend
go run main.go
# Open http://localhost:8080
```

### Phase 3: Desktop
```bash
cd phase3-desktop/TuringMachineSimulator
TuringMachineSimulator.exe
```

## Features Comparison

| Feature | Core | Web | Desktop |
|---------|------|-----|---------|
| CLI Interface | ✅ | ❌ | ❌ |
| Web Interface | ❌ | ✅ | ✅ |
| Desktop App | ❌ | ❌ | ✅ |
| API Endpoints | ❌ | ✅ | ✅ |
| Menu System | ❌ | ❌ | ✅ |
| Portable | ✅ | ❌ | ✅ |
| Cross-platform | ✅ | ✅ | ✅ |

## Technology Stack

- **Core Logic**: Go
- **Web Framework**: Gin
- **Frontend**: HTML/CSS/JavaScript
- **Desktop**: Electron
- **Build**: electron-builder

## Requirements

- Go 1.16+
- Node.js 16+ (for desktop)
- Modern web browser (for web)

## Development

Each phase can be developed independently:

1. **Core**: Focus on algorithm and logic
2. **Web**: Add user interface and API
3. **Desktop**: Package as native application

## License

MIT License - see LICENSE file in each phase.

## Academic Context

This project demonstrates:
- Software architecture evolution
- Multi-platform development
- API design principles
- User interface design
- Application packaging and distribution