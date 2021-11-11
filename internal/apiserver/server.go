package apiserver

import (
	"github.com/RaulZuo/deep/internal/apiserver/config"
	"github.com/RaulZuo/deep/internal/apiserver/store"
	"github.com/RaulZuo/deep/internal/apiserver/store/mysql"
	genericoptions "github.com/RaulZuo/deep/internal/pkg/options"
	genericapiserver "github.com/RaulZuo/deep/internal/pkg/server"
)

type apiServer struct {
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}
// ExtraConfig defines extra configuration for the deep-apiserver.
type ExtraConfig struct {
	Addr		 string
	mysqlOptions *genericoptions.MySQLOptions
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	// 配置其他选项，例如数据库等
	extraConfig.complete().New()

	server := &apiServer{
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	// TODO
	//s.initRedisStore()

	//s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
	//	mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
	//	if mysqlStore != nil {
	//		return mysqlStore.Close()
	//	}
	//
	//	s.gRPCAPIServer.Close()
	//	s.genericAPIServer.Close()
	//
	//	return nil
	//}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	// TODO
	//go s.gRPCAPIServer.Run()

	//// start shutdown managers
	//if err := s.gs.Start(); err != nil {
	//	log.Fatalf("start shutdown manager failed: %s", err.Error())
	//}

	return s.genericAPIServer.Run()
}

type completedExtraConfig struct {
	*ExtraConfig
}

// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

func (c *completedExtraConfig) New() (interface {}, error) {
	storeIns, _ := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	store.SetClient(storeIns)
	return nil, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	// TODO
	//if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}

	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		mysqlOptions: cfg.MySQLOptions,
	}, nil
}
