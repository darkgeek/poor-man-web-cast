package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"net/http"
	"os/exec"
	"time"
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

func callAsync(exe string, params ...string) {
	exeCmd := cmd.NewCmd(exe, params...)
	statusChan := exeCmd.Start()

	go func() {
		<-time.After(10 * time.Second)
		select {
		case cmdStatus := <-statusChan:
			fmt.Println("exited after 10 secs: ", cmdStatus.Exit)
		default:
			status := exeCmd.Status()
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
	call("notify-send", "正在启动firefox...")
	callAsync("firefox", "--kiosk", urlToOpen)
}

func openMpv(w http.ResponseWriter, req *http.Request) {
	urlToOpen := getParamFromReq(req, "url")
	call("killall", "-9", "mpv")
	call("notify-send", "正在解析资源并启动mpv播放器...")
	callAsync("./run-mpv.sh", urlToOpen)
}

func openChromiumInKioskMode(w http.ResponseWriter, req *http.Request) {
	urlToOpen := getParamFromReq(req, "url")
	call("killall", "-9", "/usr/lib/chromium-browser/chromium-browser")
	call("notify-send", "正在启动chromium...")
	callAsync("chromium-browser", "--kiosk", urlToOpen)
}

func main() {

	http.HandleFunc("/firefox-kiosk", openFirefoxInKioskMode)
	http.HandleFunc("/mpv", openMpv)
	http.HandleFunc("/chromium-kiosk", openChromiumInKioskMode)

	http.ListenAndServe(":8090", nil)
}
