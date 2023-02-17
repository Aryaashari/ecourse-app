CREATE TABLE admin (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR,
    email VARCHAR,
    password VARCHAR,
    PRIMARY KEY (id)
) ENGINE=InnoDB;