package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/saguywalker/sitcompetence/api"
	"github.com/saguywalker/sitcompetence/app"
)

/*
type authenMiddleware struct {
	tokenUsers map[string]*model.User
}

func (amw *authenMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		token = strings.TrimSpace(token)

		if len(token) == 0 {
			http.Error(w, "Missing auth token", http.StatusForbidden)
			return
		}

		if user, found := amw.tokenUsers[token]; found {
			log.Printf("Authenticated user %s\n", user.User)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "No user in session", http.StatusForbidden)
			return
		}
	})
}
*/
func serveAPI(ctx context.Context, api *api.API) {
	// var adminEntry = "ui/admin-sc/dist/admin/index.html"
	// var adminStatic = "ui/admin-sc/dist/"

	router := mux.NewRouter()
	api.Init(router.PathPrefix("/api").Subrouter())

	// router.PathPrefix("/admin").Handler(http.FileServer(http.Dir(adminStatic)))
	// router.PathPrefix("/admin").HandlerFunc(IndexHandler(adminEntry))

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8082"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(router)

	// amw := authenMiddleware{}

	// router.Use(amw.Middleware)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", api.Config.Port),
		Handler:     corsHandler,
		ReadTimeout: 2 * time.Minute,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("serving api at http://127.0.0.1:%d", api.Config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New()
		if err != nil {
			return err
		}
		defer a.Close()

		api, err := api.New(a)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			logrus.Info("signal caught. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer cancel()
			serveAPI(ctx, api)
		}()

		wg.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
