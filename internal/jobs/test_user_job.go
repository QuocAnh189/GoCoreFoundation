package jobs

import (
	"context"
	"log"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/user"
)

type UserJob struct {
	name    string
	id      int
	userSvc *user.Service
}

func NewUserJob(userSvc *user.Service) *UserJob {
	return &UserJob{
		name:    "user-job",
		id:      2,
		userSvc: userSvc,
	}
}

func (j *UserJob) Name() string {
	return j.name
}

func (j *UserJob) TickInterval() int {
	return 30
}

func (j *UserJob) Run(ctx context.Context) error {
	now := time.Now()
	log.Printf("[%s] [id=%d] %s", j.name, j.id, now)

	data := &user.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Email:     "join@gmail.com",
		Role:      enum.RoleUser,
	}

	_, _, err := j.userSvc.CreateUser(ctx, data)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}
