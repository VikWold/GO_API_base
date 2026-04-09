# Go base project

This is a project for a simple api for the go backend language. This is to be used to quickly set up a project without the hassel to deal with the setup parts, making sure everything to start the server goes correct   

Start by changing the **module** name in the **go.mod** file

## What is missing here

As of now this don't include any form of docker compose files, as they have to be placed other placess aswell. This might be a fix down the line to make that part easier aswell.

## Files and what you need

As this project it made it includes a simple handler and datalayer to deal with users. Here you can easily insert a user with a password, that gets hashed. Not all of this is nessesary for you, and can in that case be removed. Under here you find the files/folders you can remove for this reason.

>‼️ Also make sure you fix the imports when you coppy over this file and have changed the **go.mod**, as they are made for this root project as of now. In general you should make your go project, and then past in this base project for a better experience

- Auth package.
	- This packages is were we have our password hasher, and check to see if the password that comes in are correct
	- Only needed for a safe loging system as of now, or for users to verify themself to the api

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