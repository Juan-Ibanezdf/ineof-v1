-- Criar um novo equipamento
-- name: CreateEquipamento :one
INSERT INTO equipamentos (id, nome, descricao, tipo, numero_serie, modelo, fabricante, frequencia, data_calibracao, data_ultima_manutencao, responsavel_manutencao, data_fabricacao, data_aquisicao, tipos_dados, notas, data_expiracao_garantia, status_operacional, localizacao, imagem)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, ST_GeomFromText($18, 4326), $19)
RETURNING *;

-- Buscar um equipamento pelo ID
-- name: GetEquipamentoByID :one
SELECT * FROM equipamentos WHERE id = $1;

-- Listar todos os equipamentos
-- name: ListEquipamentos :many
SELECT * FROM equipamentos ORDER BY nome ASC;

-- Atualizar um equipamento existente
-- name: UpdateEquipamento :one
UPDATE equipamentos
SET 
    nome = $2,
    descricao = $3,
    tipo = $4,
    numero_serie = $5,
    modelo = $6,
    fabricante = $7,
    frequencia = $8,
    data_calibracao = $9,
    data_ultima_manutencao = $10,
    responsavel_manutencao = $11,
    data_fabricacao = $12,
    data_aquisicao = $13,
    tipos_dados = $14,
    notas = $15,
    data_expiracao_garantia = $16,
    status_operacional = $17,
    localizacao = ST_GeomFromText($18, 4326),
    imagem = $19
WHERE id = $1
RETURNING *;

-- Deletar um equipamento
-- name: DeleteEquipamento :exec
DELETE FROM equipamentos WHERE id = $1;
