package routes

import (
	"strconv"
	"net/http"
)

func queryParamToInt64(r *http.Request, param string) (int64, error) {
	strParam := r.URL.Query().Get(param)
	n, err := strconv.ParseInt(strParam, 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}