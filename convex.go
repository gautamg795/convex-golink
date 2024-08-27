package golink

import (
	"context"
	"fmt"
	"io/fs"
	"time"

	convex "github.com/jordanhunt22/golink_convex_client"
)

type ConvexDB struct {
	token  string
	client *convex.APIClient
}

func NewConvexDB(host *string, token string) *ConvexDB {
	config := convex.NewConfiguration()
	if host != nil {
		config.Servers = convex.ServerConfigurations{
			{
				URL:         *host,
				Description: "Convex URL that serves requests",
			},
		}
	}
	client := convex.NewAPIClient(config)

	return &ConvexDB{token: token, client: client}
}

func (c *ConvexDB) LoadAll() ([]*Link, error) {
	request := *convex.NewRequestLoadLoadAll(*convex.NewRequestClearDefaultArgs(c.token))
	resp, httpRes, err := c.client.QueryAPI.ApiRunLoadLoadAllPost(context.Background()).RequestLoadLoadAll(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return nil, validationErr
	}

	var links []*Link
	for _, doc := range resp.Value {
		link := Link{
			Short:    doc.Short,
			Long:     doc.Long,
			Created:  time.Unix(int64(doc.Created), 0),
			LastEdit: time.Unix(int64(doc.LastEdit), 0),
			Owner:    doc.Owner,
		}
		links = append(links, &link)
	}

	return links, nil
}

func (c *ConvexDB) Load(short string) (*Link, error) {
	request := *convex.NewRequestLoadLoadOne(*convex.NewRequestLoadLoadOneArgs(linkID(short), c.token))
	resp, httpRes, err := c.client.QueryAPI.ApiRunLoadLoadOnePost(context.Background()).RequestLoadLoadOne(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return nil, validationErr
	}

	linkDoc := resp.Value.Get()
	if linkDoc == nil {
		err := fs.ErrNotExist
		return nil, err
	}
	link := Link{
		Short:    linkDoc.Short,
		Long:     linkDoc.Long,
		Created:  time.Unix(int64(linkDoc.Created), 0),
		LastEdit: time.Unix(int64(linkDoc.LastEdit), 0),
		Owner:    linkDoc.Owner,
	}

	return &link, nil
}

func (c *ConvexDB) Save(link *Link) error {
	linkDoc := convex.RequestStoreDefaultArgsLink{
		Short:        link.Short,
		Long:         link.Long,
		Owner:        link.Owner,
		NormalizedId: linkID(link.Short),
		Created:      float32(link.Created.Unix()),
		LastEdit:     float32(link.LastEdit.Unix()),
	}
	request := *convex.NewRequestStoreDefault(*convex.NewRequestStoreDefaultArgs(linkDoc, c.token))
	resp, httpRes, err := c.client.MutationAPI.ApiRunStoreDefaultPost(context.Background()).RequestStoreDefault(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return validationErr
	}

	return nil
}

func (c *ConvexDB) LoadStats() (ClickStats, error) {
	request := *convex.NewRequestStatsLoadStats(*convex.NewRequestClearDefaultArgs(c.token))
	resp, httpRes, err := c.client.QueryAPI.ApiRunStatsLoadStatsPost(context.Background()).RequestStatsLoadStats(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return nil, validationErr
	}

	value, ok := resp.Value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response from convex: %v", resp)
	}

	clicks := make(ClickStats)
	for k, v := range value {
		num, ok := v.(float64)
		if !ok {
			return nil, err
		}
		clicks[k] = int(num)
	}

	return clicks, nil
}

func (c *ConvexDB) SaveStats(stats ClickStats) error {
	mungedStats := make(map[string]interface{})
	for id, clicks := range stats {
		mungedStats[linkID(id)] = clicks
	}

	request := *convex.NewRequestStatsSaveStats(*convex.NewRequestStatsSaveStatsArgs(mungedStats, c.token))
	resp, httpRes, err := c.client.MutationAPI.ApiRunStatsSaveStatsPost(context.Background()).RequestStatsSaveStats(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return validationErr
	}

	return nil
}

func clear(c *ConvexDB) error {
	request := *convex.NewRequestClearDefault(*convex.NewRequestClearDefaultArgs(c.token))
	resp, httpRes, err := c.client.MutationAPI.ApiRunClearDefaultPost(context.Background()).RequestClearDefault(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return validationErr
	}

	return nil
}

func validateResponse(statusCode int, err error, convexStatus string) error {
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return err
	}
	if convexStatus == "error" {
		return fmt.Errorf("error from convex")
	}

	return nil
}
