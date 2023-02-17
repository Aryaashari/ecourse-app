CREATE TABLE courses(
    id BIGINT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    course_category_id BIGINT NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(course_category_id) REFERENCES course_categories(id)
)ENGINE=InnoDB;