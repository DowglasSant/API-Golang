insert into users (user_name, nick, email, user_password)
values
("User 1", "Usuario_1", "usuario1@gmail.com", "$2a$10$ciaX/x7LUAzVRGF2AN4oOuQa6CI16ui3LjENkDld8xVbfCT/6kutG"),
("User 2", "Usuario_2", "usuario2@gmail.com", "$2a$10$ciaX/x7LUAzVRGF2AN4oOuQa6CI16ui3LjENkDld8xVbfCT/6kutG"),
("User 3", "Usuario_3", "usuario3@gmail.com", "$2a$10$ciaX/x7LUAzVRGF2AN4oOuQa6CI16ui3LjENkDld8xVbfCT/6kutG");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts (title, content, author_id)
values
("User 1 Post", "That's my post!", 1),
("User 2 Post", "That's my post!", 2),
("User 3 Post", "That's my post!", 3)