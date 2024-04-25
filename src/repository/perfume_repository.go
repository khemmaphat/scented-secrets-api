package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	"google.golang.org/api/iterator"
)

type PerfumeRepository struct {
	client *firestore.Client
}

func MakePerfumeRepository(client *firestore.Client) infRepo.IPerfumeRepository {
	return &PerfumeRepository{client: client}
}

func (r PerfumeRepository) GetPerfumeById(ctx context.Context, id string) (entities.Perfume, error) {
	var perfume entities.Perfume
	perfumeDoc, err := r.client.Collection("perfumes").Doc(id).Get(ctx)

	if err != nil {
		return perfume, err
	}

	if err := perfumeDoc.DataTo(&perfume); err != nil {
		return perfume, err
	}

	return perfume, nil
}

func (r PerfumeRepository) GetPerfumeByCosineValue(ctx context.Context, cosineValue float64) (string, entities.Perfume, error) {
	var perfume entities.Perfume
	var perfumeId string
	perfumeQuery := r.client.Collection("perfumes").Where("CosineValue", ">=", cosineValue).Limit(1)
	perfumeDoc := perfumeQuery.Documents(ctx)

	for {
		doc, err := perfumeDoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", perfume, err
		}

		if err := doc.DataTo(&perfume); err != nil {
			return "", perfume, err
		}

		perfumeId = doc.Ref.ID
	}

	return perfumeId, perfume, nil
}

func (r PerfumeRepository) SearchPerfume(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error) {
	query := r.client.Collection("perfumes").Where("CosineValue", ">=", req.CosineValue)

	if req.Search != "" {
		query = r.client.Collection("perfumes").Where("Name", ">=", req.Search)
	}

	var perfumes []entities.PerfumeResponse
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var perfume entities.PerfumeResponse
		if err := doc.DataTo(&perfume); err != nil {
			return nil, err
		}
		perfume.PerfumeId = doc.Ref.ID
		perfumes = append(perfumes, perfume)
	}
	return perfumes, nil
}

func (r PerfumeRepository) SearchPerfumeWithGender(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error) {
	query := r.client.Collection("perfumes").Where("CosineValue", ">=", req.CosineValue).Where("Gender", "in", []string{req.Gender, "Unisex"})

	if req.Search != "" {
		query = r.client.Collection("perfumes").Where("Name", ">=", req.Search).Where("Gender", "in", []string{req.Gender, "Unisex"})
	}

	var perfumes []entities.PerfumeResponse
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var perfume entities.PerfumeResponse
		if err := doc.DataTo(&perfume); err != nil {
			return nil, err
		}
		perfume.PerfumeId = doc.Ref.ID
		perfumes = append(perfumes, perfume)
	}
	return perfumes, nil
}

func (r PerfumeRepository) SearchPerfumeWithName(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error) {
	query := r.client.Collection("perfumes").Where("CosineValue", ">=", req.CosineValue)

	if req.Search != "" {
		query = r.client.Collection("perfumes").Where("Name", ">=", req.Search)
	}

	var perfumes []entities.PerfumeResponse
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var perfume entities.PerfumeResponse
		if err := doc.DataTo(&perfume); err != nil {
			return nil, err
		}
		perfume.PerfumeId = doc.Ref.ID
		perfumes = append(perfumes, perfume)
	}
	return perfumes, nil
}

func (r PerfumeRepository) SearchPerfumeWithBrand(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error) {
	query := r.client.Collection("perfumes").Where("CosineValue", ">=", req.CosineValue)

	if req.Search != "" {
		query = r.client.Collection("perfumes").Where("Brand", ">=", req.Search)
	}

	var perfumes []entities.PerfumeResponse
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var perfume entities.PerfumeResponse
		if err := doc.DataTo(&perfume); err != nil {
			return nil, err
		}
		perfume.PerfumeId = doc.Ref.ID
		perfumes = append(perfumes, perfume)
	}
	return perfumes, nil
}

func (r PerfumeRepository) SearchPerfumeNameWithBrand(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error) {
	query := r.client.Collection("perfumes").Where("Brand", "==", req.SearchGroup)

	if req.Search != "" {
		query = r.client.Collection("perfumes").Where("Name", ">=", req.Search).Where("Brand", "==", req.SearchGroup)
	}

	var perfumes []entities.PerfumeResponse
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var perfume entities.PerfumeResponse
		if err := doc.DataTo(&perfume); err != nil {
			return nil, err
		}
		perfume.PerfumeId = doc.Ref.ID
		perfumes = append(perfumes, perfume)
	}
	return perfumes, nil
}

func (r PerfumeRepository) AddPerfumeData(perfumes entities.Perfume) error {
	_, _, err := r.client.Collection("perfumes").Add(context.Background(), perfumes)
	return err
}

func (r PerfumeRepository) AddNoteData(note entities.Note) error {
	_, _, err := r.client.Collection("notes").Add(context.Background(), note)
	return err
}

func (r PerfumeRepository) GetAllNotes(ctx context.Context) ([]entities.Note, error) {
	var notes []entities.Note

	iter := r.client.Collection("notes").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var note entities.Note
		if err := doc.DataTo(&note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r PerfumeRepository) GetNotesFromArray(ctx context.Context, inputNotes []string) ([]entities.Note, error) {
	var notes []entities.Note
	iter := r.client.Collection("notes").Where("Name", "in", inputNotes).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var note entities.Note
		if err := doc.DataTo(&note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r PerfumeRepository) GetPerfumeComment(ctx context.Context, perfumeId string) ([]entities.PerfumeComment, error) {
	var comments []entities.PerfumeComment
	iter := r.client.Collection("comments").Where("perfumeId", "==", perfumeId).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var comment entities.PerfumeComment
		if err := doc.DataTo(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r PerfumeRepository) GetAllPerfumeId() ([]string, error) {
	iter := r.client.Collection("perfumes").Documents(context.Background())
	var pathArray []string

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		pathArray = append(pathArray, doc.Ref.ID)
	}
	return pathArray, nil
}
