// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"jike_todo/gateway/cmd/api/internal/logic"
	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Update task status
func UpdateTaskStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTaskStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从URL路径参数中获取任务ID
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			// 尝试从路径中解析
			pathVars := r.Context().Value("path_params")
			if pathVars != nil {
				if params, ok := pathVars.(map[string]string); ok {
					idStr = params["id"]
				}
			}
		}

		l := logic.NewUpdateTaskStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTaskStatus(&req, idStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
