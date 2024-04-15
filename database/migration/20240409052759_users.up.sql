CREATE TABLE users
(
    id INT
    AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR
    (255) NOT NULL,
    username VARCHAR
    (255) NOT NULL UNIQUE,
    password VARCHAR
    (255) NOT NULL,
    role ENUM
    ('admin', 'teacher', 'student') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON
    UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    NULL
) ENGINE=InnoDB;