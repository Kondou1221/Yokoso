#使うDockerイメージ
FROM mysql:8.0-debian

#MySQL設定ファイルをイメージ内にコピー
COPY ./my.cnf /etc/mysql/conf.d/my.cnf

#日本語設定
RUN apt-get update && apt-get install -y locales \
    && sed -i -e 's/# \(ja_JP.UTF-8\)/\1/' /etc/locale.gen \
    && locale-gen \
    && update-locale LANG=ja_JP.UTF-8

#docker runに実行される
CMD ["mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci"]