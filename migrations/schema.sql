CREATE TABLE `anime` (
    `aid` int NOT NULL AUTO_INCREMENT COMMENT '动画 id',
    `image` varchar(255) DEFAULT NULL COMMENT '动画封面图',
    `title` varchar(255) DEFAULT NULL COMMENT '动画标题',
    `sorce` float DEFAULT NULL COMMENT '推荐指数',
    `desc` text COMMENT '简介',
    `tag` varchar(255) DEFAULT NULL COMMENT '标签',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`aid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;