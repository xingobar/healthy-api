CREATE TABLE `weights` (
    `id` BIGINT PRIMARY  KEY AUTO_INCREMENT,
    `device_id` VARCHAR(255) NOT NULL COMMENT '設備編號',
    `number` FLOAT NOT NULL COMMENT '體重',
    `created_at` timestamp COMMENT '新增時間',
    `updated_at` timestamp COMMENT '更新時間',

    INDEX(`device_id`)
)