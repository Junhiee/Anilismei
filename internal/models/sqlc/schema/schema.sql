/* 动画信息表 */
CREATE TABLE `animations` (
	-- 主键
	`anime_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	-- 动画类型表id
	`genre_id` INT NOT NULL,
	-- 制作公司表id
	`studio_id` INT NOT NULL,
	-- 标题
	`title` VARCHAR(255) NOT NULL,
	`country` VARCHAR(255) NOT NULL,
	-- 图片
	`image_url` VARCHAR(255),
	-- 简介
	`evaluate` TEXT(65535),
	-- 更新日期
	`update_time` TIMESTAMP,
	-- 发布日期
	`release_date` TIMESTAMP,
	-- 动画类型
	`anime_status` ENUM("coming soon", "airing", "completed", "paused"),
	-- 评分
	`rating` FLOAT,
	PRIMARY KEY(`anime_id`)
);


CREATE INDEX `anime_id_index`
ON `animations` (`anime_id`);
/* 动画类型表 */
CREATE TABLE `genres` (
	`genre_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	-- 类型名称
	`genre_name` VARCHAR(255) NOT NULL,
	PRIMARY KEY(`genre_id`)
);


CREATE INDEX `genre_id_index`
ON `genres` (`genre_id`);

/* 动画信息-动画类型关联表*/
CREATE TABLE animation_genres (
    `anime_id` BIGINT,
    `genre_id` INT,
    PRIMARY KEY (`anime_id`, `genre_id`),
    FOREIGN KEY (`anime_id`) REFERENCES animations(`anime_id`),
    FOREIGN KEY (`genre_id`) REFERENCES genres(`genre_id`)
);

/* 制作公司 */
CREATE TABLE `studios` (
	`studio_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	`studio_name` VARCHAR(255) NOT NULL,
	`studio_staff` VARCHAR(255) NOT NULL,
	PRIMARY KEY(`studio_id`)
);


CREATE INDEX `studio_id_index`
ON `studios` (`studio_id`);

/* 用户表 */
CREATE TABLE `users` (
	`user_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`user_name` VARCHAR(255) NOT NULL UNIQUE,
	`email` VARCHAR(255) NOT NULL,
	`user_pwd` VARCHAR(255) NOT NULL,
	`avatar_url` VARCHAR(255),
	PRIMARY KEY(`user_id`)
);


CREATE INDEX `user_id_index`
ON `users` (`user_id`);
/* 用户订阅表 */
CREATE TABLE `subscriptions` (
	`subscription_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	`user_id` BIGINT NOT NULL,
	`anime_id` BIGINT NOT NULL,
	`subscription_date` TIMESTAMP NOT NULL,
	PRIMARY KEY(`subscription_id`)
);


CREATE INDEX `subscription_id_index`
ON `subscriptions` (`subscription_id`);
/* 用户评论表 */
CREATE TABLE `comment` (
	`comment_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`anime_id` BIGINT NOT NULL,
	`comment_text` TEXT(65535) NOT NULL,
	`comment_date` TIMESTAMP NOT NULL,
	PRIMARY KEY(`comment_id`)
);


CREATE INDEX `comment_id_index`
ON `comment` (`comment_id`);
