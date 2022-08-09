# pxier_web
`pixer_web` is the web-api interface for [Pxier](https://github.com/JobberRT/pxier)

## Configuration
`echo.listen`: address which `pxier_web` will listen to
`echo.rate_limit`: `pxier_web` http rate limit(NOTICE: this will also infect `report` api)
`echo.max_get_number`: how many proxies client can get for one request

## API and Params
- `/require`
  - description: get proxies for client
  - params:
    - `num`: how many proxies you want. 
    - `provider`: which kinds of proxies you want. You can choose form these(case not sensitive):
      - cpl (From: https://github.com/clarketm/proxy-list)
      - tsx (From: https://github.com/TheSpeedX/PROXY-List)
      - str (From: https://github.com/ShiftyTR/Proxy-List)
      - ihuan (From: https://ip.ihuan.me/ti.html)
      - mix (all kinds)


## How to use
Recommend to use [Pxier](https://github.com/JobberRT/pxier) README's docker-compose file to deploy. Otherwise, you can compile and change the configuration and rename the `config.example.yaml` to `config.yaml`, then you can start the executable.