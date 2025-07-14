# 🕵️‍♂️ Go Network Scan

A simple CLI and web-based tool to scan a subnet for available ports on reachable devices.

## 📦 Features

- 🔍 Scans a subnet for IPs with open ports
- 🌐 Web interface for user-friendly scanning
- 💻 CLI mode for quick terminal use

<img width="2560" height="1600" alt="localhost_8009_scan(Nest Hub Max)" src="https://github.com/user-attachments/assets/72db26b8-76c4-4d84-8aea-1177aadca983" />

## 🚀 Usage

### CLI Mode

Run with 2 arguments: `<port> <subnet>`

```bash
go run main.go 8000 1
```
This will scan 192.168.1.0/24 for anything broadcasting on port 8000.

### Web Mode
Start the server:

```bash
go run main.go
```
Then open your browser and navigate to:
http://localhost:8009

Use the form to scan a given port and subnet (e.g., port 8000, subnet 1).

## 🧠 Examples

CLI Output:

```
Running in CLI mode...
Valid addresses:
192.168.1.12
192.168.1.21
```

Web Output:

Results are shown as a list of clickable IP addresses.

## 🔧 Configuration

|Parameter|Description|
|---|---|
|port|Port to check (e.g. 8000)|
|subnet|Subnet octet (e.g. 1 for 192.168.1.x)|

## 📜 License

MIT — use freely, contribute gladly!
```
Let me know if you want badges, contribution guidelines, or Docker instructions added.
```
