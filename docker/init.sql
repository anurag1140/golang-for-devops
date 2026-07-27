-------------------------------------------------------
-- USERS
-------------------------------------------------------

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,

    username VARCHAR(100) UNIQUE NOT NULL,

    password_hash TEXT NOT NULL,

    role VARCHAR(30) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-------------------------------------------------------
-- BOOKS
-------------------------------------------------------

CREATE TABLE IF NOT EXISTS books
(
    id INTEGER PRIMARY KEY,

    title VARCHAR(255) NOT NULL,

    author VARCHAR(255) NOT NULL,

    isbn VARCHAR(50) UNIQUE NOT NULL,

    available BOOLEAN DEFAULT TRUE
);

-------------------------------------------------------
-- REFRESH TOKENS
-------------------------------------------------------

CREATE TABLE IF NOT EXISTS refresh_tokens
(
    token TEXT PRIMARY KEY,

    username VARCHAR(100) NOT NULL,

    expires_at TIMESTAMP NOT NULL,

    revoked BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_refresh_user
        FOREIGN KEY(username)
        REFERENCES users(username)
        ON DELETE CASCADE
);

-------------------------------------------------------
-- LOANS
-------------------------------------------------------

CREATE TABLE IF NOT EXISTS loans
(
    id SERIAL PRIMARY KEY,

    book_id INTEGER NOT NULL,

    username VARCHAR(100) NOT NULL,

    borrowed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    due_date TIMESTAMP NOT NULL,

    returned_at TIMESTAMP,

    CONSTRAINT fk_loan_book
        FOREIGN KEY(book_id)
        REFERENCES books(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_loan_user
        FOREIGN KEY(username)
        REFERENCES users(username)
        ON DELETE CASCADE
);

INSERT INTO users(username,password_hash,role)
VALUES
(
'admin',
'$2a$10$592La6yVrX8fIflhQO/KbeBXSqaPN9Mb6PFhHSreikO02XS0xi/16',
'admin'
)
ON CONFLICT(username) DO NOTHING;

INSERT INTO users(username,password_hash,role)
VALUES
(
'librarian',
'$2a$10$592La6yVrX8fIflhQO/KbeBXSqaPN9Mb6PFhHSreikO02XS0xi/16',
'librarian'
)
ON CONFLICT(username) DO NOTHING;

INSERT INTO users(username,password_hash,role)
VALUES
(
'member',
'$2a$10$592La6yVrX8fIflhQO/KbeBXSqaPN9Mb6PFhHSreikO02XS0xi/16',
'member'
)
ON CONFLICT(username) DO NOTHING;

INSERT INTO books(id,title,author,isbn,available)
VALUES
(1,'Clean Code','Robert Martin','9780132350884',TRUE),
(2,'Go Programming','Alan Donovan','9780134190440',TRUE),
(3,'Designing Data Intensive Applications','Martin Kleppmann','9781449373320',TRUE)
ON CONFLICT(id) DO NOTHING;