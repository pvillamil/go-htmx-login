package database

var UserTable = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);
`

var CreateUser = `
INSERT INTO users (username, password) VALUES (?, ?);
`

var GetUser = `
SELECT id, username, password FROM users WHERE username = ?;
`

var DeleteUser = `
DELETE FROM users WHERE id = ?;
`

var LoggedInUserTable = `
CREATE TABLE IF NOT EXISTS logged_in_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL UNIQUE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
`

var CreateLoggedInUser = `
INSERT INTO logged_in_users (user_id, token) VALUES (?, ?);
`

var GetLoggedInUser = `
SELECT id, user_id, token FROM logged_in_users WHERE user_id = ? AND token = ?;
`

var UpdateLoggedInUser = `
UPDATE logged_in_users SET token = ? WHERE id = ?;
`

var DeleteLoggedInUser = `
DELETE FROM logged_in_users WHERE id = ?;
`
