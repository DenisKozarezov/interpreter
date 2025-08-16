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
> изучение темы распознавания исходного кода и его последующие лексинг, парсинг и выполнение абстрактного 
> синтаксического дерева (AST).

## Установка

### Установка `ipret` в Linux <img src="https://logo.svgcdn.com/d/linux-original.png" width=25 height=25>

#### Скачивание готового бинарного файла `ipret` через wget с *GitHub*

1. Откройте раздел [Releases](https://github.com/DenisKozarezov/interpreter/releases) и выберете нужную версию утилиты `ipret`.
2. Скачайте архив для платформы Linux с помощью команды `wget` для требуемой архитектуры:
```shell
wget https://github.com/DenisKozarezov/interpreter/releases/latest/ipret-linux-arm64.tar.gz
```
3. Распакуйте архив и перенесите скачанный бинарный файл в исполняемую директорию текущего пользователя:
```shell
mkdir ipret-temp && tar -xvzf ipret-linux-amd64.tar.gz -C ipret-temp
sudo mv ipret-temp/ipret /usr/local/bin
sudo rm -rf ipret-temp
```
4. Проверьте работоспособность утилиты `ipret` с помощью команды:
```shell
ipret --version
```

### Установка `ipret` в Windows <img src="https://logo.svgcdn.com/l/microsoft-windows-icon.png" width=25 height=25>

#### Скачивание готового бинарного файла `ipret` напрямую с *GitHub*

1. Откройте раздел [Releases](https://github.com/DenisKozarezov/interpreter/releases) и выберете нужную версию утилиты `ipret`.
2. Скачайте архив для платформы Windows для требуемой архитектуры:
- `ipret-windows-amd64.zip`
- `ipret-windows-386.zip`
3. Распакуйте zip-архив при помощи любого доступного архиватора.
4. Проверьте работоспособность утилиты `ipret` с помощью команды:
```shell
ipret.exe --version
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
# Запустить программу из текстового файла
ipret run -f ./someFile.txt

# Запустить программу из текстового файла и показать результаты бенчмарка
ipret run --filename=someFile.irt --bench
```