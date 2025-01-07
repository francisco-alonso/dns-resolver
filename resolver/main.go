package main

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
    msg := new(dns.Msg)
    msg.SetReply(r)
    msg.Compress = false

    domain := msg.Question[0].Name
    log.Printf("Received query for domain: %s", domain)

    // Placeholder response for now
    msg.Answer = []dns.RR{
        &dns.A{
            Hdr: dns.RR_Header{
                Name:   domain,
                Rrtype: dns.TypeA,
                Class:  dns.ClassINET,
                Ttl:    600,
            },
            A: net.IPv4(127, 0, 0, 1),
        },
    }

    w.WriteMsg(msg)
}

func main() {
    server := &dns.Server{Addr: ":53", Net: "udp"}

    dns.HandleFunc(".", handleDNSRequest)

    fmt.Println("Starting DNS Resolver on :53")
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
