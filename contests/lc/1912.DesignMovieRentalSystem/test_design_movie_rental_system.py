import importlib.util
import unittest
from pathlib import Path


_MODULE_PATH = Path(__file__).with_name("design-movie-rental-system.py")
_MODULE_SPEC = importlib.util.spec_from_file_location("design_movie_rental_system", _MODULE_PATH)
_MODULE = importlib.util.module_from_spec(_MODULE_SPEC)
assert _MODULE_SPEC and _MODULE_SPEC.loader  # for mypy/static checkers
_MODULE_SPEC.loader.exec_module(_MODULE)
MovieRentingSystem = _MODULE.MovieRentingSystem


class MovieRentingSystemTest(unittest.TestCase):
    def test_example_flow(self) -> None:
        entries = [
            [0, 1, 5],
            [0, 2, 6],
            [0, 3, 7],
            [1, 1, 4],
            [1, 2, 7],
            [2, 1, 5],
        ]
        system = MovieRentingSystem(3, entries)

        self.assertEqual([1, 0, 2], system.search(1))

        system.rent(0, 1)
        self.assertEqual([1, 2], system.search(1))

        system.drop(0, 1)
        self.assertEqual([1, 0, 2], system.search(1))

        system.rent(1, 2)
        system.rent(2, 1)
        self.assertEqual([[2, 1], [1, 2]], system.report())

    def test_report_orders_by_price_shop_and_movie(self) -> None:
        entries = [
            [0, 1, 5],
            [0, 2, 5],
            [1, 3, 4],
            [2, 4, 4],
            [3, 5, 4],
            [4, 6, 6],
        ]
        system = MovieRentingSystem(7, entries)

        system.rent(1, 3)
        system.rent(2, 4)
        system.rent(3, 5)
        system.rent(0, 1)
        system.rent(0, 2)
        system.rent(4, 6)

        self.assertEqual(
            [[1, 3], [2, 4], [3, 5], [0, 1], [0, 2]],
            system.report(),
        )

    def test_search_prefers_lower_shop_on_price_tie(self) -> None:
        entries = [
            [0, 1, 3],
            [1, 1, 3],
        ]
        system = MovieRentingSystem(2, entries)

        self.assertEqual([0, 1], system.search(1))

        system.rent(0, 1)
        self.assertEqual([1], system.search(1))

        system.drop(0, 1)
        self.assertEqual([0, 1], system.search(1))


if __name__ == "__main__":
    unittest.main()
