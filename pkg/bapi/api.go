package bapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{url}
}

func (a *API) GetConsumerList(ctx context.Context, source string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s?source=%s", a.URL, "api/consumer/list", source)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//首先创建并设置当前跨度的信息和标签内容
	span, newCtx := opentracing.StartSpanFromContext(ctx, "HTTP GET:"+url, opentracing.Tag{
		Key:   string(ext.Component),
		Value: "HTTP",
	})
	span.SetTag("url", url)
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	//传递http的ctx
	req = req.WithContext(newCtx)
	client := http.Client{
		Timeout: time.Second * 60,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer span.Finish()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
