package controller

import (
	"context"

	"github.com/umenerineri/hai-viewer-backend/application/usecase"
	impl_repository "github.com/umenerineri/hai-viewer-backend/infrastructure/repository"

	ogen "github.com/umenerineri/hai-viewer-backend/ogen"
)

func (h *HaiHandler) APIHandlerViewGet(ctx context.Context) (ogen.APIHandlerViewGetRes, error) {
	historyRepository, err := impl_repository.NewImplHistoryRepository(ctx)
	if err != nil {
		return &ogen.APIHandlerViewGetBadRequest{Error: ogen.NewOptString("failed to get history repository")}, err
	}

	currentHistory, err := historyRepository.FindLatest()
	if err != nil {
		return &ogen.APIHandlerViewGetBadRequest{Error: ogen.NewOptString("failed to get current history")}, err
	}

	userRepository, err := impl_repository.NewImplUserRepository(ctx, currentHistory.GetHistoryId())
	if err != nil {
		return &ogen.APIHandlerViewGetBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	viewerUsecase, err := usecase.NewGetViewDataUsecase(userRepository)
	if err != nil {
		return &ogen.APIHandlerViewGetBadRequest{Error: ogen.NewOptString("failed to get usecase")}, err
	}

	arr, err := viewerUsecase.GetViewData()
	if err != nil {
		return &ogen.APIHandlerViewGetBadRequest{Error: ogen.NewOptString("failed to get view data")}, err
	}

	var resArr []ogen.APIHandlerViewGetOKResultItem
	for i := 0; i < len(arr); i++ {
		resItem := &ogen.APIHandlerViewGetOKResultItem{
			Position: ogen.APIHandlerViewGetOKResultItemPosition{
				X: arr[i].GetPosition().GetX(),
				Y: arr[i].GetPosition().GetY(),
			},
			URL: arr[i].GetUrl(),
		}
		resArr = append(resArr, *resItem)
	}

	return &ogen.APIHandlerViewGetOK{Result: resArr}, nil
}
