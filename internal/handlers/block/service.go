package block

import (
	"context"
	"database/sql"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ListUsers(ctx context.Context, req *ListBlockRequest) (status.Code, *ListBlockResponse, error) {
	result, err := s.repo.ListBlocks(ctx, req)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	if result != nil && len(result.Items) == 0 {
		result.Items = []*Block{}
	}

	return status.SUCCESS, result, nil
}

func (s *Service) CreateBlock(ctx context.Context, block *CreateBlockReq) (status.Code, error) {
	if statusCode, err := ValidateCreateBlockReq(block); err != nil {
		return statusCode, err
	}

	dto := BuildCeateBlockDTO(block)

	err := s.repo.StoreBlock(ctx, nil, dto)
	if err != nil {
		return status.INTERNAL, err
	}

	return status.SUCCESS, nil
}

func (s *Service) CreateBlockByValue(ctx context.Context, block *CreateBlockByValueReq) (status.Code, error) {
	handler := func(tx *sql.Tx) error {
		for _, item := range block.Items {
			dto := BuildCreateBlockByValueDTO(item.Type, item.Value)

			err := s.repo.StoreBlock(ctx, tx, dto)
			if err != nil {
				return err
			}
		}
		return nil
	}

	err := s.repo.StoreMultipleBlocks(ctx, handler)
	if err != nil {
		return status.INTERNAL, err
	}

	return status.SUCCESS, nil
}
