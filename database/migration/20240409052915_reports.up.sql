CREATE TABLE reports
(
    id INT
    AUTO_INCREMENT PRIMARY KEY,
    student_id INT NOT NULL,
    tingkatan_kelas VARCHAR
    (255) NOT NULL,
    semester VARCHAR
    (255) NOT NULL,
    rapot VARCHAR
    (255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON
    UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY
    (student_id) REFERENCES students
    (id)
) ENGINE=InnoDB;
