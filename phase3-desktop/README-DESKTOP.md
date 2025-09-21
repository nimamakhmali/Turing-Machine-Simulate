# Turing Machine Simulator - Desktop App

A desktop application for simulating Turing machines, built with Electron and Go.

## Features

- **Desktop Application**: Native desktop app experience
- **3 Built-in Examples**: Simple, Palindrome, Binary Counter
- **Interactive UI**: Modern, responsive interface
- **Step-by-step Execution**: Visualize each step of computation
- **JSON Editor**: Customize machine configurations
- **Cross-platform**: Works on Windows, Mac, and Linux

## Installation

### Prerequisites

- Node.js (v16 or higher)
- Go (v1.16 or higher)
- Git

### Development Setup

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd Turing-Machine-Go
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Run in development mode**
   ```bash
   npm run dev
   ```

### Building the App

1. **Build for current platform**
   ```bash
   npm run build
   ```

2. **Or use the build script (Windows)**
   ```bash
   build.bat
   ```

The built application will be in the `dist` folder.

## Usage

### Running the App

1. **Development mode**: `npm run dev`
2. **Production mode**: `npm start`
3. **Built app**: Run the executable from `dist` folder

### Using the App

1. **Choose an example** from the dropdown menu
2. **Modify the tape** input if needed
3. **Click "Run Simulation"** for complete execution
4. **Click "Step Through"** for step-by-step execution
5. **View results** in the simulation panel

### Menu Options

- **File > New Simulation**: Reset to default state
- **File > Load Example**: Quick access to examples
- **View > Developer Tools**: Open debugging tools
- **Help > About**: App information

## Project Structure

```
Turing-Machine-Go/
├── main.js                 # Electron main process
├── package.json            # Node.js dependencies
├── frontend/               # Web interface
│   ├── templates/
│   └── static/
├── backend/                # Go server
│   └── main.go
├── turing/                 # Turing machine logic
├── input/                  # Example configurations
└── assets/                 # App icons
```

## Building for Different Platforms

### Windows
```bash
npm run build -- --win
```

### macOS
```bash
npm run build -- --mac
```

### Linux
```bash
npm run build -- --linux
```

## Troubleshooting

### Common Issues

1. **Port 8080 already in use**
   - Close other applications using port 8080
   - Or modify the port in `backend/main.go`

2. **Go dependencies missing**
   ```bash
   cd backend
   go mod tidy
   ```

3. **Electron app not starting**
   - Check if Go is installed and in PATH
   - Verify all dependencies are installed

### Development Tips

- Use `Ctrl+Shift+I` to open Developer Tools
- Check the console for error messages
- The Go server runs on `localhost:8080`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

MIT License - see LICENSE file for details.
