package googlecloudsupport

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func EnableServicesCmd() []string {
	var commands []string
	services := []string{
		"cloudresourcemanager", "compute", "container", "datastore", "dns", "sqladmin", "storage-component",
	}
	for _, s := range services {
		commands = append(commands, fmt.Sprintf("gcloud services enable %s.googleapis.com", s))
	}
	return commands
}

func ConfigureCmd(projectId, zone, clusterName string) string {
	return fmt.Sprintf("gcloud container clusters get-credentials '%s' --zone '%v' --project '%v'\n", clusterName, zone, projectId)
}

func CreateClusterCmd(resourcesDirectory, projectId, zone, clusterName string) string {
	name := "google_cloud_cluster"
	data := struct {
		ProjectID   string
		Zone        string
		ClusterName string
	}{
		ProjectID:   projectId,
		Zone:        zone,
		ClusterName: clusterName,
	}
	return templatesupport.Parse(resourcesDirectory, name, data)
}
