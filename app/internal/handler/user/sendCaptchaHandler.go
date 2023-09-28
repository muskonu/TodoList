package user

import (
	"TodoList/common/response"
	"net/http"

	"TodoList/app/internal/logic/user"
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := svcCtx.Validate.StructCtx(r.Context(), req)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := user.NewSendCaptchaLogic(r.Context(), svcCtx)
		err = l.SendCaptcha(&req)
		response.Response(r, w, nil, err)
	}
}
