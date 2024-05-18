package service

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/go-nlp/tfidf"
	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	infServ "github.com/khemmaphat/scented-secrets-api/src/service/infService"
	"github.com/xtgo/set"
	"gorgonia.org/tensor"
)

type QuestionService struct {
	questionRepository infRepo.IQuestionRepository
	perfumeRepository  infRepo.IPerfumeRepository
}

func MakeQuestionService(questionRepository infRepo.IQuestionRepository, perfumeRepository infRepo.IPerfumeRepository) infServ.IQuestionService {
	return &QuestionService{questionRepository: questionRepository, perfumeRepository: perfumeRepository}
}

func (s QuestionService) GetQuestions(ctx context.Context) ([]entities.Question, error) {
	questions, err := s.questionRepository.GetQuestions(ctx)
	return questions, err
}

func (s QuestionService) GetResultQuestion(ctx context.Context, answered string) (entities.ResultQuestion, error) {
	var resultQuestion entities.ResultQuestion

	perfumes, err := s.perfumeRepository.GetAllPerfume(ctx)
	if err != nil {
		return resultQuestion, err
	}

	corpus, _ := MakeCorpusForQuestions(perfumes)
	docs := MakeDocumentsForQuestions(perfumes, corpus)
	tf := tfidf.New()

	for _, doc := range docs {
		tf.Add(doc)
	}
	tf.CalculateIDF()

	test := doc{corpus[answered]}
	testScore := tf.Score(test)

	testDocs, testVec := contains(test, docs, tf)
	testRes := cosineSimilarity(testScore, testDocs, testVec)
	sort.Sort(sort.Reverse(testRes))

	fmt.Printf("\tID   : %d\n\tScore: %1.3f\n\tDoc  : %q\n", testRes[0].id, testRes[0].score, perfumes[testRes[0].id])

	name := perfumes[testRes[0].id].Name

	perfumeId, perfume, err := s.perfumeRepository.GetPerfumeByName(ctx, name)
	resultQuestion.PerfumeId = perfumeId
	resultQuestion.Name = perfume.Name
	resultQuestion.Brand = perfume.Brand

	PerfumeNotesTemp := strings.Split(perfume.Notes, ",")
	var PerfumeNotes entities.Notes
	for i, pt := range PerfumeNotesTemp {
		if i < 2 {
			PerfumeNotes.BaseNotes = append(PerfumeNotes.BaseNotes, pt)
		} else if i >= 2 && i <= 4 {
			PerfumeNotes.TopNotes = append(PerfumeNotes.TopNotes, pt)
		} else {
			PerfumeNotes.MiddleNotes = append(PerfumeNotes.MiddleNotes, pt)
		}

	}
	resultQuestion.Notes = PerfumeNotes
	resultQuestion.Description = perfume.Description
	resultQuestion.ImgUrl = perfume.ImgUrl

	return resultQuestion, nil
}

type doc []int

// IDs implements tfidf.Document.
func (d doc) IDs() []int {
	return []int(d)
}

func MakeCorpusForQuestions(perfumes []entities.Perfume) (map[string]int, []string) {
	retVal := make(map[string]int)
	invRetVal := make([]string, 0)
	var id int
	for _, perfume := range perfumes {
		for _, des := range strings.Fields(perfume.Description) {
			des = strings.ToLower(des)
			if _, ok := retVal[des]; !ok {
				retVal[des] = id
				invRetVal = append(invRetVal, des)
				id++
			}
		}
	}
	return retVal, invRetVal
}

func MakeDocumentsForQuestions(perfumes []entities.Perfume, corpus map[string]int) []tfidf.Document {

	retVal := make([]tfidf.Document, 0, len(perfumes))

	for _, perfume := range perfumes {
		var ts []int
		for _, des := range strings.Fields(perfume.Description) {
			des = strings.ToLower(des)
			id := corpus[des]
			ts = append(ts, id)
		}
		retVal = append(retVal, doc(ts))
	}
	return retVal
}

type docScore struct {
	id    int
	score float64
}

type docScores []docScore

func (ds docScores) Len() int           { return len(ds) }
func (ds docScores) Less(i, j int) bool { return ds[i].score < ds[j].score }
func (ds docScores) Swap(i, j int) {
	ds[i].score, ds[j].score = ds[j].score, ds[i].score
	ds[i].id, ds[j].id = ds[j].id, ds[i].id
}

func cosineSimilarity(queryScore []float64, docIDs []int, relVec []float64) docScores {
	// special case
	if len(docIDs) == 1 {
		// even more special case!
		if len(queryScore) == 1 {
			return docScores{
				{docIDs[0], queryScore[0] * relVec[0]},
			}
		}

		q := tensor.New(tensor.WithBacking(queryScore))
		m := tensor.New(tensor.WithBacking(relVec))
		score, err := q.Inner(m)
		if err != nil {
			panic(err)
		}
		return docScores{
			{docIDs[0], score.(float64)},
		}
	}

	m := tensor.New(tensor.WithShape(len(docIDs), len(queryScore)), tensor.WithBacking(relVec))
	q := tensor.New(tensor.WithShape(len(queryScore)), tensor.WithBacking(queryScore))
	dp, err := m.MatVecMul(q)
	if err != nil {
		panic(err)
	}

	m2, err := tensor.Square(m)
	if err != nil {
		panic(err)
	}

	normDocs, err := tensor.Sum(m2, 1)
	if err != nil {
		panic(err)
	}

	normDocs, err = tensor.Sqrt(normDocs)
	if err != nil {
		panic(err)
	}

	q2, err := tensor.Square(q)
	if err != nil {
		panic(err)
	}
	normQt, err := tensor.Sum(q2)
	if err != nil {
		panic(err)
	}
	normQ := normQt.Data().(float64)
	normQ = math.Sqrt(normQ)

	norms, err := tensor.Mul(normDocs, normQ)
	if err != nil {
		panic(err)
	}

	cosineSim, err := tensor.Div(dp, norms)
	if err != nil {
		panic(err)
	}
	csData := cosineSim.Data().([]float64)

	var ds docScores
	for i, id := range docIDs {
		score := csData[i]
		ds = append(ds, docScore{id: id, score: score})
	}
	return ds

}

func contains(query tfidf.Document, in []tfidf.Document, tf *tfidf.TFIDF) (docIDs []int, relVec []float64) {
	q := query.IDs()
	q = set.Ints(q) // unique words only
	for i := range in {
		doc := in[i].IDs()

		var count int
		var relevant []float64
		for _, wq := range q {
		inner:
			for _, wd := range doc {
				if wq == wd {
					count++
					break inner
				}
			}
		}
		if count == len(q) {
			// calculate the score of the doc
			score := tf.Score(in[i])
			// get the  scores of the relevant words
			for _, wq := range q {
			inner2:
				for j, wd := range doc {
					if wd == wq {
						relevant = append(relevant, score[j])
						break inner2
					}
				}
			}
			docIDs = append(docIDs, i)
			relVec = append(relVec, relevant...)
		}
	}
	return
}
