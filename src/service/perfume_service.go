package service

import (
	"context"
	"strings"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
	"github.com/khemmaphat/scented-secrets-api/src/model"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	infServ "github.com/khemmaphat/scented-secrets-api/src/service/infService"
)

type PerfumeService struct {
	perfumeRepository infRepo.IPerfumeRepository
}

func MakePerfumeService(perfumeRepository infRepo.IPerfumeRepository) infServ.IPerfumeService {
	return &PerfumeService{perfumeRepository: perfumeRepository}
}

func (r PerfumeService) GetPerfumeById(ctx context.Context, id string) (model.PerfumeDetail, error) {
	perfume, err := r.perfumeRepository.GetPerfumeById(ctx, id)
	PerfumeNotesTemp := strings.Split(perfume.Notes, ",")
	var PerfumeNotes model.Notes

	for i, pt := range PerfumeNotesTemp {
		if i < 2 {
			PerfumeNotes.BaseNotes = append(PerfumeNotes.BaseNotes, pt)
		} else if i >= 2 && i <= 4 {
			PerfumeNotes.TopNotes = append(PerfumeNotes.TopNotes, pt)
		} else {
			PerfumeNotes.MiddleNotes = append(PerfumeNotes.MiddleNotes, pt)
		}

	}

	PerfumeDetail := model.PerfumeDetail{
		Name:        perfume.Name,
		Brand:       perfume.Brand,
		Description: perfume.Description,
		HowTo:       perfume.Description,
		Notes:       PerfumeNotes,
		Gender:      perfume.Gender,
		ImgUrl:      perfume.ImgUrl,
		CosineValue: perfume.CosineValue,
		PerfumeType: []string{"Relax", "Fresh", "Fruit"},
	}
	return PerfumeDetail, err
}

func (s PerfumeService) AddPerfumeData(perfumes []entities.Perfume) error {
	for _, perfume := range perfumes {
		err := s.perfumeRepository.AddPerfumeData(perfume)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s PerfumeService) AddNoteData(notes []entities.Note) error {
	for _, note := range notes {
		err := s.perfumeRepository.AddNoteData(note)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s PerfumeService) SearchPerfumePagination(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, int, error) {
	var res []entities.PerfumeResponse
	var err error

	switch req.SearchType {
	case "":
		res, err = s.perfumeRepository.SearchPerfume(ctx, req)
	case "Gender":
		res, err = s.perfumeRepository.SearchPerfumeWithGender(ctx, req)
	case "Name":
		res, err = s.perfumeRepository.SearchPerfumeWithName(ctx, req)
	case "Brand":
		res, err = s.SearchPerfumePaginationWithBrand(ctx, req)
	case "Group":
		res, err = s.perfumeRepository.SearchPerfumeNameWithBrand(ctx, req)
	}

	if err != nil {
		return nil, 0, err
	}

	if req.Search != "" {
		var filteredRes []entities.PerfumeResponse
		for _, val := range res {
			if strings.Contains(val.Name, req.Search) {
				filteredRes = append(filteredRes, val)
			}
		}

		res = filteredRes
	}

	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := req.PageNum * req.PageSize

	if endIndex > len(res) {
		endIndex = len(res)
	}

	total := len(res)

	return res[startIndex:endIndex], total, nil
}

func (s PerfumeService) SearchPerfumePaginationWithBrand(ctx context.Context, req entities.PerfumePaginationRequest) ([]entities.PerfumeResponse, error) {
	res, err := s.perfumeRepository.SearchPerfumeWithBrand(ctx, req)
	if err != nil {
		return nil, err
	}

	uniqueBrands := make(map[string]bool)
	var uniqueBrandList []entities.PerfumeResponse

	for _, perfume := range res {
		if _, found := uniqueBrands[perfume.Brand]; !found {
			uniqueBrands[perfume.Brand] = true
			uniqueBrandList = append(uniqueBrandList, perfume)
		}
	}

	return uniqueBrandList, nil
}

func (s PerfumeService) GetAllNoteGroup(ctx context.Context) ([]model.GroupNotes, error) {
	notes, err := s.perfumeRepository.GetAllNotes(ctx)
	if err != nil {
		return nil, err
	}

	groupMap := make(map[string]*model.GroupNotes)
	for _, note := range notes {
		group, ok := groupMap[note.Group]
		if !ok {
			group = &model.GroupNotes{Name: note.Group}
			groupMap[note.Group] = group
		}
		group.Notes = append(group.Notes, note)
	}

	var groupNotes []model.GroupNotes
	for _, group := range groupMap {
		groupNotes = append(groupNotes, *group)
	}

	for i := range groupNotes {
		switch groupNotes[i].Name {
		case "Citrus":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/1410/1410689.png"
		case "Fruits":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/601/601984.png"
		case "Flowers":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/678/678100.png"
		case "White Flowers":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/6900/6900993.png"
		case "Herbs":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/616/616742.png"
		case "Spices":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/5862/5862835.png"
		case "Sweets":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/2454/2454279.png"
		case "Woods":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/3105/3105220.png"
		case "Resins":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/14435/14435222.png"
		case "Musk":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/4277/4277682.png"
		case "Berverages":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/4682/4682208.png"
		case "Netural":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/2823/2823511.png"
		}
	}

	return groupNotes, nil
}
