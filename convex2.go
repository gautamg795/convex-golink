package golink

import (
	"context"
	"encoding/json"
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
	if err != nil {
		return nil, err
	}
	if httpRes.StatusCode != 200 {
		return nil, err
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
	if err != nil {
		return nil, err
	}
	if httpRes.StatusCode != 200 {
		return nil, err
	}
	if !resp.HasValue() {
		err := fs.ErrNotExist
		return nil, err
	}

	linkDoc := resp.Value.Get()
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
	linkDoc := convexclient.ResponseLoadLoadAllValueInner{
		Short:        link.Short,
		Long:         link.Long,
		Owner:        link.Owner,
		NormalizedId: linkID(link.Short),
		Created:      float32(link.Created.Unix()),
		LastEdit:     float32(link.LastEdit.Unix()),
	}
	request := *convexclient.NewRequestStoreDefault(*convexclient.NewRequestStoreDefaultArgs(linkDoc, c.token))
	_, httpRes, err := c.client.MutationAPI.ApiRunStoreDefaultPost(context.Background()).RequestStoreDefault(request).Execute()

	if err != nil {
		return err
	}
	if httpRes.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *ConvexDB2) LoadStats() (ClickStats, error) {
	request := *convexclient.NewRequestStatsLoadStats(*convexclient.NewRequestClearDefaultArgs(c.token))
	resp, httpRes, err := c.client.QueryAPI.ApiRunStatsLoadStatsPost(context.Background()).RequestStatsLoadStats(request).Execute()
	if err != nil {
		return nil, err
	}
	if httpRes.StatusCode != 200 {
		return nil, err
	}
	value, ok := resp.Value.(map[string]interface{})
	if !ok {
		return nil, fs.ErrNotExist
	}
	clicks := make(ClickStats)
	for k, v := range value {
		num, err := v.(json.Number).Float64()
		if err != nil {
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
	_, httpRes, err := c.client.MutationAPI.ApiRunStatsSaveStatsPost(context.Background()).RequestStatsSaveStats(request).Execute()

	if err != nil {
		return err
	}
	if httpRes.StatusCode != 200 {
		return err
	}
	return nil
}
