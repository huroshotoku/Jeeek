// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package client

import (
	"encoding/json"
	"fmt"

	activity "github.com/tonouchi510/Jeeek/gen/activity"
)

// BuildReflectionActivityPayload builds the payload for the Activity
// Reflection activity endpoint from CLI flags.
func BuildReflectionActivityPayload(activityReflectionActivityBody string, activityReflectionActivityToken string) (*activity.ActivityPostPayload, error) {
	var err error
	var body ReflectionActivityRequestBody
	{
		err = json.Unmarshal([]byte(activityReflectionActivityBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"Attributes\": [\n         {\n            \"uid\": \"Et et beatae sunt quia nihil.\"\n         },\n         {\n            \"uid\": \"Et et beatae sunt quia nihil.\"\n         },\n         {\n            \"uid\": \"Et et beatae sunt quia nihil.\"\n         }\n      ],\n      \"Data\": \"QXBlcmlhbSB0ZW5ldHVyIGV4ZXJjaXRhdGlvbmVtLg==\"\n   }'")
		}
	}
	var token *string
	{
		if activityReflectionActivityToken != "" {
			token = &activityReflectionActivityToken
		}
	}
	v := &activity.ActivityPostPayload{
		Data: body.Data,
	}
	if body.Attributes != nil {
		v.Attributes = make([]*activity.Attributes, len(body.Attributes))
		for i, val := range body.Attributes {
			v.Attributes[i] = marshalAttributesRequestBodyToActivityAttributes(val)
		}
	}
	v.Token = token
	return v, nil
}