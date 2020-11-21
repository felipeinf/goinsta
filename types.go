package goinsta

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ConfigFile is a structure to store the session information so that can be exported or imported.
type ConfigFile struct {
	ID        int64          `json:"id"`
	User      string         `json:"username"`
	DeviceID  string         `json:"device_id"`
	UUID      string         `json:"uuid"`
	RankToken string         `json:"rank_token"`
	Token     string         `json:"token"`
	PhoneID   string         `json:"phone_id"`
	Cookies   []*http.Cookie `json:"cookies"`
}

// School is void structure (yet).
type School struct {
}

// PicURLInfo repre
type PicURLInfo struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

// ErrorN is general instagram error
type ErrorN struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

// Error503 is instagram API error
type Error503 struct {
	Message string
}

func (e Error503) Error() string {
	return e.Message
}

func (e ErrorN) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Status, e.Message, e.ErrorType)
}

// Error400 is error returned by HTTP 400 status code.
type Error400 struct {
	ChallengeError
	Action     string `json:"action"`
	StatusCode string `json:"status_code"`
	Payload    struct {
		ClientContext string `json:"client_context"`
		Message       string `json:"message"`
	} `json:"payload"`
	Status string `json:"status"`
}

func (e Error400) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Payload.Message)
}

// ChallengeError is error returned by HTTP 400 status code.
type ChallengeError struct {
	Message   string `json:"message"`
	Challenge struct {
		URL               string `json:"url"`
		APIPath           string `json:"api_path"`
		HideWebviewHeader bool   `json:"hide_webview_header"`
		Lock              bool   `json:"lock"`
		Logout            bool   `json:"logout"`
		NativeFlow        bool   `json:"native_flow"`
	} `json:"challenge"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

func (e ChallengeError) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Message)
}

// Nametag is part of the account information.
type Nametag struct {
	Mode          int64       `json:"mode"`
	Gradient      json.Number `json:"gradient,Number"`
	Emoji         string      `json:"emoji"`
	SelfieSticker json.Number `json:"selfie_sticker,Number"`
}

type friendResp struct {
	Status     string     `json:"status"`
	Friendship Friendship `json:"friendship_status"`
}

// Location stores media location information.
type Location struct {
	Pk               int64   `json:"pk"`
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	City             string  `json:"city"`
	ShortName        string  `json:"short_name"`
	Lng              float64 `json:"lng"`
	Lat              float64 `json:"lat"`
	ExternalSource   string  `json:"external_source"`
	FacebookPlacesID int64   `json:"facebook_places_id"`
}

// SuggestedUsers stores the information about user suggestions.
type SuggestedUsers struct {
	Type        int `json:"type"`
	Suggestions []struct {
		User            User          `json:"user"`
		Algorithm       string        `json:"algorithm"`
		SocialContext   string        `json:"social_context"`
		Icon            string        `json:"icon"`
		Caption         string        `json:"caption"`
		MediaIds        []interface{} `json:"media_ids"`
		ThumbnailUrls   []interface{} `json:"thumbnail_urls"`
		LargeUrls       []interface{} `json:"large_urls"`
		MediaInfos      []interface{} `json:"media_infos"`
		Value           float64       `json:"value"`
		IsNewSuggestion bool          `json:"is_new_suggestion"`
	} `json:"suggestions"`
	LandingSiteType  string `json:"landing_site_type"`
	Title            string `json:"title"`
	ViewAllText      string `json:"view_all_text"`
	LandingSiteTitle string `json:"landing_site_title"`
	NetegoType       string `json:"netego_type"`
	UpsellFbPos      string `json:"upsell_fb_pos"`
	AutoDvance       string `json:"auto_dvance"`
	ID               string `json:"id"`
	TrackingToken    string `json:"tracking_token"`
}

// Friendship stores the details of the relationship between two users.
type Friendship struct {
	IncomingRequest bool `json:"incoming_request"`
	FollowedBy      bool `json:"followed_by"`
	OutgoingRequest bool `json:"outgoing_request"`
	Following       bool `json:"following"`
	Blocking        bool `json:"blocking"`
	IsPrivate       bool `json:"is_private"`
	Muting          bool `json:"muting"`
	IsMutingReel    bool `json:"is_muting_reel"`
}

// SavedMedia stores the information about media being saved before in my account.
type SavedMedia struct {
	Items []struct {
		Media Item `json:"media"`
	} `json:"items"`
	NumResults          int    `json:"num_results"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Status              string `json:"status"`
}

// Images are different quality images
type Images struct {
	Versions []Candidate `json:"candidates"`
}

// GetBest returns the URL of the image with the best quality.
func (img Images) GetBest() string {
	best := ""
	var mh, mw int
	for _, v := range img.Versions {
		if v.Width > mw || v.Height > mh {
			best = v.URL
			mh, mw = v.Height, v.Width
		}
	}
	return best
}

// Candidate is something that I really have no idea what it is.
type Candidate struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// Tag is the information of an user being tagged on any media.
type Tag struct {
	In []struct {
		User                  User        `json:"user"`
		Position              []float64   `json:"position"`
		StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
		DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
	} `json:"in"`
}

// Caption is media caption
type Caption struct {
	ID              int64  `json:"pk"`
	UserID          int64  `json:"user_id"`
	Text            string `json:"text"`
	Type            int    `json:"type"`
	CreatedAt       int64  `json:"created_at"`
	CreatedAtUtc    int64  `json:"created_at_utc"`
	ContentType     string `json:"content_type"`
	Status          string `json:"status"`
	BitFlags        int    `json:"bit_flags"`
	User            User   `json:"user"`
	DidReportAsSpam bool   `json:"did_report_as_spam"`
	MediaID         int64  `json:"media_id"`
	HasTranslation  bool   `json:"has_translation"`
}

// Mentions is a user being mentioned on media.
type Mentions struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int64   `json:"z"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	IsPinned int     `json:"is_pinned"`
	User     User    `json:"user"`
}

// Video are different quality videos
type Video struct {
	Type   int    `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
	ID     string `json:"id"`
}

type timeStoryResp struct {
	Status string       `json:"status"`
	Media  []StoryMedia `json:"tray"`
}

type trayResp struct {
	Reels  map[string]StoryMedia `json:"reels"`
	Status string                `json:"status"`
}

// Tray is a set of story media received from timeline calls.
type Tray struct {
	Stories []StoryMedia `json:"tray"`
	Lives   struct {
		LiveItems []LiveItems `json:"post_live_items"`
	} `json:"post_live"`
	StoryRankingToken    string      `json:"story_ranking_token"`
	Broadcasts           []Broadcast `json:"broadcasts"`
	FaceFilterNuxVersion int         `json:"face_filter_nux_version"`
	HasNewNuxStory       bool        `json:"has_new_nux_story"`
	Status               string      `json:"status"`
}

func (tray *Tray) set(inst *Instagram, url string) {
	for i := range tray.Stories {
		tray.Stories[i].inst = inst
		tray.Stories[i].endpoint = url
		tray.Stories[i].setValues()
	}
	for i := range tray.Lives.LiveItems {
		tray.Lives.LiveItems[i].User.inst = inst
		for j := range tray.Lives.LiveItems[i].Broadcasts {
			tray.Lives.LiveItems[i].Broadcasts[j].BroadcastOwner.inst = inst
		}
	}
	for i := range tray.Broadcasts {
		tray.Broadcasts[i].BroadcastOwner.inst = inst
	}
}

// LiveItems are Live media items
type LiveItems struct {
	ID                  string      `json:"pk"`
	User                User        `json:"user"`
	Broadcasts          []Broadcast `json:"broadcasts"`
	LastSeenBroadcastTs float64     `json:"last_seen_broadcast_ts"`
	RankedPosition      int64       `json:"ranked_position"`
	SeenRankedPosition  int64       `json:"seen_ranked_position"`
	Muted               bool        `json:"muted"`
	CanReply            bool        `json:"can_reply"`
	CanReshare          bool        `json:"can_reshare"`
}

// Broadcast is live videos.
type Broadcast struct {
	ID                   int64  `json:"id"`
	BroadcastStatus      string `json:"broadcast_status"`
	DashManifest         string `json:"dash_manifest"`
	ExpireAt             int64  `json:"expire_at"`
	EncodingTag          string `json:"encoding_tag"`
	InternalOnly         bool   `json:"internal_only"`
	NumberOfQualities    int    `json:"number_of_qualities"`
	CoverFrameURL        string `json:"cover_frame_url"`
	BroadcastOwner       User   `json:"broadcast_owner"`
	PublishedTime        int64  `json:"published_time"`
	MediaID              string `json:"media_id"`
	BroadcastMessage     string `json:"broadcast_message"`
	OrganicTrackingToken string `json:"organic_tracking_token"`
}

// BlockedUser stores information about a used that has been blocked before.
type BlockedUser struct {
	// TODO: Convert to user
	UserID        int64  `json:"user_id"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	ProfilePicURL string `json:"profile_pic_url"`
	BlockAt       int64  `json:"block_at"`
}

// Unblock unblocks blocked user.
func (b *BlockedUser) Unblock() error {
	u := User{ID: b.UserID}
	return u.Unblock()
}

type blockedListResp struct {
	BlockedList []BlockedUser `json:"blocked_list"`
	PageSize    int           `json:"page_size"`
	Status      string        `json:"status"`
}

// InboxItemMedia is inbox media item
type InboxItemMedia struct {
	ClientContext              string `json:"client_context"`
	ExpiringMediaActionSummary struct {
		Count     int    `json:"count"`
		Timestamp int64  `json:"timestamp"`
		Type      string `json:"type"`
	} `json:"expiring_media_action_summary"`
	ItemID     string `json:"item_id"`
	ItemType   string `json:"item_type"`
	RavenMedia struct {
		MediaType int64 `json:"media_type"`
	} `json:"raven_media"`
	ReplyChainCount int           `json:"reply_chain_count"`
	SeenUserIds     []interface{} `json:"seen_user_ids"`
	Timestamp       int64         `json:"timestamp"`
	UserID          int64         `json:"user_id"`
	ViewMode        string        `json:"view_mode"`
}

//InboxItemLike is the heart sent during a conversation.
type InboxItemLike struct {
	ItemID    string `json:"item_id"`
	ItemType  string `json:"item_type"`
	Timestamp int64  `json:"timestamp"`
	UserID    int64  `json:"user_id"`
}

type respLikers struct {
	Users     []User `json:"users"`
	UserCount int64  `json:"user_count"`
	Status    string `json:"status"`
}

type threadResp struct {
	Conversation Conversation `json:"thread"`
	Status       string       `json:"status"`
}

type ErrChallengeProcess struct {
	StepName string
}

func (ec ErrChallengeProcess) Error() string {
	return ec.StepName
}

type UnlogedData struct {
	Graphql struct {
		User struct {
			inst       *Instagram
			Biography  string `json:"biography"`
			FollowedBy struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			Follow struct {
				Count int `json:"count"`
			} `json:"edge_follow"`
			FullName           string `json:"full_name"`
			HighlightReelCount int    `json:"highlight_reel_count"`
			ID                 string `json:"id"`
			IsPrivate          bool   `json:"is_private"`
			IsVerified         bool   `json:"is_verified"`
			ProfilePicURL      string `json:"profile_pic_url"`
			ProfilePicURLHd    string `json:"profile_pic_url_hd"`
			Username           string `json:"username"`
			IGTv               struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Code       string `json:"shortcode"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL            string `json:"display_url"`
						EdgeMediaToTaggedUser struct {
							Edges []interface{} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
						MediaPreview string `json:"media_preview"`
						Owner        struct {
							ID       string `json:"id"`
							Username string `json:"username"`
						} `json:"owner"`
						IsVideo        bool   `json:"is_video"`
						HasAudio       bool   `json:"has_audio"`
						TrackingToken  string `json:"tracking_token"`
						VideoURL       string `json:"video_url"`
						VideoViewCount int    `json:"video_view_count"`
						Comments       struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						CommentsDisabled bool `json:"comments_disabled"`
						Likes            int  `json:"taken_at_timestamp"`
						EdgeLikedBy      struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
						} `json:"edge_media_preview_like"`
						Location           interface{} `json:"location"`
						ThumbnailSrc       string      `json:"thumbnail_src"`
						ThumbnailResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`

						IsPublished   bool    `json:"is_published"`
						ProductType   string  `json:"product_type"`
						Title         string  `json:"title"`
						VideoDuration float64 `json:"video_duration"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_felix_video_timeline"`
			Publications struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Shortcode  string `json:"shortcode"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL string `json:"display_url"`
						Tags       struct {
							Edges []struct {
								Node struct {
									User struct {
										FullName      string `json:"full_name"`
										ID            string `json:"id"`
										IsVerified    bool   `json:"is_verified"`
										ProfilePicURL string `json:"profile_pic_url"`
										Username      string `json:"username"`
									} `json:"user"`
									X float64 `json:"x"`
									Y float64 `json:"y"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
						Owner struct {
							ID       string `json:"id"`
							Username string `json:"username"`
						} `json:"owner"`
						IsVideo              bool   `json:"is_video"`
						AccessibilityCaption string `json:"accessibility_caption"`
						DashInfo             struct {
							IsDashEligible    bool   `json:"is_dash_eligible"`
							VideoDashManifest string `json:"video_dash_manifest"`
							NumberOfQualities int    `json:"number_of_qualities"`
						} `json:"dash_info"`
						HasAudio       bool    `json:"has_audio"`
						TrackingToken  string  `json:"tracking_token"`
						VideoURL       string  `json:"video_url"`
						VideoViewCount float64 `json:"video_view_count"`
						Caption        struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						Comments struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						CommentsDisabled bool  `json:"comments_disabled"`
						TakenAtTimestamp int64 `json:"taken_at_timestamp"`
						Likes            struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
						} `json:"edge_media_preview_like"`
						Location struct {
							ID            string `json:"id"`
							HasPublicPage bool   `json:"has_public_page"`
							Name          string `json:"name"`
							Slug          string `json:"slug"`
						} `json:"location"`
						ThumbnailSrc       string `json:"thumbnail_src"`
						ThumbnailResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`
						Carrousel struct {
							Edges []struct {
								Node struct {
									Typename   string `json:"__typename"`
									ID         string `json:"id"`
									Shortcode  string `json:"shortcode"`
									Dimensions struct {
										Height int `json:"height"`
										Width  int `json:"width"`
									} `json:"dimensions"`
									DisplayURL string `json:"display_url"`
									Tags       struct {
										Edges []struct {
											Node struct {
												User struct {
													FullName      string `json:"full_name"`
													ID            string `json:"id"`
													IsVerified    bool   `json:"is_verified"`
													ProfilePicURL string `json:"profile_pic_url"`
													Username      string `json:"username"`
												} `json:"user"`
												X float64 `json:"x"`
												Y float64 `json:"y"`
											} `json:"node"`
										} `json:"edges"`
									} `json:"edge_media_to_tagged_user"`
									FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
									FactCheckInformation   interface{} `json:"fact_check_information"`
									GatingInfo             interface{} `json:"gating_info"`
									SharingFrictionInfo    struct {
										ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
										BloksAppURL               interface{} `json:"bloks_app_url"`
									} `json:"sharing_friction_info"`
									MediaOverlayInfo interface{} `json:"media_overlay_info"`
									MediaPreview     string      `json:"media_preview"`
									Owner            struct {
										ID       string `json:"id"`
										Username string `json:"username"`
									} `json:"owner"`
									IsVideo              bool   `json:"is_video"`
									AccessibilityCaption string `json:"accessibility_caption"`
									DashInfo             struct {
										IsDashEligible    bool   `json:"is_dash_eligible"`
										VideoDashManifest string `json:"video_dash_manifest"`
										NumberOfQualities int    `json:"number_of_qualities"`
									} `json:"dash_info"`
									HasAudio       bool    `json:"has_audio"`
									TrackingToken  string  `json:"tracking_token"`
									VideoURL       string  `json:"video_url"`
									VideoViewCount float64 `json:"video_view_count"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_sidecar_to_children"`
					} `json:"node,omitempty"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"graphql"`
}
