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

func serveAPI(ctx context.Context, api *api.API) {
	router := mux.NewRouter()
	api.Init(router.PathPrefix("/api").Subrouter())

	// CORS middleware
	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"*"},
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(router)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", api.Config.Port),
		Handler: corsHandler,
		// Handler: router,
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
