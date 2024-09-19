package acc_tasks

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/url"
	"server/api/tencent"
	"server/common"
	"strconv"
	"time"
)

func (x *Controller) Invoke(ctx context.Context, c *app.RequestContext) {
	r, err := x.AccTasksX.Invoke(ctx)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, r)
}

func (x *Service) Invoke(ctx context.Context) (r M, err error) {
	u, _ := url.Parse(x.V.AccelerateAddress)
	timestamp := time.Now().Unix()
	headerx := map[string]string{
		"Content-Type":  "application/json",
		"Host":          u.Host,
		"X-Scf-Cam-Uin": x.V.CamUin,
	}
	body := M{}
	if _, err = common.HttpClient(u.String()).R().SetContext(ctx).
		SetHeaders(headerx).
		SetHeader("X-Scf-Cam-Timestamp", strconv.FormatInt(timestamp, 10)).
		SetHeader("Authorization", x.TencentX.TC3Authorization(tencent.TC3Option{
			Service:   "scf",
			Headers:   headerx,
			Timestamp: timestamp,
			Body:      body,
		})).
		SetBody(body).
		SetSuccessResult(&r).
		Post(""); err != nil {
		return
	}

	return
}
