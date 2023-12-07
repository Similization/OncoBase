-- CREATE DATABASE med_base;

CREATE SCHEMA IF NOT EXISTS onco_base;

CREATE TABLE onco_base.external_user
(
    email    VARCHAR(60)  NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role     VARCHAR(30)  NOT NULL,

    PRIMARY KEY (email)
);

CREATE TABLE IF NOT EXISTS onco_base.admin
(
    id          SERIAL,
    first_name  VARCHAR(30),
    middle_name VARCHAR(30),
    last_name   VARCHAR(30),
    email       VARCHAR(60) NOT NULL UNIQUE,
    password    VARCHAR(30) NOT NULL,
--     role        VARCHAR(15) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.patient
(
    id          SERIAL,
    first_name  VARCHAR(30),
    middle_name VARCHAR(30),
    last_name   VARCHAR(30),
    birth_date  DATE,
    sex         VARCHAR(10),
    snils       VARCHAR(12) UNIQUE,
    email       VARCHAR(60),
    phone       VARCHAR(12) UNIQUE,

    PRIMARY KEY (id),
    FOREIGN KEY (email) REFERENCES onco_base.external_user (email)
);

-- INSERT INTO onco_base.patient (email, password)
-- VALUES ('dan18b@yandex.ru', '12qw12wq');
--
-- SELECT id
-- FROM onco_base.patient
-- WHERE email = 'dan18b@yandex.ru'
--   AND password = '12qw12wq';

-- SELECT email
-- FROM onco_base.external_user
-- WHERE email = 'dan18b@yandex.ru'
--   AND password = '355c6e24644ef1ccc0112f3388cf53cbd8631785fb1704c09ffd23a9685d2dad52f78d1fcc571200e98d723eb92cab6521361868157f4966bd4ba54fe06e12fb';

CREATE TABLE IF NOT EXISTS onco_base.doctor
(
    id            SERIAL,
    first_name    VARCHAR(30) NOT NULL,
    middle_name   VARCHAR(30) NOT NULL,
    last_name     VARCHAR(30) NOT NULL,
    qualification VARCHAR(300),
    email         VARCHAR(60),
    phone         VARCHAR(12) UNIQUE,

    PRIMARY KEY (id),
    FOREIGN KEY (email) REFERENCES onco_base.external_user (email)
);

CREATE TABLE IF NOT EXISTS onco_base.doctor_patient
(
    patient INT NOT NULL,
    doctor  INT NOT NULL,

    PRIMARY KEY (patient, doctor),

    FOREIGN KEY (patient) REFERENCES onco_base.patient (id),
    FOREIGN KEY (doctor) REFERENCES onco_base.doctor (id)
);

CREATE TABLE IF NOT EXISTS onco_base.diagnosis
(
    id          VARCHAR(10)  NOT NULL,
    description VARCHAR(300) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.drug
(
    id                 VARCHAR(10) NOT NULL,
    name               VARCHAR(60) NOT NULL UNIQUE,
    dosage_form        VARCHAR(30) NOT NULL,
    active_ingredients VARCHAR(60) NOT NULL,
    country            VARCHAR(30),
    manufacturer       VARCHAR(45),
    prescribing_order  VARCHAR(30),
    description        VARCHAR(300),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.unit_measure
(
    id        VARCHAR(15) NOT NULL,
    shorthand VARCHAR(15) NOT NULL,
    full_text VARCHAR(30) UNIQUE,
    global    VARCHAR(15),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.course
(
    id           VARCHAR(30) NOT NULL,
    period       INT         NOT NULL,
    frequency    FLOAT       NOT NULL,
    dose         FLOAT       NOT NULL,
    drug         VARCHAR(10) NOT NULL,
    measure_code VARCHAR(15) NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (drug) REFERENCES onco_base.drug (id),
    FOREIGN KEY (measure_code) REFERENCES onco_base.unit_measure (id)
);

CREATE TABLE IF NOT EXISTS onco_base.blood_count
(
    id               VARCHAR(15) NOT NULL,
    min_normal_value FLOAT       NOT NULL,
    max_normal_value FLOAT       NOT NULL,
    min_value        FLOAT       NOT NULL,
    max_value        FLOAT       NOT NULL,
    measure_code     VARCHAR(15) NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (measure_code) REFERENCES onco_base.unit_measure (id)
);

CREATE TABLE IF NOT EXISTS onco_base.disease
(
    id          VARCHAR(15) NOT NULL,
    description VARCHAR(300),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.patient_disease
(
    patient   INT NOT NULL,
    disease   VARCHAR(15),
    stage     VARCHAR(10),
    diagnosis VARCHAR(10),


    PRIMARY KEY (patient, disease),

    FOREIGN KEY (patient) REFERENCES onco_base.patient (id),
    FOREIGN KEY (disease) REFERENCES onco_base.disease (id),
    FOREIGN KEY (diagnosis) REFERENCES onco_base.diagnosis (id)
);

CREATE TABLE IF NOT EXISTS onco_base.patient_course
(
    id         INT         NOT NULL,
    patient    INT         NOT NULL,
    disease    VARCHAR(15),
    course     VARCHAR(30) NOT NULL,
    doctor     INT         NOT NULL,
    begin_date DATE        NOT NULL,
    end_date   DATE,
    diagnosis  VARCHAR(10),

    PRIMARY KEY (id),

    FOREIGN KEY (patient, disease) REFERENCES onco_base.patient_disease (patient, disease),
    FOREIGN KEY (course) REFERENCES onco_base.course (id),
    FOREIGN KEY (doctor) REFERENCES onco_base.doctor (id),
    FOREIGN KEY (diagnosis) REFERENCES onco_base.diagnosis (id)
);

CREATE TABLE IF NOT EXISTS onco_base.blood_count_value
(
    disease     VARCHAR(15) NOT NULL,
    blood_count VARCHAR(15) NOT NULL,
    coefficient FLOAT       NOT NULL,
    description VARCHAR(300),

    PRIMARY KEY (disease, blood_count),

    FOREIGN KEY (disease) REFERENCES onco_base.disease (id),
    FOREIGN KEY (blood_count) REFERENCES onco_base.blood_count (id)
);

CREATE TABLE IF NOT EXISTS onco_base.course_procedure
(
    patient_course INT         NOT NULL,
    blood_count    VARCHAR(15) NOT NULL,
    begin_date     DATE        NOT NULL,
    period         INT         NOT NULL,
    result         FLOAT,
    measure_code   VARCHAR(15),

    PRIMARY KEY (patient_course, blood_count),

    FOREIGN KEY (patient_course) REFERENCES onco_base.patient_course (id),
    FOREIGN KEY (blood_count) REFERENCES onco_base.blood_count (id),
    FOREIGN KEY (measure_code) REFERENCES onco_base.unit_measure (id)
);
