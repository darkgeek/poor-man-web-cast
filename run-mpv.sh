#!/bin/sh

mpv -fs --ytdl-format=ytdl --ytdl-raw-options=cookies-from-browser=chromium "$1" || notify-send "无法解析资源，请使用播放器打开"
