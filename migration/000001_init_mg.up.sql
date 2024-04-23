CREATE TABLE client (
    ID INT AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    surname VARCHAR(30) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    number VARCHAR(20) NOT NULL UNIQUE,
    hash_name VARCHAR(30) DEFAULT NULL,
    hash_surname VARCHAR(30) DEFAULT NULL,
    hash_email VARCHAR(50) DEFAULT NULL,
    hash_number VARCHAR(20) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (ID)
);

CREATE INDEX idx_client_id ON client (ID);
CREATE INDEX idx_client_deleted_at ON client (deleted_at);

INSERT INTO client (name, surname, email, number) VALUES
('John1', 'Doe1', 'john1@doe.test', '100000001'),
('John2', 'Doe2', 'john2@doe.test', '100000002'),
('John3', 'Doe3', 'john3@doe.test', '100000003'),
('John4', 'Doe4', 'john4@doe.test', '100000004'),
('John5', 'Doe5', 'john5@doe.test', '100000005'),
('John6', 'Doe6', 'john6@doe.test', '100000006'),
('John7', 'Doe7', 'john7@doe.test', '100000007'),
('John8', 'Doe8', 'john8@doe.test', '100000008'),
('John9', 'Doe9', 'john9@doe.test', '100000009'),
('John10', 'Doe10', 'john10@doe.test', '100000010'),
('John11', 'Doe11', 'john11@doe.test', '100000011'),
('John12', 'Doe12', 'john12@doe.test', '100000012'),
('John13', 'Doe13', 'john13@doe.test', '100000013'),
('John14', 'Doe14', 'john14@doe.test', '100000014'),
('John15', 'Doe15', 'john15@doe.test', '100000015'),
('John16', 'Doe16', 'john16@doe.test', '100000016'),
('John17', 'Doe17', 'john17@doe.test', '100000017'),
('John18', 'Doe18', 'john18@doe.test', '100000018'),
('John19', 'Doe19', 'john19@doe.test', '100000019'),
('John20', 'Doe20', 'john20@doe.test', '100000020');