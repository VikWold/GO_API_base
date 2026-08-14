# Go base project

This is a project for a simple api for the go backend language. This is to be used to quickly set up a project without the hassel to deal with the setup parts, making sure everything to start the server goes correct.

## 🚀 Quick Setup — Search & Replace

To configure the project for your use, use your editor's **Search and Replace** to find and replace each keyword below across the **entire workspace**.

| Search for | Replace with | Description |
|---|---|---|
| `CHANGE_APP_NAME` | Your app/service name | Docker container name in `docker-compose.yaml` |
| `CHANGE_USERNAME` | Your PostgreSQL username | Database user in `.env` |
| `CHANGE_PASSWORD` | Your PostgreSQL password | Database password in `.env` |
| `CHANGE_DATABASE_NAME` | Your database name | Database name in `.env` |
| `CHANGE_NAME_DB_DSN` | Your chosen DSN env var name | The env var name used for the DB connection string — used in `docker-compose.yaml`, `Makefile`, and `config.go` |
| `go_base_project` | Your Go module name | Module name in `go.mod` — **also update all import paths** |

> ⚠️ After replacing `go_base_project`, make sure all `import` statements across the project are updated to use your new module name. Most editors will do this automatically via search and replace.

---

Start by changing the **module** name in the **go.mod** file

## Files and what you need

As this project it made it includes a simple handler and datalayer to deal with users. Here you can easily insert a user with a password, that gets hashed. Not all of this is nessesary for you, and can in that case be removed. Under here you find the files/folders you can remove for this reason.

- Auth package.
	- This packages is were we have our password hasher, and check to see if the password that comes in are correct
	- Only needed for a safe authentication system as of now, or for users to verify themself to the api

- Users datalayer
	- This is a file located in the internal/data folder. This is were you want to place your datalayer, or what is talking with the database directly.

- User Handlerlayer
	- This is the file located in the cmd folder. Here you should make sure the names of the handlers reflect what is in the routes file, as that is how the request gets sendt correctly.
	- Make sure that the you have diffrent structures for the returns from the api, and the password is not send here aswell if you are taking this part in use.

- Rest package
	- If you find yourself needing more helper functions or returns for errors. Then you should add those in the rest package. it is already placed some there that will help you get started, and to make this work out of the box

- Migrations
	- The migrations you can find here is made for the Users table. This folder is were you should put the migrations in. Make sure to keep the numbers after eachother to you can more easily run them in a sequence.
	- This project as of now is buildt on PostgreSQL, so if you want to swap database, you will need to change the migrations and also the datafiles that is taken in use as of now.