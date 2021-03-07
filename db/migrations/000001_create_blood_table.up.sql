CREATE TABLE `bloods` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `device_id` VARCHAR (255) NOT NULL COMMENT '設備編號',
    `pulse` INT NOT NULL COMMENT '脈搏',
    `diastolic` FLOAT NOT NULL COMMENT '舒張壓',
    `systolic` FLOAT NOT NULL COMMENT '收縮壓',
    `created_at` timestamp NULL COMMENT '創建時間',
    `updated_at` timestamp null comment '更新時間',

    INDEX (`device_id`)
)