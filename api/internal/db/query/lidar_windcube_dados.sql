-- Inserir um novo dado do LIDAR WindCube
-- name: CreateLidarWindcubeDado :one
INSERT INTO lidar_windcube_dados (
    campanha_id, equipamento_id, timestamp, posicao, temperatura, wiper_count,
    cnr_40m, radial_wind_speed_40m, wind_speed_disp_40m, wind_speed_40m, wind_direction_40m, x_wind_40m, y_wind_40m, z_wind_40m,
    cnr_50m, radial_wind_speed_50m, wind_speed_disp_50m, wind_speed_50m, wind_direction_50m, x_wind_50m, y_wind_50m, z_wind_50m,
    cnr_60m, radial_wind_speed_60m, wind_speed_disp_60m, wind_speed_60m, wind_direction_60m, x_wind_60m, y_wind_60m, z_wind_60m,
    cnr_70m, radial_wind_speed_70m, wind_speed_disp_70m, wind_speed_70m, wind_direction_70m, x_wind_70m, y_wind_70m, z_wind_70m,
    cnr_80m, radial_wind_speed_80m, wind_speed_disp_80m, wind_speed_80m, wind_direction_80m, x_wind_80m, y_wind_80m, z_wind_80m,
    cnr_90m, radial_wind_speed_90m, wind_speed_disp_90m, wind_speed_90m, wind_direction_90m, x_wind_90m, y_wind_90m, z_wind_90m,
    cnr_100m, radial_wind_speed_100m, wind_speed_disp_100m, wind_speed_100m, wind_direction_100m, x_wind_100m, y_wind_100m, z_wind_100m,
    cnr_110m, radial_wind_speed_110m, wind_speed_disp_110m, wind_speed_110m, wind_direction_110m, x_wind_110m, y_wind_110m, z_wind_110m,
    cnr_120m, radial_wind_speed_120m, wind_speed_disp_120m, wind_speed_120m, wind_direction_120m, x_wind_120m, y_wind_120m, z_wind_120m,
    cnr_130m, radial_wind_speed_130m, wind_speed_disp_130m, wind_speed_130m, wind_direction_130m, x_wind_130m, y_wind_130m, z_wind_130m,
    cnr_140m, radial_wind_speed_140m, wind_speed_disp_140m, wind_speed_140m, wind_direction_140m, x_wind_140m, y_wind_140m, z_wind_140m,
    cnr_150m, radial_wind_speed_150m, wind_speed_disp_150m, wind_speed_150m, wind_direction_150m, x_wind_150m, y_wind_150m, z_wind_150m,
    cnr_160m, radial_wind_speed_160m, wind_speed_disp_160m, wind_speed_160m, wind_direction_160m, x_wind_160m, y_wind_160m, z_wind_160m,
    cnr_170m, radial_wind_speed_170m, wind_speed_disp_170m, wind_speed_170m, wind_direction_170m, x_wind_170m, y_wind_170m, z_wind_170m,
    cnr_180m, radial_wind_speed_180m, wind_speed_disp_180m, wind_speed_180m, wind_direction_180m, x_wind_180m, y_wind_180m, z_wind_180m,
    cnr_190m, radial_wind_speed_190m, wind_speed_disp_190m, wind_speed_190m, wind_direction_190m, x_wind_190m, y_wind_190m, z_wind_190m,
    cnr_200m, radial_wind_speed_200m, wind_speed_disp_200m, wind_speed_200m, wind_direction_200m, x_wind_200m, y_wind_200m, z_wind_200m,
    cnr_220m, radial_wind_speed_220m, wind_speed_disp_220m, wind_speed_220m, wind_direction_220m, x_wind_220m, y_wind_220m, z_wind_220m,
    cnr_250m, radial_wind_speed_250m, wind_speed_disp_250m, wind_speed_250m, wind_direction_250m, x_wind_250m, y_wind_250m, z_wind_250m,
    cnr_300m, radial_wind_speed_300m, wind_speed_disp_300m, wind_speed_300m, wind_direction_300m, x_wind_300m, y_wind_300m, z_wind_300m
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11, $12, $13,
    $14, $15, $16, $17, $18, $19, $20,
    $21, $22, $23, $24, $25, $26, $27,
    $28, $29, $30, $31, $32, $33, $34,
    $35, $36, $37, $38, $39, $40, $41,
    $42, $43, $44, $45, $46, $47, $48,
    $49, $50, $51, $52, $53, $54, $55,
    $56, $57, $58, $59, $60, $61, $62,
    $63, $64, $65, $66, $67, $68, $69,
    $70, $71, $72, $73, $74, $75, $76,
    $77, $78, $79, $80, $81, $82, $83,
    $84, $85, $86, $87, $88, $89, $90,
    $91, $92, $93, $94, $95, $96, $97,
    $98, $99, $100, $101, $102, $103, $104,
    $105, $106, $107, $108, $109, $110, $111,
    $112, $113, $114, $115, $116, $117, $118,
    $119, $120, $121, $122, $123, $124, $125,
    $126, $127, $128, $129, $130, $131, $132,
    $133, $134, $135, $136, $137, $138, $139,
    $140, $141, $142, $143, $144, $145, $146,
    $147, $148, $149, $150, $151, $152, $153,
    $154, $155, $156, $157, $158, $159, $160,
    $161, $162, $163, $164, $165, $166
) RETURNING *;

-- Buscar um dado do LIDAR pelo ID
-- name: GetLidarWindcubeDadoByID :one
SELECT * FROM lidar_windcube_dados WHERE id = $1;

-- Listar dados do LIDAR por campanha
-- name: ListLidarWindcubeDadosByCampanha :many
SELECT * FROM lidar_windcube_dados WHERE campanha_id = $1 ORDER BY timestamp DESC;

-- name: ListLidarWindcubeDadosCustom :many
SELECT 
    id,
    campanha_id,
    equipamento_id,
    timestamp,
    CASE WHEN sqlc.arg(include_posicao)::boolean THEN posicao ELSE NULL END AS posicao,
    CASE WHEN sqlc.arg(include_temperatura)::boolean THEN temperatura ELSE NULL END AS temperatura,
    CASE WHEN sqlc.arg(include_wiper_count)::boolean THEN wiper_count ELSE NULL END AS wiper_count,
    CASE WHEN sqlc.arg(include_wind_speed_40m)::boolean THEN wind_speed_40m ELSE NULL END AS wind_speed_40m,
    CASE WHEN sqlc.arg(include_wind_direction_40m)::boolean THEN wind_direction_40m ELSE NULL END AS wind_direction_40m,
    CASE WHEN sqlc.arg(include_wind_speed_50m)::boolean THEN wind_speed_50m ELSE NULL END AS wind_speed_50m,
    CASE WHEN sqlc.arg(include_wind_direction_50m)::boolean THEN wind_direction_50m ELSE NULL END AS wind_direction_50m,
    CASE WHEN sqlc.arg(include_wind_speed_60m)::boolean THEN wind_speed_60m ELSE NULL END AS wind_speed_60m,
    CASE WHEN sqlc.arg(include_wind_direction_60m)::boolean THEN wind_direction_60m ELSE NULL END AS wind_direction_60m,
    CASE WHEN sqlc.arg(include_wind_speed_70m)::boolean THEN wind_speed_70m ELSE NULL END AS wind_speed_70m,
    CASE WHEN sqlc.arg(include_wind_direction_70m)::boolean THEN wind_direction_70m ELSE NULL END AS wind_direction_70m,
    CASE WHEN sqlc.arg(include_wind_speed_80m)::boolean THEN wind_speed_80m ELSE NULL END AS wind_speed_80m,
    CASE WHEN sqlc.arg(include_wind_direction_80m)::boolean THEN wind_direction_80m ELSE NULL END AS wind_direction_80m,
    CASE WHEN sqlc.arg(include_wind_speed_90m)::boolean THEN wind_speed_90m ELSE NULL END AS wind_speed_90m,
    CASE WHEN sqlc.arg(include_wind_direction_90m)::boolean THEN wind_direction_90m ELSE NULL END AS wind_direction_90m,
    CASE WHEN sqlc.arg(include_wind_speed_100m)::boolean THEN wind_speed_100m ELSE NULL END AS wind_speed_100m,
    CASE WHEN sqlc.arg(include_wind_direction_100m)::boolean THEN wind_direction_100m ELSE NULL END AS wind_direction_100m,
    CASE WHEN sqlc.arg(include_wind_speed_110m)::boolean THEN wind_speed_110m ELSE NULL END AS wind_speed_110m,
    CASE WHEN sqlc.arg(include_wind_direction_110m)::boolean THEN wind_direction_110m ELSE NULL END AS wind_direction_110m,
    CASE WHEN sqlc.arg(include_wind_speed_120m)::boolean THEN wind_speed_120m ELSE NULL END AS wind_speed_120m,
    CASE WHEN sqlc.arg(include_wind_direction_120m)::boolean THEN wind_direction_120m ELSE NULL END AS wind_direction_120m,
    CASE WHEN sqlc.arg(include_wind_speed_130m)::boolean THEN wind_speed_130m ELSE NULL END AS wind_speed_130m,
    CASE WHEN sqlc.arg(include_wind_direction_130m)::boolean THEN wind_direction_130m ELSE NULL END AS wind_direction_130m,
    CASE WHEN sqlc.arg(include_wind_speed_140m)::boolean THEN wind_speed_140m ELSE NULL END AS wind_speed_140m,
    CASE WHEN sqlc.arg(include_wind_direction_140m)::boolean THEN wind_direction_140m ELSE NULL END AS wind_direction_140m,
    CASE WHEN sqlc.arg(include_wind_speed_150m)::boolean THEN wind_speed_150m ELSE NULL END AS wind_speed_150m,
    CASE WHEN sqlc.arg(include_wind_direction_150m)::boolean THEN wind_direction_150m ELSE NULL END AS wind_direction_150m,
    CASE WHEN sqlc.arg(include_wind_speed_160m)::boolean THEN wind_speed_160m ELSE NULL END AS wind_speed_160m,
    CASE WHEN sqlc.arg(include_wind_direction_160m)::boolean THEN wind_direction_160m ELSE NULL END AS wind_direction_160m,
    CASE WHEN sqlc.arg(include_wind_speed_170m)::boolean THEN wind_speed_170m ELSE NULL END AS wind_speed_170m,
    CASE WHEN sqlc.arg(include_wind_direction_170m)::boolean THEN wind_direction_170m ELSE NULL END AS wind_direction_170m,
    CASE WHEN sqlc.arg(include_wind_speed_180m)::boolean THEN wind_speed_180m ELSE NULL END AS wind_speed_180m,
    CASE WHEN sqlc.arg(include_wind_direction_180m)::boolean THEN wind_direction_180m ELSE NULL END AS wind_direction_180m,
    CASE WHEN sqlc.arg(include_wind_speed_190m)::boolean THEN wind_speed_190m ELSE NULL END AS wind_speed_190m,
    CASE WHEN sqlc.arg(include_wind_direction_190m)::boolean THEN wind_direction_190m ELSE NULL END AS wind_direction_190m,
    CASE WHEN sqlc.arg(include_wind_speed_200m)::boolean THEN wind_speed_200m ELSE NULL END AS wind_speed_200m,
    CASE WHEN sqlc.arg(include_wind_direction_200m)::boolean THEN wind_direction_200m ELSE NULL END AS wind_direction_200m,
    CASE WHEN sqlc.arg(include_wind_speed_220m)::boolean THEN wind_speed_220m ELSE NULL END AS wind_speed_220m,
    CASE WHEN sqlc.arg(include_wind_direction_220m)::boolean THEN wind_direction_220m ELSE NULL END AS wind_direction_220m,
    CASE WHEN sqlc.arg(include_wind_speed_250m)::boolean THEN wind_speed_250m ELSE NULL END AS wind_speed_250m,
    CASE WHEN sqlc.arg(include_wind_direction_250m)::boolean THEN wind_direction_250m ELSE NULL END AS wind_direction_250m,
    CASE WHEN sqlc.arg(include_wind_speed_300m)::boolean THEN wind_speed_300m ELSE NULL END AS wind_speed_300m,
    CASE WHEN sqlc.arg(include_wind_direction_300m)::boolean THEN wind_direction_300m ELSE NULL END AS wind_direction_300m  
FROM lidar_windcube_dados
WHERE campanha_id = sqlc.arg(campanha_id)
ORDER BY timestamp DESC;


-- Deletar um dado do LIDAR
-- name: DeleteLidarWindcubeDado :exec
DELETE FROM lidar_windcube_dados WHERE id = $1;