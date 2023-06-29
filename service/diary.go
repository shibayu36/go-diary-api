package service

import (
	"context"
	"errors"
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
	repo   *repository.Repository
}

// NewDiary returns the diary service implementation.
func NewDiary(logger *log.Logger, repo *repository.Repository) diary.Service {
	return &diarysrvc{logger, repo}
}

func (s *diarysrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	unauthorizedError := diary.MakeUnauthorized(errors.New("invalid token"))

	apiKey, err := s.repo.FindApiKeyByApiKey(key)
	if err != nil {
		return ctx, unauthorizedError
	}

	user, err := s.repo.FindUserByID(apiKey.UserID)
	if err != nil {
		return ctx, unauthorizedError
	}

	ctx = context.WithValue(ctx, "user", user)

	return ctx, nil
}

// UserSignup implements UserSignup.
func (s *diarysrvc) UserSignup(ctx context.Context, p *diary.UserSignupPayload) (err error) {
	_, err = s.repo.CreateUser(p.Email, p.Name)

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
	u, err := s.repo.FindUserByEmail(p.Email)
	if err != nil {
		if repository.IsNotFound(err) {
			return "", diary.MakeBadRequest(errors.New("email is invalid"))
		}
		return "", err
	}

	key, err := s.repo.CreateApiKeyByUser(u)
	if err != nil {
		return "", err
	}

	return key.ApiKey, nil
}

func (s *diarysrvc) CreateDiary(ctx context.Context, p *diary.CreateDiaryPayload) (err error) {
	user := ctx.Value("user").(*model.User)
	s.logger.Print(user)
	s.logger.Print("diary.CreateDiary")
	return
}
