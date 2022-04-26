package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/googlecloudsupport"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(servicesCmd)
	rootCmd.AddCommand(clusterCmd)
	rootCmd.AddCommand(configureCmd)
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Enable google cloud services",
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range googlecloudsupport.EnableServicesCmd() {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s\n", s))
		}
	},
}

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")

		_, _ = fmt.Fprintf(cmd.OutOrStderr(), googlecloudsupport.CreateClusterCmd(resourcesDirectory, projectID, zone, clusterName))
	},
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")

		_, _ = fmt.Fprintf(cmd.OutOrStderr(), googlecloudsupport.ConfigureCmd(projectID, zone, clusterName))
	},
}

func must(variable string) string {
	if f := os.Getenv(variable); f == "" {
		panic(fmt.Sprintf("please set the %v environemnt variable.", variable))
	} else {
		return f
	}
}
