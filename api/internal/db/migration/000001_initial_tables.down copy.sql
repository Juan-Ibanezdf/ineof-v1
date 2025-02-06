
-- ===============================
-- Removendo as Constraints
-- ===============================

-- 🔹 Removendo chaves estrangeiras da tabela de relacionamento Campanha x Equipamento
ALTER TABLE campanha_equipamento DROP CONSTRAINT IF EXISTS fk_campanha_equipamento_campanha;
ALTER TABLE campanha_equipamento DROP CONSTRAINT IF EXISTS fk_campanha_equipamento_equipamento;

-- 🔹 Removendo chaves estrangeiras da tabela de dados do LIDAR
ALTER TABLE lidar_windcube_dados DROP CONSTRAINT IF EXISTS fk_lidar_windcube_campanha;
ALTER TABLE lidar_windcube_dados DROP CONSTRAINT IF EXISTS fk_lidar_windcube_equipamento;

-- ===============================
-- Excluindo as Tabelas na Ordem Correta
-- ===============================

DROP TABLE IF EXISTS lidar_windcube_dados;
DROP TABLE IF EXISTS campanha_equipamento;
DROP TABLE IF EXISTS equipamentos;
DROP TABLE IF EXISTS campanhas;