# ETL (Extract, Transform, Load) for Events

This project implements an ETL pipeline to fetch event data from an external API (KudaGo) and store it in a database using GORM.

## Project Structure

```bash
ETL/
├── cmd/                            #
│   ├── parse/                      #
│   │   └── main.go                 # Main ETL entry point
├── internal/                       #Internal logic
│   ├── db/                         # Database initialization and migration
│   │   └── init.go                 #
│   │   └── migrate.go              #
│   ├── models/                     # Data models + custom JSON unmarshal logic
│   │   └── models.go               #
│   ├── static/                     # Constants like FirstUrl and TotalRequests
│   │   └── constants.go            #
│   └── web/                        # HTTP request utilities
│   │   └── fetcher.go              #
└── README.md                       #
├── go.mod                          # Go module file
└── .gitignore                      # Git ignore file
```
## Installation

To install and run this project locally:

1. Clone the repository:

``` bash
git clone https://github.com/MrPixik/ETL.git
cd ETL
```
2. Install dependencies:

```bash
go mod tidy
```
3. Set up dsn parameters and run:

```bash
go run cmd/parse/main.go
```