
-- ----------------------------
-- Table structure for comment
-- ----------------------------
CREATE TABLE `comment` (
    `comment_id` bigint NOT NULL,
    `content` text NOT NULL,
    `post_id` bigint NOT NULL,
    `author_id` bigint NOT NULL,
    `parent_id` bigint DEFAULT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`comment_id`),
    KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;