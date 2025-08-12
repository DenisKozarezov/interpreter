<div align="center">

  [![en](https://img.shields.io/badge/lang-en-green.svg)](https://github.com/DenisKozarezov/interpreter/blob/master/README.md)
  [![ru](https://img.shields.io/badge/lang-ru-red.svg)](https://github.com/DenisKozarezov/interpreter/blob/master/README-ru.md)

  [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/DenisKozarezov/interpreter.svg)](https://github.com/DenisKozarezov)

  <h1>Интерпретатор скриптового языка программирования</h1>
</div>

## Описание

> [!WARNING]
> Данный репозиторий является пет-проектом и преследует исключительно академический интерес. Разрабатываемый 
> язык программирования (т.н. **guest language**) не является существующим языком. Целями проекта выступают
> изучение темы распознавания текста и его последующие лексинг, парсинг и выполнение абстрактного 
> синтаксического дерева (AST).

> [!NOTE]
> 📚 Приложенная к коду **[документация]()**

## Установка

### Установка `ipret` в Linux <img src="https://logo.svgcdn.com/d/linux-original.png" width=25 height=25>

#### Скачивание готового бинарного файла `ipret` через wget с *GitHub*

1. Откройте раздел [Releases](https://github.com/DenisKozarezov/interpreter/releases) и выберете последнюю версию утилиты `ipret`.
2. Скачайте бинарный файл для платформы Linux с помощью команды `wget` для требуемой архитектуры:
```shell
wget https://github.com/DenisKozarezov/interpreter/releases/latest/ipret-linux-arm64
```
3. Перенесите скачанный файл в исполняемую директорию текущего пользователя:
```shell
sudo mv ./ipret /usr/local/bin
```
4. Проверьте работоспособность утилиты `ipret` с помощью команды:
```shell
ipret --version
```

### Установка `ipret` в Windows <img src="https://logo.svgcdn.com/l/microsoft-windows-icon.png" width=25 height=25>

#### Скачивание готового бинарного файла `ipret` напрямую с *GitHub*

1. Откройте раздел [Releases](https://github.com/DenisKozarezov/interpreter/releases) и выберете последнюю версию утилиты `ipret`.
2. Скачайте бинарный файл для платформы Windows для требуемой архитектуры:
- `ipret-windows-x64.exe`
- `ipret-windows-x86_64.exe`

3. Перенесите скачанный файл в исполняемую директорию текущего пользователя:
```shell
sudo mv ./ipret /usr/local/bin
```
4. Проверьте работоспособность утилиты `ipret` с помощью команды:
```shell
ipret --version
```

## Применение

Для вывода подсказки применяется команда `--help` (`-h`):
```shell
# Использование --help в корневой команде
ipret --help

# Использование --help в команде запуска интерпрератора
ipret run --help

# Общий вид команды --help
ipret <comand> <subcommand> --help
```

```shell
ipret run -f ./someFile.txt
```