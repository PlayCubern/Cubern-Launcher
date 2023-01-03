
# Cubern Launcher

Cubern is an online sandbox where the creativity never ends. Users can customize their avatars as they want. Players can also make there own games using our editor built in Godot game Engine! But, to run these a client is required which also requires a URL Protocol! 

## Authors

- [@OoIks](https://www.github.com/OoIks)


## Run Locally

In order to run this locally, we are required to install go programming language on our computer! 

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd path/to/cloned/repository
```

Compile application:

```bash
  go build
```



## Features

- Install client and unzip it from server
- Install new updates with ease
- Install custom url protocols
- Install required things


## API Reference

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.

