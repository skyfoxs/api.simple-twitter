package data

import "github.com/skyfoxs/api.simple-twitter/pkg/idata"

type FollowingRequest struct {
	ID string `json:"id"`
}

type FollowingResponse struct {
	Following []ProfileResponse `json:"following"`
}

func NewFollowingResponse(l []idata.Profile) FollowingResponse {
	f := []ProfileResponse{}
	for _, v := range l {
		f = append(f, NewProfileResponse(&v))
	}
	return FollowingResponse{
		Following: f,
	}
}
