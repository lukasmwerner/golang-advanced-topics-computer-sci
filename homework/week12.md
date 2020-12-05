# Week 12

Plan: The plan is to start on a sample project. This sample project is porting adafruit-io to golang because the client lib hasnt been updated in 4 years. (go 1.2 we are now on go 1.15) The project will be on [lwerner-lshigh/io-client-go](https://github.com/lwerner-lshigh/io-client-go) so i can issue a pull request for the main project so i can contribute to open source!

Status: Begun porting to the v2 HTTP api, and have begun writing a MQTT client as well as much of the HTTP api is already implemented and only requires a small amount of changes to the v2 api. After writing some of the client types I realized that the client needed to have datatypes that were independent of the client type (HTTP vs MQTT) as they are currently intertwined.