package models

// Subreddit represents a subreddit
type Subreddit struct {
	UserFlairBackgroundColor   string   `json:"user_flair_background_color"`
	SubmitTextHTML             string   `json:"submit_text_html"`
	RestrictPosting            bool     `json:"restrict_posting"`
	UserIsBanned               bool     `json:"user_is_banned"`
	FreeFormReports            bool     `json:"free_form_reports"`
	WikiEnabled                bool     `json:"wiki_enabled"`
	UserIsMuted                bool     `json:"user_is_muted"`
	UserCanFlairInSr           string   `json:"user_can_flair_in_sr"`
	DisplayName                string   `json:"display_name"`
	HeaderImg                  string   `json:"header_img"`
	Title                      string   `json:"title"`
	IconSize                   []int    `json:"icon_size"`
	PrimaryColor               string   `json:"primary_color"`
	ActiveUserCount            int      `json:"active_user_count"`
	IconImg                    string   `json:"icon_img"`
	AccountsActive             int      `json:"accounts_active"`
	PublicTraffic              bool     `json:"public_traffic"`
	Subscribers                int      `json:"subscribers"`
	UserFlairRichtext          string   `json:"user_flair_richtext"`
	VideostreamLinksCount      int      `json:"videostream_links_count"`
	Name                       RedditID `json:"name"`
	Quarantine                 bool     `json:"quarantine"`
	HideAds                    bool     `json:"hide_ads"`
	EmojisEnabled              bool     `json:"emojis_enabled"`
	AdvertiserCategory         string   `json:"advertiser_category"`
	PublicDescription          string   `json:"public_description"`
	CommentScoreHideMins       int      `json:"comment_score_hide_mins"`
	UserHasFavorited           bool     `json:"user_has_favorited"`
	UserFlairTemplateID        string   `json:"user_flair_template_id"`
	CommunityIcon              string   `json:"community_icon"`
	BannerBackgroundImage      string   `json:"banner_background_image"`
	OriginalContentTagEnabled  bool     `json:"original_content_tag_enabled"`
	SubmitText                 string   `json:"submit_text"`
	DescriptionHTML            string   `json:"description_html"`
	SpoilersEnabled            bool     `json:"spoilers_enabled"`
	HeaderTitle                string   `json:"header_title"`
	HeaderSize                 string   `json:"header_size"`
	UserFlairPosition          string   `json:"user_flair_position"`
	AllOriginalContent         bool     `json:"all_original_content"`
	HasMenuWidget              bool     `json:"has_menu_widget"`
	IsEnrolledInNewModmail     bool     `json:"is_enrolled_in_new_modmail"`
	KeyColor                   string   `json:"key_color"`
	EventPostsEnabled          bool     `json:"event_posts_enabled"`
	CanAssignUserFlair         bool     `json:"can_assign_user_flair"`
	Created                    float64  `json:"created"`
	ShowMediaPreview           bool     `json:"show_media_preview"`
	SubmissionType             string   `json:"submission_type"`
	UserIsSubscriber           bool     `json:"user_is_subscriber"`
	DisableContributorRequests bool     `json:"disable_contributor_requests"`
	AllowVideoGIFs             bool     `json:"allow_videogifs"`
	UserFlairType              string   `json:"user_flair_type"`
	CollapseDeletedComments    bool     `json:"collapse_deleted_comments"`
	EmojisCustomSize           string   `json:"emojis_custom_size"`
	PublicDescriptionHTML      string   `json:"public_description_html"`
	AllowVideos                bool     `json:"allow_videos"`
	NotificationLevel          string   `json:"notification_level"`
	CanAssignLinkFlair         bool     `json:"can_assign_link_flair"`
	AccountsActiveIsFuzzed     bool     `json:"accounts_active_is_fuzzed"`
	SubmitTextLabel            string   `json:"submit_text_label"`
	LinkFlairPosition          string   `json:"link_flair_position"`
	UserSrFlairEnabled         bool     `json:"user_sr_flair_enabled"`
	UserFlairEnabledInSr       bool     `json:"user_flair_enabled_in_sr"`
	AllowDiscovery             bool     `json:"allow_discovery"`
	UserSrThemeEnabled         bool     `json:"user_sr_theme_enabled"`
	LinkFlairEnabled           bool     `json:"link_flair_enabled"`
	SubredditType              string   `json:"subreddit_type"`
	SuggestedCommentSort       string   `json:"suggested_comment_sort"`
	BannerImg                  string   `json:"banner_img"`
	UserFlairText              string   `json:"user_flair_text"`
	BannerBackgroundColor      string   `json:"banner_background_color"`
	ShowMedia                  bool     `json:"show_media"`
	ID                         string   `json:"id"`
	UserIsModerator            bool     `json:"user_is_moderator"`
	Over18                     bool     `json:"over18"`
	Description                string   `json:"description"`
	SubmitLinkLabel            string   `json:"submit_link_label"`
	UserFlairTextColor         string   `json:"user_flair_text_color"`
	RestrictCommenting         bool     `json:"restrict_commenting"`
	UserFlairCSSClass          string   `json:"user_flair_css_class"`
	AllowImages                bool     `json:"allow_images"`
	Lang                       string   `json:"lang"`
	WhitelistStatus            string   `json:"whitelist_status"`
	URL                        string   `json:"url"`
	CreatedUTC                 float64  `json:"created_utc"`
	BannerSize                 []int64  `json:"banner_size"`
	MobileBannerImage          string   `json:"mobile_banner_image"`
	UserIsContributor          bool     `json:"user_is_contributor"`
}
