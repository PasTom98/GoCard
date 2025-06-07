import {app, BrowserWindow} from 'electron';
import * as path from "node:path";


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

    win.loadFile(path.join(app.getAppPath() + '/dist-react/index.html'));

})