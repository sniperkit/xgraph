package bolt

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"context"
)

// get all quads from the db
// func (a *AdminService) Quads() ([]interface{}, error) {
func (a *AdminService) Quads() ([]quad.Quad, error) {
	results, err := readAllQuads(a.Store)

	if err != nil {
		return results, err
	}

	return results, nil
}

func readAllQuads(store *cayley.Handle) ([]quad.Quad, error) {
	var results []quad.Quad
	it := store.QuadsAllIterator()
	defer it.Close()

	ctx := context.TODO()

	for it.Next(ctx) {
		results = append(results, store.Quad(it.Result()))
	}

	return results, nil
}
