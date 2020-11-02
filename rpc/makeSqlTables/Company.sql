CREATE TABLE IF NOT EXISTS `data`.`Company` (
  `id` BIGINT(20) NOT NULL,
  `name` VARCHAR(64) NULL,
  `legalForm` VARCHAR(64) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

insert into data.Company (id, name, legalForm) values (1, 'Linus', 'lf1');
insert into data.Company (id, name, legalForm) values (2, 'Bill', 'lf2');
insert into data.Company (id, name, legalForm) values (3, 'Richard', 'lf3');
insert into data.Company (id, name, legalForm) values (4, 'Steve', 'lf3');
insert into data.Company (id, name, legalForm) values (5, 'Ivan', 'lf4');