CREATE TABLE IF NOT EXISTS achievements (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  description TEXT
);

CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(256) PRIMARY KEY,
  name TEXT NOT NULL,
  avatar_url TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_achievements (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id VARCHAR(256) NOT NULL REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  achievement_id UUID NOT NULL REFERENCES achievements(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS chills (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  ended_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS user_chills (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id VARCHAR(256) NOT NULL REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  chill_id UUID NOT NULL REFERENCES chills(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS trace_points (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  chill_id UUID NOT NULL REFERENCES chills(id) ON UPDATE CASCADE ON DELETE CASCADE,
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS photos (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  chill_id UUID NOT NULL REFERENCES chills(id) ON UPDATE CASCADE ON DELETE CASCADE,
  timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  url TEXT NOT NULL
);