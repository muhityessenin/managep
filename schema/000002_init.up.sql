CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name           varchar(255) not null,
    email               varchar(255) not null unique,
    registration_date   date,
    role                varchar(50) not null
);
CREATE TABLE projects
(
    id                          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name                        varchar(255) not null,
    description                 text,
    start_date                  date not null,
    end_date                    date,
    manager_id                  uuid,
    foreign key (manager_id)    references users(id)
);
CREATE TABLE tasks
(
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name            varchar(255) not null,
    description             text,
    priority                varchar(50) not null,
    state                   varchar(50),
    responsible_person_id   UUID,
    foreign key             (responsible_person_id) references users(id),
    project_id              UUID,
    foreign key             (project_id) references projects(id),
    start_date              date not null,
    end_date                date
);



