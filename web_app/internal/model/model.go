package model

type Solution struct {
    ID string `json:"id"`
    Label string `json:"label"`
    Description string `json:"description"`
}

type DeployState int

const (
    StateUnknown DeployState = iota
    StateRunning
    StateFailed
    StateSuccess
)

type DeployStatus struct {
    ID string
    State DeployState
}

var stateName = map[DeployState]string{
    StateUnknown: "unknown",
    StateRunning: "running",
    StateFailed: "failed",
    StateSuccess: "success",
}

func (ss DeployState) String() string {
    return stateName[ss]
}