package compose

import (
	"fmt"
	"github.com/HafslundEcoVannkraft/samplesystem/internal/model"
	composetypes "github.com/compose-spec/compose-go/v2/types"
	"path"
)

func applyApp(cCtx composeContext, app model.App) error {
	srv := &appService{
		app:           app,
		envVars:       map[string]string{},
		publishedPort: cCtx.nextPort(),
	}

	cCtx.addApp(app.Name, srv)
	return nil
}

type appService struct {
	app           model.App
	envVars       map[string]string
	publishedPort int
}

func (a *appService) name() string {
	return a.app.Name
}

func (a *appService) addEnvVars(env map[string]string) {
	for k, v := range env {
		a.envVars[k] = v
	}
}

func (a *appService) emitEnvVars(app composableApp) {
}

func (a *appService) serviceConfig() composetypes.ServiceConfig {
	env := map[string]*string{}
	for k, v := range a.envVars {
		env[k] = &v
	}

	return composetypes.ServiceConfig{
		Name:        a.name(),
		Environment: env,
		Build: &composetypes.BuildConfig{
			Context:    ".",
			Dockerfile: path.Join(a.app.Directory, a.app.Dockerfile),
		},
		Ports: []composetypes.ServicePortConfig{
			{
				Target:    uint32(a.app.Port),
				Published: fmt.Sprintf("%d", a.publishedPort),
			},
		},
	}
}
