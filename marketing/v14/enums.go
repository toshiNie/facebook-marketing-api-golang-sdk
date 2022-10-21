package v14

type BidStrategy string

const (
	BidStrategyLowestConstWithoutCap BidStrategy = "LOWEST_COST_WITHOUT_CAP"
	BidStrategyLowestConstWithBidCap BidStrategy = "LOWEST_COST_WITH_BID_CAP"
	BidStrategyCostCap               BidStrategy = "COST_CAP"
)

// status
type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusPaused   Status = "PAUSED"
	StatusDeleted  Status = "DELETED"
	StatusArchived Status = "ARCHIVED"
)

type BuyingType string

const (
	BuyingTypeAuction  BuyingType = "AUCTION"
	BuyingTypeReserved BuyingType = "RESERVED"
)

// objective
type Objective string

const (
	ObjectiveAppInstalls         Objective = "APP_INSTALLS"
	ObjectiveBrandAwareness      Objective = "BRAND_AWARENESS"
	ObjectiveConversions         Objective = "CONVERSIONS"
	ObjectiveEventResponses      Objective = "EVENT_RESPONSES"
	ObjectiveLeadGeneration      Objective = "LEAD_GENERATION"
	ObjectiveLinkClicks          Objective = "LINK_CLICKS"
	ObjectiveLocalAwareness      Objective = "LOCAL_AWARENESS"
	ObjectiveMessages            Objective = "MESSAGES"
	ObjectiveOfferClaims         Objective = "OFFER_CLAIMS"
	ObjectivePageLikes           Objective = "PAGE_LIKES"
	ObjectivePostEngagement      Objective = "POST_ENGAGEMENT"
	ObjectiveProductCatalogSales Objective = "PRODUCT_CATALOG_SALES"
	ObjectiveReach               Objective = "REACH"
	ObjectiveStoreVisits         Objective = "STORE_VISITS"
	ObjectiveVideoViews          Objective = "VIDEO_VIEWS"
)

type SpecialAdCategories string

const (
	SpecialAdCategoriesNone                    SpecialAdCategories = "NONE"
	SpecialAdCategoriesEmployment              SpecialAdCategories = "EMPLOYMENT"
	SpecialAdCategoriesHousing                 SpecialAdCategories = "HOUSING"
	SpecialAdCategoriesCredit                  SpecialAdCategories = "CREDIT"
	SpecialAdCategoriesIssuesElectionsPolitics SpecialAdCategories = "ISSUES_ELECTIONS_POLITICS"
	SpecialAdCategoriesOnlineGamblingAndGaming SpecialAdCategories = "ONLINE_GAMBLING_AND_GAMING"
)

type BillingEvent string

const (
	BillingEventAppInstalls        BillingEvent = "APP_INSTALLS"
	BillingEventClicks             BillingEvent = "CLICKS"
	BillingEventImpressions        BillingEvent = "IMPRESSIONS"
	BillingEventLinkClicks         BillingEvent = "LINK_CLICKS"
	BillingEventNone               BillingEvent = "NONE"
	BillingEventOfferClaims        BillingEvent = "OFFER_CLAIMS"
	BillingEventPageLikes          BillingEvent = "PAGE_LIKES"
	BillingEventPostEngagement     BillingEvent = "POST_ENGAGEMENT"
	BillingEventThruplay           BillingEvent = "THRUPLAY"
	BillingEventPurchase           BillingEvent = "PURCHASE"
	BillingEventListingInteraction BillingEvent = "LISTING_INTERACTION"
)

type DestinationType string

const (
	DestinationTypeUndefined         DestinationType = "UNDEFINED"
	DestinationTypeWebsite           DestinationType = "WEBSITE"
	DestinationTypeApp               DestinationType = "APP"
	DestinationTypeMessenger         DestinationType = "MESSENGER"
	DestinationTypeApplinksAutomatic DestinationType = "APPLINKS_AUTOMATIC"
	DestinationTypeFacebook          DestinationType = "FACEBOOK"
)

type OptimizationGoal string

const (
	OptimizationGoalNone                             OptimizationGoal = "NONE"
	OptimizationGoalAppInstalls                      OptimizationGoal = "APP_INSTALLS"
	OptimizationGoalAdRecallLift                     OptimizationGoal = "AD_RECALL_LIFT"
	OptimizationGoalEngagedUsers                     OptimizationGoal = "ENGAGED_USERS"
	OptimizationGoalEventResponses                   OptimizationGoal = "EVENT_RESPONSES"
	OptimizationGoalImpressions                      OptimizationGoal = "IMPRESSIONS"
	OptimizationGoalLeadGeneration                   OptimizationGoal = "LEAD_GENERATION"
	OptimizationGoalQualityLead                      OptimizationGoal = "QUALITY_LEAD"
	OptimizationGoalLinkClicks                       OptimizationGoal = "LINK_CLICKS"
	OptimizationGoalOffsiteConversions               OptimizationGoal = "OFFSITE_CONVERSIONS"
	OptimizationGoalPageLikes                        OptimizationGoal = "PAGE_LIKES"
	OptimizationGoalPostEngagement                   OptimizationGoal = "POST_ENGAGEMENT"
	OptimizationGoalQualityCall                      OptimizationGoal = "QUALITY_CALL"
	OptimizationGoalReach                            OptimizationGoal = "REACH"
	OptimizationGoalLandingPageViews                 OptimizationGoal = "LANDING_PAGE_VIEWS"
	OptimizationGoalVisitInstagramProfile            OptimizationGoal = "VISIT_INSTAGRAM_PROFILE"
	OptimizationGoalValue                            OptimizationGoal = "VALUE"
	OptimizationGoalThruplay                         OptimizationGoal = "THRUPLAY"
	OptimizationGoalDerivedEvents                    OptimizationGoal = "DERIVED_EVENTS"
	OptimizationGoalAppInstallsAndOffsiteConversions OptimizationGoal = "APP_INSTALLS_AND_OFFSITE_CONVERSIONS"
	OptimizationGoalConversations                    OptimizationGoal = "CONVERSATIONS"
	OptimizationGoalInAppValue                       OptimizationGoal = "IN_APP_VALUE"
	OptimizationGoalMessagingPurchaseConversion      OptimizationGoal = "MESSAGING_PURCHASE_CONVERSION"
)

type OptimizationSubEvent string

const (
	OptimizationSubEventNone                            OptimizationSubEvent = "NONE"
	OptimizationSubEventVideoSoundOn                    OptimizationSubEvent = "VIDEO_SOUND_ON"
	OptimizationSubEventTripConsideration               OptimizationSubEvent = "TRIP_CONSIDERATION"
	OptimizationSubEventTravelIntent                    OptimizationSubEvent = "TRAVEL_INTENT"
	OptimizationSubEventTravelIntentNoDestinationIntent OptimizationSubEvent = "TRAVEL_INTENT_NO_DESTINATION_INTENT"
	OptimizationSubEventTravelIntentBucket01            OptimizationSubEvent = "TRAVEL_INTENT_BUCKET_01"
	OptimizationSubEventTravelIntentBucket02            OptimizationSubEvent = "TRAVEL_INTENT_BUCKET_02"
	OptimizationSubEventTravelIntentBucket03            OptimizationSubEvent = "TRAVEL_INTENT_BUCKET_03"
	OptimizationSubEventTravelIntentBucket04            OptimizationSubEvent = "TRAVEL_INTENT_BUCKET_04"
	OptimizationSubEventTravelIntentBucket05            OptimizationSubEvent = "TRAVEL_INTENT_BUCKET_05"
)

type TuneForCategory string

const (
	TuneForCategoryNone                    TuneForCategory = "NONE"
	TuneForCategoryEmployment              TuneForCategory = "EMPLOYMENT"
	TuneForCategoryHousing                 TuneForCategory = "HOUSING"
	TuneForCategoryCredit                  TuneForCategory = "CREDIT"
	TuneForCategoryIssuesElectionsPolitics TuneForCategory = "ISSUES_ELECTIONS_POLITICS"
	TuneForCategoryOnlineGamblingAndGaming TuneForCategory = "ONLINE_GAMBLING_AND_GAMING"
)

type EffectiveStatus string

const (
	EffectiveStatusActive         EffectiveStatus = "ACTIVE"
	EffectiveStatusPaused         EffectiveStatus = "PAUSED"
	EffectiveStatusDeleted        EffectiveStatus = "DELETED"
	EffectiveStatusCampaignPaused EffectiveStatus = "CAMPAIGN_PAUSED"
	EffectiveStatusArchived       EffectiveStatus = "ARCHIVED"
	EffectiveStatusInProcess      EffectiveStatus = "IN_PROCESS"
	EffectiveStatusWithIssues     EffectiveStatus = "WITH_ISSUES"
)

type AdCreativeLinkDataCallToActionType string

const (
	AdCreativeLinkDataCallToActionTypeOpenLink            AdCreativeLinkDataCallToActionType = "OPEN_LINK"
	AdCreativeLinkDataCallToActionTypeLikePage            AdCreativeLinkDataCallToActionType = "LIKE_PAGE"
	AdCreativeLinkDataCallToActionTypeShopNow             AdCreativeLinkDataCallToActionType = "SHOP_NOW"
	AdCreativeLinkDataCallToActionTypePlayGame            AdCreativeLinkDataCallToActionType = "PLAY_GAME"
	AdCreativeLinkDataCallToActionTypeInstallApp          AdCreativeLinkDataCallToActionType = "INSTALL_APP"
	AdCreativeLinkDataCallToActionTypeUseApp              AdCreativeLinkDataCallToActionType = "USE_APP"
	AdCreativeLinkDataCallToActionTypeCall                AdCreativeLinkDataCallToActionType = "CALL"
	AdCreativeLinkDataCallToActionTypeCallMe              AdCreativeLinkDataCallToActionType = "CALL_ME"
	AdCreativeLinkDataCallToActionTypeVideoCall           AdCreativeLinkDataCallToActionType = "VIDEO_CALL"
	AdCreativeLinkDataCallToActionTypeInstallMobileApp    AdCreativeLinkDataCallToActionType = "INSTALL_MOBILE_APP"
	AdCreativeLinkDataCallToActionTypeUseMobileApp        AdCreativeLinkDataCallToActionType = "USE_MOBILE_APP"
	AdCreativeLinkDataCallToActionTypeMobileDownload      AdCreativeLinkDataCallToActionType = "MOBILE_DOWNLOAD"
	AdCreativeLinkDataCallToActionTypeBookTravel          AdCreativeLinkDataCallToActionType = "BOOK_TRAVEL"
	AdCreativeLinkDataCallToActionTypeListenMusic         AdCreativeLinkDataCallToActionType = "LISTEN_MUSIC"
	AdCreativeLinkDataCallToActionTypeWatchVideo          AdCreativeLinkDataCallToActionType = "WATCH_VIDEO"
	AdCreativeLinkDataCallToActionTypeLearnMore           AdCreativeLinkDataCallToActionType = "LEARN_MORE"
	AdCreativeLinkDataCallToActionTypeSignUp              AdCreativeLinkDataCallToActionType = "SIGN_UP"
	AdCreativeLinkDataCallToActionTypeDownload            AdCreativeLinkDataCallToActionType = "DOWNLOAD"
	AdCreativeLinkDataCallToActionTypeWatchMore           AdCreativeLinkDataCallToActionType = "WATCH_MORE"
	AdCreativeLinkDataCallToActionTypeNoButton            AdCreativeLinkDataCallToActionType = "NO_BUTTON"
	AdCreativeLinkDataCallToActionTypeVisitPagesFeed      AdCreativeLinkDataCallToActionType = "VISIT_PAGES_FEED"
	AdCreativeLinkDataCallToActionTypeCallNow             AdCreativeLinkDataCallToActionType = "CALL_NOW"
	AdCreativeLinkDataCallToActionTypeApplyNow            AdCreativeLinkDataCallToActionType = "APPLY_NOW"
	AdCreativeLinkDataCallToActionTypeContact             AdCreativeLinkDataCallToActionType = "CONTACT"
	AdCreativeLinkDataCallToActionTypeBuyNow              AdCreativeLinkDataCallToActionType = "BUY_NOW"
	AdCreativeLinkDataCallToActionTypeGetOffer            AdCreativeLinkDataCallToActionType = "GET_OFFER"
	AdCreativeLinkDataCallToActionTypeGetOfferView        AdCreativeLinkDataCallToActionType = "GET_OFFER_VIEW"
	AdCreativeLinkDataCallToActionTypeBuyTickets          AdCreativeLinkDataCallToActionType = "BUY_TICKETS"
	AdCreativeLinkDataCallToActionTypeUpdateApp           AdCreativeLinkDataCallToActionType = "UPDATE_APP"
	AdCreativeLinkDataCallToActionTypeGetDirections       AdCreativeLinkDataCallToActionType = "GET_DIRECTIONS"
	AdCreativeLinkDataCallToActionTypeBuy                 AdCreativeLinkDataCallToActionType = "BUY"
	AdCreativeLinkDataCallToActionTypeMessagePage         AdCreativeLinkDataCallToActionType = "MESSAGE_PAGE"
	AdCreativeLinkDataCallToActionTypeDonate              AdCreativeLinkDataCallToActionType = "DONATE"
	AdCreativeLinkDataCallToActionTypeSubscribe           AdCreativeLinkDataCallToActionType = "SUBSCRIBE"
	AdCreativeLinkDataCallToActionTypeSayThanks           AdCreativeLinkDataCallToActionType = "SAY_THANKS"
	AdCreativeLinkDataCallToActionTypeSellNow             AdCreativeLinkDataCallToActionType = "SELL_NOW"
	AdCreativeLinkDataCallToActionTypeShare               AdCreativeLinkDataCallToActionType = "SHARE"
	AdCreativeLinkDataCallToActionTypeDonateNow           AdCreativeLinkDataCallToActionType = "DONATE_NOW"
	AdCreativeLinkDataCallToActionTypeGetQuote            AdCreativeLinkDataCallToActionType = "GET_QUOTE"
	AdCreativeLinkDataCallToActionTypeContactUs           AdCreativeLinkDataCallToActionType = "CONTACT_US"
	AdCreativeLinkDataCallToActionTypeOrderNow            AdCreativeLinkDataCallToActionType = "ORDER_NOW"
	AdCreativeLinkDataCallToActionTypeStartOrder          AdCreativeLinkDataCallToActionType = "START_ORDER"
	AdCreativeLinkDataCallToActionTypeAddToCart           AdCreativeLinkDataCallToActionType = "ADD_TO_CART"
	AdCreativeLinkDataCallToActionTypeVideoAnnotation     AdCreativeLinkDataCallToActionType = "VIDEO_ANNOTATION"
	AdCreativeLinkDataCallToActionTypeMoments             AdCreativeLinkDataCallToActionType = "MOMENTS"
	AdCreativeLinkDataCallToActionTypeRecordNow           AdCreativeLinkDataCallToActionType = "RECORD_NOW"
	AdCreativeLinkDataCallToActionTypeReferFriends        AdCreativeLinkDataCallToActionType = "REFER_FRIENDS"
	AdCreativeLinkDataCallToActionTypeRequestTime         AdCreativeLinkDataCallToActionType = "REQUEST_TIME"
	AdCreativeLinkDataCallToActionTypeGetShowtimes        AdCreativeLinkDataCallToActionType = "GET_SHOWTIMES"
	AdCreativeLinkDataCallToActionTypeListenNow           AdCreativeLinkDataCallToActionType = "LISTEN_NOW"
	AdCreativeLinkDataCallToActionTypeWoodhengeSupport    AdCreativeLinkDataCallToActionType = "WOODHENGE_SUPPORT"
	AdCreativeLinkDataCallToActionTypeSottoSubscribe      AdCreativeLinkDataCallToActionType = "SOTTO_SUBSCRIBE"
	AdCreativeLinkDataCallToActionTypeFollowUser          AdCreativeLinkDataCallToActionType = "FOLLOW_USER"
	AdCreativeLinkDataCallToActionTypeRaiseMoney          AdCreativeLinkDataCallToActionType = "RAISE_MONEY"
	AdCreativeLinkDataCallToActionTypeEventRsvp           AdCreativeLinkDataCallToActionType = "EVENT_RSVP"
	AdCreativeLinkDataCallToActionTypeWhatsappMessage     AdCreativeLinkDataCallToActionType = "WHATSAPP_MESSAGE"
	AdCreativeLinkDataCallToActionTypeFollowNewsStoryline AdCreativeLinkDataCallToActionType = "FOLLOW_NEWS_STORYLINE"
	AdCreativeLinkDataCallToActionTypeSeeMore             AdCreativeLinkDataCallToActionType = "SEE_MORE"
	AdCreativeLinkDataCallToActionTypeFindAGroup          AdCreativeLinkDataCallToActionType = "FIND_A_GROUP"
	AdCreativeLinkDataCallToActionTypeFindYourGroups      AdCreativeLinkDataCallToActionType = "FIND_YOUR_GROUPS"
	AdCreativeLinkDataCallToActionTypePayToAccess         AdCreativeLinkDataCallToActionType = "PAY_TO_ACCESS"
	AdCreativeLinkDataCallToActionTypePurchaseGiftCards   AdCreativeLinkDataCallToActionType = "PURCHASE_GIFT_CARDS"
	AdCreativeLinkDataCallToActionTypeFollowPage          AdCreativeLinkDataCallToActionType = "FOLLOW_PAGE"
	AdCreativeLinkDataCallToActionTypeSendAGift           AdCreativeLinkDataCallToActionType = "SEND_A_GIFT"
	AdCreativeLinkDataCallToActionTypeSwipeUpShop         AdCreativeLinkDataCallToActionType = "SWIPE_UP_SHOP"
	AdCreativeLinkDataCallToActionTypeSwipeUpProduct      AdCreativeLinkDataCallToActionType = "SWIPE_UP_PRODUCT"
	AdCreativeLinkDataCallToActionTypeSendGiftMoney       AdCreativeLinkDataCallToActionType = "SEND_GIFT_MONEY"
	AdCreativeLinkDataCallToActionTypePlayGameOnFacebook  AdCreativeLinkDataCallToActionType = "PLAY_GAME_ON_FACEBOOK"
	AdCreativeLinkDataCallToActionTypeGetStarted          AdCreativeLinkDataCallToActionType = "GET_STARTED"
)

type AdCreativeStatus string

const (
	AdCreativeStatusActive     AdCreativeStatus = "ACTIVE"
	AdCreativeStatusInProcess  AdCreativeStatus = "IN_PROCESS"
	AdCreativeStatusWithIssues AdCreativeStatus = "WITH_ISSUES"
	AdCreativeStatusDeleted    AdCreativeStatus = "DELETED"
)

type PublisherPlatform string

const (
	PublisherPlatformFacebook        PublisherPlatform = "facebook"
	PublisherPlatformInstagram       PublisherPlatform = "instagram"
	PublisherPlatformMessenger       PublisherPlatform = "messenger"
	PublisherPlatformAudienceNetwork PublisherPlatform = "audience_network"
)
