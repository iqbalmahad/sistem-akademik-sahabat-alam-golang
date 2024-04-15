    CREATE TABLE graduation_year_sds
(
    id INT
    AUTO_INCREMENT PRIMARY KEY,
    student_id INT NOT NULL UNIQUE,
    year VARCHAR
    (255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON
    UPDATE CURRENT_TIMESTAMP,
deleted_at TIMESTAMP
    NULL,
    FOREIGN KEY
    (student_id) REFERENCES students
    (id)
)ENGINE = InnoDB;