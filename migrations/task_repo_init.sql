CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(500),
    planning_date DATE,
    status BOOLEAN
);
