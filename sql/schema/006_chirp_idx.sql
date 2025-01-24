-- +goose Up
CREATE INDEX idx_chirps_created_at ON chirps(created_at);
CREATE INDEX idx_chirps_user_id ON chirps(user_id);
-- Note: No need to index 'id' as PostgreSQL automatically indexes primary keys

-- +goose Down
DROP INDEX idx_chirps_created_at;
DROP INDEX idx_chirps_user_id;