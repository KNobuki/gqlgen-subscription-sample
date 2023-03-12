DROP TABLE IF EXISTS `smart_mat`;

CREATE TABLE `smart_mat` (
    `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
    `current_weight` DOUBLE NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);
