DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
     `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `username` char(20) NOT NULL DEFAULT '' COMMENT '用户名',
     `password` varchar(80) NOT NULL DEFAULT '' COMMENT '登录密码',
     `salt` char(16) DEFAULT '' COMMENT '随机码',
     `is_activited` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否激活：0=否，1=是',
     `activated_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '激活时间',
     `admin_type` tinyint(4) NOT NULL COMMENT '1管理员 2超级管理员',
     `permissions` text COMMENT '权限，json存储',
     `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
     `update_time` int(10) NOT NULL DEFAULT '0' COMMENT '最后更新时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT = 2 DEFAULT  CHARSET=utf8mb4 COMMENT='管理员';

INSERT INTO `admin` (`id`, `username`, `password`, `salt`, `is_activited`, `activated_time`, `admin_type`, `permissions`, `create_time`, `update_time`) VALUES
    (1, 'admin', 'eb5e043ef52340c6966bba42c2792513', 'BpLnfgDsc2WD8F2q', 1, 1556101521, 0, NULL, 1533279764, 1625078889);

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
        `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '微信openid',
        `openid_xchx` varchar(32) NOT NULL DEFAULT '' COMMENT '小程序openid',
        `unionid` varchar(32) NOT NULL DEFAULT '' COMMENT '微信unionId',
        `session_key` varchar(32) NOT NULL DEFAULT '' COMMENT 'session_key',
        `subscribe` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否关注：0=否，1=是',
        `nickname` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '微信昵称',
        `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户的性别，1=男性，2=女性，0=未知',
        `city` varchar(20) NOT NULL DEFAULT '' COMMENT '用户所在城市',
        `country` varchar(20) NOT NULL DEFAULT '' COMMENT '用户所在国家',
        `province` varchar(20) NOT NULL DEFAULT '' COMMENT '用户所在省份',
        `language` varchar(20) NOT NULL DEFAULT '' COMMENT '用户的语言，简体中文为zh_CN',
        `headimgurl` varchar(256) NOT NULL DEFAULT '' COMMENT '微信头像',
        `phone_number` varchar(16) NOT NULL COMMENT '手机号，带区号',
        `pure_phone_number` varchar(16) NOT NULL COMMENT '手机号，没区号',
        `country_code` varchar(8) NOT NULL COMMENT '国家码',
        `subscribe_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户关注时间,如果用户曾多次关注，则取最后关注时间',
        `subscribe_scene` varchar(64) NOT NULL DEFAULT '' COMMENT '返回用户关注的渠道来源',
        `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
        `is_banned` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户是否被禁用，0否，1是',
        `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
        `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后更新时间',
        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET=utf8 COMMENT='用户微信信息';

DROP TABLE IF EXISTS `help`;
CREATE TABLE `help`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title`       varchar(256) default '' not null comment '标题',
    `h5_link`     varchar(256) default '' not null comment 'H5链接地址',
    `sort`  int(10) DEFAULT '0' COMMENT '排序 大的在前',
    `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ='帮助文档';