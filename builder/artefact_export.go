package profitbricks

import (
	"fmt"
)

// dummy Artifact implementation - does nothing
type PBArtifact struct {
	// The name of the snapshot
	snapshotData string
}

func (*PBArtifact) BuilderId() string {
	return BuilderId
}

func (a *PBArtifact) Files() []string {
	return []string{}
}

func (*PBArtifact) Id() string {
	return "Null"
}

func (a *PBArtifact) String() string {
	return fmt.Sprintf("A snapshot was created: '%v'", a.snapshotData)
}

func (a *PBArtifact) State(name string) interface{} {
	return nil
}

func (a *PBArtifact) Destroy() error {
	return nil
}
