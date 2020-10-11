# yumfood-go
Backend REST API server implemented in golang. Yumfood can order various dishes
from various restaurants(aka vendors). API included CRUD for vendor and dish.
## Getting Started
### Clone repo
```bash
git clone https://github.com/aeramu/nim-itb-scraper
```
### Using Docker
This project included dockerfile, so you can use docker instead if it's easier.
```bash
sudo docker build -t yumfood
sudo docker run -p 8080:8080 yumfood
```
### Using Go
#### Prerequisites
Install Golang
```bash
sudo pacman -S go
```
#### Running
Just run it with Go. Golang will automatically install the dependency
```bash
go run ./cmd/yumfood
```
Or if you want to build the binary file first
```bash
go build ./cmd/yumfood
```