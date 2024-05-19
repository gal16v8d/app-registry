CREATE TABLE `app-registry`.repo (
	id BIGINT auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	main_tech varchar(100) NOT NULL,
	CONSTRAINT repo_unique_name UNIQUE KEY (name),
	CONSTRAINT repo_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;