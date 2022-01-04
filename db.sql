CREATE SCHEMA goclock AUTHORIZATION root;

create table if not exists goclock.tbl_user (
	id SERIAL,		
	password VARCHAR(255) not null CHECK (char_length(password) >= 6),	
	email VARCHAR(255) not null UNIQUE,

	primary key(id)
);

create table if not exists goclock.tbl_subject (
	id SERIAL,
	user_id integer not null,

	subject_name VARCHAR(255) not null,
	primary key(id),
	foreign key (user_id) references tbl_user(id)
);

create table if not exists goclock.tbl_event (
	id SERIAL,
	user_id integer not null,

	event_name VARCHAR(255) not null UNIQUE,
	primary key(id),
	foreign key (user_id) references tbl_user(id)
);

create table if not exists goclock.tbl_event_subjects (
	event_id integer not null,
	subject_id integer not null,

	primary key (event_id, subject_id),
	foreign key (event_id) references tbl_event(id),
	foreign key (subject_id) references tbl_subject(id)
);

create table if not exists goclock.tbl_entry (
	id SERIAL,
	type int2 not null default 0,
	entry_desc VARCHAR,
	user_id integer not null,
	subject_id integer not null,
	event_id integer,	
	dt_start timestamp not null CHECK (dt_start > '2021-01-01'),
	dt_end timestamp not null CHECK (dt_end > dt_start),

	primary key(id),
	foreign key (user_id) references tbl_user(id),
	foreign key (event_id) references tbl_event(id),
	foreign key (subject_id) references tbl_subject(id)
);

create table if not exists goclock.tbl_task (
	id SERIAL,
	task_name VARCHAR,
	user_id integer not null,

	primary key (id),
	foreign key (user_id) references tbl_user(id)
);

create table if not exists goclock.tbl_task_completion (
	id SERIAL,
	task_id integer not null,
	completed_at timestamp,

	primary key (id),
	foreign key (task_id) references tbl_task(id)
);

create table if not exists goclock.tbl_course (
	id SERIAL,
	duration integer not null default 0,
	user_id integer not null,
	course_name varchar not null,

	primary key (id),
	unique(user_id, course_name)
);


create table if not exists goclock.tbl_study_plan (
	id SERIAL,
	study_plan_name varchar not null,
	current_course int not null default 0,
	user_id int not null,

	primary key (id),
	unique(user_id, study_plan_name)
);

create table if not exists goclock.tbl_study_plan_course (
	course_id int not null,
	study_plan_id int not null,
	study_plan_course_order int not null default 0,

	primary key (course_id, study_plan_id),
	unique(study_plan_id, study_plan_course_order)
);