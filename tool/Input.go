package tool

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func IsPost(r *http.Request) bool {
	return strings.ToUpper(r.Method) == "POST"
}

func IsAjax(r *http.Request) bool {
	return r.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func UserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}

func GetIp(r *http.Request) string {
	ip := r.Header.Get("REMOTE_ADDR")
	if ip == "" {
		var err error
		ip, _, err = net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("userip: %q is not IP:port", r.RemoteAddr)
			return ""
		}

		userIP := net.ParseIP(ip)
		if userIP == nil {
			//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
			log.Println("userip: %s is not IP:port", r.RemoteAddr)
			return ""
		}
	}
	return ip
}

// Get the values by string from querystring
func Get(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetInt(r *http.Request, key string) (j int, err error) {
	str := Get(r, key)
	if str != "" && len(str) > 0 {
		j, err = strconv.Atoi(str)
	}
	return
}
