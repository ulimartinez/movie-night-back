package movies

import "github.com/gin-gonic/gin"

type SubmissionSerializer struct {
	c          *gin.Context
	Submission MovieSubmissionModel
}

type SubmissionResponse struct {
	ID    uint   `json:"id"`
	Movie string `json:"title"`
	Votes uint   `json:"votes"`
}

func (selfr *SubmissionSerializer) Response() SubmissionResponse {
	submissionModel := selfr.Submission
	movieModel, _ := GetMovie(MovieModel{ID: submissionModel.MovieID})
	submissionResponse := SubmissionResponse{
		ID:    submissionModel.ID,
		Movie: movieModel.Title,
		Votes: submissionModel.Votes,
	}
	return submissionResponse
}

type SubmissionsSerializer struct {
	c    *gin.Context
	subs []MovieSubmissionModel
}

func (selfr *SubmissionsSerializer) Response() []SubmissionResponse {
	response := []SubmissionResponse{}
	for _, sub := range selfr.subs {
		serializer := SubmissionSerializer{selfr.c, sub}
		response = append(response, serializer.Response())
	}
	return response
}
