# WebSocket Echo App using Gin Framework

This is a simple example of a WebSocket echo app using the Gin framework in Go. The app serves an HTML page that allows users to send messages to the server over a WebSocket connection, and the server echoes those messages back to the client.

## Requirements

- Go 1.16 or higher
- Gin framework
- Gorilla WebSocket package

## Installation

1. Clone the repository:

  git clone https://github.com/royyanwibisono/testwebsocket.git


2. Install the Gin framework:

  go get -u github.com/gin-gonic/gin


## Usage

1. Start the server:

  go run main.go

2. Open `http://localhost:7000` in your browser.

3. Type a message in the text input and click "Send". You should see your message echoed back on the page.

4. Open the browser console to see log messages for WebSocket connections and messages received and sent.

## Credits

- Gin framework: https://github.com/gin-gonic/gin
- Gorilla WebSocket package: https://github.com/gorilla/websocket

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
