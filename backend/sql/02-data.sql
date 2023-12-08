INSERT INTO users (username, discord_username, password_hash)
VALUES
('admin', 'admin', 'c93ccd78b2076528346216b3b2f701e6'), --password: admin1234
('user', 'user', 'b5b73fae0d87d8b4e2573105f8fbe7bc'); --password: user1234


INSERT INTO admins 
VALUES
(1);

INSERT INTO skills VALUES 
('JavaScript programming'),
('Golang programming'),
('Skateboarding'),
('Snowboarding');

INSERT INTO categories VALUES
('Programming'),
('Computers'),
('Sports'),
('Winter'),
('Summer');

INSERT INTO skill_categories VALUES
('JavaScript programming', 'Programming'),
('JavaScript programming', 'Computers'),
('Golang programming', 'Programming'),
('Golang programming', 'Computers'),
('Skateboarding', 'Sports'),
('Skateboarding', 'Summer'),
('Snowboarding', 'Sports'),
('Snowboarding', 'Winter');

INSERT INTO offers (user_id, skill, description) VALUES
(1, 'Golang programming', 'I am a senior go dev, want to help you'),
(2, 'Skateboarding', 'Do you want to learn some kickflips?');

INSERT INTO report_reasons VALUES
('Being Rude'),
('Threatening');

INSERT INTO reports (reporting_user_id, reported_user_id, reported_offer_id, reason, description, status) VALUES
(2, 1, 1, 'Threatening', 'This user threatened, that he will destroy my motherboard! :(', 'TODO');

INSERT INTO reviews (reviewing_user_id, reviewed_user_id, star_count, review) VALUES
(1, 2, 5, 'Now I am able to do a kickflip!');