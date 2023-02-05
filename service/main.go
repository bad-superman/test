package main

import (
	"net/http"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		supports_search: false,
		supports_group_request: true,
		supported_resolutions: ["1m", "5m", "15m", "30m", "60m", "1D", "1W", "1M"],
		supports_marks: false,
		supports_time: true
	}`))
}

func symbolInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		symbol: ["000828"],
		description: ["泰达转型机遇股票A"],
		exchange-listed: "泰达",
		exchange-traded: "泰达",
		minmov: 1,
		minmov2: 0,
		pricescale: [0.01, 0.1, 1],
		has-dwm: true,
		has-intraday: true,
		has-no-volume: [true]
		type: ["stock"],
		ticker: ["000828"],
		timezone: “Asia/Shanghai”,
		session-regular: “”,
	 }`))
}

func main() {
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/symbol_info", symbolInfoHandler)
	http.HandleFunc("/symbols", symbolInfoHandler)
	http.ListenAndServe("127.0.0.1:8011", nil)
}
