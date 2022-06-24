![cplusgo](https://user-images.githubusercontent.com/62110361/175657654-c564061e-6ed1-4375-acd7-246b82b95e5d.png)

# GoDet
Экзаменационная работа по основам программирования за второй семестр.

---

Exam work on the basics of programming for the second semester.

## Описание
Мне было интересно для чего может потребоваться Web-интерфейс для программы, написанной на C++.
Так как на занятиях мы делали математическую библиотеку, связанную с матричными операциями, я подумала, что было бы интересно сделать
простую библиотеку на C++, которая бы вычисляла определитель матрицы, полученной от пользователя через web страницу. 
Связующим звеном между frontend областью и библиотекой стал язык GoLang.
Идея была в том, что C++ - самый быстрый из современных компилируемых языков и он отлично подходит для вычислений, 
а GoLang - относительно простой и современный язык для написания серверной части программы.
![illustration](https://user-images.githubusercontent.com/62110361/175657637-c6d45cd6-de1b-4d8f-aced-591352019434.png)

---

## About
I was wondering why a Web interface might be needed for a program written in C++.
Since in the classroom we were making a mathematical library related to matrix operations, I thought it would be interesting to do
a simple C++ library that would calculate the determinant of a matrix received from a user through a web page.
The link between the frontend area and the library was the GoLang language.
The idea was that C++ is the fastest of today's compiled languages ​​and is great for computing,
and GoLang is a relatively simple and modern language for writing a server-side program.

## Ход работы
Для начала, имеющуюся функцию вычисления определителя, мне необходимо было упаковать в библиотеку, чтобы ее можно было поключить к Go.
Здесь я столкнулась с следующими проблемами: GoLang и C++ не могут взаимодействовать напрямую, и код, написанный на языке C++ необходимо было конвертировать на язык C.
У меня не получилось подключить имеющуюся математическую библиотеку с матрицами, так как на любые теги <include> GoLang ругался из-за несовместимости
(выходила ошибка о xxx.dll is not a valid Win32 application).
Несмотря на большое количество затраченного времени на поиски решения, желаемого результата я так и не добилась.
Это перечеркнуло все плюсы создания библиотеки, но мне стало жаль потраченного времени а также мне было интересно изучить GoLang и в целом всю backend часть.
В итоге я решила не начинать другой проект, а написать простую функцию на C++ не импортируя дополнительные и даже стандартную std библиотеки.
На всем пути выполнения данной работы я сталкивалась с другими ошибками, которые, к счастью, не оказались нерешаемыми.
  
---
  
## Working Process
To begin with, I needed to pack the existing function for calculating the determinant into a library so that it could be connected to Go.
Here I ran into the following problems: GoLang and C++ cannot communicate directly, and code written in C++ had to be converted to C.
I was unable to connect the existing mathematical library with matrices, since GoLang cursed any <include> tags due to incompatibility
(there was an error about xxx.dll is not a valid Win32 application).
Despite the large amount of time spent looking for a solution, I did not achieve the desired result.
This crossed out all the advantages of creating a library, but I felt sorry for the time spent, and I was also interested in learning GoLang and, in general, the entire backend part.
In the end, I decided not to start another project, but to write a simple function in C ++ without importing additional and even standard std libraries.
Throughout the course of this work, I encountered other errors, which, fortunately, did not turn out to be unsolvable.
  
## Итог
> Неудачный результат - тоже результат

Выполнить хорошо данную работу на моем текущем уровне знаний не представлялось возможным.
Но, несмотря на сложившуюся ситуацию с языком C++, я узнала много нового из части backend разработки, освежила знания HTML и CSS.
  
---
  
## Result
> An unsuccessful result is also a result
  
It was not possible to do this job well at my current level of knowledge.
But, despite the current situation with the C ++ language, I learned a lot from the backend development part, refreshed my knowledge of HTML and CSS.

  

