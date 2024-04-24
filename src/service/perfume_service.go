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
	if err != nil {
		return model.PerfumeDetail{}, err
	}

	perfumeNotesArray := strings.Split(perfume.Notes, ", ")
	perfumeNotesDetailArray, err := r.perfumeRepository.GetNotesFromArray(ctx, perfumeNotesArray)
	if err != nil {
		return model.PerfumeDetail{}, err
	}

	groupPerfumeNotesDetail := model.GroupNotePerfumeDetail{}
	for i, pt := range perfumeNotesDetailArray {
		if i < 2 {
			groupPerfumeNotesDetail.BaseNotes = append(groupPerfumeNotesDetail.BaseNotes, pt)
		} else if i >= 2 && i <= 4 {
			groupPerfumeNotesDetail.TopNotes = append(groupPerfumeNotesDetail.TopNotes, pt)
		} else {
			groupPerfumeNotesDetail.MiddleNotes = append(groupPerfumeNotesDetail.MiddleNotes, pt)
		}

	}

	perfumeDetail := model.PerfumeDetail{
		Name:        perfume.Name,
		Brand:       perfume.Brand,
		Notes:       groupPerfumeNotesDetail,
		PerfumeType: []string{"Fresh", "New", perfume.Gender},
		Description: perfume.Description,
		HowTo:       perfume.HowTo,
		Gender:      perfume.Gender,
		ImgUrl:      perfume.ImgUrl,
		CosineValue: perfume.CosineValue,
	}

	return perfumeDetail, nil
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

	//search by brand
	if req.SearchType != "Brand" && req.SearchGroup == "" {
		var filteredRes []entities.PerfumeResponse
		for _, val := range res {
			if strings.HasPrefix(val.Brand, req.Search) {
				filteredRes = append(filteredRes, val)
			}
		}

		res = filteredRes
	}

	//search by name
	if req.Search != "" && req.SearchType != "Brand" {
		var filteredRes []entities.PerfumeResponse
		for _, val := range res {
			if strings.HasPrefix(val.Name, req.Search) {
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
		case "CITRUS":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/1410/1410689.png"
		case "FRUITS, VEGETABLES AND NUTS":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/601/601984.png"
		case "FLOWERS":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/678/678100.png"
		case "MUSK, AMBER, ANIMALIC":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/672/672716.png"
		case "GREENS, HERBS AND FOUGERES":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/616/616742.png"
		case "SPICES":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/5862/5862835.png"
		case "SWEETS AND GOURMAND":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/2454/2454279.png"
		case "WOODS AND MOSSES":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/3105/3105220.png"
		case "RESINS AND BALSAMS":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/14435/14435222.png"
		case "MUSK, AMBER, ANIMALIC SMELLS":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/4277/4277682.png"
		case "BEVERAGES":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/4682/4682208.png"
		case "NATURAL AND SYNTHETIC, POPULAR AND WEIRD":
			groupNotes[i].ImgGroupUrl = "https://cdn-icons-png.flaticon.com/512/2823/2823511.png"
		}
	}

	return groupNotes, nil
}

func (s PerfumeService) GetResultMixedPerfume(ctx context.Context, answered string) (model.ResultMixedPerfume, error) {
	notes := entities.Notes{
		TopNotes:    []string{"Lemon", "Ginger"},
		MiddleNotes: []string{"Peony", "rose accord"},
		BaseNotes:   []string{"Musk"},
	}

	result := model.ResultMixedPerfume{
		Notes:       notes,
		Description: "Suitable for women who desire a cool, fresh scent followed by floral notes. Ideal for those who do not prefer overly sweet fragrances, can be sprayed for a casual day.",
	}

	return result, nil
}

func (s PerfumeService) GetPerfumeComment(ctx context.Context, perfumeId string) (model.AveragePerfumeComment, error) {
	comments, err := s.perfumeRepository.GetPerfumeComment(ctx, perfumeId)

	allSum := 0
	likeSum := 0
	okSum := 0
	dislikeSum := 0

	for _, comment := range comments {

		if comment.Rating > 3 {
			likeSum++
		} else if comment.Rating >= 2 && comment.Rating <= 3 {
			okSum++
		} else {
			dislikeSum++
		}

		allSum += comment.Rating
	}

	avgRating := float32(allSum) / float32(len(comments))
	avgLikeRating := float32(likeSum) / float32(len(comments))
	avgOkRating := float32(okSum) / float32(len(comments))
	avgDislikeRating := float32(dislikeSum) / float32(len(comments))

	avgRatingModel := model.AveragePerfumeComment{
		AverageRating: avgRating,
		LikeRating:    avgLikeRating,
		OkRating:      avgOkRating,
		DislikeRating: avgDislikeRating,
		Comments:      comments,
	}

	return avgRatingModel, err
}

func (s PerfumeService) GetPerfumePath() ([]entities.PerfumePath, error) {
	pathArray, err := s.perfumeRepository.GetAllPerfumeId()
	if err != nil {
		return nil, err
	}

	var perfumePaths []entities.PerfumePath

	for _, path := range pathArray {
		perfumePath := entities.PerfumePath{Path: path}
		perfumePaths = append(perfumePaths, perfumePath)
	}

	return perfumePaths, nil
}
