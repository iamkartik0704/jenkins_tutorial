# 🎨 Go Lang Portfolio Web App

This repository contains a modern, high-performance **Go (Golang) web portfolio** for **Kartik Chawla** (Software Engineer & B.Tech student at IIT Patna).

It is built using the standard Go `net/http` and `html/template` libraries with `embed.FS` for a zero-dependency, self-contained binary and a professional dark neon UI.

---

## 🚀 Features
- **Go-Powered Backend**: Built with standard library `net/http` (zero external dependencies).
- **Embedded Assets**: Templates and static assets embedded directly into the executable using Go's `embed.FS`.
- **Responsive Dark Neon UI**: Modern aesthetic with glowing borders and dynamic gradients.
- **Interactive Sections**:
  - Hero profile section with glowing portrait and links to LinkedIn, GitHub, and personal portfolio
  - Curated DevOps & Cloud Playlist cards (Jenkins, Docker, GitHub, AWS)
  - Embedded YouTube playlist players
  - Smooth auto-rotating carousel ("Life & Tech at IIT Patna")
- **Health Check Endpoint**: `/healthz` for container orchestrator readiness/liveness probes.
- **Graceful Shutdown**: Handles OS signals (`SIGINT`, `SIGTERM`) cleanly.
- **Unit Tested**: Full test coverage in `main_test.go` covering routes, static files, and error handling.
- **CI/CD Ready**: Multi-stage `Dockerfile` and declarative `Jenkinsfile`.

---

## 🧠 Tech Stack
- **Go (Golang)** — Standard Library (`net/http`, `html/template`, `embed`)
- **HTML5, CSS3, JavaScript** — Frontend
- **Docker** — Multi-stage minimal Alpine container
- **Jenkins** — Automated Build, Test, Docker Build & Deploy Pipeline

---

## ⚙️ Run Locally

### Prerequisites
- Go 1.22+ installed on your system.

### Steps
1. Navigate to the project directory:
   ```bash
   cd go_portfolio/go-portfolio
   ```

2. Run tests to verify everything passes:
   ```bash
   go test -v ./...
   ```

3. Start the application:
   ```bash
   go run main.go
   ```

4. Open your browser:
   [http://localhost:5000](http://localhost:5000)

*(Optional) You can customize the port via the `PORT` environment variable:*
```bash
PORT=8080 go run main.go
```

---

## 🐳 Run with Docker

Build and run using Docker (creates a lightweight ~20MB image):

```bash
# Build the Docker image
docker build -t iamkartik0704/go-portfolio .

# Run the container
docker run -p 5000:5000 iamkartik0704/go-portfolio
```

Access at [http://localhost:5000](http://localhost:5000).

---

## 🔄 Translation Summary: Flask vs. Go

| Feature | Flask (Python) | Go (Golang) |
| :--- | :--- | :--- |
| **Framework** | `Flask` (`requirements.txt`) | Standard Library (`net/http`) — zero dependencies |
| **Templates** | Jinja2 (`{% extends %}`, `{% block %}`) | Go `html/template` with block overrides |
| **CSS Safety** | Raw injection | `template.CSS` for secure dynamic CSS gradients |
| **Asset Delivery** | Disk-based `/static` folder | `embed.FS` compiled directly into the binary |
| **Testing** | Flake8 linting | Native `go test -v ./...` + `go vet` |
| **Docker Build** | Single-stage Python image (~160MB+) | Multi-stage build with Alpine runtime (~20MB) |
| **Performance** | Python WSGI | Native compiled binary with goroutine concurrency |

---

## 📁 Folder Structure

```
go-portfolio/
│
├── main.go            # HTTP server, routing, handlers, embedded FS
├── main_test.go       # Unit tests for all endpoints and assets
├── go.mod             # Go module definition
├── Dockerfile         # Multi-stage production container build
├── Jenkinsfile        # CI/CD pipeline for Jenkins
├── README.md          # Documentation
│
├── templates/
│   ├── base.html      # Base HTML layout with header, nav & footer
│   ├── index.html     # Homepage with hero, playlists & carousel
│   ├── playlist.html  # YouTube iframe player view
│   └── 404.html       # Themed 404 error page
│
└── static/
    ├── style.css      # Dark neon UI theme styling
    └── img/
        ├── hero.jpg   # Profile image
        ├── slide1.jpg # Carousel slide 1
        ├── slide2.jpg # Carousel slide 2
        ├── slide3.jpg # Carousel slide 3
        └── slide4.jpg # Carousel slide 4
```

---

## 👨‍💻 Author
**Kartik Chawla**  
Software Engineer | Backend Developer | IIT Patna  
🔗 [LinkedIn](https://www.linkedin.com/in/kartik-chawla-189430203/)  
💻 [GitHub](https://github.com/iamkartik0704)  
🌐 [Portfolio](https://kartikchawla.in)  
✉️ [Email](mailto:iamkartik0704@gmail.com)

---

## 📜 License
MIT License © 2025 Kartik Chawla
