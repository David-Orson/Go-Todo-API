create database todos;

create table todos(
  Id serial Primary key,
  Body varchar(255),
  completed boolean
);

insert into todos (Id,Body,completed) values (1, 'work', false);

insert into todos (Body,completed) values ('workout', false);

select * from todos;

create user david with password 'OrsonDC';

grant all privileges on database todos to david;

alter user david with superuser;
