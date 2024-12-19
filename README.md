# pxconv — конвертер экранных единиц (dp, sp, px, inch, mm)

`pxconv` — пакет на Go для работы с экранными единицами измерения
(`dp`, `sp`, `px`, `inch`, `mm`) с учётом плотности экрана.

## Особенности

- Поддержка основных единиц: `dp`, `sp`, `px`, `inch`, `mm`.
- Настраиваемая плотность экрана (`PxPerDp`, `PxPerSp`).
- Конвертация между `dp`, `sp`, `px`, `inch`, `mm`.
- Обработка некорректных значений (по умолчанию используется значение `1`).

## Установка

```zsh
go get github.com/MiCkEyZzZ/pxconv
```

## Пример использования

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv"
)

func main() {
	m := pxconv.Metric{PxPerDp: 2, PxPerSp: 1.5, Dpi: 96}
	fmt.Println(m.DpToPx(pxconv.Dp(10))) // Конвертация dp в px
}
```

## API

Пакет предоставляет следующие основные методы для преобразования между dp, sp, px, inch и mm.
Подробности в [документации](https://pkg.go.dev/github.com/MiCkEyZzZ/pxconv).

## Лицензия

Этот пакет распространяется под лицензией MIT. Полный текст лицензии доступен в
файле [ЛИЦЕНЗИЯ](./LICENSE).
