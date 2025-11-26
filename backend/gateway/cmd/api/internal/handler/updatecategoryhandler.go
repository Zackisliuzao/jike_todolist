// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jike_todo/gateway/cmd/api/internal/logic"
	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"
)

// Update category
func UpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从URL路径参数中获取分类ID
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		categoryIdStr := ""
		if len(pathParts) > 0 {
			// 从路径末尾获取ID，路径格式为 /categories/:id
			categoryIdStr = pathParts[len(pathParts)-1]
		}

		l := logic.NewUpdateCategoryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCategory(categoryIdStr, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
