package converter

import (
	"github.com/amzdll/backend_2_go/src/internal/model"
	"math"
	"net/http"
	"strconv"
)

func RequestToPaginationModel(r *http.Request) (model.Pagination, error) {
	limit, offset := math.MaxInt, 0
	var err error

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return model.Pagination{}, err
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return model.Pagination{}, err
		}
	}

	return model.Pagination{
		Limit:  limit,
		Offset: offset,
	}, nil
}
