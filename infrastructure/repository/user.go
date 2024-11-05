package impl_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	config "github.com/umenerineri/hai-viewer-backend/config"
	domain "github.com/umenerineri/hai-viewer-backend/domain"
	"github.com/umenerineri/hai-viewer-backend/domain/entity/user"
	"github.com/umenerineri/hai-viewer-backend/infrastructure/impl/database"
)

type ImplUserRepository struct {
	HistoryId string
	Client    *firestore.Client
}

func NewImplUserRepository(ctx context.Context, historyId string) (*ImplUserRepository, error) {
	app, err := config.InitializeApp()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	return &ImplUserRepository{HistoryId: historyId, Client: client}, nil
}

func (ur *ImplUserRepository) Create(user user.User) error {
	ctx := context.Background()
	userData := ConvertUserToData(user)

	err := database.CreateUser(ur.Client, ctx, ur.HistoryId, *userData)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ImplUserRepository) FindById(userId user.UserId) (*user.User, error) {
	ctx := context.Background()

	userData, err := database.FindById(ur.Client, ctx, ur.HistoryId, userId.ToId())
	if err != nil {
		return nil, err
	}

	return ConvertDataToUser(*userData)
}

func (ur *ImplUserRepository) FindByPos(pos user.Position) (*user.User, error) {
	ctx := context.Background()

	userData, err := database.FindByPos(ur.Client, ctx, ur.HistoryId, pos.GetX(), pos.GetY())
	if err != nil {
		return nil, err
	}

	return ConvertDataToUser(*userData)
}

func (ur *ImplUserRepository) FindLatest() (*user.User, error) {
	ctx := context.Background()

	userData, err := database.FindLatest(ur.Client, ctx, ur.HistoryId)
	if err != nil {
		return nil, err
	}

	return ConvertDataToUser(*userData)
}

func (ur *ImplUserRepository) GetLatestArray() ([]user.User, error) {
	ctx := context.Background()

	targetPosition := user.NewPosition(0, 0)
	var userArray []user.User

	for {
		targetX := targetPosition.GetX()
		targetY := targetPosition.GetY()
		latestData, err := database.FindByPos(ur.Client, ctx, ur.HistoryId, targetX, targetY)

		if errors.Is(err, domain.ErrNoLatestUser) || !latestData.IsDrawn {
			break
		} else if err != nil {
			return nil, err
		}

		latestUser, err := ConvertDataToUser(*latestData)
		if err != nil {
			return nil, fmt.Errorf("failed to convert latest data by pos (%d, %d) : %w", targetX, targetY, err)
		}
		userArray = append(userArray, *latestUser)

		targetPosition, err = targetPosition.GetNext()
		if err != nil {
			return nil, fmt.Errorf("failed to get next of pos (%d, %d) : %w", targetX, targetY, err)
		}
	}

	return userArray, nil
}

func (ur *ImplUserRepository) Update(user user.User) error {
	ctx := context.Background()
	userData := ConvertUserToData(user)

	err := database.Update(ur.Client, ctx, ur.HistoryId, *userData)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ImplUserRepository) Delete(userId user.UserId) error {
	ctx := context.Background()

	err := database.Delete(ur.Client, ctx, ur.HistoryId, userId.ToId())
	if err != nil {
		return err
	}

	return nil
}

func ConvertDataToUser(data database.UserData) (*user.User, error) {
	id, err := uuid.Parse(data.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to uuid: %w", err)
	}

	userId, err := user.NewUserId(id)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}

	position := user.NewPosition(data.PosX, data.PosY)
	user := user.NewUser(*userId, *position, data.Url, data.IsDrawn, data.CreatedAt, data.UpdatedAt)
	return user, nil
}

func ConvertUserToData(user user.User) *database.UserData {
	now := time.Now()
	return &database.UserData{
		UserId:    user.GetId().ToId(),
		PosX:      user.GetPosition().GetX(),
		PosY:      user.GetPosition().GetY(),
		Url:       user.GetUrl(),
		IsDrawn:   user.IsDrawn(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: now,
	}
}
