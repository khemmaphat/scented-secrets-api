package infService

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
	"github.com/khemmaphat/scented-secrets-api/src/model"
)

type IPerfumeService interface {
	GetPerfumeById(ctx context.Context, id string) (model.PerfumeDetail, error)
	AddPerfumeData(perfume []entities.Perfume) error
	SearchPerfumePagination(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, int, error)
	AddNoteData(notes []entities.Note) error
	GetAllNoteGroup(ctx context.Context) ([]model.GroupNotes, error)
}
