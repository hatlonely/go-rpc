package service

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-article/api/gen/go/api"
)

func split(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		for _, s := range []rune("，。、？！； ,.") {
			if r == s {
				return true
			}
		}
		return false
	})
}

func (s *ArticleService) SearchArticle(ctx context.Context, req *api.SearchArticleReq) (*api.SearchArticleRes, error) {
	query := elastic.NewBoolQuery()
	for _, val := range split(req.Keyword) {
		if len(val) == 0 {
			continue
		}
		q := elastic.NewBoolQuery()
		q.Should(elastic.NewMatchPhrasePrefixQuery("title", val))
		q.Should(elastic.NewMatchPhrasePrefixQuery("author", val))
		q.Should(elastic.NewMatchPhrasePrefixQuery("tags", val))
		q.Should(elastic.NewMatchPhrasePrefixQuery("content", val))
		query.Must(q)
	}

	esCtx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()
	res, err := s.esCli.Search().Index("article").Query(query).From(int(req.Offset)).Size(int(req.Limit)).Do(esCtx)
	if err != nil {
		return nil, errors.Wrap(err, "EsCli.Search failed")
	}

	var ancients []*api.Article
	for _, item := range res.Each(reflect.TypeOf(api.Article{})) {
		if t, ok := item.(api.Article); ok {
			ancients = append(ancients, &t)
		}
	}

	return &api.SearchArticleRes{Articles: ancients}, nil
}
