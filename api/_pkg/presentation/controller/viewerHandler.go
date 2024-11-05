package controller

import (
	"context"

	"github.com/umenerineri/hai-viewer-backend/api/_pkg/application/usecase"
	impl_repository "github.com/umenerineri/hai-viewer-backend/api/_pkg/infrastructure/repository"

	ogen "github.com/umenerineri/hai-viewer-backend/api/_pkg/ogen"
)

func (h *HaiHandler) ViewGet(ctx context.Context) (ogen.ViewGetRes, error) {
	historyRepository, err := impl_repository.NewImplHistoryRepository(ctx)
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get history repository")}, err
	}

	currentHistory, err := historyRepository.FindLatest()
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get current history")}, err
	}

	userRepository, err := impl_repository.NewImplUserRepository(ctx, currentHistory.GetHistoryId())
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	viewerUsecase, err := usecase.NewGetViewDataUsecase(userRepository)
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get usecase")}, err
	}

	arr, err := viewerUsecase.GetViewData()
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get view data")}, err
	}

	var resArr []ogen.ViewGetOKResultItem
	for i := 0; i < len(arr); i++ {
		resItem := &ogen.ViewGetOKResultItem{
			Position: ogen.ViewGetOKResultItemPosition{
				X: arr[i].GetPosition().GetX(),
				Y: arr[i].GetPosition().GetY(),
			},
			URL: arr[i].GetUrl(),
		}
		resArr = append(resArr, *resItem)
	}

	return &ogen.ViewGetOK{Result: resArr}, nil
}
