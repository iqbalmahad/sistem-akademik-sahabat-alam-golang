CREATE TABLE students
(
    id INT
    AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    class_id INT NULL,
    name VARCHAR
    (255) NOT NULL,
    nis VARCHAR
    (255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON
    UPDATE CURRENT_TIMESTAMP,
deleted_at TIMESTAMP
    NULL,
    FOREIGN KEY
    (user_id) REFERENCES users
    (id),
    FOREIGN KEY
    (class_id) REFERENCES classes
    (id)
)ENGINE = InnoDB;