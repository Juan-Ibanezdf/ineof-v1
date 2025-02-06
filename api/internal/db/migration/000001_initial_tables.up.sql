-- ===============================
-- Ativar a extensão pgcrypto para suportar gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS postgis;
-- ===============================

-- ===============================
--  Tabela de Campanhas
--  (Relacionamento: N:N com Equipamentos)
-- ===============================
CREATE TABLE campanhas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(255) NOT NULL,
    imagem TEXT,
    data_inicio TIMESTAMP,
    data_fim TIMESTAMP,
    equipe VARCHAR(255),
    localizacao GEOMETRY(Point, 4326), -- Coordenadas geográficas
    objetivos TEXT,
    contato VARCHAR(255),
    status VARCHAR(50) CHECK (status IN ('Em Andamento', 'Finalizada', 'Em Planejamento','Paralisada')),
    notas TEXT,
    descricao TEXT
);

-- ===============================
--  Tabela de Equipamentos
--  (Relacionamento: N:N com Campanhas)
-- ===============================
CREATE TABLE equipamentos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(255) NOT NULL,
    descricao TEXT,
    tipo VARCHAR(100),
    numero_serie VARCHAR(100) UNIQUE,
    modelo VARCHAR(100),
    fabricante VARCHAR(100),
    frequencia FLOAT,
    data_calibracao TIMESTAMP,
    data_ultima_manutencao TIMESTAMP,
    responsavel_manutencao VARCHAR(255),
    data_fabricacao TIMESTAMP,
    data_aquisicao TIMESTAMP,
    tipos_dados TEXT,
    notas TEXT,
    data_expiracao_garantia TIMESTAMP,
    status_operacional VARCHAR(50) CHECK (status_operacional IN ('Operacional', 'Manutenção', 'Desativado','Parado')),
    localizacao GEOMETRY(Point, 4326), -- Coordenadas do equipamento
    imagem TEXT
);

-- ===============================
--  Tabela de Relacionamento Campanha x Equipamento (N:N)
--  (Cada campanha pode ter vários equipamentos e vice-versa)
-- ===============================
CREATE TABLE campanha_equipamento (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    campanha_id UUID NOT NULL,
    equipamento_id UUID NOT NULL,
    CONSTRAINT unique_campanha_equipamento UNIQUE (campanha_id, equipamento_id)
);

-- ===============================
--  Tabela de Dados do LIDAR WindCube
--  (Relacionamento: 1:N com Campanhas e Equipamentos)
-- ===============================
CREATE TABLE lidar_windcube_dados (
    id SERIAL PRIMARY KEY,
    campanha_id UUID NOT NULL,
    equipamento_id UUID NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    posicao VARCHAR(50),
    temperatura FLOAT,
    wiper_count INT,
    -- Adicionando todas as alturas de medição do LIDAR
    cnr_40m FLOAT, radial_wind_speed_40m FLOAT, wind_speed_disp_40m FLOAT, wind_speed_40m FLOAT, wind_direction_40m FLOAT, x_wind_40m FLOAT, y_wind_40m FLOAT, z_wind_40m FLOAT,
    cnr_50m FLOAT, radial_wind_speed_50m FLOAT, wind_speed_disp_50m FLOAT, wind_speed_50m FLOAT, wind_direction_50m FLOAT, x_wind_50m FLOAT, y_wind_50m FLOAT, z_wind_50m FLOAT,
    cnr_60m FLOAT, radial_wind_speed_60m FLOAT, wind_speed_disp_60m FLOAT, wind_speed_60m FLOAT, wind_direction_60m FLOAT, x_wind_60m FLOAT, y_wind_60m FLOAT, z_wind_60m FLOAT,
    cnr_70m FLOAT, radial_wind_speed_70m FLOAT, wind_speed_disp_70m FLOAT, wind_speed_70m FLOAT, wind_direction_70m FLOAT, x_wind_70m FLOAT, y_wind_70m FLOAT, z_wind_70m FLOAT,
    cnr_80m FLOAT, radial_wind_speed_80m FLOAT, wind_speed_disp_80m FLOAT, wind_speed_80m FLOAT, wind_direction_80m FLOAT, x_wind_80m FLOAT, y_wind_80m FLOAT, z_wind_80m FLOAT,
    cnr_90m FLOAT, radial_wind_speed_90m FLOAT, wind_speed_disp_90m FLOAT, wind_speed_90m FLOAT, wind_direction_90m FLOAT, x_wind_90m FLOAT, y_wind_90m FLOAT, z_wind_90m FLOAT,
    cnr_100m FLOAT, radial_wind_speed_100m FLOAT, wind_speed_disp_100m FLOAT, wind_speed_100m FLOAT, wind_direction_100m FLOAT, x_wind_100m FLOAT, y_wind_100m FLOAT, z_wind_100m FLOAT,
    cnr_110m FLOAT, radial_wind_speed_110m FLOAT, wind_speed_disp_110m FLOAT, wind_speed_110m FLOAT, wind_direction_110m FLOAT, x_wind_110m FLOAT, y_wind_110m FLOAT, z_wind_110m FLOAT,
    cnr_120m FLOAT, radial_wind_speed_120m FLOAT, wind_speed_disp_120m FLOAT, wind_speed_120m FLOAT, wind_direction_120m FLOAT, x_wind_120m FLOAT, y_wind_120m FLOAT, z_wind_120m FLOAT,
    cnr_130m FLOAT, radial_wind_speed_130m FLOAT, wind_speed_disp_130m FLOAT, wind_speed_130m FLOAT, wind_direction_130m FLOAT, x_wind_130m FLOAT, y_wind_130m FLOAT, z_wind_130m FLOAT,
    cnr_140m FLOAT, radial_wind_speed_140m FLOAT, wind_speed_disp_140m FLOAT, wind_speed_140m FLOAT, wind_direction_140m FLOAT, x_wind_140m FLOAT, y_wind_140m FLOAT, z_wind_140m FLOAT,
    cnr_150m FLOAT, radial_wind_speed_150m FLOAT, wind_speed_disp_150m FLOAT, wind_speed_150m FLOAT, wind_direction_150m FLOAT, x_wind_150m FLOAT, y_wind_150m FLOAT, z_wind_150m FLOAT,
    cnr_160m FLOAT, radial_wind_speed_160m FLOAT, wind_speed_disp_160m FLOAT, wind_speed_160m FLOAT, wind_direction_160m FLOAT, x_wind_160m FLOAT, y_wind_160m FLOAT, z_wind_160m FLOAT,
    cnr_170m FLOAT, radial_wind_speed_170m FLOAT, wind_speed_disp_170m FLOAT, wind_speed_170m FLOAT, wind_direction_170m FLOAT, x_wind_170m FLOAT, y_wind_170m FLOAT, z_wind_170m FLOAT,
    cnr_180m FLOAT, radial_wind_speed_180m FLOAT, wind_speed_disp_180m FLOAT, wind_speed_180m FLOAT, wind_direction_180m FLOAT, x_wind_180m FLOAT, y_wind_180m FLOAT, z_wind_180m FLOAT,
    cnr_190m FLOAT, radial_wind_speed_190m FLOAT, wind_speed_disp_190m FLOAT, wind_speed_190m FLOAT, wind_direction_190m FLOAT, x_wind_190m FLOAT, y_wind_190m FLOAT, z_wind_190m FLOAT,
    cnr_200m FLOAT, radial_wind_speed_200m FLOAT, wind_speed_disp_200m FLOAT, wind_speed_200m FLOAT, wind_direction_200m FLOAT, x_wind_200m FLOAT, y_wind_200m FLOAT, z_wind_200m FLOAT,
    cnr_220m FLOAT, radial_wind_speed_220m FLOAT, wind_speed_disp_220m FLOAT, wind_speed_220m FLOAT, wind_direction_220m FLOAT, x_wind_220m FLOAT, y_wind_220m FLOAT, z_wind_220m FLOAT,
    cnr_250m FLOAT, radial_wind_speed_250m FLOAT, wind_speed_disp_250m FLOAT, wind_speed_250m FLOAT, wind_direction_250m FLOAT, x_wind_250m FLOAT, y_wind_250m FLOAT, z_wind_250m FLOAT,
    cnr_300m FLOAT, radial_wind_speed_300m FLOAT, wind_speed_disp_300m FLOAT, wind_speed_300m FLOAT, wind_direction_300m FLOAT, x_wind_300m FLOAT, y_wind_300m FLOAT, z_wind_300m FLOAT
);

-- ===============================
--  Adicionando Chaves Estrangeiras e Índices
-- ===============================
ALTER TABLE campanha_equipamento
ADD CONSTRAINT fk_campanha_equipamento_campanha FOREIGN KEY (campanha_id) REFERENCES campanhas(id) ON DELETE CASCADE;

ALTER TABLE campanha_equipamento
ADD CONSTRAINT fk_campanha_equipamento_equipamento FOREIGN KEY (equipamento_id) REFERENCES equipamentos(id) ON DELETE CASCADE;

ALTER TABLE lidar_windcube_dados
ADD CONSTRAINT fk_lidar_windcube_campanha FOREIGN KEY (campanha_id) REFERENCES campanhas(id) ON DELETE CASCADE;

ALTER TABLE lidar_windcube_dados
ADD CONSTRAINT fk_lidar_windcube_equipamento FOREIGN KEY (equipamento_id) REFERENCES equipamentos(id) ON DELETE CASCADE;

-- Índices para melhorar desempenho
CREATE INDEX idx_campanha_equipamento_campanha ON campanha_equipamento(campanha_id);
CREATE INDEX idx_campanha_equipamento_equipamento ON campanha_equipamento(equipamento_id);
CREATE INDEX idx_lidar_windcube_timestamp ON lidar_windcube_dados(timestamp);