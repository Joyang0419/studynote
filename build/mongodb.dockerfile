FROM mongo:latest

# 设置环境变量以配置root用户名和密码
ENV MONGO_INITDB_ROOT_USERNAME=root
ENV MONGO_INITDB_ROOT_PASSWORD=root

# Expose the MongoDB port
EXPOSE 27017