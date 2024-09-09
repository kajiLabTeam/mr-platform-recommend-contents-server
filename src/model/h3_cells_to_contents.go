package model

import (
	"context"
	"fmt"

	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/lib"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/uber/h3-go/v4"
)

func H3CellsToContentIds(cells []h3.Cell) ([]string, error) {
	ctx, driver, err := lib.ConnectNeo4j()
	if err != nil {
		return nil, err
	}
	defer func() { err = lib.HandleClose(ctx, driver, err) }()
	session := driver.NewSession(ctx, neo4jSessionConfig)
	defer session.Close(ctx)

	var cellsToContentIds []string

	for _, cell := range cells {
		cellToContentIds, err := h3CellToContentIds(ctx, session, cell)
		if err != nil {
			return nil, err
		}
		cellsToContentIds = append(cellsToContentIds, cellToContentIds...)
	}

	return cellsToContentIds, nil
}

func h3CellToContentIds(ctx context.Context, session neo4j.SessionWithContext, cell h3.Cell) ([]string, error) {
	query := fmt.Sprintf(`MATCH (:H3_Cell_%d {cell: $cell})-[:H3CellToContentId]->(c:Content) RETURN c`, cell.Resolution())
	params := map[string]interface{}{
		"cell": cell.String(),
	}

	contentIds, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		txResult, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		records, err := txResult.Collect(ctx)
		if err != nil {
			return nil, err
		}

		// cyhper のキーの定義
		key := "c"
		var contentIds []string

		for _, record := range records {
			rawContentId, found := record.Get(key)
			if !found {
				return nil, fmt.Errorf("record does not have key %s", key)
			}

			contentNode, ok := rawContentId.(neo4j.Node)
			if !ok {
				return nil, fmt.Errorf("expected node but got %T", rawContentId)
			}

			contentId, found := contentNode.Props["content"]
			if !found {
				return nil, fmt.Errorf("node does not have key content")
			}

			contentIds = append(contentIds, contentId.(string))
		}
		return contentIds, nil
	})
	if err != nil {
		return nil, err
	}

	return contentIds.([]string), nil
}
