# Phase 2: Web Application

This is the web-based version of the Turing Machine Simulator.

## Structure

```
phase2-web/
├── backend/              # Go web server
│   ├── main.go          # Web server with Gin framework
│   ├── go.mod           # Go dependencies
│   └── go.sum           # Go dependencies
├── frontend/            # Web interface
│   ├── templates/
│   │   └── index.html   # Main HTML page
│   └── static/
│       ├── css/
│       │   └── style.css # Styling
│       └── js/
│           └── app.js   # JavaScript logic
├── turing/              # Core Turing machine logic
│   ├── machine.go       # Machine implementation
│   ├── types.go         # Data structures
│   └── utils.go         # Utility functions
├── input/               # Example configurations
│   ├── example_tm.json  # Simple example
│   └── palindrome_tm.json # Palindrome checker
└── README.md            # This file
```

## Features

- **Web interface** - Modern, responsive UI
- **Real-time simulation** - Step-by-step execution
- **Interactive tape display** - Visual tape representation
- **Example library** - Built-in examples
- **JSON editor** - Custom machine definitions
- **RESTful API** - Backend API endpoints

## Usage

### Development

```bash
# Start the web server
cd backend
go run main.go

# Open in browser
http://localhost:8080
```

### API Endpoints

- `GET /` - Main web interface
- `POST /api/simulate` - Run simulation
- `GET /api/examples` - List examples
- `GET /api/examples/:name` - Get specific example

## Requirements

- Go 1.16 or higher
- Gin framework (auto-installed)

## Installation

```bash
cd backend
go mod tidy
```

## Technologies

- **Backend**: Go + Gin framework
- **Frontend**: HTML + CSS + JavaScript
- **API**: RESTful endpoints
- **Data**: JSON format

This phase provides a web interface for the core Turing machine simulator.
