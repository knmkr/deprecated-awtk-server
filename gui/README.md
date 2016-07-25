#

Install

```
$ npm install
```

Run

```
$ electron .
```

Build

```
$ electron-packager . --platform=darwin --arch=x64 --version=$(npm view electron-prebuilt version) --icon=icon/app.icns
```

Logs

- OSX: `~/Library/Logs/wgx/log.log`
- Linux: `~/.config/wgx/log.log`
- Windows: `$HOME/AppData/Roaming/wgx/log.log`
