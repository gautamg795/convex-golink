package golink

import (
	"context"
	"fmt"
	"io/fs"
	"time"

	convexclient "github.com/jordanhunt22/golink_convex_client"
)

type ConvexDB2 struct {
	token  string
	client *convexclient.APIClient
}

func NewConvexDB2(token string) *ConvexDB2 {
	configuration := convexclient.NewConfiguration()
	client := convexclient.NewAPIClient(configuration)
	return &ConvexDB2{token: token, client: client}
}

func (c *ConvexDB2) LoadAll() ([]*Link, error) {
	request := *convexclient.NewRequestLoadLoadAll(*convexclient.NewRequestClearDefaultArgs(c.token))
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

func (c *ConvexDB2) Load(short string) (*Link, error) {
	request := *convexclient.NewRequestLoadLoadOne(*convexclient.NewRequestLoadLoadOneArgs(short, c.token))
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

func (c *ConvexDB2) Save(link *Link) error {
	linkDoc := convexclient.RequestStoreDefaultArgsLink{
		Short:        link.Short,
		Long:         link.Long,
		Owner:        link.Owner,
		NormalizedId: linkID(link.Short),
		Created:      float32(link.Created.Unix()),
		LastEdit:     float32(link.LastEdit.Unix()),
	}
	request := *convexclient.NewRequestStoreDefault(*convexclient.NewRequestStoreDefaultArgs(linkDoc, c.token))
	resp, httpRes, err := c.client.MutationAPI.ApiRunStoreDefaultPost(context.Background()).RequestStoreDefault(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return validationErr
	}

	return nil
}

func (c *ConvexDB2) LoadStats() (ClickStats, error) {
	request := *convexclient.NewRequestStatsLoadStats(*convexclient.NewRequestClearDefaultArgs(c.token))
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

func (c *ConvexDB2) SaveStats(stats ClickStats) error {
	mungedStats := make(map[string]int)
	for id, clicks := range stats {
		mungedStats[linkID(id)] = clicks
	}

	request := *convexclient.NewRequestStatsSaveStats(*convexclient.NewRequestStatsSaveStatsArgs(map[string]interface{}{"stats": mungedStats}, c.token))
	resp, httpRes, err := c.client.MutationAPI.ApiRunStatsSaveStatsPost(context.Background()).RequestStatsSaveStats(request).Execute()
	validationErr := validateResponse(httpRes.StatusCode, err, resp.Status)
	if validationErr != nil {
		return validationErr
	}

	if err != nil {
		return err
	}
	if httpRes.StatusCode != 200 {
		return err
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
