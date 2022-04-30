package googlecloudsupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/googlecloudsupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestEnableServices(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	servicesCmd := googlecloudsupport.EnableServicesCmd(resourcesDirectory)
	assert.Contains(t, servicesCmd[0], "gcloud services enable container.googleapis.com --quiet")
}

func TestCreateClusterCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := googlecloudsupport.CreateClustersCmd(resourcesDirectory, "aProject", "aZone", "aClusterName")
	expected := `gcloud container clusters create aClusterName --zone aZone --num-nodes 4`
	assert.Equal(t, clusterCmd[0], expected)
}

func TestListClustersCmdCmd(t *testing.T) {
	cmd := googlecloudsupport.ListClustersCmd("aProject", "aZone")
	assert.Equal(t, "gcloud container clusters list --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}

func TestConfigureCmd(t *testing.T) {
	cmd := googlecloudsupport.ConfigureCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters get-credentials 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}

func TestDeleteClustersCmdCmd(t *testing.T) {
	cmd := googlecloudsupport.DeleteClustersCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters delete 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}
