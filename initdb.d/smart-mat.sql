CREATE TABLE IF NOT EXISTS `smart_mats` (
    `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
    `current_weight` DOUBLE NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);
