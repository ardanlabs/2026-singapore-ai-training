// This example shows you how to use an LLM to create vector embeddings and
// get the same results from the handcrafted solution.
//
// # Running the example
//
//	$ make example02
//
// # Optional environment overrides
//
//	EMBED_SERVER  embeddings endpoint   (default: http://localhost:11435/v1/embeddings)
//	EMBED_MODEL   embedding model name  (default: embeddinggemma-300m-qat-Q8_0)
//
// # Prerequisites
//
//	$ make kronk-up

// Example 02 — Embeddings
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ardanlabs/2026-singapore-ai-training/foundation/client"
	"github.com/ardanlabs/2026-singapore-ai-training/foundation/vector"
)

var (
	embedURL   = "http://localhost:11435/v1/embeddings"
	embedModel = "embeddinggemma-300m-qat-Q8_0"
)

func init() {
	if v := os.Getenv("EMBED_SERVER"); v != "" {
		embedURL = v
	}

	if v := os.Getenv("EMBED_MODEL"); v != "" {
		embedModel = v
	}
}

// =============================================================================

type data struct {
	Name      string
	Text      string
	Embedding []float64 // The vector where the data is embedded in space.
}

// Vector can convert the specified data into a vector.
func (d data) Vector() []float64 {
	return d.Embedding
}

// =============================================================================

// =============================================================================

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Construct the llm client for access the model server.
	embedClient := client.NewLLM(embedURL, embedModel)

	// -------------------------------------------------------------------------
	// Embed data points
	// region Embed data points

	// Old way of representing this data with our own vector data points.
	// dataPoints := []vector.Data{
	// 	data{Name: "Horse   ", Authority: 0.0, Animal: 1.0, Human: 0.0, Rich: 0.0, Gender: +1.0},
	// 	data{Name: "Man     ", Authority: 0.0, Animal: 0.0, Human: 1.0, Rich: 0.0, Gender: -1.0},
	// 	data{Name: "Woman   ", Authority: 0.0, Animal: 0.0, Human: 1.0, Rich: 0.0, Gender: +1.0},
	// 	data{Name: "King    ", Authority: 1.0, Animal: 0.0, Human: 1.0, Rich: 1.0, Gender: -1.0},
	// 	data{Name: "Queen   ", Authority: 1.0, Animal: 0.0, Human: 1.0, Rich: 1.0, Gender: +1.0},
	// }

	// Apply the feature vectors to the handcrafted data points.
	// This time you need to use words since we are using a word based model.
	dataPoints := []vector.Data{
		data{Name: "Horse   ", Text: "Animal, Female"},
		data{Name: "Man     ", Text: "Human,  Male,   Pants, Poor, Worker"},
		data{Name: "Woman   ", Text: "Human,  Female, Dress, Poor, Worker"},
		data{Name: "King    ", Text: "Human,  Male,   Pants, Rich, Ruler"},
		data{Name: "Queen   ", Text: "Human,  Female, Dress, Rich, Ruler"},
	}

	fmt.Print("\n")

	// Iterate over each data point and use the LLM to generate the vector
	// embedding related to the model.
	for i, dp := range dataPoints {
		dataPoint := dp.(data)

		vect, err := embedClient.EmbedText(ctx, dataPoint.Text)
		if err != nil {
			return fmt.Errorf("embedding: %w", err)
		}

		dataPoint.Embedding = vect
		dataPoints[i] = dataPoint

		fmt.Printf("Vector: Name(%s) len(%d) %v...%v\n", dataPoint.Name, len(vect), vect[0:2], vect[len(vect)-2:])
	}

	fmt.Print("\n")

	// endregion

	return nil
}
