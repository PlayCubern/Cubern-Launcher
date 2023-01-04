
# Cubern Launcher

Cubern is an online sandbox where the creativity never ends. Users can customize their avatars as they want. Players can also make there own games using our editor built in Godot game Engine! But, to run these a client is required which also requires a URL Protocol! 

## Authors

- [@OoIks](https://www.github.com/OoIks)

Helping the tool become more optimizations could earn your name being here, open a [pull request](https://github.com/PlayCubern/Cubern-Launcher/pulls) if you want to suggest a change!


## Running Locally

In order to run this locally, we are required to install go programming language on our computer! 

Clone the project by either using git or by clicking the green code button and downloading as zip

Go to the project directory where you extracted it

```bash
  cd path/to/extracted/repository
```

Compile application:

```bash
  go build
```



## Features

- Install client and unzip it from server
- Install custom url protocols
- Standalone build

## Dependencies
The installer uses [Gowin](https://github.com/luisiturrios1/gowin) in order to create or modify windows registry files, it could be installed by simply running:

```bash
  go get github.com/luisiturrios/gowin
```

## API Reference

#### • downloadFile(url, string)
Gets a online file and downloads on the computer using the "net/http" package

#### • unzipFile(src string, dest string)
Unzips a local file

## Why?
The reason behind making the installer open source is none other then to prevent suspicious activities. But how this would help refrain you may ask. Well, as the installer requires administrative permissions to do its things, it can look sketchy like a virus.

