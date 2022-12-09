Open-source framework for IoT edge computing
---
EdgeX Foundry is a vendor-neutral, highly flexible and scalable open-source framework hosted by The Linux Foundation. It enables developers to build apps that run at the edge, acting as a middleware between the things and the cloud.

It provides the components to develop microservices responsible for data acquisition, data analytics and cloud connectors, exposing a rich set of APIs to allow full control of the system and configuration. The reference architecture allows scaling out to thousands of devices and sensors.

This snap is a part of EdgeX Foundry and contains the following reference services:

**Core Services**
* Core Data
* Core Command
* Core Metadata
* Registry and Config Provider (Consul)
* Message Bus (Redis)

**Supporting Services**
* Scheduling
* Alerts & Notifications

**Security Services**
* API Gateway a.k.a. Reverse Proxy (Kong)
* Secret Store (Vault)

**Management Services**
* Management Service Agent - Deprecated

Note that not all the above services are enabled and started by default.

=============

**Further Reading**
* Getting started with snaps: https://docs.edgexfoundry.org/latest/getting-started/Ch-GettingStartedSnapUsers
* EdgeX documentation: https://docs.edgexfoundry.org
* Source code: https://github.com/edgexfoundry/edgex-go

====================

**Development Tool Snaps**
* EdgeX UI: https://snapcraft.io/edgex-ui
* EdgeX CLI: https://snapcraft.io/edgex-cli

**Other Supported EdgeX Snaps**
* eKuiper: https://snapcraft.io/edgex-ekuiper
* App Service Configurable: https://snapcraft.io/edgex-app-service-configurable
* App RFID LLRP Inventory: https://snapcraft.io/edgex-app-rfid-llrp-inventory
* Device Camera: https://snapcraft.io/edgex-device-camera
* Device GPIO: https://snapcraft.io/edgex-device-gpio
* Device Modbus: https://snapcraft.io/edgex-device-modbus
* Device MQTT: https://snapcraft.io/edgex-device-mqtt
* Device ONVIF Camera: https://snapcraft.io/edgex-device-onvif-camera
* Device REST: https://snapcraft.io/edgex-device-rest
* Device RFID LLRP: https://snapcraft.io/edgex-device-rfid-llrp
* Device SNMP: https://snapcraft.io/edgex-device-snmp
* Device USB Camera: https://snapcraft.io/edgex-device-usb-camera
* Device Virtual: https://snapcraft.io/edgex-device-virtual
* Device Grove: https://snapcraft.io/edgex-device-grove

**Additional References**
* Redis: https://github.com/redis/redis
* Consul: https://github.com/hashicorp/consul
* eKuiper: https://github.com/lf-edge/ekuiper
* Kong: https://github.com/Kong/kong
* Vault: https://github.com/hashicorp/vault
