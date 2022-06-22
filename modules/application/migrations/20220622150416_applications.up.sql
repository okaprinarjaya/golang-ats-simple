CREATE TABLE applications (
    id VARCHAR(64) PRIMARY KEY,
    applicant_id VARCHAR(64) NOT NULL,
    job_id VARCHAR(64) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);
