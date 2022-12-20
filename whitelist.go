
package main

import (
    "net/http"
    "strings"
)

// Whitelist is a slice of strings containing the IP addresses
// that are on the whitelist.
var Whitelist []string

func main() {
    // Set up the HTTP handler for the "/whitelist" endpoint.
    http.HandleFunc("/whitelist", whitelistHandler)

    // Start the HTTP server.
    http.ListenAndServe(":8080", nil)
}

func whitelistHandler(w http.ResponseWriter, r *http.Request) {
    // Check the request method.
    switch r.Method {
    case "POST":
        // If the request is a POST request, add the IP address
        // from the request body to the whitelist.
        ip := strings.TrimSpace(r.FormValue("ip"))
        Whitelist = append(Whitelist, ip)

        // Return a success message.
        w.Write([]byte("IP address added to whitelist"))

    case "GET":
        // If the request is a GET request, get the IP address
        // from the URL parameters and check if it is on the
        // whitelist.
        ip := r.URL.Query().Get("ip")
        if contains(Whitelist, ip) {
            w.Write([]byte("true"))
        } else {
            w.Write([]byte("false"))
        }

    default:
        // If the request method is not supported, return an
        // error message.
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Method not allowed"))
}
