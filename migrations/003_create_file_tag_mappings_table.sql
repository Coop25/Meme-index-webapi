CREATE TABLE IF NOT EXISTS file_tag_mappings (
    file_id UUID NOT NULL,
    tag_id UUID NOT NULL,
    PRIMARY KEY (file_id, tag_id),
    FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);