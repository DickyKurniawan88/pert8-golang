-- PENTING: GANTI ANGKA 12345 DENGAN NPM KALIAN
CREATE DATABASE IF NOT EXISTS event_realm_12345;
USE event_realm_12345; 

CREATE TABLE IF NOT EXISTS events (
    id_event INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    organizer VARCHAR(255) NOT NULL,
    description TEXT,
    date DATETIME NOT NULL,
    location VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    capacity INT NOT NULL,
    image_url VARCHAR(255),
    status ENUM('upcoming', 'ongoing', 'completed', 'cancelled') DEFAULT 'upcoming',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO events (title, organizer, description, date, location, price, capacity, image_url, status) VALUES
('BlissCon 2025', 'Blissard Entertainment', 'The ultimate gaming experience for Blissard fans.', '2025-10-25 09:00:00', 'Anaheim Convention Center, California, USA', 3880000.00, 35000, 'placeholder.jpg', 'upcoming'),
('Tokyo Game Show 2025', 'Gamer''z Association', 'The world''s biggest gaming expo in Asia.', '2025-09-18 10:00:00', 'Makuhari Messe, Chiba, Japan', 2030000.00, 250000, 'https://events.nikkeibp.co.jp/tgs/2025/jp/exhibitor/ogp.png', 'upcoming'),
('Utaite Dream Festival', 'Utattemitayo Productions', 'A two-day festival featuring top Japanese utaite and virtual singers.', '2025-05-15 17:30:00', 'Tokyo Dome, Tokyo, Japan', 1850000.00, 55000, 'https://img.youtube.com/vi/QMGvSE9CyaY/maxresdefault.jpg', 'upcoming');

SELECT * FROM events;