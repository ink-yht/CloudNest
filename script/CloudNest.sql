-- 创建用户表
CREATE TABLE user (
                      id bigint AUTO_INCREMENT NOT NULL COMMENT '用户ID',
                      create_time bigint DEFAULT NULL COMMENT '创建时间',
                      update_time bigint DEFAULT NULL COMMENT '更新时间',
                      email varchar(100) NOT NULL COMMENT '邮箱，使用唯一键',
                      user_id bigint DEFAULT NULL COMMENT '用户名，使用唯一键',
                      password varchar(64) DEFAULT NULL COMMENT '加密密码',
                      UNIQUE KEY (email),
                      UNIQUE KEY (user_id),
                      PRIMARY KEY (id)
) COMMENT '用户表';

-- 创建用户信息表
CREATE TABLE user_info (
                           id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                           create_time bigint DEFAULT NULL COMMENT '创建时间',
                           update_time bigint DEFAULT NULL COMMENT '更新时间',
                           user_id bigint DEFAULT NULL COMMENT '用户ID',
                           nickname varchar(30) DEFAULT NULL COMMENT '昵称',
                           signature varchar(100) DEFAULT NULL COMMENT '个性签名',
                           avatar varchar(255) DEFAULT NULL COMMENT '头像',
                           birthday bigint DEFAULT 1737830400000 COMMENT '生日 默认2025-01-26',
                           sex tinyint DEFAULT 3 COMMENT '性别 1 男，2 女，3 未选择 默认未选择(3)',
                           points_remaining int DEFAULT 0 COMMENT '剩余积分',
                           number_of_resources smallint DEFAULT 0 COMMENT '资源数',
                           number_of_posts smallint DEFAULT 0 COMMENT '帖子数',
                           PRIMARY KEY (id)
) COMMENT '用户信息表';

-- 创建用户角色表
CREATE TABLE user_role (
                           id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                           create_time bigint DEFAULT NULL COMMENT '创建时间',
                           update_time bigint DEFAULT NULL COMMENT '更新时间',
                           user_id bigint DEFAULT NULL COMMENT '用户ID',
                           is_admin tinyint DEFAULT 2 COMMENT '是否管理员 1 是管理员，2 不是管理员 默认不是(2)',
                           is_disabled tinyint DEFAULT 2 COMMENT '是否禁用 1 禁用，2 未禁用 默认未禁用(2)',
                           is_order tinyint DEFAULT 2 COMMENT '是否有订单 1 有订单，2 无订单 默认无订单(2)',
                           PRIMARY KEY (id)
) COMMENT '用户角色表';

-- 创建栏目表
CREATE TABLE category (
                          id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                          create_time bigint DEFAULT NULL COMMENT '创建时间',
                          update_time bigint DEFAULT NULL COMMENT '更新时间',
                          category_name varchar(50) NOT NULL COMMENT '栏目名称',
                          category_type tinyint DEFAULT NULL COMMENT '栏目类型 1 帖子，2 资源，3 视频',
                          PRIMARY KEY (id)
) COMMENT '栏目表';

-- 创建资源表
CREATE TABLE resources (
                           id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                           create_time bigint DEFAULT NULL COMMENT '创建时间',
                           update_time bigint DEFAULT NULL COMMENT '更新时间',
                           user_id bigint DEFAULT NULL COMMENT '用户ID',
                           category_id bigint DEFAULT NULL COMMENT '栏目ID',
                           resources_name varchar(100) NOT NULL COMMENT '资源名称',
                           simple_description varchar(200) DEFAULT NULL COMMENT '简易描述',
                           rich_text_descriptions text DEFAULT NULL COMMENT '富文本描述',
                           picture_address varchar(255) DEFAULT NULL COMMENT '图片地址',
                           file_address varchar(255) DEFAULT NULL COMMENT '文件地址',
                           is_comment tinyint DEFAULT 2 COMMENT '是否评论 1 是，2 否 默认否(2)',
                           number_of_views smallint DEFAULT 0 COMMENT '浏览量',
                           number_of_likes smallint DEFAULT 0 COMMENT '点赞量',
                           number_of_favorites smallint DEFAULT 0 COMMENT '收藏量',
                           number_of_comments smallint DEFAULT 0 COMMENT '评论量',
                           number_of_downloads smallint DEFAULT 0 COMMENT '下载量',
                           points_required int DEFAULT 0 COMMENT '所需积分',
                           points_reward int DEFAULT 0 COMMENT '奖励积分',
                           is_disabled tinyint DEFAULT 2 COMMENT '是否禁用 1 禁用，2 未禁用 默认未禁用(2)',
                           is_review tinyint DEFAULT 2 COMMENT '是否审核 1 审核通过，2 未审核 默认未审核(2)',
                           PRIMARY KEY (id)
) COMMENT '资源表';

-- 创建下载记录表
CREATE TABLE download_the_record_sheet (
                                           id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                                           create_time bigint DEFAULT NULL COMMENT '创建时间',
                                           update_time bigint DEFAULT NULL COMMENT '更新时间',
                                           user_id bigint DEFAULT NULL COMMENT '用户ID',
                                           resources_id bigint DEFAULT NULL COMMENT '资源ID',
                                           points_required smallint DEFAULT 0 COMMENT '所需积分',
                                           is_download tinyint DEFAULT 2 COMMENT '是否下载 1 下载，2 未下载 默认未下载(2)',
                                           PRIMARY KEY (id)
) COMMENT '下载记录表';

-- 创建帖子表
CREATE TABLE post (
                      id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                      create_time bigint DEFAULT NULL COMMENT '创建时间',
                      update_time bigint DEFAULT NULL COMMENT '更新时间',
                      user_id bigint DEFAULT NULL COMMENT '用户ID',
                      category_id bigint DEFAULT NULL COMMENT '栏目ID',
                      post_name varchar(100) NOT NULL COMMENT '帖子名称',
                      simple_description varchar(200) DEFAULT NULL COMMENT '简易描述',
                      rich_text_descriptions text DEFAULT NULL COMMENT '富文本描述',
                      picture_address varchar(255) DEFAULT NULL COMMENT '图片地址',
                      file_address varchar(255) DEFAULT NULL COMMENT '文件地址',
                      number_of_views smallint DEFAULT 0 COMMENT '浏览量',
                      number_of_likes smallint DEFAULT 0 COMMENT '点赞量',
                      number_of_favorites smallint DEFAULT 0 COMMENT '收藏量',
                      number_of_comments smallint DEFAULT 0 COMMENT '评论量',
                      is_disabled tinyint DEFAULT 2 COMMENT '是否禁用 1 禁用，2 未禁用 默认未禁用(2)',
                      is_review tinyint DEFAULT 2 COMMENT '是否审核 1 审核通过，2 未审核 默认未审核(2)',
                      PRIMARY KEY (id)
) COMMENT '帖子表';

-- 创建视频表
CREATE TABLE video (
                       id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                       create_time bigint DEFAULT NULL COMMENT '创建时间',
                       update_time bigint DEFAULT NULL COMMENT '更新时间',
                       category_id bigint DEFAULT NULL COMMENT '栏目ID',
                       video_name varchar(100) NOT NULL COMMENT '视频名称',
                       simple_description varchar(200) DEFAULT NULL COMMENT '简易描述',
                       rich_text_descriptions text DEFAULT NULL COMMENT '富文本描述',
                       picture_address varchar(255) DEFAULT NULL COMMENT '图片地址',
                       number_of_views smallint DEFAULT 0 COMMENT '浏览量',
                       number_of_likes smallint DEFAULT 0 COMMENT '点赞量',
                       number_of_favorites smallint DEFAULT 0 COMMENT '收藏量',
                       number_of_comments smallint DEFAULT 0 COMMENT '评论量',
                       name_of_the_video_group varchar(50) DEFAULT NULL COMMENT '视频分组名称',
                       tencent_cloud_id varchar(50) DEFAULT NULL COMMENT '腾讯云ID',
                       tencent_cloud_video_id varchar(50) DEFAULT NULL COMMENT '腾讯云视频ID',
                       tencent_cloud_video_address varchar(255) DEFAULT NULL COMMENT '腾讯云视频地址',
                       video_permissions tinyint DEFAULT 2 COMMENT '视频权限 1 需要会员，2 公开 默认公开(2)',
                       is_disabled tinyint DEFAULT 2 COMMENT '是否禁用 1 禁用，2 未禁用 默认未禁用(2)',
                       PRIMARY KEY (id)
) COMMENT '视频表';

-- 创建视频浏览记录表
CREATE TABLE video_browsing_history (
                                        id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                                        create_time bigint DEFAULT NULL COMMENT '创建时间',
                                        update_time bigint DEFAULT NULL COMMENT '更新时间',
                                        user_id bigint DEFAULT NULL COMMENT '用户ID',
                                        video_id bigint DEFAULT NULL COMMENT '视频ID',
                                        watch_time bigint DEFAULT NULL COMMENT '观看时长',
                                        PRIMARY KEY (id)
) COMMENT '视频浏览记录表';

-- 创建积分变化表
CREATE TABLE table_of_points_changes (
                                         id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                                         create_time bigint DEFAULT NULL COMMENT '创建时间',
                                         update_time bigint DEFAULT NULL COMMENT '更新时间',
                                         raw_points int DEFAULT NULL COMMENT '原始积分',
                                         this_points int DEFAULT NULL COMMENT '当前积分',
                                         after_points int DEFAULT NULL COMMENT '消费后积分',
                                         user_id bigint DEFAULT NULL COMMENT '用户ID',
                                         type tinyint DEFAULT NULL COMMENT '类型 1 下载资源 2 资源奖励，3 视频资源，4 充值资源，5邀请用户赠送，6 其他',
                                         PRIMARY KEY (id)
) COMMENT '积分变化表';

-- 创建支付记录表
CREATE TABLE payment_history (
                                 id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                                 create_time bigint DEFAULT NULL COMMENT '创建时间',
                                 update_time bigint DEFAULT NULL COMMENT '更新时间',
                                 user_id bigint DEFAULT NULL COMMENT '用户ID',
                                 the_amount_to_be_paid int DEFAULT NULL COMMENT '支付金额',
                                 order_number bigint DEFAULT NULL COMMENT '订单号',
                                 is_successful tinyint DEFAULT 2 COMMENT '是否支付成功 1 没有成功，2 成功',
                                 successful_payment_callback varchar(255) DEFAULT NULL COMMENT '支付成功回调',
                                 PRIMARY KEY (id)
) COMMENT '支付记录表';

-- 创建评论表
CREATE TABLE comments (
                          id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                          create_time bigint DEFAULT NULL COMMENT '创建时间',
                          update_time bigint DEFAULT NULL COMMENT '更新时间',
                          user_id bigint DEFAULT NULL COMMENT '用户ID',
                          parent_id bigint DEFAULT NULL COMMENT '父级ID (资源ID/帖子ID/视频ID)',
                          comments_type tinyint DEFAULT NULL COMMENT '评论类型 1 资源评论，2 帖子评论，3 视频评论',
                          comment_content text DEFAULT NULL COMMENT '评论内容',
                          PRIMARY KEY (id)
) COMMENT '评论表';

-- 创建举报表
CREATE TABLE report (
                        id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                        create_time bigint DEFAULT NULL COMMENT '创建时间',
                        update_time bigint DEFAULT NULL COMMENT '更新时间',
                        user_id bigint DEFAULT NULL COMMENT '用户ID',
                        report_type tinyint DEFAULT NULL COMMENT '举报类型 1 资源举报，2 帖子举报，3 视频举报，4 用户，5 评论',
                        be_reported_id bigint DEFAULT NULL COMMENT '被举报的ID',
                        reason_for_reporting text DEFAULT NULL COMMENT '举报原因',
                        PRIMARY KEY (id)
) COMMENT '举报表';

-- 创建私信表
CREATE TABLE private_messages (
                                  id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                                  create_time bigint DEFAULT NULL COMMENT '创建时间',
                                  update_time bigint DEFAULT NULL COMMENT '更新时间',
                                  sender_id bigint DEFAULT NULL COMMENT '发送者ID (当前账号的用户ID)',
                                  receiver_id bigint DEFAULT NULL COMMENT '接收者ID (对方账号的用户ID)',
                                  content text DEFAULT NULL COMMENT '内容',
                                  PRIMARY KEY (id)
) COMMENT '私信表';

-- 创建点赞表
CREATE TABLE thumbs_up (
                           id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                           create_time bigint DEFAULT NULL COMMENT '创建时间',
                           update_time bigint DEFAULT NULL COMMENT '更新时间',
                           user_id bigint DEFAULT NULL COMMENT '用户ID',
                           thumbs_up_type tinyint DEFAULT NULL COMMENT '点赞类型 1 资源，2 帖子，3 视频',
                           liked_id bigint DEFAULT NULL COMMENT '被点赞的ID (资源ID/帖子ID/视频ID)',
                           PRIMARY KEY (id)
) COMMENT '点赞表';

-- 创建收藏表
CREATE TABLE collect (
                         id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                         create_time bigint DEFAULT NULL COMMENT '创建时间',
                         update_time bigint DEFAULT NULL COMMENT '更新时间',
                         user_id bigint DEFAULT NULL COMMENT '用户ID',
                         collect_type tinyint DEFAULT NULL COMMENT '收藏类型 1 资源，2 帖子，3 视频',
                         favorite_id bigint DEFAULT NULL COMMENT '被收藏的ID (资源ID/帖子ID/视频ID)',
                         PRIMARY KEY (id)
) COMMENT '收藏表';

-- 创建友链表
CREATE TABLE friend_chains (
                               id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                               create_time bigint DEFAULT NULL COMMENT '创建时间',
                               update_time bigint DEFAULT NULL COMMENT '更新时间',
                               chains_name varchar(50) NOT NULL COMMENT '友链名称',
                               chains_url varchar(255) NOT NULL COMMENT '友链地址',
                               is_disabled tinyint DEFAULT 2 COMMENT '是否禁用 1 禁用，2 未禁用 默认未禁用(2)',
                               PRIMARY KEY (id)
) COMMENT '友链表';

-- 创建广告表
CREATE TABLE advert (
                        id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                        create_time bigint DEFAULT NULL COMMENT '创建时间',
                        update_time bigint DEFAULT NULL COMMENT '更新时间',
                        advert_name varchar(50) NOT NULL COMMENT '广告名称',
                        advert_url varchar(255) NOT NULL COMMENT '广告地址',
                        advert_avatar varchar(255) DEFAULT NULL COMMENT '广告图片',
                        is_disabled tinyint DEFAULT 2 COMMENT '是否禁用 1 禁用，2 未禁用 默认未禁用(2)',
                        PRIMARY KEY (id)
) COMMENT '广告表';

-- 创建网站信息表
CREATE TABLE website_info (
                              id bigint AUTO_INCREMENT NOT NULL COMMENT 'ID',
                              create_time bigint DEFAULT NULL COMMENT '创建时间',
                              update_time bigint DEFAULT NULL COMMENT '更新时间',
                              the_website_icp_filing_number varchar(50) NOT NULL COMMENT '网站备案号',
                              the_domain_name_of_the_website varchar(100) NOT NULL COMMENT '网站域名',
                              the_name_of_the_website varchar(50) NOT NULL COMMENT '网站名称',
                              login_mode tinyint DEFAULT NULL COMMENT '登录模式 1 只能同时登录一个账号，2 可以同时多点登录',
                              alipay_pid varchar(50) DEFAULT NULL COMMENT '支付宝开放平台账号',
                              alipay_rsa_key varchar(255) DEFAULT NULL COMMENT '支付宝RSA私钥',
                              tencent_cloud_video_key varchar(50) DEFAULT NULL COMMENT '腾讯云视频密钥',
                              PRIMARY KEY (id)
) COMMENT '网站信息表';