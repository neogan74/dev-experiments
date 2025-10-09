import threading


class Foo:
    def __init__(self) -> None:
        self._first_done = threading.Event()
        self._second_done = threading.Event()

    def first(self, print_first) -> None:
        print_first()
        self._first_done.set()

    def second(self, print_second) -> None:
        self._first_done.wait()
        print_second()
        self._second_done.set()

    def third(self, print_third) -> None:
        self._second_done.wait()
        print_third()


def main() -> None:
    foo = Foo()

    output_lock = threading.Lock()
    result = []

    def safe_append(value: str) -> None:
        with output_lock:
            result.append(value)

    threads = [
        threading.Thread(target=foo.third, args=(lambda: safe_append("third"),)),
        threading.Thread(target=foo.second, args=(lambda: safe_append("second"),)),
        threading.Thread(target=foo.first, args=(lambda: safe_append("first"),)),
    ]

    for thread in threads:
        thread.start()

    for thread in threads:
        thread.join()

    print("".join(result))


if __name__ == "__main__":
    main()
