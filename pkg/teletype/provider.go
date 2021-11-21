package teletype

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jaytaylor/html2text"
)

type Provider struct {
	client  Client
	cleaner Cleaner
}

func NewProvider(client Client, cleaner Cleaner) *Provider {
	return &Provider{
		client:  client,
		cleaner: cleaner,
	}
}

func (p *Provider) Each(ctx context.Context, blogID int64, fn func(a Article) error) error {
	var lastArticleID int64

	for {
		u, err := p.articlesUrl(blogID, batchSize, lastArticleID)
		if err != nil {
			return err
		}

		var a articles
		if err = p.client.JSON(ctx, u, &a); err != nil {
			return err
		}

		for _, article := range a.Articles {
			if err = p.getArticleContent(ctx, &article); err != nil {
				return err
			}

			if err = fn(article); err != nil {
				return fmt.Errorf(`something went wrong: %w`, err)
			}

			lastArticleID = article.ID
		}

		if len(a.Articles) < batchSize {
			break
		}
	}

	return nil
}

func (p *Provider) getArticleContent(ctx context.Context, article *Article) (err error) {
	u := p.articleUrl(article.ID)
	if err = p.client.JSON(ctx, u, article); err != nil {
		return fmt.Errorf(`error getting article text: %w`, err)
	}

	article.Text, err = p.cleaner.Clean(article.Text, "tags")
	if err != nil {
		return fmt.Errorf(`error cleaning article html: %w`, err)
	}

	article.Text, err = html2text.FromString(article.Text)
	if err != nil {
		return fmt.Errorf(`error extracting article text from html: %w`, err)
	}

	return
}

func (p *Provider) articlesUrl(blogID, limit, lastArticleID int64) (string, error) {
	u, err := url.Parse(fmt.Sprintf(urlArticlesPattern, blogID))
	if err != nil {
		return "", fmt.Errorf(`error parsing url: %w`, err)
	}

	v := url.Values{}
	v.Set("limit", fmt.Sprintf(`%d`, limit))

	if lastArticleID != 0 {
		v.Set("last_article", fmt.Sprintf(`%d`, lastArticleID))
	}

	u.RawQuery = v.Encode()

	return u.String(), nil
}

func (p *Provider) articleUrl(articleID int64) string {
	return fmt.Sprintf(urlArticlePattern, articleID)
}
