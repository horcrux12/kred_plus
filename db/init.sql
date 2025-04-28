-- membuat native password agar lebih secure tanpa perlu public key retrieval dan compatible dengan banyak driver
ALTER USER 'user'@'%' IDENTIFIED WITH mysql_native_password BY 'secretuser';
FLUSH PRIVILEGES;