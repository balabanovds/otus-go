# Реализовать двусвязанный список

Что такое двусвязный список: https://en.wikipedia.org/wiki/Doubly_linked_list

Ожидаемые типы (псевдокод):

```go
List      // тип контейнер
  Len()   // длинна списка
  First() *Item // первый Item
  Last() *Item  // последний Item
  PushFront(v interface{}) *Item // добавить значение в начало
  PushBack(v interface{}) *Item  // добавить значение в конец
  // дополнительно
  GetNth(n int) *Item // получить n-ный Item списка
  InsertAfterNth(n int, v interface{}) // вставить в середину списка после n-нного индекса


Item   // элемент списка
  Value() interface{}  // возвращает значение
  Next() *Item          // следующий Item
  Prev() *Item         // предыдущий
  Remove()             // удалить Item из списка
```
