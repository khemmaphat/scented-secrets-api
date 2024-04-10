package infRepo

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
)

type IPerfumeRepository interface {
	GetPerfumeById(ctx context.Context, id string) (entities.Perfume, error)
	AddPerfumeData(perfume entities.Perfume) error
	SearchPerfume(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error)
	SearchPerfumeWithGender(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error)
	SearchPerfumeWithName(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error)
	SearchPerfumeWithBrand(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error)
	SearchPerfumeNameWithBrand(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error)
	AddNoteData(note entities.Note) error
	GetAllNotes(ctx context.Context) ([]entities.Note, error)
}
