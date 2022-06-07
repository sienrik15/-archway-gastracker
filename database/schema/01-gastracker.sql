CREATE TABLE metadata
(
    id              SERIAL                   PRIMARY KEY,
    contract        TEXT                     NOT NULL,
    developer       TEXT                     NOT NULL DEFAULT '',
    reward_addr     TEXT                     NOT NULL DEFAULT '',
    gas_rebate      BOOLEAN                  NOT NULL DEFAULT FALSE,
    premium         BOOLEAN                  NOT NULL DEFAULT FALSE,
    premium_percent INTEGER                  NOT NULL DEFAULT 0,
    creation_time   TEXT                     NOT NULL DEFAULT '',
    index           INTEGER                  NOT NULL,
    height          BIGINT                   NOT NULL
);

CREATE INDEX metadata_contract_index ON metadata (contract);
CREATE INDEX metadata_developer_index ON metadata (developer);

CREATE TABLE block_rewards
(
    id              SERIAL                   PRIMARY KEY,
    contract        TEXT                     NOT NULL,
    gas_consumed    BIGINT                   NOT NULL,
    inflation       BIGINT                   NOT NULL,
    rewards         BIGINT                   NOT NULL,
    height          BIGINT                   NOT NULL
);

CREATE INDEX rewards_contract_index ON block_rewards (contract);

CREATE TABLE total_rewads
(
    contract        TEXT                     NOT NULL PRIMARY KEY,
    rewards         BIGINT                   NOT NULL
);