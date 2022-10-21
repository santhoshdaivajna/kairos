package state

import (
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/block"
)

type MountPathState struct {
	Available, Mounted, ReadOnly bool
	Mountpoint                   string
}

type Runtime struct {
	Persistent block.Partition
	OEM        block.Partition
	State      block.Partition
}

func detectRuntimeState(r *Runtime) error {
	blockDevices, err := block.New(ghw.WithDisableTools(), ghw.WithDisableWarnings())
	if err != nil {
		return err
	}
	for _, d := range blockDevices.Disks {
		for _, part := range d.Partitions {
			switch part.Name {
			case "COS_PERSISTENT":
				r.Persistent = *part
			case "COS_OEM":
				r.OEM = *part
			case "COS_STATE":
				r.State = *part
			}
		}
	}
	return nil
}

func NewRuntime() (Runtime, error) {
	runtime := Runtime{}
	err := detectRuntimeState(&runtime)
	return runtime, err
}
