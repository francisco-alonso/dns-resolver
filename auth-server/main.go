package main

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func handleAuthRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Compress = false

	question := msg.Question[0]
	domain := question.Name

	log.Printf("Authoritative server received query for domain: %s", domain)

	// Domain records
	domainRecords := map[string]string{
		"example.com.": "93.184.216.34", // Final IP for example.com
	}

	// Check if domain exists
	if ip, exists := domainRecords[domain]; exists {
		msg.Authoritative = true
		msg.Answer = []dns.RR{
			&dns.A{
				Hdr: dns.RR_Header{
					Name:   domain,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    3600,
				},
				A: net.ParseIP(ip),
			},
		}
	} else {
		msg.Rcode = dns.RcodeNameError // NXDOMAIN if domain is not found
	}

	w.WriteMsg(msg)
}

func main() {
	server := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc(".", handleAuthRequest)

	log.Println("Starting Authoritative Server on :53")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start Authoritative Server: %s", err)
	}
}
