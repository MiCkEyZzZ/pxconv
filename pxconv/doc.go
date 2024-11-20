// Пакет pxconv предоставляет инструменты для работы с единицами измерения
// пользовательского интерфейса: независимыми от устройства пикселями (dp), единицами для шрифтов (sp)
// и их преобразованием в пиксели (px) с учетом плотности экрана.
//
// # Типы Dp и Sp
//
// Dp представляет собой независимые от устройства пиксели, которые используются для задания размеров
// элементов интерфейса, таких как кнопки, поля и отступы. Эти единицы гарантируют адаптивность
// интерфейса на устройствах с разной плотностью пикселей.
//
// Sp представляет собой единицы для шрифтов, которые, помимо плотности экрана, учитывают
// предпочтения пользователя, такие как настройки размера шрифта. Это делает sp подходящими
// для задания текстовых размеров.
//
// Оба типа используют float32 для хранения значений.
//
// Пример:
//
//	var size pxconv.Dp = 16    // Размер в dp
//	var textSize pxconv.Sp = 14 // Размер текста в sp
//
// # Тип Metric
//
// Metric предоставляет методы для конвертации значений между dp, sp и пикселями (px).
// Основные поля структуры Metric:
//   - PxPerDp: количество пикселей, соответствующее одному dp.
//   - PxPerSp: количество пикселей, соответствующее одному sp.
//
// # Создание Metric
//
// Вы можете создать экземпляр Metric двумя способами:
// 1. **С помощью конструктора `NewMetric`:**
//
// Конструктор проверяет входные значения. Если они равны 0 или отрицательны, они заменяются на 1.
//
// Пример:
//
//	metric := pxconv.NewMetric(2.5, 3.0) // Плотность: 2.5 px/dp и 3.0 px/sp.
//
// 2. **Создание вручную:**
//
// Вы можете объявить структуру Metric вручную, указав значения полей.
// Методы структуры Metric автоматически заменяют некорректные значения полей (0 или отрицательные)
// на значение по умолчанию 1, что делает ручное создание безопасным.
//
// Пример:
//
//	metric := pxconv.Metric{PxPerDp: 0, PxPerSp: -1.5}
//	px := metric.DpToPx(10) // Результат: 10 (PxPerDp заменяется на 1 по умолчанию).
//
// # Основные методы
//
// Методы конверсии:
//   - DpToPx: Конвертирует dp в пиксели (px), округляя до ближайшего целого.
//   - SpToPx: Конвертирует sp в пиксели (px), округляя до ближайшего целого.
//   - PxToDp: Конвертирует пиксели (px) в dp.
//   - PxToSp: Конвертирует пиксели (px) в sp.
//   - DpToSp: Конвертирует dp в sp.
//   - SpToDp: Конвертирует sp в dp.
//
// Пример работы с методами:
//
//	metric := pxconv.NewMetric(2.0, 1.5) // Плотность экрана
//	pxFromDp := metric.DpToPx(10)       // Результат: 20 px
//	pxFromSp := metric.SpToPx(10)       // Результат: 15 px
//	dpFromPx := metric.PxToDp(20)       // Результат: 10 dp
//	spFromPx := metric.PxToSp(15)       // Результат: 10 sp
//	spFromDp := metric.DpToSp(10)       // Результат: 13.33 sp
//	dpFromSp := metric.SpToDp(15)       // Результат: 10 dp
//
// Метод для получения текущих плотностей:
//   - GetDensity: Возвращает значения плотности `PxPerDp` и `PxPerSp`. Полезно для проверки
//     текущих настроек или отладки.
//
// Пример:
//
//	metric := pxconv.NewMetric(2.5, 1.8)
//	pxPerDp, pxPerSp := metric.GetDensity()
//	fmt.Printf("PxPerDp: %.1f, PxPerSp: %.1f\n", pxPerDp, pxPerSp) // PxPerDp: 2.5, PxPerSp: 1.8
//
// # Умолчания
//
// Если значения плотности экрана заданы как 0 или отрицательные, конструктор `NewMetric`
// и методы `Metric` автоматически заменяют их на 1. Это позволяет избежать деления на ноль
// и обеспечивает стабильную работу.
//
// Пример:
//
//	metric := pxconv.NewMetric(0, -5) // Значения PxPerDp и PxPerSp будут автоматически заменены на 1.
//
// # Преимущества
//
// Пакет pxconv полезен для адаптации пользовательских интерфейсов под экраны
// с разной плотностью, что делает его подходящим для мобильных устройств, настольных приложений
// и других графических систем. Конверсия между единицами позволяет создавать масштабируемые
// и отзывчивые интерфейсы.
package pxconv
