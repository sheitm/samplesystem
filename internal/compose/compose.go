package compose

import (
	composetypes "github.com/compose-spec/compose-go/v2/types"
	"gopkg.in/yaml.v2"
)

func Compose() ([]byte, error) {
	proj := composetypes.Project{
		Name:             "",
		WorkingDir:       "",
		Services:         nil,
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

	proj, err := applyLGTM(proj)
	if err != nil {
		return nil, err
	}

	b, err := yaml.Marshal(proj)
	if err != nil {
		return nil, err
	}

	return b, nil
}
