package user_dao

// User 用户
type User struct {
	ID         int64    `json:"id"`          // ID
	CreateTime int64    `json:"create_time"` // 创建时间
	UpdateTime int64    `json:"update_time"` // 更新时间
	Email      string   `json:"email"`       // 邮箱，使用唯一键
	UserID     int64    `json:"user_id"`     // 用户ID 使用唯一键
	Password   string   `json:"password"`    // 加密密码
	UserInfo   UserInfo `json:"user_info"`
	UserRole   UserRole `json:"user_role"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID                int64  `json:"id"`                  // ID
	CreateTime        int64  `json:"create_time"`         // 创建时间
	UpdateTime        int64  `json:"update_time"`         // 更新时间
	UserID            int64  `json:"user_id"`             // 用户ID
	Nickname          string `json:"nickname"`            // 昵称
	Signature         string `json:"signature"`           // 个性签名
	Avatar            string `json:"avatar"`              // 头像
	Birthday          int64  `json:"birthday"`            //生日 默认2025-01-26
	Sex               int8   `json:"sex"`                 // 性别 1 男，2 女，3 未选择 默认未选择(3)
	PointsRemaining   int32  `json:"points_remaining"`    // 剩余积分
	NumberOfResources int16  `json:"number_of_resources"` // 资源数
	NumberOfPosts     int16  `json:"number_of_posts"`     // 帖子数
}

// UserRole 用户角色
type UserRole struct {
	ID         int64 `json:"id"`          // ID
	CreateTime int64 `json:"create_time"` // 创建时间
	UpdateTime int64 `json:"update_time"` // 更新时间
	UserID     int64 `json:"user_id"`     // 用户ID
	IsAdmin    int8  `json:"is_admin"`    // 是否管理员 1 是管理员，2 不是管理员 默认不是(2)
	IsDisabled int8  `json:"is_disabled"` // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
	IsOrder    int8  `json:"is_order"`    // 是否有订单 1 有订单，2 无订单 默认无订单(2)
}

// Category 栏目表
type Category struct {
	ID           int64  `json:"id"`            // ID
	CreateTime   int64  `json:"create_time"`   // 创建时间
	UpdateTime   int64  `json:"update_time"`   // 更新时间
	CategoryName string `json:"category_name"` // 栏目名称
	CategoryType int8   `json:"category_type"` // 栏目类型 1 帖子，2 资源，3 视频
}

// Resources 资源表
type Resources struct {
	ID                   int64  `json:"id"`                     // ID
	CreateTime           int64  `json:"create_time"`            // 创建时间
	UpdateTime           int64  `json:"update_time"`            // 更新时间
	UserID               int64  `json:"user_id"`                // 用户ID
	CategoryID           int64  `json:"category_id"`            // 栏目ID
	ResourcesName        string `json:"resources_name"`         // 资源名称
	SimpleDescription    string `json:"simple_description"`     // 简易描述
	RichTextDescriptions string `json:"rich_text_descriptions"` // 富文本描述
	PictureAddress       string `json:"picture_address"`        // 图片地址
	FileAddress          string `json:"file_address"`           // 文件地址
	IsComment            int8   `json:"is_comment"`             // 是否评论 1 是，2 否 默认否(2)
	NumberOfViews        int16  `json:"number_of_views"`        // 浏览量
	NumberOfLikes        int16  `json:"number_of_likes"`        // 点赞量
	NumberOfFavorites    int16  `json:"number_of_favorites"`    // 收藏量
	NumberOfComments     int16  `json:"number_of_comments"`     // 评论量
	NumberOfDownloads    int16  `json:"number_of_downloads"`    // 下载量
	PointsRequired       int32  `json:"points_required"`        // 所需积分
	PointsReward         int32  `json:"points_reward"`          // 奖励积分
	IsDisabled           int8   `json:"is_disabled"`            // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
	IsReview             int8   `json:"is-review"`              // 是否审核 1 审核通过，2 未审核 默认未审核(2)
}

// DownloadTheRecordSheet 下载记录表
type DownloadTheRecordSheet struct {
	ID             int64 `json:"id"`              // ID
	CreateTime     int64 `json:"create_time"`     // 创建时间
	UpdateTime     int64 `json:"update_time"`     // 更新时间
	UserID         int64 `json:"user_id"`         // 用户ID
	ResourcesID    int64 `json:"resources_id"`    // 资源ID
	PointsRequired int16 `json:"points_required"` // 所需积分
	IsDownload     int8  `json:"is_download"`     // 是否下载 1 下载，2 未下载 默认未下载(2)
}

// Post 帖子表
type Post struct {
	ID                   int64  `json:"id"`                     // ID
	CreateTime           int64  `json:"create_time"`            // 创建时间
	UpdateTime           int64  `json:"update_time"`            // 更新时间
	UserID               int64  `json:"user_id"`                // 用户ID
	CategoryID           int64  `json:"category_id"`            // 栏目ID
	PostName             string `json:"post_name"`              // 帖子名称
	SimpleDescription    string `json:"simple_description"`     // 简易描述
	RichTextDescriptions string `json:"rich_text_descriptions"` // 富文本描述
	PictureAddress       string `json:"picture_address"`        // 图片地址
	FileAddress          string `json:"file_address"`           // 文件地址
	NumberOfViews        int16  `json:"number_of_views"`        // 浏览量
	NumberOfLikes        int16  `json:"number_of_likes"`        // 点赞量
	NumberOfFavorites    int16  `json:"number_of_favorites"`    // 收藏量
	NumberOfComments     int16  `json:"number_of_comments"`     // 评论量
	IsDisabled           int8   `json:"is_disabled"`            // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
	IsReview             int8   `json:"is_review"`              // 是否审核 1 审核通过，2 未审核 默认未审核(2)
}

// Video 视频表
type Video struct {
	ID                       int64  `json:"id"`                          // ID
	CreateTime               int64  `json:"create_time"`                 // 创建时间
	UpdateTime               int64  `json:"update_time"`                 // 更新时间
	CategoryID               int64  `json:"category_id"`                 // 栏目ID
	VideoName                string `json:"video_name"`                  // 视频名称
	SimpleDescription        string `json:"simple_description"`          // 简易描述
	RichTextDescriptions     string `json:"rich_text_descriptions"`      // 富文本描述
	PictureAddress           string `json:"picture_address"`             // 图片地址
	NumberOfViews            int16  `json:"number_of_views"`             // 浏览量
	NumberOfLikes            int16  `json:"number_of_likes"`             // 点赞量
	NumberOfFavorites        int16  `json:"number_of_favorites"`         // 收藏量
	NumberOfComments         int16  `json:"number_of_comments"`          // 评论量
	NameOfTheVideoGroup      string `json:"name_of_the_video_group"`     // 视频分组名称
	TencentCloudID           string `json:"tencent_cloud_id"`            // 腾讯云ID
	TencentCloudVideoID      string `json:"tencent_cloud_video_id"`      // 腾讯云视频ID
	TencentCloudVideoAddress string `json:"tencent_cloud_video_address"` // 腾讯云视频地址
	VideoPermissions         int8   `json:"video_permissions"`           // 视频权限 1 需要会员，2 公开 默认公开(2)
	IsDisabled               int8   `json:"is_disabled"`                 // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
}

// VideoBrowsingHistory 视频浏览记录表
type VideoBrowsingHistory struct {
	ID         int64 `json:"id"`          // ID
	CreateTime int64 `json:"create_time"` // 创建时间
	UpdateTime int64 `json:"update_time"` // 更新时间
	UserID     int64 `json:"user_id"`     // 用户ID
	VideoID    int64 `json:"video_id"`    // 视频ID
	WatchTime  int64 `json:"watch_time"`  // 观看时长
}

// TableOfPointsChanges 积分变化表
type TableOfPointsChanges struct {
	ID          int64 `json:"id"`           // ID
	CreateTime  int64 `json:"create_time"`  // 创建时间
	UpdateTime  int64 `json:"update_time"`  // 更新时间
	RawPoints   int32 `json:"raw_points"`   // 原始积分
	ThisPoints  int32 `json:"this_points"`  // 当前积分
	AfterPoints int32 `json:"after_points"` // 消费后积分
	UserID      int64 `json:"user_id"`      // 用户ID
	Type        int8  `json:"type"`         // 类型 1 下载资源 2 资源奖励，3 视频资源，4 充值资源，5邀请用户赠送，6 其他
}

// PaymentHistory 支付记录表
type PaymentHistory struct {
	ID                        int64  `json:"id"`                          // ID
	CreateTime                int64  `json:"create_time"`                 // 创建时间
	UpdateTime                int64  `json:"update_time"`                 // 更新时间
	UserID                    int64  `json:"user_id"`                     // 用户ID
	TheAmountToBePaid         int32  `json:"the_amount_to_be_paid"`       // 支付金额
	OrderNumber               int64  `json:"order_number"`                // 订单号
	IsSuccessful              int8   `json:"is_successful"`               // 是否支付成功 1 没有成功，2 成功
	SuccessfulPaymentCallback string `json:"successful_payment_callback"` // 支付成功回调
}

// Comments 评论表
type Comments struct {
	ID             int64  `json:"id"`              // ID
	CreateTime     int64  `json:"create_time"`     // 创建时间
	UpdateTime     int64  `json:"update_time"`     // 更新时间
	UserID         int64  `json:"user_id"`         // 用户ID
	ParentID       int64  `json:"parent_id"`       // 父级ID (资源ID/帖子ID/视频ID)
	CommentsType   int8   `json:"comments_type"`   // 评论类型 1 资源评论，2 帖子评论，3 视频评论
	CommentContent string `json:"comment_content"` // 评论内容
}

// Report 举报表
type Report struct {
	ID                 int64  `json:"id"`                   // ID
	CreateTime         int64  `json:"create_time"`          // 创建时间
	UpdateTime         int64  `json:"update_time"`          // 更新时间
	UserID             int64  `json:"user_id"`              // 用户ID
	ReportType         int8   `json:"report_type"`          // 举报类型 1 资源举报，2 帖子举报，3 视频举报，4 用户，5 评论
	BeReportedID       int64  `json:"be_reported_id"`       // 被举报的ID
	ReasonForReporting string `json:"reason_for_reporting"` // 举报原因
}

// PrivateMessages 私信表
type PrivateMessages struct {
	ID         int64  `json:"id"`          // ID
	CreateTime int64  `json:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time"` // 更新时间
	SenderID   int64  `json:"sender_id"`   // 发送者ID (当前账号的用户ID)
	ReceiverID int64  `json:"receiver_id"` // 接收者ID (对方账号的用户ID)
	Content    string `json:"content"`     // 内容
}

// ThumbsUp 点赞表
type ThumbsUp struct {
	ID           int64 `json:"id"`             // ID
	CreateTime   int64 `json:"create_time"`    // 创建时间
	UpdateTime   int64 `json:"update_time"`    // 更新时间
	UserID       int64 `json:"user_id"`        // 用户ID
	ThumbsUpType int8  `json:"thumbs_up_type"` // 点赞类型 1 资源，2 帖子，3 视频
	LikedID      int64 `json:"liked_id"`       // 被点赞的ID (资源ID/帖子ID/视频ID)
}

// Collect 收藏表
type Collect struct {
	ID          int64 `json:"id"`           // ID
	CreateTime  int64 `json:"create_time"`  // 创建时间
	UpdateTime  int64 `json:"update_time"`  // 更新时间
	UserID      int64 `json:"user_id"`      // 用户ID
	CollectType int8  `json:"collect_type"` // 收藏类型 1 资源，2 帖子，3 视频
	FavoriteID  int64 `json:"favorite_id"`  // 被收藏的ID (资源ID/帖子ID/视频ID)
}

// FriendChains 友链表
type FriendChains struct {
	ID         int64  `json:"id"`          // ID
	CreateTime int64  `json:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time"` // 更新时间
	ChainsName string `json:"chains_name"` // 友链名称
	ChainsUrl  string `json:"chains_url"`  // 友链地址
	IsDisabled int8   `json:"is_disabled"` // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
}

// Advert 广告表
type Advert struct {
	ID           int64  `json:"id"`            // ID
	CreateTime   int64  `json:"create_time"`   // 创建时间
	UpdateTime   int64  `json:"update_time"`   // 更新时间
	AdvertName   string `json:"advert_name"`   // 广告名称
	AdvertUrl    string `json:"advert_url"`    // 广告地址
	AdvertAvatar string `json:"advert_avatar"` // 广告图片
	IsDisabled   int8   `json:"is_disabled"`   // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
}

// WebsiteInfo 网站信息表
type WebsiteInfo struct {
	ID                        int64  `json:"id"`                             // ID
	CreateTime                int64  `json:"create_time"`                    // 创建时间
	UpdateTime                int64  `json:"update_time"`                    // 更新时间
	TheWebsiteICPFilingNumber string `json:"the_website_icp_filing_number"`  // 网站备案号
	TheDomainNameOfTheWebsite string `json:"the_domain_name_of_the_website"` // 网站域名
	TheNameOfTheWebsite       string `json:"the_name_of_the_website"`        // 网站名称
	LoginMode                 int8   `json:"login_mode"`                     // 登录模式 1 只能同时登录一个账号，2 可以同时多点登录
	AlipayPID                 string `json:"alipay_pid"`                     // 支付宝开放平台账号
	AlipayRSAKey              string `json:"alipay_rsa_key"`                 // 支付宝RSA私钥
	TencentCloudVideoKey      string `json:"tencent_cloud_video_key"`        // 腾讯云视频密钥
}
