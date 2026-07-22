
CREATE SCHEMA IF NOT EXISTS biblioteca;

SET search_path TO biblioteca;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password_hash TEXT NOT NULL,
    phone VARCHAR(30),
    role VARCHAR(30) NOT NULL DEFAULT 'user',
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_users_role CHECK (role IN ('admin', 'user'))
);

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE publishers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    website VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    publisher_id BIGINT,
    title VARCHAR(200) NOT NULL,
    publication_year INT,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_books_publisher FOREIGN KEY (publisher_id) REFERENCES publishers(id)
);

CREATE TABLE book_authors (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL,
    author_id INT NOT NULL,

    CONSTRAINT fk_book_authors_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    CONSTRAINT fk_book_authors_author FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);

CREATE TABLE book_genres (
    id SERIAL PRIMARY KEY,
    book_id BIGINT NOT NULL,
    genre_id BIGINT NOT NULL,

    CONSTRAINT fk_book_genres_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,

    CONSTRAINT fk_book_genres_genre FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);

CREATE TABLE book_copies (
    id SERIAL PRIMARY KEY,
    book_id BIGINT NOT NULL,
    barcode VARCHAR(100) NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'available',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_book_copies_book
        FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,

    CONSTRAINT chk_book_copies_status
        CHECK (status IN ('available', 'loaned', 'reserved', 'maintenance', 'lost'))
);

CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    user_id INT,
    book_copy_id INT NOT NULL,
    loan_date TIMESTAMP NOT NULL DEFAULT NOW(),
    due_date TIMESTAMP NOT NULL,
    returned_at TIMESTAMP,
    status VARCHAR(30) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_loans_user
        FOREIGN KEY (user_id) REFERENCES users(id),

    CONSTRAINT fk_loans_book_copy
        FOREIGN KEY (book_copy_id) REFERENCES book_copies(id),

    CONSTRAINT chk_loans_status
        CHECK (status IN ('active', 'returned', 'overdue'))
);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    position INT,
    status VARCHAR(30) NOT NULL DEFAULT 'waiting',
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_reservations_user
        FOREIGN KEY (user_id) REFERENCES users(id),

    CONSTRAINT fk_reservations_book
        FOREIGN KEY (book_id) REFERENCES books(id),

    CONSTRAINT chk_reservations_status
        CHECK (status IN ('waiting', 'available', 'cancelled', 'completed'))
);

CREATE TABLE fines (
    id SERIAL PRIMARY KEY,
    loan_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    reason TEXT,
    paid BOOLEAN DEFAULT FALSE,
    paid_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_fines_loan
        FOREIGN KEY (loan_id) REFERENCES loans(id),

    CONSTRAINT fk_fines_user
        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

CREATE INDEX idx_books_publisher_id ON books(publisher_id);

CREATE INDEX idx_book_authors_book_id ON book_authors(book_id);
CREATE INDEX idx_book_authors_author_id ON book_authors(author_id);

CREATE INDEX idx_book_genres_book_id ON book_genres(book_id);
CREATE INDEX idx_book_genres_genre_id ON book_genres(genre_id);

CREATE INDEX idx_book_copies_book_id ON book_copies(book_id);
CREATE INDEX idx_book_copies_status ON book_copies(status);
CREATE INDEX idx_book_copies_barcode ON book_copies(barcode);

CREATE INDEX idx_loans_user_id ON loans(user_id);
CREATE INDEX idx_loans_book_copy_id ON loans(book_copy_id);
CREATE INDEX idx_loans_status ON loans(status);
CREATE INDEX idx_loans_due_date ON loans(due_date);

CREATE INDEX idx_reservations_user_id ON reservations(user_id);
CREATE INDEX idx_reservations_book_id ON reservations(book_id);
CREATE INDEX idx_reservations_status ON reservations(status);
CREATE INDEX idx_reservations_created_at ON reservations(created_at);

CREATE INDEX idx_fines_user_id ON fines(user_id);
CREATE INDEX idx_fines_loan_id ON fines(loan_id);
CREATE INDEX idx_fines_paid ON fines(paid);

CREATE UNIQUE INDEX IF NOT EXISTS uq_fines_loan_id ON biblioteca.fines (loan_id);
