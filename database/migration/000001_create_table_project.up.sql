CREATE TABLE db.mst_projects (
    id INT auto_increment NOT NULL,
	title varchar(100) NOT NULL,
	client varchar(100) NOT NULL,
	`date` DATE NOT NULL,
	date_completion DATE NULL,
	description TEXT NULL,
	`position` varchar(100) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	created_by varchar(100) DEFAULT 'SYSTEM' NOT NULL,
	updated_at TIMESTAMP NULL,
	updated_by varchar(100) NULL,
	deleted_at TIMESTAMP NULL,
	CONSTRAINT mst_projects_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE db.mst_medias (
	id int auto_increment NOT NULL,
    uuid varchar(100) NOT NULL,
	`path` varchar(100) NOT NULL,
	`size` INT NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP  NOT NULL,
	created_by varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'SYSTEM' NOT NULL,
	updated_at timestamp NULL,
	updated_by varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
	deleted_at timestamp NULL,
	CONSTRAINT mst_medias_pk PRIMARY KEY (id),
	CONSTRAINT mst_medias_unique UNIQUE KEY (uuid)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE db.mst_projects ADD media_uuid varchar(100) NULL;
ALTER TABLE db.mst_projects CHANGE media_uuid media_uuid varchar(100) NULL AFTER date_completion;
ALTER TABLE db.mst_projects ADD CONSTRAINT mst_projects_mst_medias_FK FOREIGN KEY (media_uuid) REFERENCES db.mst_medias(uuid);
