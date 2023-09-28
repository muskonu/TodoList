package todo

import (
	"TodoList/common/response"
	"net/http"

	"TodoList/app/internal/logic/todo"
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CompleteTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteTodoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := todo.NewCompleteTodoLogic(r.Context(), svcCtx)
		err := l.CompleteTodo(&req)
		response.Response(r, w, nil, err)
	}
}
