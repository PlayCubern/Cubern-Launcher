
# Cubern Launcher

Cubern is an online sandbox where the creativity never ends. Users can customize their avatars as they want. Players can also make there own games using our editor built in Godot game Engine! But, to run these a client is required which also requires a URL Protocol! 

## Authors

- [@OoIks](https://www.github.com/OoIks)

Helping the tool become more optimizations could earn your name being here, open a [pull request](https://github.com/PlayCubern/Cubern-Launcher/pulls) if you want to suggest a change!


## Running Locally

As this project is available for you to use for your sandbox or even your game, here is a detailed and a easy guide on setting it up:

1 - We first install golang from the official (website)[https://go.dev/doc/install] <br>
2 - Run the installer and then clone or download the repository <br>
3 - Open the extracted repository in visual studio code <br>
4 - Now open the terminl and run:

```bash
go version
``` 

if it successfully returns the version, go is installed on your computer! <br>
5 - now we install [Gowin](https://github.com/luisiturrios1/gowin) which is required to run the essential part, building the protocols. We do that by running this into the visual studio code's console: 
```bash
   go get github.com/luisiturrios/gowin
``` 
<br>
6 - We have now installed the dependencies! Now we edit the code to make it work with your project! At line 478 we check if the client is installed, replace Cubern.exe with your .exe's name such as: mygame.exe <br>

7 - Now we install the file from online server at line 485 replace it with your url! <br>

8 - Now you are to configure your registry which is done at line 504 to 505! Such as instead of:

```bash
KeyString := "\"" + filepath.Join(cwd, "Cubern.exe") + "\" \"%1\""
``` 
you make it to:
```bash
KeyString := "\"" + filepath.Join(cwd, "MyGame.exe") + "\" \"%1\""
```
please remember to make sure that the .exe is same as the exe name in the online zip!
And at line 505, you replace Cubern with yourgame and after configuring, save the file and using visual studio code's terminal, build it using:
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

## Our Discord!
Join our discord  [here](https://dsc.gg/Cubern) <br>
![Dsc](https://img.shields.io/discord/1000052563711901767)
