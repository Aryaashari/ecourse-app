CREATE TABLE user_courses(
    id BIGINT NOT NULL AUTO_INCREMENT,
    users_id BIGINT NOt NULL,
    course_id BIGINT NOt NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(users_id) REFERENCES users(id),
    FOREIGN KEY(course_id) REFERENCES courses(id)
)ENGINE=InnoDB;