package useragent

import (
	"log"
	"regexp"
	"strings"
)

type UserAgent struct {
	UserAgent_str string
	Platforms     string
	Browser
	Mobile
}

type Mobile struct {
	IsMobile bool
	Name     string
}

type Browser struct {
	Name    string
	Version string
}

var platforms = make(map[string]string)
var browsers = make(map[string]string)
var mobiles = make(map[string]string)
var browsers_slice []string

func init() {
	config()
}

func NewUserAgent() *UserAgent{
	ua := &UserAgent{}
	return ua
}
func config() {
	platforms["windows nt 6.0"] = "Windows Longhorn"
	platforms["windows nt 5.2"] = "Windows 2003"
	platforms["windows nt 5.0"] = "Windows 2000"
	platforms["windows nt 5.1"] = "Windows XP"
	platforms["windows nt 4.0"] = "Windows NT 4.0"
	platforms["winnt4.0"] = "Windows NT 4.0"
	platforms["winnt 4.0"] = "Windows NT"
	platforms["winnt"] = "Windows NT"
	platforms["windows 98"] = "Windows 98"
	platforms["win98"] = "Windows 98"
	platforms["windows 95"] = "Windows 95"
	platforms["win95"] = "Windows 95"
	platforms["windows"] = "Unknown Windows OS"
	platforms["os x"] = "Mac OS X"
	platforms["ppc mac"] = "Power PC Mac"
	platforms["freebsd"] = "FreeBSD"
	platforms["ppc"] = "Macintosh"
	platforms["linux"] = "Linux"
	platforms["debian"] = "Debian"
	platforms["sunos"] = "Sun Solaris"
	platforms["beos"] = "BeOS"
	platforms["apachebench"] = "ApacheBench"
	platforms["aix"] = "AIX"
	platforms["irix"] = "Irix"
	platforms["osf"] = "DEC OSF"
	platforms["hp-ux"] = "HP-UX"
	platforms["netbsd"] = "NetBSD"
	platforms["bsdi"] = "BSDi"
	platforms["openbsd"] = "OpenBSD"
	platforms["gnu"] = "GNU/Linux"
	platforms["unix"] = "Unknown Unix OS"

	mobiles["mobileexplorer"] = "Mobile Explorer"
	mobiles["palmsource"] = "Palm"
	mobiles["palmscape"] = "Palmscape"

	// Phones and Manufacturers
	mobiles["motorola"] = "Motorola"
	mobiles["nokia"] = "Nokia"
	mobiles["palm"] = "Palm"
	mobiles["iphone"] = "Apple iPhone"
	mobiles["ipod"] = "Apple iPod Touch"
	mobiles["sony"] = "Sony Ericsson"
	mobiles["ericsson"] = "Sony Ericsson"
	mobiles["blackberry"] = "BlackBerry"
	mobiles["cocoon"] = "O2 Cocoon"
	mobiles["blazer"] = "Treo"
	mobiles["lg"] = "LG"
	mobiles["amoi"] = "Amoi"
	mobiles["xda"] = "XDA"
	mobiles["mda"] = "MDA"
	mobiles["vario"] = "Vario"
	mobiles["htc"] = "HTC"
	mobiles["samsung"] = "Samsung"
	mobiles["sharp"] = "Sharp"
	mobiles["sie-"] = "Siemens"
	mobiles["alcatel"] = "Alcatel"
	mobiles["benq"] = "BenQ"
	mobiles["ipaq"] = "HP iPaq"
	mobiles["mot-"] = "Motorola"
	mobiles["playstation portable"] = "PlayStation Portable"
	mobiles["hiptop"] = "Danger Hiptop"
	mobiles["nec-"] = "NEC"
	mobiles["panasonic"] = "Panasonic"
	mobiles["philips"] = "Philips"
	mobiles["sagem"] = "Sagem"
	mobiles["sanyo"] = "Sanyo"
	mobiles["spv"] = "SPV"
	mobiles["zte"] = "ZTE"
	mobiles["sendo"] = "Sendo"

	// Operating Systems
	mobiles["symbian"] = "Symbian"
	mobiles["Symbianos"] = "SymbianOS"
	mobiles["elaine"] = "Palm"
	mobiles["palm"] = "Palm"
	mobiles["series60"] = "Symbian S60"
	mobiles["windows ce"] = "Windows CE"

	// browsers
	mobiles["obigo"] = "Obigo"
	mobiles["netfront"] = "Netfront Browser"
	mobiles["openwave"] = "Openwave Browser"
	mobiles["mobilexplorer"] = "Mobile Explorer"
	mobiles["operamini"] = "Opera Mini"
	mobiles["opera mini"] = "Opera Mini"

	// Other
	mobiles["digital paths"] = "Digital Paths"
	mobiles["avantgo"] = "AvantGo"
	mobiles["xiino"] = "Xiino"
	mobiles["novarra"] = "Novarra Transcoder"
	mobiles["vodafone"] = "Vodafone"
	mobiles["docomo"] = "NTT DoCoMo"
	mobiles["o2"] = "O2"

	// Fallback
	// mobiles["mobile"] = "Generic Mobile"
	mobiles["wireless"] = "Generic Mobile"
	mobiles["j2me"] = "Generic Mobile"
	mobiles["midp"] = "Generic Mobile"
	mobiles["cldc"] = "Generic Mobile"
	mobiles["up.link"] = "Generic Mobile"
	mobiles["up.browser"] = "Generic Mobile"
	mobiles["smartphone"] = "Generic Mobile"
	mobiles["cellphone"] = "Generic Mobile"

	browsers_slice = []string{"Flock", "Chrome", "Opera",
		"MSIE", "Internet Explorer", "ipad", "Shiira", "Firefox",
		"Chimera", "Phoenix", "Firebird", "Camino",
		"Netscape", "OmniWeb", "Safari", "Mozilla",
		"Konqueror", "icab", "Lynx", "Links", "hotjava", "amaya", "IBrowse"}

	browsers["Flock"] = "Flock"
	browsers["Chrome"] = "Chrome"
	browsers["Opera"] = "Opera"
	browsers["MSIE"] = "Internet Explorer"
	browsers["Internet Explorer"] = "Internet Explorer"
	browsers["ipad"] = "iPad"
	browsers["Shiira"] = "Shiira"
	browsers["Firefox"] = "Firefox"
	browsers["Chimera"] = "Chimera"
	browsers["Phoenix"] = "Phoenix"
	browsers["Firebird"] = "Firebird"
	browsers["Camino"] = "Camino"
	browsers["Netscape"] = "Netscape"
	browsers["OmniWeb"] = "OmniWeb"
	browsers["Safari"] = "Safari"
	browsers["Mozilla"] = "Mozilla"
	browsers["Konqueror"] = "Konqueror"
	browsers["icab"] = "iCab"
	browsers["Lynx"] = "Lynx"
	browsers["Links"] = "Links"
	browsers["hotjava"] = "HotJava"
	browsers["amaya"] = "Amaya"
	browsers["IBrowse"] = "IBrowse"

	
}

func (ua *UserAgent) SetUseragent(useragent string) {
	ua.UserAgent_str = useragent
	ua.set_browser()
	ua.set_mobile()
}

func (ua *UserAgent) set_mobile() {
	user_agent_lower := strings.ToLower(ua.UserAgent_str)
	for k, v := range mobiles {
		if strings.Contains(user_agent_lower, k) {
			if strings.Index(user_agent_lower, k) != -1 {
				ua.Mobile.IsMobile = true
				ua.Mobile.Name = v
				break
			}
		}
	}
}

func (ua *UserAgent) set_browser() {
	var matched_slice [][]string
	for _, v := range browsers_slice {

		regexp, err := regexp.Compile(`(?i)` + v + `.*?([0-9\.]+)`)
		if err != nil {
			log.Panic(err)
		}

		matched := regexp.FindStringSubmatch(ua.UserAgent_str)
		if matched != nil {
			matched_slice = append(matched_slice, matched)
		}
	}
	if matched_slice != nil {
		key := strings.Split(matched_slice[0][0], "/")
		ua.Browser.Name = key[0]
		ua.Browser.Version = matched_slice[0][1]
	}
}

func (ua *UserAgent) set_platform() {

	for k, _ := range platforms {

		regexp, err := regexp.Compile(`(?i)` + k)
		if err != nil {
			log.Panic(err)
		}

		platforms_str := strings.ToLower(regexp.FindString(ua.UserAgent_str))
		if platforms_str != "" {
			ua.Platforms = platforms[platforms_str]
		}
	}
}
