const { app, BrowserWindow, Menu } = require('electron');
const path = require('path');
const { spawn } = require('child_process');

let mainWindow;
let serverProcess;

function createWindow() {
    // Create the browser window
    mainWindow = new BrowserWindow({
        width: 1400,
        height: 900,
        minWidth: 1000,
        minHeight: 700,
        webPreferences: {
            nodeIntegration: false,
            contextIsolation: true,
            enableRemoteModule: false
        },
        icon: path.join(__dirname, 'assets', 'icon.png'),
        titleBarStyle: 'default',
        show: false
    });

    // Load the app
    mainWindow.loadURL('http://localhost:8080');

    // Show window when ready
    mainWindow.once('ready-to-show', () => {
        mainWindow.show();
    });

    // Handle window closed
    mainWindow.on('closed', () => {
        mainWindow = null;
    });

    // Create menu
    createMenu();
}

function createMenu() {
    const template = [
        {
            label: 'File',
            submenu: [
                {
                    label: 'New Simulation',
                    accelerator: 'CmdOrCtrl+N',
                    click: () => {
                        mainWindow.webContents.send('new-simulation');
                    }
                },
                {
                    label: 'Load Example',
                    submenu: [
                        {
                            label: 'Simple Example',
                            click: () => {
                                mainWindow.webContents.send('load-example', 'simple');
                            }
                        },
                        {
                            label: 'Palindrome Check',
                            click: () => {
                                mainWindow.webContents.send('load-example', 'palindrome');
                            }
                        },
                        {
                            label: 'Binary Counter',
                            click: () => {
                                mainWindow.webContents.send('load-example', 'binary');
                            }
                        }
                    ]
                },
                { type: 'separator' },
                {
                    label: 'Quit',
                    accelerator: process.platform === 'darwin' ? 'Cmd+Q' : 'Ctrl+Q',
                    click: () => {
                        app.quit();
                    }
                }
            ]
        },
        {
            label: 'View',
            submenu: [
                { role: 'reload' },
                { role: 'forceReload' },
                { role: 'toggleDevTools' },
                { type: 'separator' },
                { role: 'resetZoom' },
                { role: 'zoomIn' },
                { role: 'zoomOut' },
                { type: 'separator' },
                { role: 'togglefullscreen' }
            ]
        },
        {
            label: 'Help',
            submenu: [
                {
                    label: 'About Turing Machine Simulator',
                    click: () => {
                        require('electron').dialog.showMessageBox(mainWindow, {
                            type: 'info',
                            title: 'About',
                            message: 'Turing Machine Simulator',
                            detail: 'A desktop application for simulating Turing machines.\n\nVersion 1.0.0\nBuilt with Electron and Go'
                        });
                    }
                }
            ]
        }
    ];

    const menu = Menu.buildFromTemplate(template);
    Menu.setApplicationMenu(menu);
}

function startServer() {
    return new Promise((resolve, reject) => {
        // Start the Go server
        serverProcess = spawn('go', ['run', 'main.go'], {
            cwd: path.join(__dirname, 'backend'),
            stdio: 'pipe'
        });

        serverProcess.stdout.on('data', (data) => {
            console.log(`Server: ${data}`);
            if (data.toString().includes('Server starting on :8080')) {
                resolve();
            }
        });

        serverProcess.stderr.on('data', (data) => {
            console.error(`Server Error: ${data}`);
        });

        serverProcess.on('error', (error) => {
            console.error('Failed to start server:', error);
            reject(error);
        });

        // Wait a bit for server to start
        setTimeout(resolve, 2000);
    });
}

// App event handlers
app.whenReady().then(async () => {
    try {
        await startServer();
        createWindow();
    } catch (error) {
        console.error('Failed to start app:', error);
        app.quit();
    }
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    }
});

app.on('before-quit', () => {
    if (serverProcess) {
        serverProcess.kill();
    }
});

// Handle app quit
process.on('SIGINT', () => {
    if (serverProcess) {
        serverProcess.kill();
    }
    app.quit();
});
