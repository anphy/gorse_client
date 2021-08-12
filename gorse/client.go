package gorse

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type GorseClient interface {
	InsertItem(item Item) error
	InsertItems(items []Item) error
	InsertFeedback(fb Feedback) error
	GetRecommendItems(userID string) ([]Item, error)
	GetPopularItems(offset, limit int) ([]Item, error)
	GetLatestItems(offset, limit int) ([]Item, error)
}

func NewGorseClient(entry string, xapiKey string) GorseClient {
	return &gorseclient{entry: entry, apiKey: xapiKey}
}

type gorseclient struct {
	entry  string
	apiKey string
}

func (gc *gorseclient) InsertItem(item Item) error {
	body, err := json.Marshal(item)
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return err
	}
	r := Request{
		URL:    fmt.Sprintf("%s/api/item", gc.entry),
		Body:   body,
		Method: "POST",
	}
	if _, err := r.Do(); err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return err
	}
	return nil
}

func (gc *gorseclient) InsertItems(items []Item) error {
	body, err := json.Marshal(items)
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return err
	}
	r := Request{
		URL:    fmt.Sprintf("%s/api/items", gc.entry),
		Body:   body,
		Method: "POST",
	}
	if _, err := r.Do(); err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return err
	}
	return nil
}

func (gc *gorseclient) InsertFeedback(fb Feedback) error {
	body, err := json.Marshal(fb)
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return err
	}
	r := Request{
		URL:    fmt.Sprintf("%s/api/feedback", gc.entry),
		Body:   body,
		Method: "POST",
	}
	if _, err := r.Do(); err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return err
	}
	return nil
}

func (gc *gorseclient) GetRecommendItems(userID string) ([]Item, error) {
	r := Request{
		URL:    fmt.Sprintf("%s/api/recommend/%s", gc.entry, userID),
		Method: "GET",
	}
	result, err := r.Do()
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	var items []Item
	if err = json.Unmarshal(result, &items); err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	return items, nil
}

func (gc *gorseclient) GetPopularItems(offset, limit int) ([]Item, error) {
	r := Request{
		URL:    fmt.Sprintf("%s/api/popular", gc.entry),
		Method: "GET",
		Params: map[string]interface{}{},
	}
	r.Params["offset"] = offset
	r.Params["n"] = limit
	result, err := r.Do()
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	var items []Item
	if err = json.Unmarshal(result, &items); err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	return items, nil
}

func (gc *gorseclient) GetLatestItems(offset, limit int) ([]Item, error) {
	r := Request{
		URL:    fmt.Sprintf("%s/api/latest", gc.entry),
		Method: "GET",
		Params: map[string]interface{}{},
	}
	r.Params["offset"] = offset
	r.Params["n"] = limit
	result, err := r.Do()
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	var items []Item
	if err = json.Unmarshal(result, &items); err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	return items, nil
}
