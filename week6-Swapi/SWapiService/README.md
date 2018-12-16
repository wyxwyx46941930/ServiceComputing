# SWapi后端开发文档

## API 设计

- API 必须规范，请在项目文档部分给出一个简洁的说明，参考 github v3 或 v4 overview

- 选择 1-2 个 API 作为实例写在项目文档，文档格式标准，参考 github v3 或 v4

[参考](https://swapi.co/documentation#base)

## 资源来源与数据库支持

- 必须是真实数据，可以从公共网站获取
- 在项目文档中，**务必注明资源来源**

- 数据库 **只能使用 boltDB**，请 *不要使用 mysql 或 postgre 或 其他*



## API 和其他要求

- API 不能少于 6 个
- API root 能获取服务列表
- 部分资源必须授权服务（必须支持授权服务）
- 支持分页
- 支持 jsonp 输出 （仅 REST 服务）
- 提供项目文档首页的 URL。在文档中包含前后端安装指南。后端使用 go get 安装
- 队员必须提交一个相关的博客，或项目小结（请用markdown编写，存放在文档仓库中）



**认证技术提示**

- 为了方便实现用户认证，建议采用 JWT 产生 token 实现用户认证。
- 什么是 jwt？ 官网：https://jwt.io/ 中文支持：http://jwtio.com/
- 如何使用 jwt 签发用户 token ，用户验证 http://jwtio.com/introduction.html
- 各种语言工具 http://jwtio.com/index.html#debugger-io
- 案例研究：[基于 Token 的身份验证：JSON Web Token](https://ninghao.net/blog/2834)



## 实现功能

安装与使用（待测试）

- 各种依赖包

```bash
go get github.com/codegangsta/negroni
go get github.com/gorilla/mux
go get github.com/unrolled/render
go get github.com/boltdb/bolt
go get github.com/peterhellberg/swapi
go get github.com/SYSUServiceOnComputingCloud2018/SwapiService
```

- 运行

`cd src/github.com/SYSUServiceOnComputingCloud2018/SWapiService`

`go run main.go`

- 访问

`localhost:3000/api/[optional]`

-----

### 一、基本操作

1. base url（所以其他请求的url必须满足此前缀）

   `HOST/api`like`localhost:3000/api/`（服务端默认在本地3000端口监听，或者可以传入端口参数，详见main.go）

2. 请求频率限制： 10,000 API request per day

3. 支持Schema，请求`/api/<resource>/schema`会返回data包括的所有字段

4. query查询：`api/people/?search=r2`，使用case-insensitive partial matches返回所有符合条件的

5. 使用JSON格式输出

6. 支持分页输出所有资源



### 二、资源访问

首先在`SwapiService/boltdb`目录下使用`go run main.go`指令启动客户端

随后进入浏览器可按顺序访问下列指令：

#### 2.1 访问 root

指令：

`http://localhost:3000/api`

页面响应 :

```go
HTTP/1.0 200 OK
Content-Type: application/json
{
    "films": "http://localhost:3000/api/films/",
    "people": "http://localhost:3000/api/people/",
    "planets": "http://localhost:3000/api/planets/",
    "species": "http://localhost:3000/api/species/",
    "starships": "http://localhost:3000/api/starships/",
    "vehicles": "http://localhost:3000/api/vehicles/"
}
```

#### 2.2 访问 people 

1. 按id查询

- 指令 ：

`http://localhost:3000/api/people/1/`

- 页面响应 :

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "name": "Luke Skywalker",
  "height": "172",
  "mass": "77",
  "hair_color": "blond",
  "skin_color": "fair",
  "eye_color": "blue",
  "birth_year": "19BBY",
  "gender": "male",
  "homeworld": "https://swapi.co/api/planets/1/",
  "films": [
    "https://swapi.co/api/films/2/",
    "https://swapi.co/api/films/6/",
    "https://swapi.co/api/films/3/",
    "https://swapi.co/api/films/1/",
    "https://swapi.co/api/films/7/"
  ],
  "species": [
    "https://swapi.co/api/species/1/"
  ],
  "vehicles": [
    "https://swapi.co/api/vehicles/14/",
    "https://swapi.co/api/vehicles/30/"
  ],
  "starships": [
    "https://swapi.co/api/starships/12/",
    "https://swapi.co/api/starships/22/"
  ],
  "created": "2014-12-09T13:50:51.644000Z",
  "edited": "2014-12-20T21:17:56.891000Z",
  "url": "https://swapi.co/api/people/1/"
}
```



2. search

- 指令：

  `http://localhost:3000/api/people/?search=H`

- 页面响应：

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "count": 2,
  "next": "",
  "previous": "",
  "results": [
    {
      "name": "Han Solo",
      "height": "180",
      "mass": "80",
      "hair_color": "brown",
      "skin_color": "fair",
      "eye_color": "brown",
      "birth_year": "29BBY",
      "gender": "male",
      "homeworld": "https://swapi.co/api/planets/22/",
      "films": [
        "https://swapi.co/api/films/2/",
        "https://swapi.co/api/films/3/",
        "https://swapi.co/api/films/1/",
        "https://swapi.co/api/films/7/"
      ],
      "species": [
        "https://swapi.co/api/species/1/"
      ],
      "vehicles": [],
      "starships": [
        "https://swapi.co/api/starships/10/",
        "https://swapi.co/api/starships/22/"
      ],
      "created": "2014-12-10T16:49:14.582000Z",
      "edited": "2014-12-20T21:17:50.334000Z",
      "url": "https://swapi.co/api/people/14/"
    },
    {
      "name": "San Hill",
      "height": "191",
      "mass": "unknown",
      "hair_color": "none",
      "skin_color": "grey",
      "eye_color": "gold",
      "birth_year": "unknown",
      "gender": "male",
      "homeworld": "https://swapi.co/api/planets/57/",
      "films": [
        "https://swapi.co/api/films/5/"
      ],
      "species": [
        "https://swapi.co/api/species/34/"
      ],
      "vehicles": [],
      "starships": [],
      "created": "2014-12-20T17:58:17.049000Z",
      "edited": "2014-12-20T21:17:50.484000Z",
      "url": "https://swapi.co/api/people/77/"
    }
  ]
}
```



3. 返回所有的people

- 指令：

  `http://localhost:3000/api/people/`

- 支持分页输出（默认第一页）

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "count": 10,
  "next": "localhost:3000/api/people/?page=2",
  "previous": "",
  "results": [
  	...
   ]
}
```

- 指令：

  `http://localhost:3000/api/people/?page=9`

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "count": 7,
  "next": "",
  "previous": "localhost:3000/api/people/?page=8",
  "results": [
    ...
  ]
}
```



4. 查看schema

- 指令：

  `http://localhost:3000/api/people/schema`

- 页面响应：

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "required": [
    "name",
    "height",
    "mass",
    "hair_color",
    "skin_color",
    "eye_color",
    "birth_year",
    "gender",
    "homeworld",
    "films",
    "species",
    "vehicles",
    "starships",
    "url",
    "created",
    "edited"
  ],
  "title": "People",
  "properties": {
    "birth_year": {
      "description": "The birth year of this person. BBY (Before the Battle of Yavin) or ABY (After the Battle of Yavin).",
      "type": "string",
      "format": ""
    },
    "created": {
      "description": "The ISO 8601 date format of the time that this resource was created.",
      "type": "string",
      "format": "date-time"
    },
    "edited": {
      "description": "the ISO 8601 date format of the time that this resource was edited.",
      "type": "string",
      "format": "date-time"
    },
    "eye_color": {
      "description": "The eye color of this person.",
      "type": "string",
      "format": ""
    },
    "films": {
      "description": "An array of urls of film resources that this person has been in.",
      "type": "array",
      "format": ""
    },
    "gender": {
      "description": "The gender of this person (if known).",
      "type": "string",
      "format": ""
    },
    "hair_color": {
      "description": "The hair color of this person.",
      "type": "string",
      "format": ""
    },
    "height": {
      "description": "The height of this person in meters.",
      "type": "string",
      "format": ""
    },
    "homeworld": {
      "description": "The url of the planet resource that this person was born on.",
      "type": "string",
      "format": ""
    },
    "mass": {
      "description": "The mass of this person in kilograms.",
      "type": "string",
      "format": ""
    },
    "name": {
      "description": "The name of this person.",
      "type": "string",
      "format": ""
    },
    "skin_color": {
      "description": "The skin color of this person.",
      "type": "string",
      "format": ""
    },
    "species": {
      "description": "The url of the species resource that this person is.",
      "type": "array",
      "format": ""
    },
    "starships": {
      "description": "An array of starship resources that this person has piloted",
      "type": "array",
      "format": ""
    },
    "url": {
      "description": "The url of this resource",
      "type": "string",
      "format": "uri"
    },
    "vehicles": {
      "description": "An array of vehicle resources that this person has piloted",
      "type": "array",
      "format": ""
    }
  },
  "description": "A person within the Star Wars universe",
  "$schema": "http://json-schema.org/draft-04/schema",
  "type": "object"
}
```





#### 2.3 访问 planet

- 指令：

`http://localhost:3000/api/planets/1/`

- 页面响应 :

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "name": "Tatooine",
  "rotation_period": "23",
  "orbital_period": "304",
  "diameter": "10465",
  "climate": "arid",
  "gravity": "1 standard",
  "terrain": "desert",
  "surface_water": "1",
  "population": "200000",
  "residents": [
    "https://swapi.co/api/people/1/",
    "https://swapi.co/api/people/2/",
    "https://swapi.co/api/people/4/",
    "https://swapi.co/api/people/6/",
    "https://swapi.co/api/people/7/",
    "https://swapi.co/api/people/8/",
    "https://swapi.co/api/people/9/",
    "https://swapi.co/api/people/11/",
    "https://swapi.co/api/people/43/",
    "https://swapi.co/api/people/62/"
  ],
  "films": [
    "https://swapi.co/api/films/5/",
    "https://swapi.co/api/films/4/",
    "https://swapi.co/api/films/6/",
    "https://swapi.co/api/films/3/",
    "https://swapi.co/api/films/1/"
  ],
  "created": "2014-12-09T13:50:49.641000Z",
  "edited": "2014-12-21T20:48:04.175778Z",
  "url": "https://swapi.co/api/planets/1/"
}
```



#### 2.4 访问 film

- 指令

`http://localhost:3000/api/films/1/`

- 页面响应

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "title": "A New Hope",
  "episode_id": 4,
  "opening_crawl": "It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....",
  "director": "George Lucas",
  "producer": "Gary Kurtz, Rick McCallum",
  "characters": [
    "https://swapi.co/api/people/1/",
    "https://swapi.co/api/people/2/",
    "https://swapi.co/api/people/3/",
    "https://swapi.co/api/people/4/",
    "https://swapi.co/api/people/5/",
    "https://swapi.co/api/people/6/",
    "https://swapi.co/api/people/7/",
    "https://swapi.co/api/people/8/",
    "https://swapi.co/api/people/9/",
    "https://swapi.co/api/people/10/",
    "https://swapi.co/api/people/12/",
    "https://swapi.co/api/people/13/",
    "https://swapi.co/api/people/14/",
    "https://swapi.co/api/people/15/",
    "https://swapi.co/api/people/16/",
    "https://swapi.co/api/people/18/",
    "https://swapi.co/api/people/19/",
    "https://swapi.co/api/people/81/"
  ],
  "planets": [
    "https://swapi.co/api/planets/2/",
    "https://swapi.co/api/planets/3/",
    "https://swapi.co/api/planets/1/"
  ],
  "starships": [
    "https://swapi.co/api/starships/2/",
    "https://swapi.co/api/starships/3/",
    "https://swapi.co/api/starships/5/",
    "https://swapi.co/api/starships/9/",
    "https://swapi.co/api/starships/10/",
    "https://swapi.co/api/starships/11/",
    "https://swapi.co/api/starships/12/",
    "https://swapi.co/api/starships/13/"
  ],
  "vehicles": [
    "https://swapi.co/api/vehicles/4/",
    "https://swapi.co/api/vehicles/6/",
    "https://swapi.co/api/vehicles/7/",
    "https://swapi.co/api/vehicles/8/"
  ],
  "species": [
    "https://swapi.co/api/species/5/",
    "https://swapi.co/api/species/3/",
    "https://swapi.co/api/species/2/",
    "https://swapi.co/api/species/1/",
    "https://swapi.co/api/species/4/"
  ],
  "created": "2014-12-10T14:23:31.880000Z",
  "edited": "2015-04-11T09:46:52.774897Z",
  "url": "https://swapi.co/api/films/1/"
}
```



#### 2.5 访问 vehicle

- 指令

`http://localhost:3000/api/vehicles/4/`

- 页面响应

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "name": "Sand Crawler",
  "model": "Digger Crawler",
  "manufacturer": "Corellia Mining Corporation",
  "cost_in_credits": "150000",
  "length": "36.8",
  "max_atmosphering_speed": "30",
  "crew": "46",
  "passengers": "30",
  "cargo_capacity": "50000",
  "consumables": "2 months",
  "vehicle_class": "wheeled",
  "pilots": [],
  "films": [
    "https://swapi.co/api/films/5/",
    "https://swapi.co/api/films/1/"
  ],
  "created": "2014-12-10T15:36:25.724000Z",
  "edited": "2014-12-22T18:21:15.523587Z",
  "url": "https://swapi.co/api/vehicles/4/"
}
```



#### 2.6 访问 species

- 指令

`http://localhost:3000/api/species/1/`

- 页面响应

```
HTTP/1.0 200 OK
Content-Type: application/json
{
  "name": "Human",
  "classification": "mammal",
  "designation": "sentient",
  "average_height": "180",
  "skin_colors": "caucasian, black, asian, hispanic",
  "hair_colors": "blonde, brown, black, red",
  "eye_colors": "brown, blue, green, hazel, grey, amber",
  "average_lifespan": "120",
  "homeworld": "https://swapi.co/api/planets/9/",
  "language": "Galactic Basic",
  "people": [
    "https://swapi.co/api/people/1/",
    "https://swapi.co/api/people/4/",
    "https://swapi.co/api/people/5/",
    "https://swapi.co/api/people/6/",
    "https://swapi.co/api/people/7/",
    "https://swapi.co/api/people/9/",
    "https://swapi.co/api/people/10/",
    "https://swapi.co/api/people/11/",
    "https://swapi.co/api/people/12/",
    "https://swapi.co/api/people/14/",
    "https://swapi.co/api/people/18/",
    "https://swapi.co/api/people/19/",
    "https://swapi.co/api/people/21/",
    "https://swapi.co/api/people/22/",
    "https://swapi.co/api/people/25/",
    "https://swapi.co/api/people/26/",
    "https://swapi.co/api/people/28/",
    "https://swapi.co/api/people/29/",
    "https://swapi.co/api/people/32/",
    "https://swapi.co/api/people/34/",
    "https://swapi.co/api/people/43/",
    "https://swapi.co/api/people/51/",
    "https://swapi.co/api/people/60/",
    "https://swapi.co/api/people/61/",
    "https://swapi.co/api/people/62/",
    "https://swapi.co/api/people/66/",
    "https://swapi.co/api/people/67/",
    "https://swapi.co/api/people/68/",
    "https://swapi.co/api/people/69/",
    "https://swapi.co/api/people/74/",
    "https://swapi.co/api/people/81/",
    "https://swapi.co/api/people/84/",
    "https://swapi.co/api/people/85/",
    "https://swapi.co/api/people/86/",
    "https://swapi.co/api/people/35/"
  ],
  "films": [
    "https://swapi.co/api/films/2/",
    "https://swapi.co/api/films/7/",
    "https://swapi.co/api/films/5/",
    "https://swapi.co/api/films/4/",
    "https://swapi.co/api/films/6/",
    "https://swapi.co/api/films/3/",
    "https://swapi.co/api/films/1/"
  ],
  "created": "2014-12-10T13:52:11.567000Z",
  "edited": "2015-04-17T06:59:55.850671Z",
  "url": "https://swapi.co/api/species/1/"
}

```



#### 2.7 访问 starship

- 指令 :

  `http://localhost:3000/api/starships/3/`

- 页面响应 :

```json
HTTP/1.0 200 OK
Content-Type: application/json
{
  "name": "Star Destroyer",
  "model": "Imperial I-class Star Destroyer",
  "manufacturer": "Kuat Drive Yards",
  "cost_in_credits": "150000000",
  "length": "1,600",
  "max_atmosphering_speed": "975",
  "crew": "47060",
  "passengers": "0",
  "cargo_capacity": "36000000",
  "consumables": "2 years",
  "hyperdrive_rating": "2.0",
  "MGLT": "60",
  "starship_class": "Star Destroyer",
  "pilots": [],
  "films": [
    "https://swapi.co/api/films/2/",
    "https://swapi.co/api/films/3/",
    "https://swapi.co/api/films/1/"
  ],
  "created": "2014-12-10T15:08:19.848000Z",
  "edited": "2014-12-22T17:35:44.410941Z",
  "url": "https://swapi.co/api/starships/3/"
}
```



## 具体实现

1. 注册路由
2. 数据库访问
3. JSON格式和状态码输出

...