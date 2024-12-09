package compose

import (
	"fmt"
	"net/url"
	"time"

	composetypes "github.com/compose-spec/compose-go/v2/types"
)

const (
	lgtmServiceName = "otel-lgtm"
	lgtmImageName   = "grafana/otel-lgtm"

	// Internal publishedPort numbers for the LGTM container.
	portLgtmUI   = 3000
	portLgtmGRPC = 4317
	portLgtmHTTP = 4318
)

// applyLGTM adds the LGTM stack to the compose project.
func applyLGTM(cCtx composeContext) error {
	srv := &lgtmService{
		UIHostPort:   portLgtmUI,
		GRPCHostPort: portLgtmGRPC,
		HTTPHostPort: portLgtmHTTP,
	}

	//cfg := srv.ServiceConfig()

	cCtx.addApp(lgtmServiceName, srv)
	return nil

	//if xiter.Empty2(opts.RunnableApps()) {
	//	return project, nil
	//}
	//lgtm := lgtmService{
	//	Options:      opts,
	//	UIHostPort:   opts.AssignHostPort(LGTMServiceName, portLgtmUI, DefaultGrafanaPort),
	//	HTTPHostPort: opts.AssignHostPort(LGTMServiceName, portLgtmHTTP, useDefaultHostPort),
	//	GRPCHostPort: opts.AssignHostPort(LGTMServiceName, portLgtmGRPC, useDefaultHostPort),
	//}
	//project.Services[LGTMServiceName] = lgtm.ServiceConfig()
	//
	//if project.Services == nil {
	//	project.Services = make(composetypes.Services)
	//}
	//for _, app := range opts.RunnableApps() {
	//	svc := project.Services[app.Name.String()]
	//	if svc.DependsOn == nil {
	//		svc.DependsOn = make(composetypes.DependsOnConfig)
	//	}
	//	dep := svc.DependsOn[LGTMServiceName]
	//	dep.Condition = "service_healthy"
	//	dep.Required = true
	//	svc.DependsOn[LGTMServiceName] = dep
	//
	//	if svc, ok := project.Services[app.Name.String()]; ok {
	//		if svc.Environment == nil {
	//			svc.Environment = composetypes.NewMappingWithEquals(nil)
	//		}
	//		svc.Environment.OverrideBy(composetypes.MappingWithEquals{
	//			string(runtime.EnvOTELExporterOTLPEndpoint): ptr(lgtm.HTTPConnString(opts.IsExcepted(app.Name))),
	//			string(runtime.EnvOTELMetricExportInterval): ptr("5000"),
	//			string(runtime.EnvOTELServiceName):          ptr(app.Name.String()),
	//		})
	//	}
	//
	//	project.Services[app.Name.String()] = svc
	//}
	//
	//return project, nil
}

type lgtmService struct {
	//Options      *Options
	UIHostPort   uint16
	GRPCHostPort uint16
	HTTPHostPort uint16
}

func (s *lgtmService) name() string {
	return lgtmServiceName
}

func (s *lgtmService) serviceConfig() composetypes.ServiceConfig {
	return composetypes.ServiceConfig{
		Name:  lgtmServiceName,
		Image: lgtmImageName,
		HealthCheck: &composetypes.HealthCheckConfig{
			Test: composetypes.HealthCheckTest{
				"CMD",
				"curl",
				"-f",
				fmt.Sprintf("http://localhost:%d/health", portLgtmUI),
			},
			Interval: ptr(composetypes.Duration(time.Second * 10)),
			Timeout:  ptr(composetypes.Duration(time.Second * 10)),
			Retries:  ptr[uint64](5),
		},
		Ports: []composetypes.ServicePortConfig{
			{Published: fmt.Sprintf("%d", s.UIHostPort), Target: portLgtmUI},
			{Published: fmt.Sprintf("%d", s.GRPCHostPort), Target: portLgtmGRPC},
			{Published: fmt.Sprintf("%d", s.HTTPHostPort), Target: portLgtmHTTP},
		},
	}
}

// HTTPConnString creates the connection string for the HTTP service.
func (s *lgtmService) HTTPConnString(hostNetworking bool) string {
	host := "localhost"
	port := s.HTTPHostPort

	if !hostNetworking {
		host = lgtmServiceName
		port = portLgtmHTTP
	}

	q := make(url.Values)
	q.Set("sslmode", "disable")

	u := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", host, port),
	}

	return u.String()
}

func (s *lgtmService) addEnvVars(env map[string]string) {}

func (s *lgtmService) emitEnvVars(app composableApp) {
	app.addEnvVars(s.envVars())
}

func (s *lgtmService) envVars() map[string]string {
	return map[string]string{
		"OTEL_EXPORTER_OTLP_ENDPOINT": "http://localhost:4318",
		"OTEL_METRIC_EXPORT_INTERVAL": "5000",
		"OTEL_SERVICE_NAME":           "test",
	}
}

func ptr[T any](v T) *T {
	return &v
}
