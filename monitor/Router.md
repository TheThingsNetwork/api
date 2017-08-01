# API Reference

## .router.Router

The Router service provides pure network functionality

### Methods

#### `Activate`

Gateway requests device activation

- Request: [`DeviceActivationRequest`](#routerdeviceactivationrequest)
- Response: [`DeviceActivationResponse`](#routerdeviceactivationresponse)

#### `GatewayStatus`

Gateway streams status messages to Router

- Request: stream of [`Status`](#gatewaystatus)
- Response: [`Empty`](#googleprotobufempty)

#### `Subscribe`

Gateway subscribes to downlink messages from Router
It is possible to open multiple subscriptions (but not recommended).
If you do this, you are responsible for de-duplication of downlink messages.

- Request: [`SubscribeRequest`](#routersubscriberequest)
- Response: stream of [`DownlinkMessage`](#routerdownlinkmessage)

#### `Uplink`

Gateway streams uplink messages to Router

- Request: stream of [`UplinkMessage`](#routeruplinkmessage)
- Response: [`Empty`](#googleprotobufempty)

## Messages

### `LocationMetadata`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `accuracy` | `int32` | the accuracy of the location (meters) |
| `altitude` | `int32` | the altitude (meters), where 0 is the mean sea level |
| `latitude` | `float` | the north–south position (degrees; -90 to +90), where 0 is the equator, north pole is positive, south pole is negative |
| `longitude` | `float` | the east-west position (degrees; -180 to +180), where 0 is the Prime Meridian (Greenwich), east is positive , west is negative |
| `source` | [`LocationSource`](#gatewaylocationmetadatalocationsource) |  |
| `time` | `int64` | Time (unix nanoseconds) |

### `RxMetadata`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `antennas` | _repeated_ [`Antenna`](#gatewayrxmetadataantenna) |  |
| `channel` | `uint32` |  |
| `encrypted_time` | `bytes` | Encrypted time from the Gateway FPGA |
| `frequency` | `uint64` | Frequency in Hz |
| `gateway_id` | `string` |  |
| `gateway_trusted` | `bool` | Indicates whether the gateway is trusted. Components that are able to verify gateway trust MUST do so and set this value accordingly |
| `location` | [`LocationMetadata`](#gatewaylocationmetadata) |  |
| `rf_chain` | `uint32` |  |
| `rssi` | `float` | Received signal strength in dBm |
| `snr` | `float` | Signal-to-noise-ratio in dB |
| `time` | `int64` | Time in Unix nanoseconds |
| `timestamp` | `uint32` | Timestamp (uptime of LoRa module) in microseconds with rollover |

### `Antenna`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `antenna` | `uint32` |  |
| `channel` | `uint32` |  |
| `channel_rssi` | `float` | Received channel power in dBm |
| `encrypted_time` | `bytes` | Encrypted time from the Gateway FPGA |
| `frequency_offset` | `int64` | Frequency offset (Hz) |
| `rssi` | `float` | Received signal power in dBm |
| `rssi_standard_deviation` | `float` | Standard deviation of the RSSI |
| `snr` | `float` | Signal-to-noise-ratio in dB |

### `Status`

message Status represents a status update from a Gateway.

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `boot_time` | `int64` | Boot time in Unix nanoseconds |
| `bridge` | `string` | The value of Bridge is set by the Bridge |
| `contact_email` | `string` |  |
| `description` | `string` |  |
| `dsp` | `uint32` | Version of Gateway DSP software |
| `fpga` | `uint32` | Version of Gateway FPGA |
| `frequency_plan` | `string` | The gateway's frequency plan: one of EU_863_870, US_902_928, CN_779_787, EU_433, AU_915_928, CN_470_510, AS_923, AS_920_923, AS_923_925, KR_920_923 |
| `gateway_trusted` | `bool` | Indicates whether the gateway is trusted. Components that are able to verify gateway trust MUST do so and set this value accordingly |
| `hal` | `string` | Version of gateway driver (in X.X.X format) |
| `ip` | _repeated_ `string` |  |
| `l_pps` | `uint32` | Number of lost PPS pulses |
| `lm_nw` | `uint32` | Sequence number of the last packet received from the link testing mote |
| `lm_ok` | `uint32` | Total number of packets received from link testing mote, with CRC OK |
| `lm_st` | `uint32` | Sequence number of the first packet received from the link testing mote |
| `location` | [`LocationMetadata`](#gatewaylocationmetadata) |  |
| `messages` | _repeated_ `string` | debug or warning messages from the gateway |
| `os` | [`OSMetrics`](#gatewaystatusosmetrics) |  |
| `platform` | `string` |  |
| `router` | `string` | The value of Router is set by the Router |
| `rtt` | `uint32` | Round-trip time to the server in milliseconds |
| `rx_in` | `uint32` | Total number of received uplink packets since boot |
| `rx_ok` | `uint32` | Total number of successful (correct) uplink packets since boot |
| `time` | `int64` | Time in Unix nanoseconds |
| `timestamp` | `uint32` | Timestamp (uptime of gateway) in microseconds with rollover |
| `tx_in` | `uint32` | Total number of received downlink packets since boot |
| `tx_ok` | `uint32` | Total number of successfully sent downlink packets since boot |

### `OSMetrics`

Additional metrics from the operating system

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `cpu_percentage` | `float` |  |
| `load_1` | `float` |  |
| `load_15` | `float` |  |
| `load_5` | `float` |  |
| `memory_percentage` | `float` |  |
| `temperature` | `float` |  |

### `TxConfiguration`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `frequency` | `uint64` | Frequency in Hz |
| `frequency_deviation` | `uint32` | FSK frequency deviation in Hz |
| `polarization_inversion` | `bool` | LoRa polarization inversion (basically always true for messages from gateway to node) |
| `power` | `int32` | Transmit power in dBm |
| `rf_chain` | `uint32` |  |
| `timestamp` | `uint32` | Timestamp (uptime of LoRa module) in microseconds with rollover |

### `Empty`

A generic empty message that you can re-use to avoid defining duplicated
empty messages in your APIs.

### `ActivationMetadata`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `app_eui` | `bytes` |  |
| `cf_list` | [`CFList`](#lorawancflist) |  |
| `dev_addr` | `bytes` |  |
| `dev_eui` | `bytes` |  |
| `frequency_plan` | [`FrequencyPlan`](#lorawanfrequencyplan) |  |
| `nwk_s_key` | `bytes` |  |
| `rx1_dr_offset` | `uint32` |  |
| `rx2_dr` | `uint32` |  |
| `rx_delay` | `uint32` |  |

### `Message`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `join_accept_payload` | [`JoinAcceptPayload`](#lorawanjoinacceptpayload) |  |
| `join_request_payload` | [`JoinRequestPayload`](#lorawanjoinrequestpayload) |  |
| `m_hdr` | [`MHDR`](#lorawanmhdr) |  |
| `mac_payload` | [`MACPayload`](#lorawanmacpayload) |  |
| `mic` | `bytes` |  |

### `Metadata`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `bit_rate` | `uint32` | FSK bit rate in bit/s |
| `coding_rate` | `string` | LoRa coding rate |
| `data_rate` | `string` | The DevEUI is a unique, 8 byte identifier for the device. |
| `f_cnt` | `uint32` | Store the full 32 bit FCnt (deprecated; do not use) |
| `frequency_plan` | [`FrequencyPlan`](#lorawanfrequencyplan) |  |
| `modulation` | [`Modulation`](#lorawanmodulation) | The AppEUI is a unique, 8 byte identifier for the application a device belongs to. |

### `TxConfiguration`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `bit_rate` | `uint32` | The AppID is a unique identifier for the application a device belongs to. It can contain lowercase letters, numbers, - and _. |
| `coding_rate` | `string` | The DevID is a unique identifier for the device. It can contain lowercase letters, numbers, - and _. |
| `data_rate` | `string` | The DevEUI is a unique, 8 byte identifier for the device. |
| `f_cnt` | `uint32` | The DevAddr is a dynamic, 4 byte session address for the device. |
| `modulation` | [`Modulation`](#lorawanmodulation) | The AppEUI is a unique, 8 byte identifier for the application a device belongs to. |

### `ActivationMetadata`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `lorawan` | [`ActivationMetadata`](#lorawanactivationmetadata) |  |

### `Message`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `lorawan` | [`Message`](#lorawanmessage) |  |

### `RxMetadata`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `lorawan` | [`Metadata`](#lorawanmetadata) |  |

### `TxConfiguration`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `lorawan` | [`TxConfiguration`](#lorawantxconfiguration) |  |

### `DeviceActivationRequest`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `activation_metadata` | [`ActivationMetadata`](#protocolactivationmetadata) |  |
| `app_eui` | `bytes` |  |
| `dev_eui` | `bytes` |  |
| `gateway_metadata` | [`RxMetadata`](#gatewayrxmetadata) |  |
| `message` | [`Message`](#protocolmessage) |  |
| `payload` | `bytes` |  |
| `protocol_metadata` | [`RxMetadata`](#protocolrxmetadata) |  |
| `trace` | [`Trace`](#tracetrace) |  |

### `DeviceActivationResponse`

### `DownlinkMessage`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `gateway_configuration` | [`TxConfiguration`](#gatewaytxconfiguration) |  |
| `message` | [`Message`](#protocolmessage) |  |
| `payload` | `bytes` |  |
| `protocol_configuration` | [`TxConfiguration`](#protocoltxconfiguration) |  |
| `trace` | [`Trace`](#tracetrace) |  |

### `SubscribeRequest`

### `UplinkMessage`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `gateway_metadata` | [`RxMetadata`](#gatewayrxmetadata) |  |
| `message` | [`Message`](#protocolmessage) |  |
| `payload` | `bytes` |  |
| `protocol_metadata` | [`RxMetadata`](#protocolrxmetadata) |  |
| `trace` | [`Trace`](#tracetrace) |  |

### `Trace`

Trace information

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `event` | `string` | Short event name |
| `id` | `string` | Generated ID |
| `metadata` | _repeated_ [`MetadataEntry`](#tracetracemetadataentry) | metadata for the event |
| `parents` | _repeated_ [`Trace`](#tracetrace) | Parents of the event |
| `service_id` | `string` | The ID of the component |
| `service_name` | `string` | The name of the component (router/broker/handler) |
| `time` | `int64` | Time in Unix nanoseconds |

### `MetadataEntry`

| **Name** | **Type** | **Description** |
| -------- | -------- | --------------- |
| `key` | `string` |  |
| `value` | `string` |  |

## Enums

### `LocationSource`

| **Name** | **Description** |
| -------- | --------------- |
| `CONFIG` | The location is fixed by configuration |
| `GPS` | The location is determined by GPS |
| `IP_GEOLOCATION` | The location is estimated with IP Geolocation |
| `REGISTRY` | The location is set in and updated from a registry |
| `UNKNOWN` | The source of the location is not known or not set |

### `FrequencyPlan`

| **Name** | **Description** |
| -------- | --------------- |
| `AS_920_923` |  |
| `AS_923` |  |
| `AS_923_925` |  |
| `AU_915_928` |  |
| `CN_470_510` |  |
| `CN_779_787` |  |
| `EU_433` |  |
| `EU_863_870` |  |
| `IN_865_867` |  |
| `KR_920_923` |  |
| `US_902_928` |  |

### `Modulation`

| **Name** | **Description** |
| -------- | --------------- |
| `FSK` |  |
| `LORA` |  |
