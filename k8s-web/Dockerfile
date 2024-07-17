FROM nginx:latest

# 复制 Nginx 配置文件
COPY nginx.conf /etc/nginx/nginx.conf

# 复制构建的前端静态文件到 Nginx 的静态资源目录
COPY dist /usr/share/nginx/html

# 暴露 80 端口
EXPOSE 8081

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]