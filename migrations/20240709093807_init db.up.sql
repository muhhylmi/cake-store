CREATE TABLE IF NOT EXISTS cakes (
    id int NOT NULL AUTO_INCREMENT,
    title VARCHAR(150) NOT NULL,
    description VARCHAR(255) NOT NULL,
    rating FLOAT(4),
    image VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    primary key (id)
)