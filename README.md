<!--
title: MQTT
weight: 4705
-->
# MQTT
This activity allows you to send WS messages.

## Installation

### Flogo CLI
```bash
flogo install github.com/ultrafab/ws_flogo
```

### Flogo WEB
```bash
flogo install https://github.com/ultrafab/ws_flogo
```

## Configuration

### Settings:
| Name         | Type   | Description
| :---         | :---   | :---
| server       | string | The server URL - ***REQUIRED***

### Input:

| Name        | Type   | Description
| :---        | :---   | :---
| message     | string | The message to send  
| destination | string | The destination

### Output:

| Name  | Type   | Description
| :---  | :---   | :---
| data  | string | The data recieved
