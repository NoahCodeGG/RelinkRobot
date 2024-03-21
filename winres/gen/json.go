// Package main generates winres.json
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const JS = `{
  "RT_GROUP_ICON": {
    "APP": {
      "0000": [
        "icon.png"
      ]
    }
  },
  "RT_MANIFEST": {
    "#1": {
      "0409": {
        "identity": {
          "name": "RelinkRobot-Plugin",
          "version": "%s"
        },
        "description": "",
        "minimum-os": "vista",
        "execution-level": "as invoker",
        "ui-access": false,
        "auto-elevate": false,
        "dpi-awareness": "system",
        "disable-theming": false,
        "disable-window-filtering": false,
        "high-resolution-scrolling-aware": false,
        "ultra-high-resolution-scrolling-aware": false,
        "long-path-aware": false,
        "printer-driver-isolation": false,
        "gdi-scaling": false,
        "segment-heap": false,
        "use-common-controls-v6": false
      }
    }
  },
  "RT_VERSION": {
    "#1": {
      "0000": {
        "fixed": {
          "file_version": "%s",
          "product_version": "%s",
          "timestamp": "%s"
        },
        "info": {
          "0409": {
            "Comments": "用于加/解密链接的 QQ 机器人插件",
            "CompanyName": "NoahCode",
            "FileDescription": "https://github.com/NoahCodeGG/RelinkRobot-Plugin",
            "FileVersion": "%s",
            "InternalName": "",
            "LegalCopyright": "%s",
            "LegalTrademarks": "",
            "OriginalFilename": "RRP.EXE",
            "PrivateBuild": "",
            "ProductName": "RelinkRobot-Plugin",
            "ProductVersion": "%s",
            "SpecialBuild": ""
          }
        }
      }
    }
  }
}`

const TimeFormat = `2006-01-02T15:04:05+08:00`
const Version = "v1.0.0"
const Copyright = "NoahCode"

func main() {
	f, err := os.Create("winres.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	commitcnt := strings.Builder{}
	commitcnt.WriteString(Version[1:])
	commitcnt.WriteByte('.')
	commitcntcmd := exec.Command("git", "rev-list", "--count", "HEAD")
	commitcntcmd.Stdout = &commitcnt
	err = commitcntcmd.Run()
	if err != nil {
		panic(err)
	}
	fv := commitcnt.String()[:commitcnt.Len()-1]
	_, err = fmt.Fprintf(f, JS, fv, fv, Version, time.Now().Format(TimeFormat), fv, Copyright+". All Rights Reserved.", Version)
	if err != nil {
		panic(err)
	}
}
