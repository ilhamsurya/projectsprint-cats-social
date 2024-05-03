INSERT INTO "user" ("email", "name", "password") 
VALUES 
('user1@example.com', 'John Doe', 'password123'),
('user2@example.com', 'Jane Smith', 'securepassword');
INSERT INTO "user" ("email", "name", "password") 
VALUES 
('user3@example.com', 'Emma Watson', 'password456'),
('user4@example.com', 'Michael Johnson', 'strongpassword'),
('user5@example.com', 'Sophia Garcia', 'securepassword123');


INSERT INTO "cat" ("id_user", "name", "race", "sex", "age_in_month", "description") 
VALUES 
(1, 'Molly', 'Siamese', 'Female', 24, 'Playful and affectionate'),
(1, 'Tom', 'Maine Coon', 'Male', 36, 'Gentle giant'),
(2, 'Bella', 'Persian', 'Female', 12, 'Fluffy and elegant');
INSERT INTO "cat" ("id_user", "name", "race", "sex", "age_in_month", "description") 
VALUES 
(2, 'Simba', 'Bengal', 'Male', 18, 'Energetic and playful'),
(2, 'Luna', 'Scottish Fold', 'Female', 10, 'Cute and curious'),
(3, 'Oscar', 'Ragdoll', 'Male', 24, 'Calm and affectionate'),
(3, 'Milo', 'British Shorthair', 'Male', 14, 'Friendly and gentle'),
(4, 'Cleo', 'Sphynx', 'Female', 36, 'Unique and intelligent'),
(5, 'Max', 'Abyssinian', 'Male', 12, 'Active and adventurous');


INSERT INTO "cat_image" ("id_cat", "image") 
VALUES 
(1, 'https://example.com/molly_1.jpg'),
(1, 'https://example.com/molly_2.jpg'),
(1, 'https://example.com/molly_3.jpg'),
(2, 'https://example.com/tom_1.jpg'),
(2, 'https://example.com/tom_2.jpg'),
(3, 'https://example.com/bella_1.jpg'),
(3, 'https://example.com/bella_2.jpg');
INSERT INTO "cat_image" ("id_cat", "image") 
VALUES 
(4, 'https://example.com/simba_1.jpg'),
(4, 'https://example.com/simba_2.jpg'),
(5, 'https://example.com/luna_1.jpg'),
(5, 'https://example.com/luna_2.jpg'),
(6, 'https://example.com/oscar_1.jpg'),
(6, 'https://example.com/oscar_2.jpg'),
(7, 'https://example.com/milo_1.jpg'),
(8, 'https://example.com/cleo_1.jpg'),
(9, 'https://example.com/max_1.jpg'),
(9, 'https://example.com/max_2.jpg');

INSERT INTO "match_cat" ("id_user_cat", "id_matched_cat", "is_matched")
VALUES 
(1, 2, true),
(4, 5, false);