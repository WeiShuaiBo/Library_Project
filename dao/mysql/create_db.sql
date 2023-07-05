create database library;
use library;
DROP TABLE IF EXISTS `user`;

CREATE TABLE `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `username` varchar(255) NOT NULL  COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `email` varchar(255) NOT NULL  COMMENT '邮箱',
    `phone` int(40)  NOT NULL  COMMENT '手机号',
    `age`  int(20) NOT NULL COMMENT '年龄',
    `sex`  varchar(20) NOT NULL COMMENT '性别',
    `identity` varchar(20) DEFAULT '普通用户' COMMENT '身份',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_username`(`username`) USING BTREE,
    UNIQUE KEY `id_password`(`password`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `book`;
CREATE TABLE `book`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `book_name` varchar(255) NOT NULL COMMENT '图书名字',
    `book_author` varchar(255) NOT NULL COMMENT '图书作者',
    `book_number` int(20) NOT NULL COMMENT '图书数量',
    `book_kind` varchar(255) NOT NULL COMMENT '图书种类',
    `book_brief` varchar(255) NOT NULL COMMENT  '图书简介',
     PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;


DROP TABLE IF EXISTS `user_book`;
CREATE TABLE `user_book`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户id',
    `book_id` bigint(20) NOT NULL COMMENT '图书id',
    `is_return` varchar(20) NOT NULL COMMENT '是否归还',
    `start_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '借阅时间',
    `end_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '归还时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_user` (`user_id`) USING BTREE,
    UNIQUE KEY `id_book` (`book_id`) USING BTREE
) ENGINE = InnoDB DEFAULT  CHARSET = utf8mb4 COLLATE =utf8mb4_general_ci;