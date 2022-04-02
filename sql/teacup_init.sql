DROP TABLE IF EXISTS `user_info`;
CREATE TABLE IF NOT EXISTS `user_info`  (
  `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'use id',
  `user_phone` varchar(20) NULL DEFAULT '' COMMENT 'user phone no',
  `passcode` varchar(100) NULL DEFAULT '' COMMENT 'password',
  `gender` TINYINT NOT NULL DEFAULT 1 COMMENT '用户性别 1 男 2女',
  `birthday` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  `nation` varchar(20) NOT NULL DEFAULT '' COMMENT '用户的国家',
  `user_phone_type` TINYINT NOT NULL DEFAULT 1 COMMENT '手机类型 1 iOS 2 android',
  `uuid` varchar(100) NOT NULL DEFAULT '' COMMENT '手机uuid',
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_avator` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_union` varchar(20) NOT NULL DEFAULT '' COMMENT  '用户注册所在区域 1中国 2美国',
  `user_image` varchar(256) NOT NULL default '' COMMENT '用户头像地址',
  `user_lvl` TINYINT NOT NULL DEFAULT 1 COMMENT '用户等级 1为茶客 2为茶神',
  `user_badges` varchar(256) NOT NULL DEFAULT '' COMMENT '用户徽章',
  `user_topics` varchar(256) NOT NULL DEFAULT '' COMMENT '用户感兴趣话题',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '用户更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户创建时间'
) ENGINE = INNODB COMMENT 'teacup 用户表';

DROP TABLE IF EXISTS `sms_log`;
CREATE TABLE IF NOT EXISTS `sms_log`  (
  `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'use id',
  `user_phone` varchar(20) NULL DEFAULT '' COMMENT 'user phone no',
  `nation` varchar(20) NULL DEFAULT '' COMMENT '用户的国家',
  `user_phone_type` TINYINT NOT NULL DEFAULT 1 COMMENT '手机类型 1 iOS 2 android',
  `type` TINYINT NOT NULL DEFAULT 1 COMMENT 'type=1 login type=2 register',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '用户更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户创建时间'
) ENGINE = INNODB COMMENT 'teacup 用户表';


DROP TABLE IF EXISTS `apply_to_vip`;
CREATE TABLE IF NOT EXISTS `apply_to_vip` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'apply id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `apply_comment` varchar(256) not NULL default ''  COMMENT '申请描述',
    `apply_status` TINYINT NOT NULL DEFAULT 1 COMMENT '申请状态 1提交 2通过 3否决',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '申请茶神表';

DROP TABLE IF EXISTS `fans`;
CREATE TABLE IF NOT EXISTS `fans` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'fans id',
    `user_id` BIGINT not NULL default 0  COMMENT '关注用户id',
    `attention_user_id` BIGINT not NULL default 0  COMMENT '被关注的用户id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '社区创建时间'
) ENGINE = INNODB COMMENT '粉丝表';

DROP TABLE IF EXISTS `community`;
CREATE TABLE IF NOT EXISTS `community` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'community id',
    `created_user_id` BIGINT not NULL default 0  COMMENT '创建社区用户id',
    `community_name` varchar(100) not NULL default ''  COMMENT '社区名称',
    `community_comment` varchar(256) not NULL default ''  COMMENT '社区描述',
    `community_tags` varchar(100) not NULL default ''  COMMENT '社区标签列表',
    `community_image` varchar(256) not NULL default ''  COMMENT '社区图标地址',
    `is_only_invite` TINYINT NOT NULL DEFAULT 1 COMMENT '是否只有邀请人加入 1不是 2是',
    `is_any_join` TINYINT NOT NULL DEFAULT 1 COMMENT '是否任何人都可以加入 1不可以 2可以',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '社区更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '社区创建时间'
) ENGINE = INNODB COMMENT '社区表';

DROP TABLE IF EXISTS `fans_community`;
CREATE TABLE IF NOT EXISTS `fans_community` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'fans_community id',
    `community_id` BIGINT not NULL default 0  COMMENT '社区id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '社区创建时间'
) ENGINE = INNODB COMMENT '关注社区表';

DROP TABLE IF EXISTS `member_community`;
CREATE TABLE IF NOT EXISTS `member_community` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'member_community id',
    `community_id` BIGINT not NULL default 0  COMMENT '社区id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '社区创建时间'
) ENGINE = INNODB COMMENT '社区成员表';

DROP TABLE IF EXISTS `invite_to_community`;
CREATE TABLE IF NOT EXISTS `invite_to_community` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'invite id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `community_id` BIGINT not NULL default 0  COMMENT '社区id',
    `invite_user_id` BIGINT not NULL default 0  COMMENT '邀请用户id',
    `invite_status` TINYINT NOT NULL DEFAULT 1 COMMENT '邀请状态 1提交 2通过 3否决',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '邀请进社区表';

DROP TABLE IF EXISTS `apply_to_community`;
CREATE TABLE IF NOT EXISTS `apply_to_community` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'apply community id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `community_id` BIGINT not NULL default 0  COMMENT '社区id',
    `apply_status` TINYINT NOT NULL DEFAULT 1 COMMENT '申请状态 1提交 2通过 3否决',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '申请进社区表';


DROP TABLE IF EXISTS `teacup`;
CREATE TABLE IF NOT EXISTS `teacup` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'teacup id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `community_id` BIGINT not NULL default 0  COMMENT '社区id',
    `teacup_name` varchar(256) not NULL default ''  COMMENT '茶室名称',
    `teacup_comment` varchar(256) not NULL default ''  COMMENT '茶室描述',
    `teacup_image` varchar(256) not NULL default ''  COMMENT '茶室图片',
    `teacup_tags` varchar(256) not NULL default ''  COMMENT '茶室标签列表',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室表';

DROP TABLE IF EXISTS `teacup_history`;
CREATE TABLE IF NOT EXISTS `teacup_history` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'teacup_history id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `community_id` BIGINT not NULL default 0  COMMENT '社区id',
    `teacup_start_at` timestamp not NULL default CURRENT_TIMESTAMP  COMMENT '茶室开始时间',
    `teacup_end_at` timestamp not NULL default CURRENT_TIMESTAMP  COMMENT '茶室结束时间',
    `speech_room` varchar(256) not NULL default ''  COMMENT '直播房间地址',
    `speech_channel` varchar(256) not NULL default ''  COMMENT '直播流的地址',
    `speech_status` TINYINT not NULL default 1  COMMENT '开播 1未开播 2 开播中 3 已开播',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室记录表';

DROP TABLE IF EXISTS `teacup_vip`;
CREATE TABLE IF NOT EXISTS `teacup_vip` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'teacup_vip id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT '用户id',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室嘉宾表';

DROP TABLE IF EXISTS `invite_to_teacup`;
CREATE TABLE IF NOT EXISTS `invite_to_teacup` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'teacup_visitor id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT '邀请id',
    `invited_user_id` BIGINT not NULL default 0  COMMENT '被邀请id',
    `invited_status` TINYINT NOT NULL DEFAULT 1 COMMENT '邀请状态 1提交 2通过 3否决',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室嘉宾邀请表';

DROP TABLE IF EXISTS `join_teacup`;
CREATE TABLE IF NOT EXISTS `join_teacup` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'join_teacup id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT 'user id',
    `join_status` TINYINT not NULL default 1  COMMENT '当前状态 1 在线 2 离开',
    `join_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入茶室时间',
    `leave_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '离开茶室时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室加入日志表';

DROP TABLE IF EXISTS `join_channel_teacup`;
CREATE TABLE IF NOT EXISTS `join_channel_teacup` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'join_channel_teacup id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT 'user id',
    `channel` varchar(256) NOT NULL DEFAULT ''  COMMENT '茶室演讲频道地址',
    `join_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入茶室时间',
    `leave_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '离开茶室时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE = INNODB COMMENT '茶室加入频道日志表';

DROP TABLE IF EXISTS `teacup_chat`;
CREATE TABLE IF NOT EXISTS `teacup_chat` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'join_teacup id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT '社区id',
    `chat_category` TINYINT NOT NULL DEFAULT 1 COMMENT '聊天类型 1 心情 2举手 3 发言语音',
    `chat_comment` varchar(256) NOT NULL DEFAULT '' COMMENT '聊天内容，语音地址',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室聊天记录表';

DROP TABLE IF EXISTS `teacup_mark`;
CREATE TABLE IF NOT EXISTS `teacup_mark` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'join_teacup id',
    `teacup_id` BIGINT not NULL default 0  COMMENT '茶室id',
    `user_id` BIGINT not NULL default 0  COMMENT '社区id',
    `mark_value` TINYINT NOT NULL DEFAULT 9 COMMENT '打分 0-9分',
    `mark_comment` varchar(256) NOT NULL DEFAULT '' COMMENT '评价内容',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '茶室评价表';

DROP TABLE IF EXISTS `tags`;
CREATE TABLE IF NOT EXISTS `tags` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'join_teacup id',
    `tag` varchar(50) not NULL default ''  COMMENT '标签名称',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '标签表';

DROP TABLE IF EXISTS `version`;
CREATE TABLE IF NOT EXISTS `version` (
    `id` BIGINT not NULL  AUTO_INCREMENT primary key COMMENT 'id',
    `version` varchar(10) not NULL default ''  COMMENT '当前版本',
    `phone_type` TINYINT NOT NULL DEFAULT 1 COMMENT '手机类型 1 iOS 2 android',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间'
) ENGINE = INNODB COMMENT '版本表';