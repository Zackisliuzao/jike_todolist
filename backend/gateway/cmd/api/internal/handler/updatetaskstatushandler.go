// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"
	"strings"

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
		// 路径格式为 /api/v1/tasks/:id/status，我们需要提取 :id 部分
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		idStr := ""
		// 寻找 "tasks" 后面的数字ID
		for i, part := range pathParts {
			if part == "tasks" && i+1 < len(pathParts) {
				idStr = pathParts[i+1]
				break
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
