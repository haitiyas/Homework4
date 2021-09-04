# Homework 4

Homework 4 dengan membuat suatu REST API CRUD dengan menggunakan:
-Mux
-GORM
-MySQL

## Prerequisite

To run this program, you will need

### App Dependencies

```$xslt
- Golang 1.12+
- Go mod enabled
```

## How to Run

### Setup App Config

```
file .env diisi bagian JWT_SECRET dan ditambah DBDSN
cp .env.example .env
```
### migration
```
jalankan perintah "go run main.go migration" di terminal
```

### Build

```
run dengan perintah "go run . " di terminal
```
### test menggunakan postman
```
dengan localhost:80/movies
```
## Configuration

| NAME | DESCRIPTION | TYPE | VALUE
| ------ | ------ | ------ | ------ |
| APP_NAME | Application name | string | alphabet |
| APP_PORT | Application port | int | number |
| LOG_LEVEL | Mode for log level configuration | string | debug/info |
| ENVIRONMENT | Application environment | string | development |
| JWT_SECRET | JWT Secret | string | alphabet |
