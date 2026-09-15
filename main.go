package main

import (
	"2fa/internal/handler"
	"2fa/internal/repository"
	"2fa/internal/server"
	"2fa/internal/service"
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

//go:embed web/dist
var frontendAssets embed.FS

func main() {
	port := flag.Int("port", 2026, "HTTP server port")
	dataDir := flag.String("data", "./data", "Data directory for CVE2fa.db")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("[BOOT] CVE2fa starting...")
	log.Printf("[BOOT] runtime: %s/%s", runtime.GOOS, runtime.GOARCH)

	if err := repository.InitDB(*dataDir); err != nil {
		log.Fatalf("[FATAL] database initialization failed: %v", err)
	}
	defer repository.CloseDB()

	svr, err := NewServer(*port)
	if err != nil {
		log.Fatalf("[FATAL] server creation failed: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("[SHUTDOWN] shutting down...")
		repository.CloseDB()
		os.Exit(0)
	}()

	openBrowser(*port)

	if err := svr.Start(); err != nil {
		log.Fatalf("[FATAL] server error: %v", err)
	}
}

type Server struct {
	Port     int
	Auth     *handler.AuthHandler
	Accounts *handler.AccountHandler
	Export   *handler.ExportHandler
	Import   *handler.ImportHandler
	AuthMW   *server.AuthMiddleware
}

func NewServer(port int) (*Server, error) {
	userRepo := repository.NewUserRepo()
	accountRepo := repository.NewAccountRepo()

	authService := service.NewAuthService(userRepo)
	totpService := service.NewTOTPService()
	accountService := service.NewAccountService(accountRepo)
	cryptoService := service.NewCryptoService()

	if err := authService.SeedAdmin(); err != nil {
		return nil, err
	}

	return &Server{
		Port:     port,
		Auth:     handler.NewAuthHandler(authService),
		Accounts: handler.NewAccountHandler(accountService, totpService),
		Export:   handler.NewExportHandler(accountService, cryptoService),
		Import:   handler.NewImportHandler(cryptoService, accountRepo),
		AuthMW:   server.NewAuthMiddleware(authService),
	}, nil
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth/login", s.Auth.Login)
	mux.HandleFunc("POST /api/auth/logout", s.Auth.Logout)
	mux.Handle("PUT /api/auth/password", s.AuthMW.RequireAuth(http.HandlerFunc(s.Auth.ChangePassword)))
	mux.Handle("GET /api/auth/me", s.AuthMW.RequireAuth(http.HandlerFunc(s.Auth.Me)))

	mux.Handle("GET /api/accounts", s.AuthMW.RequireAuth(http.HandlerFunc(s.Accounts.List)))
	mux.Handle("POST /api/accounts", s.AuthMW.RequireAuth(http.HandlerFunc(s.Accounts.Create)))
	mux.Handle("DELETE /api/accounts/{id}", s.AuthMW.RequireAuth(http.HandlerFunc(s.Accounts.Delete)))
	mux.Handle("GET /api/accounts/{id}/code", s.AuthMW.RequireAuth(http.HandlerFunc(s.Accounts.GetCode)))

	mux.Handle("POST /api/export", s.AuthMW.RequireAuth(http.HandlerFunc(s.Export.Export)))
	mux.Handle("POST /api/import", s.AuthMW.RequireAuth(http.HandlerFunc(s.Import.Import)))

	subFS, err := fs.Sub(frontendAssets, "web/dist")
	if err != nil {
		log.Println("[WARN] frontend dist not embedded, serving API only")
		mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>CVE2fa API</h1><p>Frontend not built. Run: cd web && npm install && npm run build</p>"))
		})
	} else {
		fileServer := http.FileServer(http.FS(subFS))
		mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api/") {
				http.NotFound(w, r)
				return
			}
			if isStaticAsset(r.URL.Path) || r.URL.Path == "/" {
				fileServer.ServeHTTP(w, r)
				return
			}
			if _, err := fs.Stat(subFS, r.URL.Path[1:]); err == nil {
				fileServer.ServeHTTP(w, r)
				return
			}
			content, err := fs.ReadFile(subFS, "index.html")
			if err != nil {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(content)
		}))
	}

	handler := server.CORSMiddleware(mux)

	addr := ":" + itoa(s.Port)
	log.Printf("[SERVER] listening on http://localhost%s", addr)
	return http.ListenAndServe(addr, handler)
}

func openBrowser(port int) {
	url := "http://localhost:" + itoa(port)

	go func() {
		switch runtime.GOOS {
		case "windows":
			execCmd("rundll32", "url.dll,FileProtocolHandler", url)
		case "darwin":
			execCmd("open", url)
		default:
			execCmd("xdg-open", url)
		}
	}()
}

func execCmd(name string, args ...string) {
	attr := &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}
	proc, err := os.StartProcess(name, append([]string{name}, args...), attr)
	if err != nil {
		log.Printf("[BROWSER] could not open browser: %v", err)
		return
	}
	proc.Release()
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	s := ""
	for i > 0 {
		s = string(rune('0'+i%10)) + s
		i /= 10
	}
	return s
}

func isStaticAsset(path string) bool {
	extensions := []string{".html", ".css", ".js", ".svg", ".png", ".ico", ".woff", ".woff2", ".ttf", ".eot"}
	for _, ext := range extensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
