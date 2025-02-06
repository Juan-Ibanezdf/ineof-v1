-- Criar uma nova campanha
-- name: CreateCampanha :one
INSERT INTO campanhas (id, nome, imagem, data_inicio, data_fim, equipe, localizacao, objetivos, contato, status, notas, descricao)
VALUES ($1, $2, $3, $4, $5, $6, ST_GeomFromText($7, 4326), $8, $9, $10, $11, $12)
RETURNING *;

-- Buscar uma campanha pelo ID
-- name: GetCampanhaByID :one
SELECT * FROM campanhas WHERE id = $1;

-- Listar todas as campanhas
-- name: ListCampanhas :many
SELECT * FROM campanhas ORDER BY data_inicio DESC;

-- Atualizar uma campanha existente
-- name: UpdateCampanha :one
UPDATE campanhas
SET 
    nome = $2,
    imagem = $3,
    data_inicio = $4,
    data_fim = $5,
    equipe = $6,
    localizacao = ST_GeomFromText($7, 4326),
    objetivos = $8,
    contato = $9,
    status = $10,
    notas = $11,
    descricao = $12
WHERE id = $1
RETURNING *;

-- Deletar uma campanha
-- name: DeleteCampanha :exec
DELETE FROM campanhas WHERE id = $1;