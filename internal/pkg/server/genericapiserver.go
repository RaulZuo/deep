package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/RaulZuo/deep/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
	"time"
)

// GenericAPIServer contains state for a deep api server.
// type GenericAPIServer gin.Engine.
type GenericAPIServer struct {
	middlewares []string
	mode        string
	// SecureServingInfo holds configuration of the TLS server.
	// TODO
	//SecureServingInfo *SecureServingInfo

	// InsecureServingInfo holds configuration of the insecure HTTP server.
	InsecureServingInfo *InsecureServingInfo

	// ShutdownTimeout is the timeout used for server shutdown. This specifies the timeout before server
	// gracefully shutdown returns.
	ShutdownTimeout time.Duration

	*gin.Engine
	healthz         bool
	// TODO
	//enableMetrics   bool
	//enableProfiling bool
	// wrapper for gin.Engine

	insecureServer, secureServer *http.Server
}

func initGenericAPIServer(s *GenericAPIServer) {
	// do some setup
	// s.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Setup()
	// TODO
	//s.InstallMiddlewares()
	s.InstallAPIs()
}

// InstallAPIs install generic apis.
func (s *GenericAPIServer) InstallAPIs() {
	// install healthz handler
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			core.WriteResponse(c, nil, map[string]string{"status": "ok"})
		})
	}

	// TODO
	//// install metric handler
	//if s.enableMetrics {
	//	prometheus := ginprometheus.NewPrometheus("gin")
	//	prometheus.Use(s.Engine)
	//}
	//
	//// install pprof handler
	//if s.enableProfiling {
	//	pprof.Register(s.Engine)
	//}

	//s.GET("/version", func(c *gin.Context) {
	//	core.WriteResponse(c, nil, version.Get())
	//})
}

// Setup do some setup work for gin engine.
func (s *GenericAPIServer) Setup() {
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		// TODO
		//log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *GenericAPIServer) InstallMiddlewares() {
	// necessary middlewares
	// TODO
	//s.Use(middleware.RequestID())

	// install custom middlewares
	for _, m := range s.middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			//log.Warnf("can not find middleware: %s", m)
			continue
		}
		//log.Infof("install middleware: %s", m)
		s.Use(mw)
	}

	s.Use(middleware.Context())
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	s.insecureServer = &http.Server{
		Addr:	 s.InsecureServingInfo.Address,
		Handler: s,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	eg.Go(func() error {
		//log.Infof("Start to listening the incoming requests on http address: %s", s.InsecureServingInfo.Address)
		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			//log.Fatal(err.Error())
			return err
		}
		//log.Infof("Server on %s stopped", s.InsecureServingInfo.Address)
		return nil
	})

	// Ping the server to make sure the router is working.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if s.healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		//log.Fatal(err.Error())
	}

	return nil
}

// ping pings the http server to make sure the router is working.
func (s *GenericAPIServer) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.InsecureServingInfo.Address)
	if strings.Contains(s.InsecureServingInfo.Address, "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.InsecureServingInfo.Address, ":")[1])
	}

	for {
		// Change NewRequest to NewRequestWithContext and pass context it
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		// Ping the server by sending a GET request to `/healthz`.
		// nolint: gosec
		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			//log.Info("The router has been deployed successfully.")

			resp.Body.Close()

			return nil
		}

		// Sleep for a second to continue the next ping.
		//log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			//log.Fatal("can not ping http server within the specified time interval.")
		default:
		}
	}
}

