// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jike_todo/gateway/cmd/api/internal/logic"
	"jike_todo/gateway/cmd/api/internal/svc"
)

// Delete task
func DeleteTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从URL路径参数中获取任务ID
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		taskIdStr := ""
		if len(pathParts) > 0 {
			// 从路径末尾获取ID，路径格式为 /tasks/:id
			taskIdStr = pathParts[len(pathParts)-1]
		}

		l := logic.NewDeleteTaskLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTask(taskIdStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
