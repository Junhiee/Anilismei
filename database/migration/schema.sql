/* 动画信息表 */
CREATE TABLE `animations` (
	-- 主键
	`anime_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	-- 标题
	`title` VARCHAR(255) NOT NULL,
	-- 简介
	`evaluate` TEXT(65535) NOT NULL,
	-- 外键-动画类型表
	`genre_id` INT NOT NULL,
	-- 更新日期
	`release_date` DATE,
	`studio_id` INT NOT NULL,
	`anime_status` ENUM("airing", "completed", "paused"),
	-- 评分
	`rating` DECIMAL,
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
/* 制作公司 */
CREATE TABLE `studios` (
	`studio_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	`studio_name` VARCHAR(255) NOT NULL,
	`studio_staff` VARCHAR(255) NOT NULL,
	PRIMARY KEY(`studio_id`)
);


CREATE INDEX `studio_id_index`
ON `studios` (`studio_id`);
/* 动画信息图片表 */
CREATE TABLE `images` (
	`image_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`anime_id` BIGINT NOT NULL,
	`image_url` VARCHAR(255) NOT NULL,
	PRIMARY KEY(`image_id`)
);


CREATE INDEX `anime_id_index`
ON `images` (`anime_id`);
/* 用户表 */
CREATE TABLE `users` (
	`user_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`user_name` VARCHAR(255) NOT NULL UNIQUE,
	`email` VARCHAR(255) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
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
	`subscription_date` DATE NOT NULL,
	PRIMARY KEY(`subscription_id`)
);


CREATE INDEX `subscription_id_index`
ON `subscriptions` (`subscription_id`);
/* 用户评论表 */
CREATE TABLE `comment` (
	`comment_id` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`anime_id` BIGINT NOT NULL,
	`comment_text` TEXT(65535) NOT NULL,
	`comment_date` DATE NOT NULL,
	PRIMARY KEY(`comment_id`)
);


CREATE INDEX `comment_id_index`
ON `comment` (`comment_id`);
