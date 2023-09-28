package todo

import (
	"TodoList/common/response"
	"net/http"

	"TodoList/app/internal/logic/todo"
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelTodoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := todo.NewDelTodoLogic(r.Context(), svcCtx)
		err := l.DelTodo(&req)
		response.Response(r, w, nil, err)
	}
}
