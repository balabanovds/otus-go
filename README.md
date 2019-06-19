# LogOtusEvent

Задача: написать функцию логирования LogOtusEvent, 
на вход которой приходят события типа HwAccepted 
(домашняя работа принята) и HwSubmitted (студент отправил дз) 
для этого - спроектировать и реализовать интерфейс OtusEvent. 
Для события HwAccepted мы хотим логировать дату, айди и грейд, 
для HwSubmitter - дату, id и комментарий, например:
  ```sh
  2019-01-01 submitted 3456 "please take a look at my homework"
  2019-01-01 accepted 3456 4
```
  ```go
package p 

import "io"

type HwAccepted struct {
     Id int
     Grade int
  }
  
  type HwSubmitted struct {
     Id int
     Code string
     Comment string
  }
  
  type OtusEvent interface {
  }
  
  func LogOtusEvent(e OtusEvent, w io.Writer) {
  
  }
```
