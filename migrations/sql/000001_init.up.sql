CREATE TABLE tag_groups
(
    group_id varchar(64) PRIMARY KEY NOT NULL UNIQUE,
    name_en  text DEFAULT NULL,
    name_ru  text DEFAULT NULL,
    name_uk  text DEFAULT NULL,
    name_be  text DEFAULT NULL
);

CREATE TABLE tags
(
    tag_id   varchar(64) PRIMARY KEY NOT NULL UNIQUE,
    emoji    text                                                           DEFAULT NULL,
    group_id varchar(64) REFERENCES tag_groups (group_id) ON DELETE CASCADE DEFAULT NULL,
    name_en  text                                                           DEFAULT NULL,
    name_ru  text                                                           DEFAULT NULL,
    name_uk  text                                                           DEFAULT NULL,
    name_be  text                                                           DEFAULT NULL
);
