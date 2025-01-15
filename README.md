# DNS Resolver Service

This project implements a custom DNS resolver and DNS servers (Root, TLD, and Authoritative servers). It uses Docker containers for modularity and ease of deployment. This guide explains how to set up and run the service using Docker Compose.

## Project Overview

The system consists of the following components:
1. **DNS Resolver**: Handles user queries and initiates recursive lookups.
2. **Root Server**: Directs queries to the appropriate TLD server.
3. **TLD Server**: Directs queries to the appropriate authoritative server.
4. **Authoritative Server**: Returns the actual IP address for a domain.

---

## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://www.docker.com/get-started) (v20.10+ recommended)
- [Docker Compose](https://docs.docker.com/compose/) (v2.0+ recommended)
- A working [Docker Hub](https://hub.docker.com/) account for pulling/pushing images

---

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/francisco-alonso/dns-resolver.git
   cd dns-resolver
   ```
2. Build the Docker images: If you don't have pre-built images, you can build them locally:
  ```bash
   docker-compose build
   ```

## Running the Service
To start the DNS service:

Launch all services using Docker Compose:
```
docker-compose up
```
Verify that all containers are running:

```
docker ps
```

Expected services:

- dns-resolver
- root-server
- tld-server
- auth-server

The DNS resolver will listen on port 53/udp by default.

