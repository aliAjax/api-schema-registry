CREATE TABLE consumers (id text primary key, name text not null, asset_id text not null, version text not null, environment text not null, strategy text not null);
CREATE TABLE changes (sequence bigserial primary key, asset_id text not null, version text not null, event_type text not null, payload jsonb not null, created_at timestamptz not null);
