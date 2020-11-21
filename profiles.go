package goinsta

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Profiles allows user function interactions
type Profiles struct {
	inst *Instagram
}

func newProfiles(inst *Instagram) *Profiles {
	profiles := &Profiles{
		inst: inst,
	}
	return profiles
}

// ByName return a *User structure parsed by username
func (prof *Profiles) ByName(name string) (*User, error) {
	if prof.inst.user == "" {
		body, err := prof.inst.sendSimpleRequest(urlUnlogedInfo, name)
		if err == nil {
			resp := UnlogedData{}
			err = json.Unmarshal(body, &resp)
			if err == nil {
				user := setUnlogedUserData(resp)
				user.feedMedia = setUnlogedFeedData(resp)
				user.inst = prof.inst
				return user, err
			}
		}
		return nil, err
	}
	body, err := prof.inst.sendSimpleRequest(urlUserByName, name)
	if err == nil {
		resp := userResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			user := &resp.User
			user.inst = prof.inst
			return user, err
		}
	}
	return nil, err
}

// ByID returns a *User structure parsed by user id
func (prof *Profiles) ByID(id int64) (*User, error) {
	data, err := prof.inst.prepareData()
	if err != nil {
		return nil, err
	}

	body, err := prof.inst.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlUserByID, id),
			Query:    generateSignature(data),
		},
	)
	if err == nil {
		resp := userResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			user := &resp.User
			user.inst = prof.inst

			if prof.inst.user == "" {
				newUser, newErr := prof.ByName(user.Username)
				return newUser, newErr
			}
			return user, err
		}
	}
	return nil, err
}

// Blocked returns a list of blocked profiles.
func (prof *Profiles) Blocked() ([]BlockedUser, error) {
	body, err := prof.inst.sendSimpleRequest(urlBlockedList)
	if err == nil {
		resp := blockedListResp{}
		err = json.Unmarshal(body, &resp)
		return resp.BlockedList, err
	}
	return nil, err
}

func setUnlogedUserData(data UnlogedData) *User {
	user := new(User)

	user.ID, _ = strconv.ParseInt(data.Graphql.User.ID, 10, 64)
	user.Username = data.Graphql.User.Username
	user.FullName = data.Graphql.User.FullName
	user.Biography = data.Graphql.User.Biography
	user.IsPrivate = data.Graphql.User.IsPrivate
	user.FollowingCount = data.Graphql.User.Follow.Count
	user.FollowerCount = data.Graphql.User.FollowedBy.Count
	user.ProfilePicURL = data.Graphql.User.ProfilePicURLHd

	return user
}

func setUnlogedFeedData(data UnlogedData) *FeedMedia {
	feed := new(FeedMedia)

	for _, item := range data.Graphql.User.Publications.Edges {
		newItem := new(Item)
		newItem.ID = item.Node.ID
		newItem.TakenAt = item.Node.TakenAtTimestamp
		newItem.Pk, _ = strconv.ParseInt(item.Node.ID, 10, 64)
		newItem.CommentsDisabled = item.Node.CommentsDisabled
		newItem.Code = item.Node.Shortcode
		newItem.Caption.Text = item.Node.Caption.Edges[0].Node.Text
		newItem.Likes = item.Node.Likes.Count
		newItem.CommentCount = item.Node.Comments.Count
		newItem.MediaType = setMediaType(item.Node.Typename)

		// set the info for video
		if item.Node.IsVideo {
			newVideo := new(Video)

			newVideo.Width = item.Node.Dimensions.Width
			newVideo.Height = item.Node.Dimensions.Height
			newVideo.URL = item.Node.VideoURL

			newItem.Videos = append(newItem.Videos, *newVideo)
			//item params related to videos
			newItem.HasAudio = item.Node.HasAudio
			newItem.ViewCount = item.Node.VideoViewCount
			// newItem.IsDashEligible = item.Node.DashInfo.IsDashEligible
			newItem.VideoDashManifest = item.Node.DashInfo.VideoDashManifest
			newItem.NumberOfQualities = item.Node.DashInfo.NumberOfQualities
		}
		// Set the best quality image
		candidate := new(Candidate)
		candidate.Width = item.Node.Dimensions.Width
		candidate.Height = item.Node.Dimensions.Height
		candidate.URL = item.Node.DisplayURL

		newItem.Images.Versions = append(newItem.Images.Versions, *candidate)
		// Set the version
		for _, image := range item.Node.ThumbnailResources {

			candidate.Width = image.ConfigWidth
			candidate.Height = image.ConfigHeight
			candidate.URL = image.Src

			newItem.Images.Versions = append(newItem.Images.Versions, *candidate)

		}

		for _, car := range item.Node.Carrousel.Edges {
			newCar := new(Item)
			newCar.Pk, _ = strconv.ParseInt(car.Node.ID, 10, 64)
			newCar.ID = car.Node.ID
			newCar.TakenAt = item.Node.TakenAtTimestamp
			newCar.MediaType = setMediaType(car.Node.Typename)

			if car.Node.IsVideo {
				newVideo := new(Video)

				newVideo.Width = car.Node.Dimensions.Width
				newVideo.Height = car.Node.Dimensions.Height
				newVideo.URL = car.Node.VideoURL

				newCar.Videos = append(newCar.Videos, *newVideo)
				//item params related to videos
				newCar.HasAudio = car.Node.HasAudio
				newCar.ViewCount = car.Node.VideoViewCount
				// newItem.IsDashEligible = item.Node.DashInfo.IsDashEligible
				newCar.VideoDashManifest = car.Node.DashInfo.VideoDashManifest
				newCar.NumberOfQualities = car.Node.DashInfo.NumberOfQualities
			}

			// Set the best quality image
			candidate := new(Candidate)
			candidate.Width = car.Node.Dimensions.Width
			candidate.Height = car.Node.Dimensions.Height
			candidate.URL = car.Node.DisplayURL

			newCar.Images.Versions = append(newCar.Images.Versions, *candidate)

			newItem.CarouselMedia = append(newItem.CarouselMedia, *newCar)
		}

		feed.Items = append(feed.Items, *newItem)
	}

	return feed
}

//funciones nuevas para manejo de la data
func setMediaType(mediaType string) int {
	switch mediaType {
	case "GraphImage":
		return 1

	case "GraphVideo":
		return 2

	case "GraphSidecar":
		return 8
	}

	return 0
}
