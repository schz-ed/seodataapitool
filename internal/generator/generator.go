package generator

import (
	"my-seo-tool/internal/config"
	"my-seo-tool/internal/utils"
	"path/filepath"
)

// GenerateKeywords creates the Cartesian product for a given template.
func GenerateKeywords(clusters map[string][]string, sequence []string) [][]string {
	if len(sequence) == 0 {
		return [][]string{}
	}
	result := [][]string{{}} // start with one empty combination
	for _, clusterName := range sequence {
		keywords := clusters[clusterName]
		var newResult [][]string
		for _, combo := range result {
			for _, kw := range keywords {
				newCombo := append([]string{}, combo...)
				newCombo = append(newCombo, kw)
				newResult = append(newResult, newCombo)
			}
		}
		result = newResult
	}
	return result
}

// GenerateAndSave generates keyword combinations for a template and writes to CSV.
func GenerateAndSave(cc *config.ClustersConfig, tmpl config.Template, outputDir string) error {
	combos := GenerateKeywords(cc.Clusters, tmpl.Sequence)
	records := [][]string{}
	for _, combo := range combos {
		records = append(records, combo)
	}
	filename := filepath.Join(outputDir, tmpl.Name+".csv")
	return utils.WriteCSV(filename, records)
}
