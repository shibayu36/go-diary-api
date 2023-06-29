package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	diary "github.com/shibayu36/go-diary-api/gen/diary"
	"github.com/shibayu36/go-diary-api/model"
	"github.com/shibayu36/go-diary-api/repository"
	"goa.design/goa/v3/security"
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

func (s *diarysrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
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
	u, err := s.repos.User.FindByEmail(p.Email)
	if err != nil {
		if repository.IsNotFound(err) {
			return "", diary.MakeBadRequest(errors.New("email is invalid"))
		}
		return "", err
	}

	key, err := s.repos.ApiKey.CreateByUser(u)
	if err != nil {
		return "", err
	}

	return key.ApiKey, nil
}

func (s *diarysrvc) CreateDiary(ctx context.Context, p *diary.CreateDiaryPayload) (err error) {
	s.logger.Print("diary.CreateDiary")
	return
}
