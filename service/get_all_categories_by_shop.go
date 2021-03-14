package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	gen "gitlab.com/cadaverine/pim-service/gen/go/api/pim-service"
	"gitlab.com/cadaverine/pim-service/models"
)

// GetAllCategoriesByShop ...
func (s *PimService) GetAllCategoriesByShop(ctx context.Context, req *gen.ShopID) (*gen.CategoriesTrees, error) {
	categories, err := s.getAllCategoriesFlat(ctx, int(req.GetShopID()))
	if err != nil {
		return nil, err
	}

	return &gen.CategoriesTrees{
		Categories: repackCategoriesToProto(categories),
	}, nil
}

func repackCategoriesToProto(src []*models.Category) []*gen.Category {
	roots := make([]*gen.Category, 0)
	itemsMap := make(map[int]*gen.Category)

	for _, item := range src {
		repacked := &gen.Category{
			ID:   int32(item.ID),
			Name: item.Title,
		}

		if itemsMap[item.ID] == nil {
			itemsMap[item.ID] = repacked
		} else {
			itemsMap[item.ID].Name = item.Title
		}

		if itemsMap[item.ParentID] == nil {
			itemsMap[item.ParentID] = &gen.Category{ID: int32(item.ParentID)}
		}
		itemsMap[item.ParentID].Children = append(itemsMap[item.ParentID].Children, repacked)

		if item.ParentID == 0 {
			roots = append(roots, repacked)
		}
	}

	return roots
}

func (s *PimService) getAllCategoriesFlat(ctx context.Context, shopID int) ([]*models.Category, error) {
	const query = `
		select item_id as id, parent_id, name as title
		from product_information.categories
		where shop_id = $1
	`

	var categories []*models.Category

	err := s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		return tx.Select(&categories, query, shopID)
	})

	return categories, err
}
