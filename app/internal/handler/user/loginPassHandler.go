package user

import (
	"TodoList/app/internal/logic/user"
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"TodoList/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginPassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailPwdLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := svcCtx.Validate.StructCtx(r.Context(), req)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := user.NewLoginPassLogic(r.Context(), svcCtx)
		resp, err := l.LoginPass(&req)
		response.Response(r, w, resp, err)
	}
}
