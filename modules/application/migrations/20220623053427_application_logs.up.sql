create table application_logs (
    id varchar(64) primary key,
    application_id varchar(64) not null,
	job_id varchar(64) not null,
	hiring_step_type varchar(64) not null,
	hiring_step_sequence int not null,
	hiring_step_status varchar(32) not null,
	hiring_step_status_closed_at timestamptz,
	hiring_step_status_closed_by varchar(64),
	hiring_step_status_closed_by_name varchar(128),
	user_type varchar(16),
    created_at timestamptz not null default (now()),
	created_by varchar(64),
	created_by_name varchar(128)
);
