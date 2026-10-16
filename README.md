# 🔌 Hex Loader

**Кроссплатформенное GUI-приложение для прошивки Arduino-совместимых плат**

![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fyne](https://img.shields.io/badge/Fyne-v2.8-1F4E79?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-4B8BBE?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

[![Build & Release](https://github.com/DFVladislav/hex-loader/actions/workflows/build.yml/badge.svg)](https://github.com/DFVladislav/hex-loader/actions/workflows/build.yml)
[![Release](https://img.shields.io/github/v/release/DFVladislav/hex-loader?style=flat-square&color=blue)](https://github.com/DFVladislav/hex-loader/releases)

---

## 📖 О проекте

**Hex Loader** — простое десктопное приложение на **Go + Fyne** для прошивки Arduino-совместимых плат готовыми `.hex` файлами. Работает поверх `arduino-cli` и не требует установки Arduino IDE.

![Hex Loader UI](docs/screenshot-app.png)

---

## ✨ Возможности

- 📁 **Выбор `.hex` файла** через нативный диалог
- 🔍 **Автоопределение serial-портов** через `arduino-cli board list`
- 🎯 **Выбор платы по FQBN** — Uno, Nano, Mega, ESP32, ESP8266 и другие
- 📜 **Встроенный лог** процесса прошивки в реальном времени
- 🌙 **Тёмная тема** оформления
- 📦 **Готовые бинарники** под Windows, Linux и macOS
- 🚀 **Автоматическая сборка и релизы** через GitHub Actions

---

## 🚀 Быстрый старт

### 1. Установи `arduino-cli`

Windows (PowerShell):

```powershell
winget install --id ArduinoSA.CLI -e
```

macOS (Homebrew):

```bash
brew install arduino-cli
```

Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | sh
```

### 2. Установи ядро платы

Для Arduino Uno, Nano, Mega:

```bash
arduino-cli core update-index
arduino-cli core install arduino:avr
```

Для ESP32:

```bash
arduino-cli config init
arduino-cli config add board_manager.additional_urls https://espressif.github.io/arduino-esp32/package_esp32_index.json
arduino-cli core update-index
arduino-cli core install esp32:esp32
```

### 3. Скачай Hex Loader

Готовые бинарники лежат на странице [Releases](https://github.com/DFVladislav/hex-loader/releases):

- 🐧 Linux — `hex-loader-linux-x64`
- 🪟 Windows — `hex-loader-windows-x64.exe`
- 🍎 macOS — `hex-loader-macos-arm64`

На macOS и Linux после скачивания сделай файл исполняемым:

```bash
chmod +x hex-loader-*
```

### 4. Запусти и прошей

1. Открой приложение
2. Нажми **Select .hex** и выбери файл прошивки
3. Выбери **serial port** (кнопка **Refresh** обновит список)
4. Выбери **board (FQBN)**
5. Нажми **Flash** и следи за логом

---

## 🛠 Сборка из исходников

### Требования

- **Go 1.23+**
- **GCC** — нужен для CGO, потому что Fyne использует OpenGL
  - Windows: [MinGW-w64](https://github.com/niXman/mingw-builds-binaries/releases) или TDM-GCC
  - macOS: `xcode-select --install`
  - Linux: `sudo apt install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev`

### Команды

```bash
git clone https://github.com/DFVladislav/hex-loader.git
cd hex-loader
go mod download
go run .
```

Собрать бинарник:

```bash
go build -o hex-loader .
```

---

## ⚙️ CI/CD

Проект использует **GitHub Actions** для автоматической сборки под три операционные системы. Каждый пуш тега вида `v*` запускает пайплайн, который собирает бинарники параллельно под Linux, Windows и macOS, а затем публикует их в GitHub Release.

Схема пайплайна:

```
Build Linux     Build Windows     Build macOS
     \               |                /
      \              |               /
       -------- Publish Release ----
```

### Как выпустить релиз

```bash
git tag v1.0.0
git push origin v1.0.0
```

Через ~12 минут на странице [Releases](https://github.com/DFVladislav/hex-loader/releases) появится релиз с тремя бинарниками.

Скриншот успешного workflow:

<img width="1522" height="763" alt="изображение" src="https://github.com/user-attachments/assets/d58a8f08-4f03-468c-a6d8-c11a24780cd2" />


Скриншот опубликованного релиза:

<img width="1298" height="678" alt="изображение" src="https://github.com/user-attachments/assets/6614d72f-ab06-46a3-9b17-c493d0dcdefc" />

Скриншот самого приложения:

<img width="632" height="533" alt="изображение" src="https://github.com/user-attachments/assets/81dbb309-7daf-40d4-8ca5-f2ae8d71110c" />

---

## 📁 Структура проекта

```
hex-loader/
├── .github/
│   └── workflows/
│       └── build.yml
├── docs/
├── gui/
│   ├── flasher.go
│   └── ui.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## 🧰 Стек

- [Go](https://go.dev/) — язык программирования
- [Fyne](https://fyne.io/) — кроссплатформенный GUI-фреймворк
- [arduino-cli](https://arduino.github.io/arduino-cli/) — CLI для работы с Arduino
- [GitHub Actions](https://github.com/features/actions) — CI/CD

---

## 🤝 Вклад

Pull request'ы приветствуются. Для крупных изменений сначала открой issue — обсудим.

1. Fork репозитория
2. Создай ветку (`git checkout -b feature/amazing-feature`)
3. Закоммить (`git commit -m "Add amazing feature"`)
4. Запушь (`git push origin feature/amazing-feature`)
5. Открой Pull Request

---

## 📄 Лицензия

Распространяется под лицензией **MIT**. Подробности — в файле [LICENSE](LICENSE).

---

Сделано с ❤️ на Go и Fyne. ⭐ Поставь звезду, если проект оказался полезен.
