# DDNS-GO

The shitty cloudflare ddns client i made for myself to use, painful stuff as a typescript writer
![929](https://github.com/mio9/ddns/assets/11581624/3ac55ae3-f408-4e69-a959-e4f213cc40cf)


## Commands

- `list zones` - List your zones
- `list records [zoneID]` - list your records


## Setup
1. Go create your cloudflare API key
2. paste that into your config cloudflare api-key
3. Run a `list zones`
4. Run a `list record [zoneID]` with the zone ID you got
5. Shove both zone ID and record ID into your your config `records`
6. Set correct name for your subdomain / domain name
7. Done

> Extremely tedious to setup right? i know lol but that's the thing of later :)

## Contribution

This project features extremely shitty go code. Pls feel free to fork this out and make some PR, my brain is too shitty to write some good go code as a 9-years JS kid
