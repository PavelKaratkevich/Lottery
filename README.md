THIS EXAMPLE IS SHAMELESSLY COPY & PASTED FROM https://www.makeareadme.com/

Dan - I would absolutely recommend having a README.md file in all you projects as they provide additional context to the
reader/reviewer and just generally look good.

# Lottery

Lottery is a simple backend of a web-app to run national lotteries.

## Setup

You would need a MySQL database to run this app.

Dan - for convenience purposes you could also add .env.example file and include the environmental variables your app needs to operate.

```bash
docker pull mysql/mysql-server:latest
docker run \
--detach \
--name=mysql_dev \
--env="MYSQL_ROOT_PASSWORD=some_difficult_password" \
--publish 3306:3306 \
mysql
```
Then simply run the app or build and run it.
```bash
go run main.go
```


## Usage
Dan - Here you could include some simple curl commands to illustrate that your app works as intended.

```bash
curl -X POST
```

[comment]: <> (TODO - add curl commands)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
Dan - not obligatory but good to have.

[MIT](https://choosealicense.com/licenses/mit/)