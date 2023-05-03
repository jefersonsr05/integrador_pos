CREATE TABLE IF NOT EXISTS `empresas` (
    `id`                varchar(36)  NOT NULL PRIMARY KEY,
    `descricao`         text         NOT NULL,
    `cnpj`              varchar(14),
    `chave_registro`    varchar(255)
);

CREATE TABLE IF NOT EXISTS `maquininhas_pos` (
    `id`                varchar(36)  NOT NULL PRIMARY KEY,
    `id_empresa`        varchar(36)  NOT NULL,
    `administradora`    varchar(255),
    `cnpj`              varchar(14),
    FOREIGN KEY (`id_empresa`) REFERENCES `empresas`(`id`)
);

CREATE TABLE IF NOT EXISTS `produtos` (
    `id`            varchar(36) NOT NULL PRIMARY KEY,
    `empresa_id`    varchar(36) NOT NULL,
    `codigo_mc`     varchar(7)  NOT NULL,   
    `descricao`     text        NOT NULL,
    `cod_barras`    varchar(20),
    `ncm`           varchar(8),
    `cest`          varchar(8),
    `cbenef`        varchar(8),
    `preco`         decimal(10,4) NOT NULL,
    FOREIGN KEY (`empresa_id`) REFERENCES `empresas`(`id`)
);

CREATE TABLE IF NOT EXISTS `estados` (
    `id`        varchar(36) NOT NULL PRIMARY KEY,
    `descricao` text        NOT NULL,
    `uf`        varchar(2)  NOT NULL,
    `cod_ibge`  TINYINT  NOT NULL
);

CREATE TABLE IF NOT EXISTS `cidades` (
    `id`        varchar(36) NOT NULL PRIMARY KEY,
    `descricao` text        NOT NULL,
    `estado_id` varchar(36) NOT NULL,
    `cod_ibge`  varchar(7)  NOT NULL,
    FOREIGN KEY (`estado_id`) REFERENCES `estados`(`id`)
);

CREATE TABLE IF NOT EXISTS `clientes` (
    `id`            varchar(36) NOT NULL PRIMARY KEY,
    `codigo_mc`     bigint         NOT NULL,
    `empresa_id`    varchar(36) NOT NULL,
    `nome`          text        NOT NULL,
    `cep`           varchar(9),
    `cidade_id`     varchar(36),
    `endereco`      varchar(255),
    `numero`        varchar(25),
    `complemento`   varchar(255),
    FOREIGN KEY (`empresa_id`) REFERENCES `empresas`(`id`),
    FOREIGN KEY (`cidade_id`)  REFERENCES `cidades`(`id`)
);

CREATE TABLE IF NOT EXISTS `pagamentos` (
    `id`                varchar(36) NOT NULL PRIMARY KEY,
    `empresa_id`        varchar(36) NOT NULL,
    `descricao`         text        NOT NULL,
    `tp_pagamento`      varchar(2)  NOT NULL,
    `ind_pagamento`     varchar(2)  NOT NULL,
    `pos_exclusivo`     boolean     NOT NULL DEFAULT false,
    `id_pos`            varchar(36),
    FOREIGN KEY (`id_pos`)     REFERENCES `maquininhas_pos`(`id`),
    FOREIGN KEY (`empresa_id`) REFERENCES `empresas`(`id`)
);

CREATE TABLE IF NOT EXISTS `vendas` (
    `id`            varchar(36)     NOT NULL PRIMARY KEY,
    `empresa_id`    varchar(36)     NOT NULL,
    `cliente_id`    varchar(36)     NOT NULL,
    `venda_mc`      varchar(10)     NOT NULL, 
    `total`         decimal(10,2)   NOT NULL,
    `data`          DATETIME        DEFAULT CURRENT_TIMESTAMP,
    `status`        boolean         NOT NULL DEFAULT false,
    FOREIGN KEY (`empresa_id`) REFERENCES `empresas`(`id`),
    FOREIGN KEY (`cliente_id`) REFERENCES `clientes`(`id`)
);

CREATE TABLE IF NOT EXISTS `itens_venda` (
    `id`          varchar(36)     NOT NULL PRIMARY KEY,
    `empresa_id`  varchar(36)     NOT NULL,
    `venda_id`    varchar(36)     NOT NULL,
    `produto_id`  varchar(36)     NOT NULL, 
    `quantidade`  decimal(10,4)   NOT NULL,
    `valor`       decimal(10,4)   NOT NULL,
    `total`       decimal(10,2)   NOT NULL,
    FOREIGN KEY (`empresa_id`)    REFERENCES `empresas`(`id`),
    FOREIGN KEY (`venda_id`)      REFERENCES `vendas`(`id`),
    FOREIGN KEY (`produto_id`)    REFERENCES `produtos`(`id`)
);

CREATE TABLE IF NOT EXISTS `pagamentos_venda` (
    `id`              varchar(36)       NOT NULL PRIMARY KEY,
    `venda_id`        varchar(36)       NOT NULL,
    `pagamento_id`    varchar(36)       NOT NULL, 
    `valor`           decimal(10,2)     NOT NULL,
    `vencimento`      DATETIME          DEFAULT CURRENT_TIMESTAMP,
    `status`          boolean           NOT NULL,
    FOREIGN KEY (`venda_id`) REFERENCES `vendas`(`id`),
    FOREIGN KEY (`pagamento_id`) REFERENCES `pagamentos`(`id`)
);

