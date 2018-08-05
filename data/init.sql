CREATE SCHEMA ki;
CREATE SEQUENCE ki.ts_index START WITH 100000000;
CREATE SEQUENCE ki.iv_index START WITH 200000000;
CREATE SEQUENCE ki.c_index START WITH 300000000;

CREATE TABLE ki.time_slots (
    id INTEGER DEFAULT nextval('ki.ts_index'::regclass) NOT NULL,
    date_from DATE NOT NULL,
    date_to DATE NOT NULL,
    CONSTRAINT pk PRIMARY KEY(id)
);

CREATE TABLE ki.interviewers (
    id INTEGER PRIMARY KEY DEFAULT nextval('ki.iv_index'::regclass) NOT NULL,
    name VARCHAR(50),
    timeslot_id INTEGER,
    CONSTRAINT fk_ts_id FOREIGN KEY(timeslot_id) REFERENCES ki.time_slots(id)
);

CREATE TABLE ki.candidates (
    id INTEGER DEFAULT nextval('ki.c_index'::regclass) NOT NULL,
    name VARCHAR(50),
    timeslot_id INTEGER,
    CONSTRAINT fk_ts_id FOREIGN KEY(timeslot_id) REFERENCES ki.time_slots(id)
);