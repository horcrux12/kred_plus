-- +migrate Up
INSERT INTO `user` (username, password, is_admin, created_by, created_at)
VALUES ('superadmin', '$2a$04$mVVUGfk3n86YIgAI9hxxGeB/MNBuIyW05PutPhK8zFxtATpM2n0Na', true, 0, now());
-- +migrate Down