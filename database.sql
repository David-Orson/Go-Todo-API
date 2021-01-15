create database todos;

create table todos(
  todo_id serial Primary key,
  description varchar(255),
  completed boolean
);

insert into todos (todo_id,description,completed) values (1, 'work', false);

insert into todos (description,completed) values ('workout', false);

select * from todos;

create user david with password 'OrsonDC';

grant all privileges on database todos to david;

alter user david with superuser;
