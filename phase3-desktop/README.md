# Phase 3: Desktop Application

This is the desktop application version of the Turing Machine Simulator built with Electron.

## Structure

```
phase3-desktop/
├── main.js                    # Electron main process
├── package.json               # Node.js dependencies
├── assets/                    # App assets
│   └── icon.svg              # App icon
├── TuringMachineSimulator/    # Complete app package
│   ├── main.js               # Electron main process
│   ├── package.json          # Dependencies
│   ├── frontend/             # Web interface
│   ├── backend/              # Go server
│   ├── turing/               # Core logic
│   ├── input/                # Examples
│   ├── node_modules/         # Node.js dependencies
│   ├── TuringMachineSimulator.exe # Executable
│   ├── launcher.bat          # Alternative launcher
│   ├── install.bat           # Install script
│   ├── create-shortcut.bat   # Desktop shortcut
│   └── README.txt            # Usage instructions
└── README-DESKTOP.md         # Detailed documentation
```

## Features

- **Desktop application** - Native desktop experience
- **Cross-platform** - Windows, Mac, Linux support
- **Menu system** - File, View, Help menus
- **Portable** - No installation required
- **Auto-start server** - Integrated Go backend
- **Modern UI** - Electron-based interface

## Usage

### Development

```bash
# Install dependencies
npm install

# Run in development mode
npm run dev

# Run in production mode
npm start
```

### Distribution

```bash
# Build executable
npm run build

# Or use the pre-built package
cd TuringMachineSimulator
TuringMachineSimulator.exe
```

## Requirements

- Node.js 16 or higher
- Go 1.16 or higher

## Installation

```bash
npm install
```

## Technologies

- **Framework**: Electron
- **Backend**: Go + Gin
- **Frontend**: HTML + CSS + JavaScript
- **Build**: electron-builder

## Executable

The `TuringMachineSimulator.exe` file is a self-contained executable that:
- Checks for Node.js and Go
- Installs dependencies automatically
- Starts the Electron application
- Launches the Go backend server

This phase provides a complete desktop application experience.
