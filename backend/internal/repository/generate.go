package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i SolutionRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i DeployRepository -o ./mocks/ -s "_minimock.go"
