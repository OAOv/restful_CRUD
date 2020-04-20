CREATE TABLE `user` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`age` INT(3) NOT NULL,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=7
;


CREATE TABLE `record` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`user_id` INT(11) NOT NULL,
	`user_name` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`subject` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`score` INT(11) NOT NULL,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=5
;
