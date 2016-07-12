var path = require("path");
var log = require('electron-log');

const electron = require('electron');
// Module to control application life.
const {app} = electron;
// Module to create native browser window.
const {BrowserWindow} = electron;

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let win;

// TODO: DRY
var addr = "localhost:1323";

function createWindow() {
    // Create the browser window.
    win = new BrowserWindow({width: 800, height: 600});

    // and load the index.html of the app.
    win.loadURL(`file://${__dirname}/index.html`);

                // Open the DevTools.
                win.webContents.openDevTools();

                // Emitted when the window is closed.
                win.on('closed', () => {
                    // Dereference the window object, usually you would store windows
                    // in an array if your app supports multi windows, this is the time
                    // when you should delete the corresponding element.
                    win = null;
                });
               }

    // This method will be called when Electron has finished
    // initialization and is ready to create browser windows.
    // Some APIs can only be used after this event occurs.
    app.on('ready', createWindow);

    // Quit when all windows are closed.
    app.on('window-all-closed', () => {
        // On macOS it is common for applications and their menu bar
        // to stay active until the user quits explicitly with Cmd + Q
        if (process.platform !== 'darwin') {
            app.quit();
        }
    });

    app.on('activate', () => {
        // On macOS it's common to re-create a window in the app when the
        // dock icon is clicked and there are no other windows open.
        if (win === null) {
            createWindow();
        }
    });

    // In this file you can include the rest of your app's specific main process
    // code. You can also put them in separate files and require them here.

    // TODO: close safely
    var child_process = require('child_process')
    runserver = child_process.execFile(path.join(__dirname, 'script', 'wgx'), ['runserver', addr], {cwd: __dirname});
    log.info(runserver);

    runserver.stdout.on('data', function (data) {
        log.info('stdout: ' + data);
    });

    runserver.stderr.on('data', function (data) {
        log.info('stderr: ' + data);
    });

    runserver.on('close', function (code) {
        log.info('child process exited with code ' + code);
    });

    //
    process.on('uncaughtException', function (err) {
        log.error('Caught exception: ' + err);
    });
