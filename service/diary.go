package service

import (
	"context"
	"errors"
	"log"

	diary "github.com/shibayu36/go-diary-api/gen/diary"
	"github.com/shibayu36/go-diary-api/model"
	"github.com/shibayu36/go-diary-api/repository"
)

// diary service example implementation.
// The example methods log the requests and return zero values.
type diarysrvc struct {
	logger *log.Logger
	repos  *repository.Repositories
}

// NewDiary returns the diary service implementation.
func NewDiary(logger *log.Logger, repos *repository.Repositories) diary.Service {
	return &diarysrvc{logger, repos}
}

// UserSignup implements UserSignup.
func (s *diarysrvc) UserSignup(ctx context.Context, p *diary.UserSignupPayload) (err error) {
	_, err = s.repos.User.Create(p.Email, p.Name)

	if err != nil {
		var validationError *model.ValidationError
		if errors.As(err, &validationError) {
			return diary.MakeUserValidationError(err)
		}

		var duplicationError *repository.DuplicationError
		if errors.As(err, &duplicationError) {
			return diary.MakeUserDuplicationError(err)
		}
	}

	return
}

func (s *diarysrvc) Signin(ctx context.Context, p *diary.SigninPayload) (res string, err error) {
	s.logger.Print("diary.Signin")
	return
}
