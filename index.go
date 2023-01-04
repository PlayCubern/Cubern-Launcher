package main

import (
    "archive/zip"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "github.com/luisiturrios/gowin"
    "strconv"
    "time"
)


func downloadFile(url string) error {
    startTime := time.Now()
  
    out, err := os.Create("file.zip")
    if err != nil {
        return err
    }
    defer out.Close()

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    n, err := io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}



func unzipFile(src string, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
        fpath := filepath.Join(dest, f.Name)

        if !strings.HasPrefix(fpath, filepath.Clean(dest) + string(os.PathSeparator)) {
            return fmt.Errorf("%s: illegal file path", fpath)
        }

        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
            continue
        }

        if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm);err != nil {
            return err
        }

        outFile,
        err := os.OpenFile(fpath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, f.Mode())
        if err != nil {
            return err
        }

        rc,
        err := f.Open()
        if err != nil {
            return err
        }

        _,
        err = io.Copy(outFile, rc)

        outFile.Close()
        rc.Close()

        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    fmt.Println("Please wait while Cubern starts to initialize!")
    
	cwd, err := os.Getwd()
  if err != nil {
      fmt.Println(err)
      return
  }

  _, err = os.Stat(filepath.Join(cwd, "Cubern.exe"))
  if err == nil {
      fmt.Println("Cubern.exe already exists, skipping download and unzip")
     return
  } else if os.IsNotExist(err) {
      fmt.Println("Cubern.exe does not exist, proceeding with download and unzip")

     err = downloadFile("https://server.cubern.xyz/Clients/Cubern.zip")
      if err != nil {
          fmt.Println(err)
          return
      }

      err = unzipFile("file.zip", cwd)
      if err != nil {
          fmt.Println(err)
      }

      os.Remove("file.zip")
  } else {
      fmt.Println(err)
      return
  }



	  KeyString := "\"" + filepath.Join(cwd, "Cubern.exe") + "\" \"%1\""
     BuildRootFolder := gowin.WriteStringReg("HKCR", `Cubern\shell\open\command`, "", KeyString)
      if BuildRootFolder != nil {
         fmt.Println(BuildRootFolder)
     } else {
         fmt.Println("Key inserted")
     }
}
