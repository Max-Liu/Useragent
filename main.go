package main

import (
	"log"
	"regexp"
	"strings"
)


var user_agent string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9) AppleWebKit/537.71 (KHTML, like Gecko) Version/7.0 Safari/537.71"

type userAgent struct {
	Platforms string
	Browser
}

type Browser struct {
	name    string
	version string
}

var uA userAgent = userAgent{}
var Platforms = make(map[string]string)
var Browsers = make(map[string]string)
var Browsers_slice []string

func init() {
	Platforms["windows nt 6.0"] = "Windows Longhorn"
	Platforms["windows nt 5.2"] = "Windows 2003"
	Platforms["windows nt 5.0"] = "Windows 2000"
	Platforms["windows nt 5.1"] = "Windows XP"
	Platforms["windows nt 4.0"] = "Windows NT 4.0"
	Platforms["winnt4.0"] = "Windows NT 4.0"
	Platforms["winnt 4.0"] = "Windows NT"
	Platforms["winnt"] = "Windows NT"
	Platforms["windows 98"] = "Windows 98"
	Platforms["win98"] = "Windows 98"
	Platforms["windows 95"] = "Windows 95"
	Platforms["win95"] = "Windows 95"
	Platforms["windows"] = "Unknown Windows OS"
	Platforms["os x"] = "Mac OS X"
	Platforms["ppc mac"] = "Power PC Mac"
	Platforms["freebsd"] = "FreeBSD"
	Platforms["ppc"] = "Macintosh"
	Platforms["linux"] = "Linux"
	Platforms["debian"] = "Debian"
	Platforms["sunos"] = "Sun Solaris"
	Platforms["beos"] = "BeOS"
	Platforms["apachebench"] = "ApacheBench"
	Platforms["aix"] = "AIX"
	Platforms["irix"] = "Irix"
	Platforms["osf"] = "DEC OSF"
	Platforms["hp-ux"] = "HP-UX"
	Platforms["netbsd"] = "NetBSD"
	Platforms["bsdi"] = "BSDi"
	Platforms["openbsd"] = "OpenBSD"
	Platforms["gnu"] = "GNU/Linux"
	Platforms["unix"] = "Unknown Unix OS"

	Browsers_slice = []string{"Flock", "Chrome", "Opera",
		"MSIE", "Internet Explorer", "ipad", "Shiira", "Firefox",
		"Chimera", "Phoenix", "Firebird", "Camino",
		"Netscape", "OmniWeb", "Safari", "Mozilla",
		"Konqueror", "icab", "Lynx", "Links", "hotjava", "amaya", "IBrowse"}

	Browsers["Flock"] = "Flock"
	Browsers["Chrome"] = "Chrome"
	Browsers["Opera"] = "Opera"
	Browsers["MSIE"] = "Internet Explorer"
	Browsers["Internet Explorer"] = "Internet Explorer"
	Browsers["ipad"] = "iPad"
	Browsers["Shiira"] = "Shiira"
	Browsers["Firefox"] = "Firefox"
	Browsers["Chimera"] = "Chimera"
	Browsers["Phoenix"] = "Phoenix"
	Browsers["Firebird"] = "Firebird"
	Browsers["Camino"] = "Camino"
	Browsers["Netscape"] = "Netscape"
	Browsers["OmniWeb"] = "OmniWeb"
	Browsers["Safari"] = "Safari"
	Browsers["Mozilla"] = "Mozilla"
	Browsers["Konqueror"] = "Konqueror"
	Browsers["icab"] = "iCab"
	Browsers["Lynx"] = "Lynx"
	Browsers["Links"] = "Links"
	Browsers["hotjava"] = "HotJava"
	Browsers["amaya"] = "Amaya"
	Browsers["IBrowse"] = "IBrowse"
}

func main() {
	set_platform()
	set_browser()
	log.Println(uA)
}

func set_browser() {
	var matched_slice [][]string
	for _, v := range Browsers_slice {

		regexp, err := regexp.Compile(`(?i)` + v + `.*?([0-9\.]+)`)
		if err != nil {
			log.Panic(err)
		}

		matched := regexp.FindStringSubmatch(user_agent)
		if matched != nil {
			matched_slice = append(matched_slice, matched)
		}
	}
	if matched_slice != nil {
		key := strings.Split(matched_slice[0][0], "/")
		uA.Browser.name = Browsers[key[0]]
		uA.Browser.version = matched_slice[0][1]
	}
}

func set_platform() {

	for k, _ := range Platforms {

		regexp, err := regexp.Compile(`(?i)` + k)
		if err != nil {
			log.Panic(err)
		}

		platforms_str := regexp.FindString(user_agent)
		if platforms_str != "" {
			uA.Platforms = platforms_str
		}
	}
}
