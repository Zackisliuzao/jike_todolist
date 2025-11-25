// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jike_todo/gateway/cmd/api/internal/logic"
	"jike_todo/gateway/cmd/api/internal/svc"
)

// Get overview statistics
func GetStatsOverviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetStatsOverviewLogic(r.Context(), svcCtx)
		resp, err := l.GetStatsOverview()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
