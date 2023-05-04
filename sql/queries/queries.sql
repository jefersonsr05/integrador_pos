### EMPRESAS ###
-- name: GetEmpresa :one
SELECT * FROM empresas WHERE id = ?;

-- name: GetEmpresaByCNPJ :one
SELECT * FROM empresas WHERE cnpj = ?;

-- name: GetEmpresaByChaveRegistro :one
SELECT * FROM empresas WHERE chave_registro = ?;

-- name: ListEmpresas :many
SELECT * FROM empresas;

-- name: CreateEmpresa :exec
INSERT INTO empresas (id, descricao, cnpj, chave_registro) VALUES (?, ?, ?, ?);

-- name: UpdateEmpresa :exec
UPDATE empresas SET descricao=?, cnpj=?, chave_registro=? WHERE id=?;

-- name: DeleteEmpresa :exec
DELETE FROM empresas WHERE id = ?;


### CIDADES ###
-- name: GetCidade :one
SELECT * FROM cidades WHERE id = ?;

-- name: GetCidadeByIBGE :one
SELECT * FROM cidades WHERE cod_ibge = ?;

-- name: GetCidadeByIbgeUF :many
SELECT * FROM cidades LEFT JOIN estados ON cidades.estado_id = estados.id WHERE estados.cod_ibge = ?;

-- name: GetCidadeByUF :many
SELECT * FROM cidades LEFT JOIN estados ON cidades.estado_id = estados.id WHERE estados.uf = ?;

-- name: ListCidades :many
SELECT * FROM cidades;

-- name: CreateCidade :exec
INSERT INTO cidades (id, descricao, estado_id, cod_ibge) VALUES (?, ?, ?, ?);

-- name: UpdateCidade :exec
UPDATE cidades SET descricao=?, estado_id=?, cod_ibge=? WHERE id=?;

-- name: DeleteCidade :exec
DELETE FROM cidades WHERE id = ?;


### CLIENTES ###
-- name: GetCliente :one
SELECT * FROM clientes WHERE id = ?;

-- name: GetClienteByCodigoMC :one
SELECT * FROM clientes WHERE codigo_mc = ?;

-- name: GetClienteByEmpresa :many
SELECT * FROM clientes WHERE empresa_id = ?;

-- name: ListClientes :many
SELECT * FROM clientes;

-- name: CreateCliente :exec
INSERT INTO clientes (id, codigo_mc, empresa_id, nome, cep, cidade_id, endereco, numero, complemento) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCliente :exec
UPDATE clientes SET codigo_mc=?, empresa_id=?, nome=?, cep=?, cidade_id=?, endereco=?, numero=?, complemento=? WHERE id=?;

-- name: DeleteCliente :exec
DELETE FROM clientes WHERE id = ?;



### ESTADOS ###
-- name: GetEstado :one
SELECT * FROM estados WHERE id = ?;

-- name: GetEstadoByCodIBGE :one
SELECT * FROM estados WHERE cod_ibge = ?;

-- name: GetEstadoByUF :one
SELECT * FROM estados WHERE uf = ?;

-- name: ListEstados :many
SELECT * FROM estados;

-- name: CreateEstado :exec
INSERT INTO estados (id, descricao, uf, cod_ibge) VALUES (?, ?, ?, ?);

-- name: UpdateEstado :exec
UPDATE estados SET descricao=?, uf=?, cod_ibge=? WHERE id=?;

-- name: DeleteEstado :exec
DELETE FROM estados WHERE id = ?;



### PRODUTOS ###
-- name: GetProduto :one
SELECT * FROM produtos WHERE id = ?;

-- name: GetProdutoByCodigoMC :one
SELECT * FROM produtos WHERE codigo_mc = ?;

-- name: GetProdutoByCodBarras :one
SELECT * FROM produtos WHERE cod_barras = ?;

-- name: GetProdutoByEmpresa :many
SELECT * FROM produtos WHERE empresa_id = ?;

-- name: ListProdutos :many
SELECT * FROM produtos;

-- name: CreateProduto :exec
INSERT INTO produtos (id, empresa_id, codigo_mc, descricao, cod_barras, ncm, cest, cbenef, preco) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateProduto :exec
UPDATE produtos SET empresa_id=?, codigo_mc=?, descricao=?, cod_barras=?, ncm=?, cest=?, cbenef=?, preco=? WHERE id=?;

-- name: DeleteProduto :exec
DELETE FROM produtos WHERE id = ?;




### PAGAMENTOS ###
-- name: GetPagamento :one
SELECT * FROM pagamentos WHERE id = ?;

-- name: ListPagamentos :many
SELECT * FROM pagamentos;

-- name: GetPagamentoByEmpresaID :many
SELECT * FROM pagamentos  WHERE empresa_id = ?;

-- name: CreatePagamento :exec
INSERT INTO pagamentos (id, empresa_id, descricao, tp_pagamento, ind_pagamento, pos_exclusivo, id_pos) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdatePagamento :exec
UPDATE pagamentos SET empresa_id=?, descricao=?, tp_pagamento=?, ind_pagamento=?, pos_exclusivo=?, id_pos=? WHERE id=?;

-- name: DeletePagamento :exec
DELETE FROM pagamentos WHERE id = ?;




### MAQUININHA_POS ###
-- name: GetMaquininhasPOS :one
SELECT * FROM maquininhas_pos WHERE id = ?;

-- name: ListMaquininhaPOS :many
SELECT * FROM maquininhas_pos;

-- name: GetMaquininhasPosByEmpresa :many
SELECT * FROM maquininhas_pos WHERE id_empresa = ?;

-- name: CreateMaquininhaPOS :exec
INSERT INTO maquininhas_pos (id, id_empresa, descricao, administradora, cnpj) VALUES (?, ?, ?, ?, ?);

-- name: UpdateMaquininhaPOS :exec
UPDATE maquininhas_pos SET id_empresa=?, descricao=?, administradora=?, cnpj=? WHERE id=?;

-- name: DeleteMaquininhaPOS :exec
DELETE FROM maquininhas_pos WHERE id = ?;




### VENDAS ###
-- name: GetVenda :one
SELECT * FROM vendas WHERE id = ?;

-- name: GetVendaByEmpresa :many
SELECT * FROM vendas WHERE empresa_id = ?;

-- name: GetVendaByStatus :many
SELECT * FROM vendas WHERE status = ?;

-- name: ListVendas :many
SELECT * FROM vendas;

-- name: CreateVenda :exec
INSERT INTO vendas (id, empresa_id, cliente_id, venda_mc, total, data, status) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateVenda :exec
UPDATE vendas SET empresa_id=?, cliente_id=?, venda_mc=?, total=?, data=?, status=? WHERE id=?;

-- name: DeleteVenda :exec
DELETE FROM vendas WHERE id = ?;




### ITENS_VENDA ###
-- name: FindAll :many
SELECT * FROM itens_venda;

-- name: GetItemVenda :one
SELECT * FROM itens_venda WHERE id = ?;

-- name: GetItensVenda :many
SELECT * FROM itens_venda WHERE venda_id = ?;

-- name: CreateItemVenda :exec
INSERT INTO itens_venda (id, empresa_id, venda_id, produto_id, quantidade, valor, total) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateItemVenda :exec
UPDATE itens_venda SET venda_id=?, empresa_id=?, produto_id=?, quantidade=?, valor=?, total=? WHERE id=?;

-- name: DeleteItemVenda :exec
DELETE FROM itens_venda WHERE id = ?;

-- name: DeleteItensDaVenda :exec
DELETE FROM itens_venda WHERE venda_id = ?;




### PAGAMENTOS_VENDA ###
-- name: GetPagamentoVenda :one
SELECT * FROM pagamentos_venda WHERE id = ?;

-- name: GetPagamentoVendaByStatus :many
SELECT * FROM pagamentos_venda WHERE status = ?;


-- name: GetPagamentoVendaByVendaID :many
SELECT * FROM pagamentos_venda WHERE venda_id = ?;

-- name: ListPagamentoVenda :many
SELECT * FROM pagamentos_venda;

-- name: CreatePagamentoVenda :exec
INSERT INTO pagamentos_venda (id, venda_id, pagamento_id, valor, vencimento, status) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdatePagamentoVenda :exec
UPDATE pagamentos_venda SET venda_id=?, pagamento_id=?, valor=?, vencimento=?, status=? WHERE id=?;

-- name: DeletePagamentoVenda :exec
DELETE FROM pagamentos_venda WHERE id = ?;

-- name: DeletePagamentosDaVenda :exec
DELETE FROM pagamentos_venda WHERE venda_id = ?;


