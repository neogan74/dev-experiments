import threading
from typing import Iterable, Tuple

import pytest

from print_in_order import Foo


def run_threads(foo: Foo, order: Tuple[int, int, int]) -> str:
    output_lock = threading.Lock()
    result: list[str] = []

    def safe_append(value: str) -> None:
        with output_lock:
            result.append(value)

    actions = {
        1: lambda: foo.first(lambda: safe_append("first")),
        2: lambda: foo.second(lambda: safe_append("second")),
        3: lambda: foo.third(lambda: safe_append("third")),
    }

    threads = [threading.Thread(target=actions[i]) for i in order]

    for thread in threads:
        thread.start()

    for thread in threads:
        thread.join()

    return "".join(result)


@pytest.mark.parametrize("order", [
    (1, 2, 3),
    (1, 3, 2),
    (2, 1, 3),
    (2, 3, 1),
    (3, 1, 2),
    (3, 2, 1),
])
def test_print_in_order(order: Tuple[int, int, int]) -> None:
    foo = Foo()
    for _ in range(10):
        output = run_threads(foo, order)
        assert output == "firstsecondthird"
        foo = Foo()
