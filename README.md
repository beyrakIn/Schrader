# Schrader C&C Framework

![Schrader](images/schrader.jpg)

### What is Schrader?

Schrader is a command&control server for the Schrader botnet. It is written in Go and C#. It is designed to be easy to
use and easy to deploy.

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

### Acknowledgements

> This interface is not executes command on the victim machine. It is just sends a command to the server then server
> sends the command to the victim machine.