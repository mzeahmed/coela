# peskio

A Symfony project running on Docker, scaffolded with StackForge.

---

## Tech Stack

### Backend

* PHP 8.4
* Symfony

### Infrastructure

* Docker Compose
* Traefik v3
* Nginx
* MariaDB
* Redis
* Mailpit

---
## Local Domains

The application uses HTTPS locally.

Available domains:

```text
https://peskio.local
https://mail.peskio.local
https://db.peskio.local
```

Traefik sits in front of the stack as a reverse proxy. It:

* Terminates HTTPS using the local certificates generated with mkcert.
* Redirects all HTTP traffic to HTTPS.
* Routes each domain to the right container based on Host rules declared as
  Docker labels in `docker-compose.yml` (e.g. `peskio.local` -> nginx, `mail.peskio.local` -> Mailpit).

This avoids exposing container ports directly and lets multiple services
share ports 80/443 under different local domains.

These domains do not resolve on their own: you need to edit `/etc/hosts` to
point them to `127.0.0.1` (see [Configure Hosts](#configure-hosts)).

---
## Requirements

* Docker
* Docker Compose
* GNU Make
* mkcert

---

## Installation

### Clone the repository

```bash
git clone <your-repository-url>

cd peskio
```

### Configure Git Hooks

```bash
make setup
```
### Install Local SSL Authority

```bash
mkcert -install
```

### Generate Local Certificates

```bash
mkdir -p certs

cd certs

mkcert peskio.local "*.peskio.local"
```

### Configure Hosts

```bash
make hosts
```

This adds the required domains to `/etc/hosts` (asks for your `sudo`
password) and is safe to run multiple times: entries already present are
left untouched.

To do it manually instead, edit `/etc/hosts` and add:

```text
127.0.0.1 peskio.local
127.0.0.1 mail.peskio.local
127.0.0.1 db.peskio.local
```
### Start the Environment

```bash
make up
```

Application:

```text
https://peskio.local
```
Traefik Dashboard:

```text
http://localhost:8080
```
---

## Make Commands

### Display Available Commands

```bash
make help
```

### Build Containers

```bash
make build
```

### Start Containers

```bash
make up
```

### Stop Containers

```bash
make down
```

### Restart Environment

```bash
make restart
```

### View Logs

```bash
make logs
```

### List Containers

```bash
make ps
```

### Access PHP Container

```bash
make bash
```

### Run Code Style Check

```bash
make pint
```

### Fix Code Style Issues

```bash
make pintf
```

### Run Static Analysis

```bash
make stan
```

---

## Database

### Create Database

```bash
php bin/console doctrine:database:create
```

### Generate Migration

```bash
php bin/console make:migration
```

### Execute Migration

```bash
php bin/console doctrine:migrations:migrate
```

---
## Mailpit

Mailpit captures all outgoing emails.

Access:

```text
https://mail.peskio.local
```

---
## Redis

Redis is used for:

* Cache
* Sessions
* Messenger transports
* Rate limiting

---
## Project Structure

```text
peskio/
├── traefik/
├── certs/
├── docker/
│   ├── nginx/
│   └── php/
│
├── docs/
│
├── app/
│   ├── config/
│   ├── public/
│   ├── src/
│   ├── templates/
│   └── vendor/
│
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## License

MIT