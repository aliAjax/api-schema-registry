CREATE TABLE namespaces (id text primary key, name text not null, owner text, created_at timestamptz not null);
CREATE TABLE assets (id text primary key, namespace_id text not null, name text not null, kind text not null, created_at timestamptz not null);
CREATE TABLE asset_versions (asset_id text not null, version text not null, status text not null, document jsonb not null, checksum text not null, primary key(asset_id,version));
