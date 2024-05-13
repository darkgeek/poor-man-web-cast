package main

import (
    "fmt"
    "net/http"
    "github.com/go-cmd/cmd"
)

func call(exe string, params ...string) {
    cmd := cmd.NewCmd(exe, params...)
    statusChan := cmd.Start()
    
    go func() {
        <-time.After(10 * time.Second)
        select {
        case cmdStatus:= <- statusChan:
            fmt.Println("exited after 10 secs")
        default:
            status := findCmd.Status()
            fmt.Println("stderr after 10 secs:")
            fmt.Println(status.Stderr)
            fmt.Println("stdout after 10 secs:")
            fmt.Println(status.Stdout)
        }
    }()
}

func getParamFromReq(req *http.Request, paramName string) string {
	value := req.FormValue(paramName)
	if len(value) == 0 {
		fmt.Println("value not present")
		return ""
	} else {
		fmt.Println("get value: " + value)
		return value
	}
}

func openFirefoxInKioskMode(w http.ResponseWriter, req *http.Request) {
	urlToOpen := getParamFromReq(req, "url")
    call("killall", "-9", "firefox")
	call("firefox", "--kiosk", urlToOpen)
}

func openMpv(w http.ResponseWriter, req *http.Request) {
	urlToOpen := getParamFromReq(req, "url")
    call("killall", "-9", "mpv")
	call("mpv", "-fs", "--ytdl-format=ytdl", "--ytdl-raw-options=cookies-from-browser=chromium", urlToOpen)
}

func openChromiumInKioskMode(w http.ResponseWriter, req *http.Request) {
	urlToOpen := getParamFromReq(req, "url")
    call("killall", "-9", "/usr/lib/chromium-browser/chromium-browser")
	call("chromium-browser", "--kiosk", urlToOpen)
}

func main() {

    http.HandleFunc("/firefox-kiosk", openFirefoxInKioskMode)
    http.HandleFunc("/mpv", openMpv)
    http.HandleFunc("/chromium-kiosk", openChromiumInKioskMode)

    http.ListenAndServe(":8090", nil)
}
