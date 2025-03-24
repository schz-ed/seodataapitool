package cluster

import "my-seo-tool/internal/config"

// GetCluster returns the keywords for a given cluster name from the provided clusters config.
func GetCluster(clusters *config.ClustersConfig, name string) ([]string, bool) {
	keywords, ok := clusters.Clusters[name]
	return keywords, ok
}

// Additional functions (e.g., AddCluster, RemoveCluster) could be implemented here as needed.
