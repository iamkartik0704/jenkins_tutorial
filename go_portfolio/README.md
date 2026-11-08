# 📁 Portfolio Web Applications (Go & Flask)

Portfolio implementations for **Kartik Chawla** (Software Engineer & B.Tech student at IIT Patna):

- **Live Reference Website**: [https://www.kartikchawla.in/](https://www.kartikchawla.in/)
- **LinkedIn**: [https://www.linkedin.com/in/iamkartik0704/](https://www.linkedin.com/in/iamkartik0704/)
- **GitHub**: [https://github.com/iamkartik0704](https://github.com/iamkartik0704)
- **Email**: [iamkartik0704@gmail.com](mailto:iamkartik0704@gmail.com)

---

## Implementations

1. **`go-portfolio/`** *(Recommended)*
   - High-performance Go implementation using standard library `net/http` and `embed.FS`.
   - Zero external dependencies.
   - Comprehensive unit tests in `main_test.go`.
   - Multi-stage Docker build producing a ~20MB container.
   - Declarative Jenkins CI/CD pipeline.

2. **`flask-portfolio/`**
   - Python Flask implementation with Jinja2 templates, Dockerfile, and Jenkins pipeline.

---

## Quick Start (Go Portfolio)

```bash
cd go-portfolio
go test -v ./...
go run main.go
```

Open [http://localhost:5000](http://localhost:5000).
