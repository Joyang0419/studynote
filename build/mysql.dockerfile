# Use the MySQL 5.7 image as the base image
FROM amd64/mysql:5.7

# Optional: Set the MySQL root password
# Note: It's not recommended to use the root user/password in production
ENV MYSQL_ROOT_PASSWORD root

# Optional: Create a database, user, and password
ENV MYSQL_DATABASE testdb