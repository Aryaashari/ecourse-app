CREATE TABLE user_courses (
    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    users_id BIGINT NOT NULL,
    course_id BIGINT NOT NULL,
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB;