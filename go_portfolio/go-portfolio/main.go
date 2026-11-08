package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//go:embed templates static
var content embed.FS

const (
	DefaultPort  = "5000"
	AuthorName   = "Kartik Chawla"
	TitleRole    = "Software Engineer • Backend Developer • IIT Patna"
	BioText      = "B.Tech student at IIT Patna passionate about backend architecture, building AI-native developer tools, and exploring Cloud & DevOps with Go."
	Location     = "Chandigarh & IIT Patna, India"
	LinkedInURL  = "https://www.linkedin.com/in/kartik-chawla-189430203/"
	GitHubURL    = "https://github.com/iamkartik0704"
	PortfolioURL = "https://www.kartikchawla.in/"
	Email        = "iamkartik0704@gmail.com"
)

// Playlist holds details for each featured learning playlist
type Playlist struct {
	Key      string       `json:"key"`
	Title    string       `json:"title"`
	ListID   string       `json:"list_id"`
	Gradient template.CSS `json:"gradient"`
	URL      string       `json:"url"`
}

// Project represents featured projects from kartikchawla.in
type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tech        string `json:"tech"`
	URL         string `json:"url"`
}

// Experience represents roles & journey
type Experience struct {
	Role         string `json:"role"`
	Organization string `json:"organization"`
	Period       string `json:"period"`
	Description  string `json:"description"`
}

// PageData is passed to HTML templates
type PageData struct {
	AuthorName   string
	TitleRole    string
	BioText      string
	Location     string
	LinkedInURL  string
	GitHubURL    string
	PortfolioURL string
	Email        string
	Skills       []string
	Projects     []Project
	Experiences  []Experience
	Playlists    map[string]Playlist
	Cards        []Playlist
	Current      *Playlist
	Year         int
}

// Server holds server dependencies
type Server struct {
	authorName     string
	titleRole      string
	bioText        string
	location       string
	linkedInURL    string
	gitHubURL      string
	portfolioURL   string
	email          string
	skills         []string
	projects       []Project
	experiences    []Experience
	playlists      map[string]Playlist
	playlistsOrder []string
	templates      map[string]*template.Template
}

// NewServer initializes and configures the portfolio server
func NewServer(fsys fs.FS) (*Server, error) {
	skills := []string{
		"Go (Golang)", "C / C++", "TypeScript", "Node.js", "Next.js",
		"Docker", "Jenkins", "DevOps", "REST APIs", "PostgreSQL", "Electron",
	}

	projects := []Project{
		{
			Title:       "⚡ Compile",
			Description: "AI-native desktop IDE featuring multi-LLM support and real-time environment compilation for developers.",
			Tech:        "Electron • React • TypeScript • AI Integrations",
			URL:         "https://github.com/iamkartik0704/ideweb",
		},
		{
			Title:       "🐸 ToadCode",
			Description: "A dynamically-typed scripting language built with a custom hand-written lexer, parser, and tree-walk interpreter.",
			Tech:        "TypeScript • Compilers • Language Engineering",
			URL:         "https://github.com/iamkartik0704",
		},
		{
			Title:       "🎟️ Celesta & TEDx Platforms",
			Description: "Full-stack event management platforms for IIT Patna featuring QR check-ins, ticket management, and SMTP authentication.",
			Tech:        "Next.js • Node.js • MongoDB • QR Scanner",
			URL:         "https://www.kartikchawla.in/",
		},
		{
			Title:       "🏢 Anwesha Accommodation",
			Description: "FCFS-based accommodation and room allotment backend platform for IIT Patna's annual cultural festival.",
			Tech:        "Node.js • Express • PostgreSQL • Zod",
			URL:         "https://www.kartikchawla.in/",
		},
	}

	experiences := []Experience{
		{
			Role:         "NJACK Sub-Coordinator @ DevOS",
			Organization: "IIT Patna",
			Period:       "June 2026 – Present",
			Description:  "Curating comprehensive web development roadmaps and mentoring freshers in open-source and modern software development.",
		},
		{
			Role:         "Web Dev Sub-Coordinator",
			Organization: "Celesta & TEDx IIT Patna",
			Period:       "April 2026 – Present",
			Description:  "Architected full-stack event platforms with real-time QR ticketing, check-in validation, and authentication.",
		},
		{
			Role:         "Web Dev Sub-Coordinator",
			Organization: "Anwesha, IIT Patna",
			Period:       "May 2026 – Present",
			Description:  "Engineered the FCFS accommodation and room allotment backend engine.",
		},
		{
			Role:         "Open Source Contributor",
			Organization: "GSSOC",
			Period:       "May 2026 – Present",
			Description:  "Implemented secure authentication, RESTful APIs, and Zod-based schema validations.",
		},
	}

	playlists := map[string]Playlist{
		"jenkins": {
			Key:      "jenkins",
			Title:    "🧩 Jenkins Pipeline",
			ListID:   "PLBr8obKbpkYvJEaPmrzhHhwx8uPj8WYbg",
			Gradient: template.CSS("linear-gradient(135deg, #ff0844 0%, #ffb199 100%)"),
			URL:      "/jenkins",
		},
		"docker": {
			Key:      "docker",
			Title:    "🐳 Docker Mastery",
			ListID:   "PLBr8obKbpkYsFtSF1XY9rM_3LH8LKRwSw",
			Gradient: template.CSS("linear-gradient(135deg, #00c6ff 0%, #0072ff 100%)"),
			URL:      "/docker",
		},
		"github": {
			Key:      "github",
			Title:    "💻 Git and GitHub",
			ListID:   "PLBr8obKbpkYt679NgO1KqZOoY_QYYLBl2",
			Gradient: template.CSS("linear-gradient(135deg, #8e2de2 0%, #4a00e0 100%)"),
			URL:      "/github",
		},
		"aws": {
			Key:      "aws",
			Title:    "☁️ AWS Cloud",
			ListID:   "PLBr8obKbpkYtn5eIK3iTgiV7MSgFwO-tr",
			Gradient: template.CSS("linear-gradient(135deg, #f7971e 0%, #ffd200 100%)"),
			URL:      "/aws",
		},
	}

	order := []string{"jenkins", "docker", "github", "aws"}

	// Custom template helper functions
	funcMap := template.FuncMap{
		"url_for": func(endpoint, filename string) string {
			if endpoint == "static" {
				return "/static/" + strings.TrimPrefix(filename, "/")
			}
			return "/" + filename
		},
		"css": func(s string) template.CSS {
			return template.CSS(s)
		},
	}

	templates := make(map[string]*template.Template)
	pages := []string{"index.html", "playlist.html", "404.html"}
	for _, page := range pages {
		tmpl, err := template.New("base.html").Funcs(funcMap).ParseFS(fsys, "templates/base.html", "templates/"+page)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template %s: %w", page, err)
		}
		templates[page] = tmpl
	}

	return &Server{
		authorName:     AuthorName,
		titleRole:      TitleRole,
		bioText:        BioText,
		location:       Location,
		linkedInURL:    LinkedInURL,
		gitHubURL:      GitHubURL,
		portfolioURL:   PortfolioURL,
		email:          Email,
		skills:         skills,
		projects:       projects,
		experiences:    experiences,
		playlists:      playlists,
		playlistsOrder: order,
		templates:      templates,
	}, nil
}

func (s *Server) cards() []Playlist {
	cards := make([]Playlist, 0, len(s.playlistsOrder))
	for _, key := range s.playlistsOrder {
		if p, ok := s.playlists[key]; ok {
			cards = append(cards, p)
		}
	}
	return cards
}

func (s *Server) pageData(current *Playlist) PageData {
	return PageData{
		AuthorName:   s.authorName,
		TitleRole:    s.titleRole,
		BioText:      s.bioText,
		Location:     s.location,
		LinkedInURL:  s.linkedInURL,
		GitHubURL:    s.gitHubURL,
		PortfolioURL: s.portfolioURL,
		Email:        s.email,
		Skills:       s.skills,
		Projects:     s.projects,
		Experiences:  s.experiences,
		Playlists:    s.playlists,
		Cards:        s.cards(),
		Current:      current,
		Year:         time.Now().Year(),
	}
}

func (s *Server) handleIndexOrPlaylist(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")

	// Root path -> render index page
	if path == "" {
		data := s.pageData(nil)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := s.templates["index.html"].ExecuteTemplate(w, "base.html", data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error rendering index: %v", err)
		}
		return
	}

	// Playlist path -> render playlist
	if p, ok := s.playlists[path]; ok {
		data := s.pageData(&p)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := s.templates["playlist.html"].ExecuteTemplate(w, "base.html", data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error rendering playlist %s: %v", path, err)
		}
		return
	}

	// Unknown path -> render 404
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := s.pageData(nil)
	_ = s.templates["404.html"].ExecuteTemplate(w, "base.html", data)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"service":   "go-portfolio",
		"author":    s.authorName,
		"portfolio": s.portfolioURL,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// loggingResponseWriter captures the status code for logging
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// loggingMiddleware logs each incoming HTTP request with timing and status code
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		log.Printf("%s %s %d (%v)", r.Method, r.URL.Path, lrw.statusCode, duration)
	})
}

// Routes sets up the HTTP router
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	// Static files from embedded filesystem
	fileServer := http.FileServer(http.FS(content))
	mux.Handle("/static/", fileServer)

	// Health check endpoint
	mux.HandleFunc("/healthz", s.handleHealth)

	// Index and playlist routes
	mux.HandleFunc("/", s.handleIndexOrPlaylist)

	return loggingMiddleware(mux)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	server, err := NewServer(content)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	addr := ":" + port
	httpServer := &http.Server{
		Addr:         addr,
		Handler:      server.Routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server run context for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Go Portfolio server starting on http://localhost:%s for %s (%s)", port, AuthorName, PortfolioURL)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped.")
}
