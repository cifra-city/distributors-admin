-- Удаление таблиц с зависимостями
DROP TABLE IF EXISTS distributors CASCADE;
DROP TABLE IF EXISTS distributors_places CASCADE;

-- Удаление типа ENUM
DROP TYPE IF EXISTS roles CASCADE;

-- Удаление индекса
DROP INDEX IF EXISTS unique_owner_per_distributor;
