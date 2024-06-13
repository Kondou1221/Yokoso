#使うDockerイメージ
FROM mysql:8.0-debian

#ポートを開ける
EXPOSE 3306

ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password
ENV MYSQL_DATABASE=yokoso
ENV MYSQL_ROOT_PASSWORD=root

#MySQL設定ファイルをイメージ内にコピー
COPY ./my.cnf /etc/mysql/conf.d/my.cnf

#docker runに実行される
CMD ["mysqld"]