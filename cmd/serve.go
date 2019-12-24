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

func serveAPI(ctx context.Context, api *api.API, dev bool) {
	router := mux.NewRouter()
	api.Init(router.PathPrefix("/api").Subrouter())

	origins := []string{
		"https://sitcompetence.ilab.sit.kmutt.ac.th:3000",
		"https://sitcompetence.ilab.sit.kmutt.ac.th:443",
		"https://sitcompetence.ilab.sit.kmutt.ac.th:443",
		"http://sitcompetence.ilab.sit.kmutt.ac.th:3000",
		"http://sitcompetence.ilab.sit.kmutt.ac.th:443",
		"http://sitcompetence.ilab.sit.kmutt.ac.th:443",
		"https://localhost:3000",
		"https://localhost:443",
		"https://localhost:443",
		"http://localhost:3000",
		"http://localhost:443",
		"http://localhost:443",
		"http://sitcompetence.ilab.sit.kmutt.ac.th",
		"https://sitcompetence.ilab.sit.kmutt.ac.th",
		"https://sitcompetence.ilab.sit.kmutt.ac.th",
		"https://localhost",
		"http://localhost",
		"http://localhost",
		fmt.Sprintf("http://%s:3000", api.Config.IPAddress),
		fmt.Sprintf("http://%s:443", api.Config.IPAddress),
		fmt.Sprintf("http://%s:80", api.Config.IPAddress),
	}

	// CORS middleware
	var c *cors.Cors
	if dev {
		// c = cors.AllowAll()
		c = cors.New(cors.Options{
			AllowedOrigins: origins,
			AllowedHeaders: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodOptions,
				http.MethodDelete,
			},
			AllowCredentials: true,
		})

	} else {
		c = cors.New(cors.Options{
			AllowedOrigins: origins,
			AllowedHeaders: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodOptions,
				http.MethodDelete,
			},
			AllowCredentials: true,
		})
	}

	corsHandler := c.Handler(router)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", api.Config.Port),
		Handler: corsHandler,
		// Handler:     router,
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

	logrus.Infof("serving api at http://%s:%d", api.Config.IPAddress, api.Config.Port)
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
			serveAPI(ctx, api, false)
		}()

		wg.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
