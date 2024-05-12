GRANT ALL PRIVILEGES ON DATABASE postgres to postgres;

CREATE SCHEMA IF NOT EXISTS onco_base;

CREATE TABLE IF NOT EXISTS onco_base.external_user
(
    id       SERIAL       NOT NULL UNIQUE,
    email    VARCHAR(60)  NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role     VARCHAR(30)  NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.internal_user
(
    id          SERIAL      NOT NULL UNIQUE,
    first_name  VARCHAR(30),
    middle_name VARCHAR(30),
    last_name   VARCHAR(30),
    email       VARCHAR(60) NOT NULL UNIQUE,
    phone       VARCHAR(12) UNIQUE,
    password    VARCHAR(30) NOT NULL,
    role        VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.diagnosis
(
    id          VARCHAR(10)  NOT NULL UNIQUE,
    description VARCHAR(300),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.patient
(
    id          SERIAL NOT NULL UNIQUE,
    first_name  VARCHAR(30),
    middle_name VARCHAR(30),
    last_name   VARCHAR(30),
    birth_date  DATE,
    sex         VARCHAR(10),
    snils       VARCHAR(12) UNIQUE,
    user_id     INT UNIQUE,
    phone       VARCHAR(12) UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES onco_base.external_user (id)
);

CREATE TABLE IF NOT EXISTS onco_base.doctor
(
    id            SERIAL      NOT NULL UNIQUE,
    first_name    VARCHAR(30) NOT NULL,
    middle_name   VARCHAR(30) NOT NULL,
    last_name     VARCHAR(30) NOT NULL,
    qualification VARCHAR(300),
    user_id       INT UNIQUE,
    phone         VARCHAR(12) UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES onco_base.external_user (id)
);

CREATE TABLE IF NOT EXISTS onco_base.doctor_patient
(
    patient INT NOT NULL,
    doctor  INT NOT NULL,
    PRIMARY KEY (patient, doctor),
    FOREIGN KEY (patient) REFERENCES onco_base.patient (id),
    FOREIGN KEY (doctor) REFERENCES onco_base.doctor (id)
);

CREATE TABLE IF NOT EXISTS onco_base.drug
(
    id                 VARCHAR(10) NOT NULL UNIQUE,
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
    id        VARCHAR(15) NOT NULL UNIQUE,
    shorthand VARCHAR(15) NOT NULL,
    full_text VARCHAR(30),
    global    VARCHAR(15),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.course
(
    id           VARCHAR(30) NOT NULL UNIQUE,
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
    id               VARCHAR(15) NOT NULL UNIQUE,
    description      VARCHAR(300),
    min_normal_value FLOAT       NOT NULL,
    max_normal_value FLOAT       NOT NULL,
    min_possible_value        FLOAT       NOT NULL,
    max_possible_value        FLOAT       NOT NULL,
    measure_code     VARCHAR(15) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (measure_code) REFERENCES onco_base.unit_measure (id)
);

CREATE TABLE IF NOT EXISTS onco_base.disease
(
    id          VARCHAR(15) NOT NULL UNIQUE,
    description VARCHAR(300),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS onco_base.patient_disease
(
    patient   INT NOT NULL,
    disease   VARCHAR(15) NOT NULL,
    stage     VARCHAR(10),
    diagnosis VARCHAR(10) NOT NULL,
    PRIMARY KEY (patient, disease),
    FOREIGN KEY (patient) REFERENCES onco_base.patient (id),
    FOREIGN KEY (disease) REFERENCES onco_base.disease (id),
    FOREIGN KEY (diagnosis) REFERENCES onco_base.diagnosis (id)
);

CREATE TABLE IF NOT EXISTS onco_base.patient_course
(
    id         SERIAL      NOT NULL UNIQUE,
    patient    INT         NOT NULL,
    disease    VARCHAR(15) NOT NULL,
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
    id             SERIAL NOT NULL UNIQUE,
    patient_course INT    NOT NULL,
    doctor         INT    NOT NULL,
    begin_date     DATE   NOT NULL,
    period         INT,
    result         VARCHAR(10),
    PRIMARY KEY (id),
    FOREIGN KEY (patient_course) REFERENCES onco_base.patient_course (id),
    FOREIGN KEY (doctor) REFERENCES onco_base.doctor (id)
);

CREATE TABLE IF NOT EXISTS onco_base.procedure_blood_count
(
    procedure    INT         NOT NULL,
    blood_count  VARCHAR(15) NOT NULL,
    value        FLOAT,
    measure_code VARCHAR(15),
    PRIMARY KEY (procedure, blood_count),
    FOREIGN KEY (procedure) REFERENCES onco_base.course_procedure (id),
    FOREIGN KEY (measure_code) REFERENCES onco_base.unit_measure (id)
);


-- INSERT INTO onco_base.external_user (email, password, role) 
-- VALUES ('sas@yandex.ru', '156brsdfgsfd6t7dghasvdh', 'doctor') RETURNING email;