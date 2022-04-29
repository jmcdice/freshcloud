package cmds

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/googlecloudsupport"
	"github.com/spf13/cobra"
)

func init() {
	clustersCmd.AddCommand(googleClusterCmd)
	googleClusterCmd.AddCommand(googleServicesCmd)
	googleClusterCmd.AddCommand(googleClustersCreateCmd)
	googleClusterCmd.AddCommand(googleClustersListCmd)
	googleClusterCmd.AddCommand(googleClustersDeleteCmd)
	googleClusterCmd.AddCommand(googleConfigureCmd)
}

var googleClusterCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Manage google cloud clusters",
}

var googleServicesCmd = &cobra.Command{
	Use:   "enable-services",
	Short: "Enable google cloud API services",
	Run: func(cmd *cobra.Command, args []string) {
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.EnableServicesCmd(resourcesDirectory))
	},
}

var googleClustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.CreateClustersCmd(resourcesDirectory, projectID, zone, clusterName))
	},
}

var googleClustersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List google cloud clusters",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ListClustersCmd(projectID, zone))
	},
}

var googleClustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.DeleteClustersCmd(projectID, zone, clusterName))
	},
}

var googleConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ConfigureCmd(projectID, zone, clusterName))
	},
}