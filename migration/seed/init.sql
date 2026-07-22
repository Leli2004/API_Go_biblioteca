
-- Seed completo para esquema biblioteca
-- Insere dados em ordem que respeita chaves estrangeiras

SET search_path TO biblioteca;

-- Users
INSERT INTO users (id, name, email, password_hash, phone, role, active)
VALUES
  (1, 'Admin User', 'admin@example.com', '$2y$10$adminhash', '000-000-0000', 'admin', TRUE),
  (2, 'Librarian One', 'librarian@example.com', '$2y$10$librarianhash', '111-111-1111', 'admin', TRUE),
  (3, 'Member One', 'member1@example.com', '$2y$10$memberhash', '222-222-2222', 'user', TRUE);

-- Genres
INSERT INTO genres (id, name, description)
VALUES
  (1, 'Fiction', 'Fictional works'),
  (2, 'Science', 'Science and technology'),
  (3, 'History', 'Historical works'),
  (4, 'Children', 'Books for children');

-- Authors
INSERT INTO authors (id, name)
VALUES
  (1, 'Gabriel García Márquez'),
  (2, 'Isaac Asimov'),
  (3, 'Yuval Noah Harari'),
  (4, 'Dr. Seuss');

-- Publishers
INSERT INTO publishers (id, name, website)
VALUES
  (1, 'Editora Exemplo', 'https://editoraxemplo.example'),
  (2, 'Science Press', 'https://sciencepress.example');

-- Books
INSERT INTO books (id, publisher_id, title, publication_year, description)
VALUES
  (1, 1, 'Cien Años de Soledad', 1967, 'A classic novel.'),
  (2, 2, 'Foundation', 1951, 'Science fiction classic.'),
  (3, 1, 'Sapiens', 2011, 'A brief history of humankind.'),
  (4, 1, 'The Cat in the Hat', 1957, 'Children''s picture book.');

-- Book Authors (many-to-many)
INSERT INTO book_authors (id, book_id, author_id)
VALUES
  (1, 1, 1),
  (2, 2, 2),
  (3, 3, 3),
  (4, 4, 4);

-- Book Genres
INSERT INTO book_genres (id, book_id, genre_id)
VALUES
  (1, 1, 1),
  (2, 2, 2),
  (3, 3, 3),
  (4, 4, 4);

-- Book Copies
INSERT INTO book_copies (id, book_id, barcode, status)
VALUES
  (1, 1, 'BC-0001', 'available'),
  (2, 1, 'BC-0002', 'loaned'),
  (3, 2, 'BC-1001', 'available'),
  (4, 3, 'BC-2001', 'available'),
  (5, 4, 'BC-3001', 'available');

-- Loans
-- Loan: user 3 has borrowed book_copy 2 (Cien Años de Soledad copy BC-0002)
INSERT INTO loans (id, user_id, book_copy_id, loan_date, due_date, returned_at, status)
VALUES
  (1, 3, 2, NOW(), NOW() + INTERVAL '14 days', NULL, 'active');

-- Reservations
-- Member 3 reserves book 2 (Foundation)
INSERT INTO reservations (id, user_id, book_id, position, status, expires_at)
VALUES
  (1, 3, 2, 1, 'waiting', NOW() + INTERVAL '7 days');

-- Fines
-- Create a fine for loan 1 (user 3)
INSERT INTO fines (id, loan_id, user_id, amount, reason, paid, paid_at)
VALUES
  (1, 1, 3, 5.00, 'Late return', FALSE, NULL);

-- Atualiza sequences para respeitar próximos inserts
SELECT setval('biblioteca.users_id_seq', 3, true);
SELECT setval('biblioteca.genres_id_seq', 4, true);
SELECT setval('biblioteca.authors_id_seq', 4, true);
SELECT setval('biblioteca.publishers_id_seq', 2, true);
SELECT setval('biblioteca.books_id_seq', 4, true);
SELECT setval('biblioteca.book_authors_id_seq', 4, true);
SELECT setval('biblioteca.book_genres_id_seq', 4, true);
SELECT setval('biblioteca.book_copies_id_seq', 5, true);
SELECT setval('biblioteca.loans_id_seq', 1, true);
SELECT setval('biblioteca.reservations_id_seq', 1, true);
SELECT setval('biblioteca.fines_id_seq', 1, true);

-- Fim do seed
