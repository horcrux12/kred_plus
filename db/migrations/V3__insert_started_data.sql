-- +migrate Up

INSERT INTO interest_setting
(tenor_months, interest_rate, description, is_active, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES
    (1, 3, 'Default rate', 1, 0, now(), null, null, null, null),
    (2, 3, 'Default rate', 1, 0, now(), null, null, null, null),
    (3, 3, 'Default rate', 1, 0, now(), null, null, null, null),
    (4, 3, 'Default rate', 1, 0, now(), null, null, null, null);

INSERT INTO customer
(nik, full_name, legal_name, place_of_birth, date_of_birth, customer_salary, identity_card_url, selfie_photo_url, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES
    ('1234567890098788', 'Budi Gunawan', 'BUDI SIGIT GUNAWAN', 'Jakarta', STR_TO_DATE('12-04-1995', '%d-%m-%Y'), 12000000, '', '', 0, now(), NULL, NULL, NULL, NULL),
    ('1234567890098789', 'Annisa Ramadhani', 'ANNISA RAMADHANI', 'Jakarta', STR_TO_DATE('12-05-1995', '%d-%m-%Y'), 15000000, '', '', 0, now(), NULL, NULL, NULL, NULL);

INSERT INTO `user`
(username, password, is_admin, customer_id, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES
    ('annisa', '$2a$04$D1Wp0DMDsG8NfcvUcphpuOBsZV93pjQt1wZ0PwFSV41.p5x1LU9Aa', 0, (SELECT id FROM customer WHERE nik = '1234567890098788'), 0, now(), NULL, NULL, NULL, NULL),
    ('budigun', '$2a$04$6bxLA59UjB3uGa7ANT9Mm.1nB4dCNspi2LESruyUWnyFLVmtZaeUu', 0, (SELECT id FROM customer WHERE nik = '1234567890098789'), 0, now(), NULL, NULL, NULL, NULL);

INSERT INTO credit_limit
(customer_id, tenor_months, limit_amount, available_limit, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES
    ((SELECT id FROM customer WHERE nik = '1234567890098789'), 1, 1000000, 1000000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098789'), 2, 1200000, 1200000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098789'), 3, 1500000, 1500000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098789'), 4, 2000000, 2000000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098788'), 1, 100000, 100000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098788'), 2, 200000, 200000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098788'), 3, 500000, 500000, 0, now(), NULL, NULL, NULL, NULL),
    ((SELECT id FROM customer WHERE nik = '1234567890098788'), 4, 700000, 700000, 0, now(), NULL, NULL, NULL, NULL);

-- +migrate Down