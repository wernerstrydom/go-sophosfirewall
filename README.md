# Sophos Firewall Go Client

## Overview

This is a Go client library for accessing the Sophos Firewall API. It is a work in progress and is not yet complete.

## Installation

```bash
go get github.com/wernerstrydom/go-sophosfirewall
```

## Usage

```go
import "github.com/wernerstrydom/go-sophosfirewall/sophosfirewall"
```

Construct a new Sophos Firewall client, then use the various services on the client to
access different parts of the Sophos Firewall API. For example:

```go
client := sophosfirewall.New("admin", "password", "https://firewall.example.com:4444/webconsole/APIController")
```
