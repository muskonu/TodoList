package todo

import (
	"TodoList/common/response"
	"net/http"

	"TodoList/app/internal/logic/todo"
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTodoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := todo.NewQueryTodoLogic(r.Context(), svcCtx)
		resp, err := l.QueryTodo(&req)
		response.Response(r, w, resp, err)
	}
}
