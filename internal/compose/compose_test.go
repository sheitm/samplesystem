package compose

import (
	"github.com/HafslundEcoVannkraft/samplesystem/internal/model"
	"github.com/bradleyjkemp/cupaloy"
	"testing"
)

type composeTestCase struct {
	name   string
	system *model.System
}

var composeTestCases = []composeTestCase{
	{
		name: "simple",
		system: &model.System{
			Name:  "test",
			Owner: "test",
			Apps: []model.App{
				{
					Name:       "simple-app",
					Port:       80,
					Directory:  "/cmd/simple-app",
					Dockerfile: "Dockerfile",
				},
			},
		},
	},
}

func TestCompose(t *testing.T) {
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))

	for _, testCase := range composeTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			b, err := Compose(testCase.system)
			if err != nil {
				t.Fatal(err)
			}

			snapshotter.SnapshotT(t, b)
		})
	}
}
