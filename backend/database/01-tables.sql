-- Create tables
create table member (
	username VARCHAR(30) primary key,
	password VARCHAR(30) not null
); 

create table invitation (
	id serial primary key,
	class VARCHAR(30) not null,
	section INT not null,
	missing_uses INT not null,
	created_at timestamp default current_timestamp not null,
	server_id text not null,
	role_id text not null
);