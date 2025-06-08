CREATE TABLE `shorturl` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `short_url` varchar(255) NOT NULL COMMENT '生成的key',
    `url` varchar(255) NOT NULL COMMENT '原始的url',
    PRIMARY KEY (`id`),
    KEY `short_url` (`short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;