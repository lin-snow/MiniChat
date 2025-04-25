CREATE TABLE chat_messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    room_number VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    cmd VARCHAR(255) NOT NULL,
    payload TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat_rooms (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_chat_rooms (
    user_id INT NOT NULL,
    chat_room_id INT NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, chat_room_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(id)
);


