CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_albums (
  "user_id" UUID NOT NULL,
  "album_id" UUID NOT NULL,
  "added_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, album_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE
);

CREATE INDEX idx_user_id_on_user_albums ON user_albums(user_id);
CREATE INDEX idx_album_id_on_user_albums ON user_albums(album_id);