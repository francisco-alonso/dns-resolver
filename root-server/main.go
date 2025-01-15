package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func handleRootRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Compress = false
	
	question := msg.Question[0]
	
	domain := question.Name //example.com
	
	log.Printf("Root server received query for domain: %s", domain)
	
	tldServers := map[string]string{
		"com.": "192.168.1.10",
		"net.": "192.168.1.11",
	}
	
	tld := domain[strings.LastIndex(domain,".")+1:] //com or net
	
	if ip, exists := tldServers[tld+"."]; exists {
		msg.Authoritative = true
		msg.Answer = []dns.RR{
			&dns.A{
				Hdr: dns.RR_Header{
					Name: tld + ".",
					Rrtype: dns.TypeA,
					Class: dns.ClassINET,
					Ttl: 3600,
				},
				A: net.ParseIP(ip),
			},
		}
	} else {
		msg.Rcode = dns.RcodeNameError // as TLD is not found
	}
	
	if err := w.WriteMsg(msg); err != nil {
		fmt.Printf("Error in response, got %s", err)
	}
}

func main() {
	server := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc(".", handleRootRequest)

	log.Println("Starting Root Server on :53")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start Root Server: %s", err)
	}
}