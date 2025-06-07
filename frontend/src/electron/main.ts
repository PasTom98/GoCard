import {app, BrowserWindow} from 'electron';
import * as path from "node:path";
import {isDev} from "./util.js";


type Test = string;

app.on('ready', () => {
    const win = new BrowserWindow(
        {
            width: 800,
            height: 600,
            webPreferences: {
                nodeIntegration: true
            }
        }
    )

    if (isDev()) {
        win.loadURL('http://localhost:5123');
        win.webContents.openDevTools();
    } else {
    }

    win.loadFile(path.join(app.getAppPath() + '/dist-react/index.html'));

})