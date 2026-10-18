package agents

import (
	"encoding/json"
	"math"
	"os"
	"sort"
)

type Chunk struct {
	Text   string
	Vector []float32
	Source string
}

type store struct {
	Chunks []Chunk
}

func (s *store) Add(c Chunk) {
	s.Chunks = append(s.Chunks, c)
}

func (s *store) Search(query []float32, k int) []Chunk {
	type scored struct {
		chunk Chunk
		score float32
	}

	scored_list := make([]scored, 0, len(s.Chunks))
	for _, c := range s.Chunks {
		scored_list = append(scored_list, scored{c, consineSimiliarity(query, c.Vector)})
	}
	sort.Slice(scored_list, func(i, j int) bool {
		return scored_list[i].score > scored_list[j].score
	})
	if k > len(scored_list) {
		k = len(scored_list)
	}
	result := make([]Chunk, k)
	for i := 0; i < k; i++ {
		result[i] = scored_list[i].chunk
	}
	return result
}

func consineSimiliarity(a, b []float32) float32 {
	var dot, normA, normB float32
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	return dot / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))
}

func (s *store) SaveToFile(path string) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
func LoadFromFile(path string) (*store, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s store
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func NewStore(c []Chunk) *store {
	return &store{Chunks: c}
}
