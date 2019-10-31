<h1 align="center">Welcome to taskhelper 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000" />
  <a href="https://twitter.com/cgardner" target="_blank">
    <img alt="Twitter: cgardner" src="https://img.shields.io/twitter/follow/cgardner.svg?style=social" />
  </a>
</p>

> Templating tool for TaskWarrior

## Install

```sh
go get github.com/cgardner/taskhelper
```

## Usage

Create the following config file in \$HOME/.config/taskhelper/taskhelper.yaml

```yaml
shortname:
  add:
    - +tag1
    - +tag2
    - project:"My Project"
    - due:tomorrow
  report:
    - project:"My Project"
    - next
```

With this config file, you can create a task using

```sh
taskhelper shortname Make something awesome
```

You can also view a report you've defined with

```sh
taskhelper shortname
```

## Author

👤 **Craig Gardner**

- Twitter: [@cgardner](https://twitter.com/cgardner)
- Github: [@cgardner](https://github.com/cgardner)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
