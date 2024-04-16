-- Database: online-streaming

-- DROP DATABASE IF EXISTS "online-streaming";

CREATE DATABASE "online-streaming"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Spanish_Spain.1252'
    LC_CTYPE = 'Spanish_Spain.1252'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

INSERT INTO public.users(
	username, password, is_enabled)
	VALUES ('admin', 'admin', true);

INSERT INTO public.films(
	title, genre, release_date, synopsis, director, duration, filmaffinity_rating, created_at, updated_at, create_user_id)
	VALUES ('Titanic', 'Romance, Drama', '1997-12-19', 'A seventeen-year-old aristocrat falls in love with a kind but poor artist aboard the luxurious, ill-fated R.M.S. Titanic.', 'James Cameron', 195, 7.8, '2024-04-16','2024-04-16', (select id from public.users where username = 'admin')),
('The Dark Knight', 'Action, Crime, Drama', '2008-07-18', 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.', 'Christopher Nolan', 152, 8.5,'2024-04-16','2024-04-16',(select id from public.users where username = 'admin')),
('Forrest Gump', 'Drama, Romance', '1994-07-06', 'The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold from the perspective of an Alabama man with an IQ of 75, whose only desire is to be reunited with his childhood sweetheart.', 'Robert Zemeckis', 142, 8.7,'2024-04-16','2024-04-16', (select id from public.users where username = 'admin')),
('Interstellar', 'Adventure, Drama, Sci-Fi', '2014-11-07', 'A team of explorers travel through a wormhole in space in an attempt to ensure humanity''s survival.', 'Christopher Nolan', 169, 8.3,'2024-04-16','2024-04-16', (select id from public.users where username = 'admin')),
('Inglourious Basterds', 'Adventure, Drama, War', '2009-08-21', 'In Nazi-occupied France during World War II, a plan to assassinate Nazi leaders by a group of Jewish U.S. soldiers coincides with a theatre owner''s vengeful plans for the same.', 'Quentin Tarantino', 153, 8.1,'2024-04-16','2024-04-16', (select id from public.users where username = 'admin')),
('The Lord of the Rings: The Return of the King', 'Action, Adventure, Drama', '2003-12-17', 'Gandalf and Aragorn lead the World of Men against Sauron''s army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.', 'Peter Jackson', 201, 8.9,'2024-04-16','2024-04-16', (select id from public.users where username = 'admin'));