// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity client
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package activity

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Activity" service client.
type Client struct {
	ManualPostOfActivityEndpoint                      goa.Endpoint
	RefreshActivitiesOfAllCooperationServicesEndpoint goa.Endpoint
	RefreshQiitaActivitiesEndpoint                    goa.Endpoint
	PickOutAllPastActivitiesOfQiitaEndpoint           goa.Endpoint
}

// NewClient initializes a "Activity" service client given the endpoints.
func NewClient(manualPostOfActivity, refreshActivitiesOfAllCooperationServices, refreshQiitaActivities, pickOutAllPastActivitiesOfQiita goa.Endpoint) *Client {
	return &Client{
		ManualPostOfActivityEndpoint:                      manualPostOfActivity,
		RefreshActivitiesOfAllCooperationServicesEndpoint: refreshActivitiesOfAllCooperationServices,
		RefreshQiitaActivitiesEndpoint:                    refreshQiitaActivities,
		PickOutAllPastActivitiesOfQiitaEndpoint:           pickOutAllPastActivitiesOfQiita,
	}
}

// ManualPostOfActivity calls the "Manual post of activity" endpoint of the
// "Activity" service.
func (c *Client) ManualPostOfActivity(ctx context.Context, p *ActivityPostPayload) (err error) {
	_, err = c.ManualPostOfActivityEndpoint(ctx, p)
	return
}

// RefreshActivitiesOfAllCooperationServices calls the "Refresh activities of
// all cooperation services" endpoint of the "Activity" service.
func (c *Client) RefreshActivitiesOfAllCooperationServices(ctx context.Context, p *SessionTokenPayload) (err error) {
	_, err = c.RefreshActivitiesOfAllCooperationServicesEndpoint(ctx, p)
	return
}

// RefreshQiitaActivities calls the "Refresh qiita activities" endpoint of the
// "Activity" service.
func (c *Client) RefreshQiitaActivities(ctx context.Context, p *SessionTokenPayload) (err error) {
	_, err = c.RefreshQiitaActivitiesEndpoint(ctx, p)
	return
}

// PickOutAllPastActivitiesOfQiita calls the "Pick out all past activities of
// qiita" endpoint of the "Activity" service.
func (c *Client) PickOutAllPastActivitiesOfQiita(ctx context.Context, p *SessionTokenPayload) (err error) {
	_, err = c.PickOutAllPastActivitiesOfQiitaEndpoint(ctx, p)
	return
}
