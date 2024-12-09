package compose

import (
	"github.com/HafslundEcoVannkraft/samplesystem/internal/model"
	composetypes "github.com/compose-spec/compose-go/v2/types"
	"gopkg.in/yaml.v2"
)

func Compose(sys *model.System) ([]byte, error) {
	cCtx := newContext(3000)

	for _, app := range sys.Apps {
		if err := applyApp(cCtx, app); err != nil {
			return nil, err
		}
	}

	err := applyLGTM(cCtx)
	if err != nil {
		return nil, err
	}

	apps := cCtx.apps()
	for name, app := range apps {
		for n, a := range apps {
			if name == n {
				continue
			}
			app.emitEnvVars(a)
		}
	}

	services := map[string]composetypes.ServiceConfig{}
	for _, app := range cCtx.apps() {
		services[app.name()] = app.serviceConfig()
	}

	proj := composetypes.Project{
		Name:             sys.Name,
		WorkingDir:       "",
		Services:         services,
		Networks:         nil,
		Volumes:          nil,
		Secrets:          nil,
		Configs:          nil,
		Extensions:       nil,
		ComposeFiles:     nil,
		Environment:      nil,
		DisabledServices: nil,
		Profiles:         nil,
	}

	b, err := yaml.Marshal(proj)
	if err != nil {
		return nil, err
	}

	return b, nil
}
