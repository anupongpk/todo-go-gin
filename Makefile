maria:
	@docker run -p 127.0.0.1:3307:3306 --name some-mariadb \
	-e MARIADB_ROOT_PASSWORD=123456 \
	-e MARIADB_DATABASE=myapp \
	-d mariadb:latest
