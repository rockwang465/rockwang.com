CREATE DATABASE IF NOT EXISTS senseguard ;
CREATE DATABASE IF NOT EXISTS uums;
CREATE USER IF NOT EXISTS 'senseguard'@'%' IDENTIFIED WITH mysql_native_password BY 'UVlY88m9suHLsthK';
GRANT ALL PRIVILEGES ON senseguard.* TO 'senseguard'@'%';
GRANT ALL PRIVILEGES ON uums.* TO 'senseguard'@'%';
FLUSH PRIVILEGES;
