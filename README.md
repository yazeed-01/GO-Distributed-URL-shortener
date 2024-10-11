# Distributed URL Shortener Service

This project is a distributed URL shortening service built with Go using the Gin framework. It features Snowflake ID generation for creating unique short URLs, QR code generation for easy sharing, and a load balancer setup using Nginx to ensure scalability and high availability.

## Features
- **Snowflake ID Generation**: Generate unique Snowflake IDs and convert them into a shortened URL with 7-8 characters using Base62 encoding.
- **URL Shortening**: Submit long URLs to get a short URL. The short URL can redirect users back to the original URL.
- **QR Code Generation**: Automatically generate a QR code for each short URL.
- **Load Balancing**: Two or more servers are set up behind an Nginx load balancer to handle traffic and ensure continuous service availability.
- **Database**: Store URL mappings in a shared PostgreSQL database, which all servers can access.
- **High Availability**: The system can handle failures on one server while another server continues serving requests.

## Architecture
- **Go with Gin**: Used for the web application and REST API.
- **PostgreSQL**: Shared database across multiple servers to store URL mappings.
- **Nginx**: Load balancer to distribute traffic between multiple servers (e.g., ports 8080 and 8081).


## Packages Used

- `github.com/bwmarrin/snowflake`
- `github.com/gin-gonic/gin`
- `github.com/fsnotify/fsnotify`
- `github.com/githubnemo/CompileDaemon`
- `github.com/itchyny/base58-go`
- `github.com/joho/godotenv`
- `github.com/jinzhu/inflection`
- `github.com/jinzhu/now`
- `github.com/skip2/go-qrcode`
- `gorm.io/gorm`
- `gorm.io/driver/postgres`
- `github.com/gin-contrib/sse`
- `github.com/go-playground/validator/v10`
- `github.com/fatih/color`
- `github.com/gabriel-vasile/mimetype`
- `github.com/go-playground/locales`
- `github.com/go-playground/universal-translator`
- `github.com/leodido/go-urn`
- `github.com/pelletier/go-toml/v2`
