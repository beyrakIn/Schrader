# Schrader C&C Framework

![Schrader](images/schrader.jpg)

### What is Schrader?

Schrader is a command&control server. It is written in Go and C# (Agent). It is designed to be easy to
use and easy to distribute.

### Server

- [x] Server written in Go
- [x] Server can be compiled for Windows, Linux, and Mac
- [x] Web interface for managing bots
- [x] Web interface for managing tasks
- [x] Multi listener support

### Client

- [x] Client written in C# and Go
- [x] Client can be compiled for Windows, Linux, and Mac
- [x] Remote shell support
- [x] File upload and download support

### How to use

1. Download the latest release.
2. Change the `conf/config.json` file to your liking. (!Don't change json format)
3. Run the server.

### How it works

- When you start the server, you will have endpoints for the clients to connect to and you manage clients and tasks from
  the web interface.
- `ws://<server_ip>:<server_port>/` is the endpoint for the clients to connect to.
- `http://<server_ip>:<server_port>/client` is the endpoint you can see and manage the clients.
- `http://<server_ip>:<server_port>/client/<client_id>` is the endpoint you can manage the specific client.
- `http://<server_ip>:<server_port>/client/<client_id>?cmd=<command>` is the endpoint you can send commands to the
  client.
- You can also see all logs and victim information on terminal screen.

### Known Issues
- [ ] When you send a command which is they don't return anything to stdout, the client will be suspended.

### TODO

- [ ] Develop a terminal interface for the management
- [ ] Add authentication to the web interface
- [ ] Add logging for actions and errors
- [ ] Develop web interface for the client
  - Design the web interface
  - Show more information about the victim
  - Develop command execution function