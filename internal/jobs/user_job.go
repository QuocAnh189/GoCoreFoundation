package jobs

import (
	"context"
	"log"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
)

type UserJob struct {
	name    string
	id      int
	userSvc *users.Service
}

func NewUserJob(userSvc *users.Service) *UserJob {
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

	data := &users.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Email:     "join@gmail.com",
		Role:      users.RoleUser,
	}

	_, _, err := j.userSvc.CreateUser(ctx, data)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}
