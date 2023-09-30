buildMysql:
	docker build -t mysql -f ./build/mysql.dockerfile .

runMysql:
	docker run -p 3306:3306 -d mysql

buildMongodb:
	docker build -t mongodb -f ./build/mongodb.dockerfile .

runMongodb:
	docker run -p 27017:27017 -d --name mongodb mongodb

