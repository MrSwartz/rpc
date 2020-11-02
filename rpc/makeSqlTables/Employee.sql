CREATE TABLE IF NOT EXISTS `data`.`Employee` (
  `id` BIGINT(20) NOT NULL,
  `name` VARCHAR(64) NULL,
  `secondName` VARCHAR(64) NULL,
  `surname` VARCHAR(64) NULL,
  `hireDate` DATE NULL,
  `position` ENUM('developer', 'manager', 'quality assurance', 'business analyst') NULL,
  `companyId` BIGINT(20) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

insert into data.Employee (id, name, secondName, surname, hireDate, position, companyId) values (1, 'Linus', 'Torvalds', 'Benedict', '2017-03-03', 'developer', 1);
insert into data.Employee (id, name, secondName, surname, hireDate, position, companyId) values (2, 'Bill', 'Gates', 'Henry', '2019-11-10', 'manager', 1134);
insert into data.Employee (id, name, secondName, surname, hireDate, position, companyId) values (2, 'Richard', 'Stallman', 'Matthew', '1980-03-02', 'developer', 22);
insert into data.Employee (id, name, secondName, surname, hireDate, position, companyId) values (2, 'Aaron', 'Hillel', 'Swartz', '1990-11-04', 'developer', 3);
insert into data.Employee (id, name, secondName, surname, hireDate, position, companyId) values (2, 'Seymour', 'Roger', 'Cray', '1970-01-01', 'manager', 5);