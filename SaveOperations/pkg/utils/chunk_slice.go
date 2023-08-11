package utils

import "github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/models"

func Chunk(r []models.Record, size int) [][]models.Record {
	var chunks [][]models.Record
	for i := 0; i < len(r); i += size {
		end := i + size
		if end > len(r) {
			end = len(r)
		}
		chunks = append(chunks, r[i:end])
	}
	return chunks
}
