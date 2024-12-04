package model

import (
	"github.com/bradleyjkemp/cupaloy"
	"testing"
)

func TestAssembleSystem(t *testing.T) {
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))

	sys, err := AssembleSystem(FromDirectory("./testdata"))
	if err != nil {
		t.Fatal(err)
	}

	snapshotter.SnapshotT(t, sys)

}
