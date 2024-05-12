package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func call(exe string, params ...string) {
    cmd := exec.Command(exe, params...)
    stdout, err := cmd.Output()
    
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(stdout))
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
