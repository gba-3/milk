use milk;

CREATE TABLE `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
    `name` varchar(255) NOT NULL COMMENT 'ユーザー名',
    `email` varchar(255) NOT NULL COMMENT 'メールアドレス',
    `password` varchar(255) NOT NULL COMMENT 'パスワード',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
    PRIMARY KEY(`id`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;