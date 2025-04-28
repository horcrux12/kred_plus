-- +migrate Up

CREATE TABLE customer (
    id                 bigint NOT NULL AUTO_INCREMENT,
    nik                varchar(20) NOT NULL,
    full_name          varchar(255),
    legal_name         varchar(255),
    place_of_birth     varchar(255),
    date_of_birth      date,
    customer_salary    float,
    identity_card_url  varchar(255),
    selfie_photo_url   varchar(255),
    created_by         bigint NOT NULL,
    created_at         timestamp NOT NULL,
    updated_by         bigint,
    updated_at         timestamp NULL,
    deleted_by         bigint,
    deleted_at         timestamp NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX (id)
);

CREATE TABLE credit_limit (
  id              bigint NOT NULL AUTO_INCREMENT,
  customer_id     bigint NOT NULL,
  tenor_months    int(11) NOT NULL,
  limit_amount    float NOT NULL,
  available_limit float NOT NULL,
  created_by      bigint NOT NULL,
  created_at      timestamp NOT NULL,
  updated_by      bigint,
  updated_at      timestamp NULL,
  deleted_by      bigint,
  deleted_at      timestamp NULL,
  PRIMARY KEY (id)
);

CREATE TABLE interest_setting (
  id            bigint NOT NULL AUTO_INCREMENT,
  tenor_months  int(10) NOT NULL,
  interest_rate decimal(5, 2) NOT NULL,
  description   varchar(255) NOT NULL,
  is_active     bit(1),
  created_by    bigint NOT NULL,
  created_at    timestamp NOT NULL,
  updated_by    bigint,
  updated_at    timestamp NULL,
  deleted_by    bigint,
  deleted_at    timestamp NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX (id)
);

CREATE TABLE transaction_loan (
   id                 bigint NOT NULL AUTO_INCREMENT,
   customer_id        bigint NOT NULL,
   contract_number    varchar(255) NOT NULL UNIQUE,
   otr_price          float NOT NULL,
   admin_fee          float NOT NULL,
   principal_amount   bigint NOT NULL,
   interest_amount    float NOT NULL,
   installment_amount float NOT NULL,
   total_loan         int(11) NOT NULL,
   tenor_months       int(11) NOT NULL,
   interest_id        bigint NOT NULL,
   asset_name         varchar(255) NOT NULL,
   platform           int(11) NOT NULL,
   created_by         bigint NOT NULL,
   created_at         timestamp NOT NULL,
   updated_by         bigint,
   updated_at         timestamp NULL,
   deleted_by         bigint,
   deleted_at         timestamp NULL,
   PRIMARY KEY (id),
   UNIQUE INDEX (id)
);

CREATE TABLE `user` (
    id          bigint NOT NULL AUTO_INCREMENT,
    username    varchar(255) NOT NULL UNIQUE,
    password    text NOT NULL,
    is_admin    BOOLEAN NOT NULL,
    customer_id bigint UNIQUE,
    created_by  bigint NOT NULL,
    created_at  timestamp NOT NULL,
    updated_by  bigint,
    updated_at  timestamp NULL,
    deleted_by  bigint,
    deleted_at  timestamp NULL,
    PRIMARY KEY (id)
);

ALTER TABLE transaction_loan ADD CONSTRAINT FKtransactio300247 FOREIGN KEY (interest_id) REFERENCES interest_setting (id);
ALTER TABLE transaction_loan ADD CONSTRAINT FKtransactio364290 FOREIGN KEY (customer_id) REFERENCES customer (id);
ALTER TABLE `user` ADD CONSTRAINT FKuser284714 FOREIGN KEY (customer_id) REFERENCES customer (id);
ALTER TABLE `user` ADD CONSTRAINT uq_user_username UNIQUE (username);
ALTER TABLE `user` ADD CONSTRAINT uq_user_customer UNIQUE (customer_id);
ALTER TABLE credit_limit ADD CONSTRAINT FKcredit_lim987830 FOREIGN KEY (customer_id) REFERENCES customer (id);
ALTER TABLE credit_limit ADD CONSTRAINT uq_creditlimit_customer_tenor UNIQUE (customer_id, tenor_months);
ALTER TABLE customer ADD CONSTRAINT uq_customer_nik UNIQUE (nik);


-- +migrate Down
ALTER TABLE `transaction_loan` DROP FOREIGN KEY FKtransactio300247;
ALTER TABLE `transaction_loan` DROP FOREIGN KEY FKtransactio364290;
ALTER TABLE `user` DROP FOREIGN KEY FKuser284714;
ALTER TABLE `user` DROP CONSTRAINT uq_user_username;
ALTER TABLE `user` DROP CONSTRAINT uq_user_customer;
ALTER TABLE credit_limit DROP FOREIGN KEY FKcredit_lim987830;
ALTER TABLE credit_limit DROP CONSTRAINT uq_creditlimit_customer_tenor;
ALTER TABLE customer DROP CONSTRAINT uq_customer_nik;
DROP TABLE IF EXISTS credit_limit;
DROP TABLE IF EXISTS customer;
DROP TABLE IF EXISTS interest_setting;
DROP TABLE IF EXISTS `transaction`;
DROP TABLE IF EXISTS `user`;
