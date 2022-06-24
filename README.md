![cplusgo](https://user-images.githubusercontent.com/62110361/175657654-c564061e-6ed1-4375-acd7-246b82b95e5d.png)

# GoDet
Экзаменационная работа по основам программирования за второй семестр.

## Описание
Мне было интересно для чего может потребоваться Web-интерфейс для программы, написанной на C++.
Так как на занятиях мы делали математическую библиотеку, связанную с матричными операциями, я подумала, что было бы интересно сделать
простую библиотеку на C++, которая бы вычисляла определитель матрицы, полученной от пользователя через web страницу. 
Связующим звеном между frontend областью и библиотекой стал язык GoLang.
Идея была в том, что C++ - самый быстрый из современных компилируемых языков и он отлично подходит для вычислений, 
а GoLang - относительно простой и современный язык для написания серверной части программы.
![illustration](https://user-images.githubusercontent.com/62110361/175657637-c6d45cd6-de1b-4d8f-aced-591352019434.png)

## Ход работы
Для начала, имеющуюся функцию вычисления определителя, мне необходимо было упаковать в библиотеку, чтобы ее можно было поключить к Go.
Здесь я столкнулась с следующими проблемами: GoLang и C++ не могут взаимодействовать напрямую, и код, написанный на языке C++ необходимо было конвертировать на язык C.
У меня не получилось подключить имеющуюся математическую библиотеку с матрицами, так как на любые теги <include> GoLang ругался из-за несовместимости
(выходила ошибка о xxx.dll is not a valid Win32 application).
Несмотря на большое количество затраченного времени на поиски решения, желаемого результата я так и не добилась.
Это перечеркнуло все плюсы создания библиотеки, но мне стало жаль потраченного времени а также мне было интересно изучить GoLang и в целом всю backend часть.
В итоге я решила не начинать другой проект, а написать простую функцию на C++ не импортируя дополнительные и даже стандартную std библиотеки.
На всем пути выполнения данной работы я сталкивалась с другими ошибками, которые, к счастью, не оказались нерешаемыми.
  
## Итог
> Неудачный результат - тоже результат
Выполнить хорошо данную работу на моем текущем уровне знаний не представлялось возможным.
Но, несмотря на сложившуюся ситуацию с языком C++, я узнала много нового из части backend разработки, освежила знания HTML и CSS.
  
