package mongodbatlas

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// ContainerService provides methods for accessing MongoDB Atlas Containers API endpoints.
type ContainerService struct {
	sling *sling.Sling
}

// PrivateIPModeService provides many needfuls
type PrivateIPModeService struct {
	sling *sling.Sling
}

// newContainerService returns a new ContainerService.
func newContainerService(sling *sling.Sling) *ContainerService {
	return &ContainerService{
		sling: sling.Path("groups/"),
	}
}

func newPrivateIPModeService(sling *sling.Sling) *PrivateIPModeService {
	return &PrivateIPModeService{
		sling: sling.Path("groups/"),
	}
}

// Container represents a Cloud Services Containers in MongoDB.
type Container struct {
	ID             string `json:"id,omitempty"`
	ProviderName   string `json:"providerName,omitempty"`
	AtlasCidrBlock string `json:"atlasCidrBlock,omitempty"`
	RegionName     string `json:"regionName,omitempty"`
	VpcID          string `json:"vpcId,omitempty"`
	GcpProjectID   string `json:"gcpProjectId,omitempty"`
	NetworkName    string `json:"networkName,omitempty"`
	Provisioned    bool   `json:"provisioned,omitempty"`
}

// PrivateIPMode does needful in terms of the other needful
type PrivateIPMode struct {
	Enabled        bool `json:"enabled,omitempty"`
}

// containerListResponse is the response from the ContainerService.List.
type containerListResponse struct {
	Results    []Container `json:"results"`
	TotalCount int         `json:"totalCount"`
}

// List all containers for the specified group.
// https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/
func (c *ContainerService) List(gid string, providerName string) ([]Container, *http.Response, error) {
	response := new(containerListResponse)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/containers?providerName=%s", gid, providerName)
	resp, err := c.sling.New().Get(path).Receive(response, apiError)
	return response.Results, resp, relevantError(err, *apiError)
}

// Get a container in the specified group.
// https://docs.atlas.mongodb.com/reference/api/vpc-get-container/
func (c *ContainerService) Get(gid string, id string) (*Container, *http.Response, error) {
	container := new(Container)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/containers/%s", gid, id)
	resp, err := c.sling.New().Get(path).Receive(container, apiError)
	return container, resp, relevantError(err, *apiError)
}

// Create a container in the specified group.
// https://docs.atlas.mongodb.com/reference/api/vpc-create-container/
func (c *ContainerService) Create(gid string, containerParams *Container) (*Container, *http.Response, error) {
	container := new(Container)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/containers", gid)
	resp, err := c.sling.New().Post(path).BodyJSON(containerParams).Receive(container, apiError)
	return container, resp, relevantError(err, *apiError)
}

// Update a container in the specified group.
// https://docs.atlas.mongodb.com/reference/api/vpc-update-container/
func (c *ContainerService) Update(gid string, id string, containerParams *Container) (*Container, *http.Response, error) {
	container := new(Container)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/containers/%s", gid, id)
	resp, err := c.sling.New().Patch(path).BodyJSON(containerParams).Receive(container, apiError)
	return container, resp, relevantError(err, *apiError)
}

// Delete a container in the specified group.
func (c *ContainerService) Delete(gid string, id string) (*http.Response, error) {
	container := new(Container)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/containers/%s", gid, id)
	resp, err := c.sling.New().Delete(path).Receive(container, apiError)
	return resp, relevantError(err, *apiError)
}

// EnablePrivateIPMode enables needfuls
// https://docs.atlas.mongodb.com/reference/api/set-private-ip-mode-for-project/
func (p *PrivateIPModeService) EnablePrivateIPMode(gid string) ( *http.Response, error) {
	privateIPMode := new(PrivateIPMode)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/privateIpMode", gid)
	params := PrivateIPMode{
		Enabled: false,
	}
	resp, err := p.sling.New().Patch(path).BodyJSON(params).Receive(privateIPMode, apiError)
	return resp, relevantError(err, *apiError)
}

// DisablePrivateIPMode enables needfuls
// https://docs.atlas.mongodb.com/reference/api/set-private-ip-mode-for-project/
func (p *PrivateIPModeService) DisablePrivateIPMode(gid string) ( *http.Response, error) {
	privateIPMode := new(PrivateIPMode)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/privateIpMode", gid)
	params := PrivateIPMode{
		Enabled: false,
	}
	resp, err := p.sling.New().Patch(path).BodyJSON(params).Receive(privateIPMode, apiError)
	return resp, relevantError(err, *apiError)
}