CREATE TABLE courses (
    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    title VARCHAR,
    course_category_id BIGINT NOT NULL,
    FOREIGN KEY (course_category_id) REFERENCES course_categories(id)
) ENGINE=InnoDB;