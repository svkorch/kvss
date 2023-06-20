# Key/Value Storage (simple)

## Техническое задание

Нужно написать простую реализацию для хранения данных по ключу.

Cчитаем, что у нас бесконечная память и нам не нужно задумываться об удалении данных.

Реализация должна удовлетворять интерфейсу:

```golang
type Storager interface {
    Add(k, v string)
    Get(k string) (v string, ok bool)
}
```

Написать тесты.

Реализация должна поддерживать работу в многопоточном режиме.

## Отчёт о проведённом тестировании

```
$ go test -v github.com/svkorch/kvss/internal/kvss
=== RUN   TestStorage
--- PASS: TestStorage (0.00s)
PASS
ok      github.com/svkorch/kvss/internal/kvss   0.001s
```

