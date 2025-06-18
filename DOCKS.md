## Инициализируем проект
```go mod init vetstore2025```

## Создаем репозиторий в гитхабе и добавляем проект в гит 

```git init```

```git remote add origin https://github.com/username/vetstore2025.git```

```git push -u origin master```

## Добавляем библиотеку gorm для работы с базой данных

``` go get -u gorm.io/gorm```

## Скачиваем драйвер базы данных, например sqlite

```go get -u gorm.io/driver/sqlite```

## Установим Fyne для создания интерфейса

```go get fyne.io/fyne/v2@latest```

