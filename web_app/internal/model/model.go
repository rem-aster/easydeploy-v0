package model

type Solution struct {
    ID string `json:"id"`
    Label string `json:"label"`
    Description string `json:"description"`
}

type DeployState int

const (
    StateUnknown DeployState = iota
    StateReady
    StateDeploying
    StateUnknownError
    StateConnectionError
)

type DeployStatus struct {
    ID string
    State DeployState
}

var stateName = map[DeployState]string{
    StateUnknown: "unknown",
    StateReady: "ready",
    StateDeploying: "deploying",
    StateUnknownError: "unknown_error",
    StateConnectionError: "connection_error",
}

func (ss DeployState) String() string {
    return stateName[ss]
}