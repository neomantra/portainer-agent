package exec

import (
	"context"

	"github.com/portainer/agent"
	libstack "github.com/portainer/docker-compose-wrapper"
	"github.com/portainer/docker-compose-wrapper/compose"
)

// DockerComposeStackService represents a service for managing stacks by using the Docker binary.
type DockerComposeStackService struct {
	deployer   libstack.Deployer
	binaryPath string
}

// NewDockerComposeStackService initializes a new DockerStackService service.
// It also updates the configuration of the Docker CLI binary.
func NewDockerComposeStackService(binaryPath string) (*DockerComposeStackService, error) {
	deployer, err := compose.NewComposeDeployer(binaryPath, "")
	if err != nil {
		return nil, err
	}

	service := &DockerComposeStackService{
		binaryPath: binaryPath,
		deployer:   deployer,
	}

	return service, nil
}

// Deploy executes the docker stack deploy command.
func (service *DockerComposeStackService) Deploy(ctx context.Context, name string, filePaths []string, options agent.DeployOptions) error {
	return service.deployer.Deploy(ctx, filePaths, libstack.DeployOptions{
		Options: libstack.Options{
			ProjectName: name,
		},
	})
}

// Pull executes the docker pull command.
func (service *DockerComposeStackService) Pull(ctx context.Context, name string, filePaths []string) error {
	return service.deployer.Pull(ctx, filePaths, libstack.Options{
		ProjectName: name,
	})
}

// Remove executes the docker stack rm command.
func (service *DockerComposeStackService) Remove(ctx context.Context, name string, filePaths []string, options agent.RemoveOptions) error {
	return service.deployer.Remove(ctx, filePaths, libstack.Options{
		ProjectName: name,
	})
}
