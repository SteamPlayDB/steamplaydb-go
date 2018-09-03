# SteamPlayDB.com Go

*Runs backround server application for SteamPlayDB.com*

# How To Install

1. Clone repo.
2. `cd` into repo.
3. Run `make` to setup and build project.


# How To Run

1. Cop `.env_example` to `.env`.
2. Read the *Server Enviroment Variables* section to see what settings are available and what the defaults are. Also make sure to set the `PUBLIC_DIR`
	to the absolute path of the directory you want.
3. Run `./bin/steamplaydb`.

# Server Enviroment Variables
*Context of variables to set in the system (docker container) or .env file.*

- `PUBLIC_DIR` = Directory of static files. *Default:* public
- `INDEX_FILE` = File to load if no file is indicated. *Default:* index.html
- `AUTOCERT_CACHE` = Directory of auto certificates to load for the server.
- `PORT` = Port to run the server on. *Default:* 80

# F.A.Q

*Coming soon...*
