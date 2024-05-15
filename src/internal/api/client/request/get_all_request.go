package request

type GetAllRequest struct {
	Limit  int `queryparam:"limit" validate:"required"`
	Offset int `queryparam:"offset" validate:"required"`
}
