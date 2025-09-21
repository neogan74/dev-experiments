import collections
import heapq
from typing import Dict, List, Optional, Tuple


class MovieRentingSystem:
    def __init__(self, n: int, entries: List[List[int]]):
        self.unrented: Dict[int, List[Tuple[int, int, int]]] = collections.defaultdict(list)
        self.price_lookup: Dict[Tuple[int, int], int] = {}
        self.unrented_id: Dict[Tuple[int, int], Optional[int]] = {}
        self.rented_heap: List[Tuple[int, int, int, int]] = []
        self.rented_id: Dict[Tuple[int, int], Optional[int]] = {}
        self._next_entry_id = 0

        for shop, movie, price in entries:
            entry_id = self._allocate_entry_id()
            heapq.heappush(self.unrented[movie], (price, shop, entry_id))
            self.unrented_id[(shop, movie)] = entry_id
            self.rented_id[(shop, movie)] = None
            self.price_lookup[(shop, movie)] = price

    def search(self, movie: int) -> List[int]:
        result: List[int] = []
        temp: List[Tuple[int, int, int]] = []
        heap = self.unrented[movie]

        while heap and len(result) < 5:
            price, shop, entry_id = heapq.heappop(heap)
            if self.unrented_id.get((shop, movie)) != entry_id:
                continue
            result.append(shop)
            temp.append((price, shop, entry_id))

        for item in temp:
            heapq.heappush(heap, item)

        return result

    def rent(self, shop: int, movie: int) -> None:
        price = self.price_lookup[(shop, movie)]
        self.unrented_id[(shop, movie)] = None
        rented_entry_id = self._allocate_entry_id()
        self.rented_id[(shop, movie)] = rented_entry_id
        heapq.heappush(self.rented_heap, (price, shop, movie, rented_entry_id))

    def drop(self, shop: int, movie: int) -> None:
        price = self.price_lookup[(shop, movie)]
        self.rented_id[(shop, movie)] = None
        unrented_entry_id = self._allocate_entry_id()
        self.unrented_id[(shop, movie)] = unrented_entry_id
        heapq.heappush(self.unrented[movie], (price, shop, unrented_entry_id))

    def report(self) -> List[List[int]]:
        result: List[List[int]] = []
        temp: List[Tuple[int, int, int, int]] = []

        while self.rented_heap and len(result) < 5:
            price, shop, movie, entry_id = heapq.heappop(self.rented_heap)
            if self.rented_id.get((shop, movie)) != entry_id:
                continue
            result.append([shop, movie])
            temp.append((price, shop, movie, entry_id))

        for item in temp:
            heapq.heappush(self.rented_heap, item)

        return result

    def _allocate_entry_id(self) -> int:
        entry_id = self._next_entry_id
        self._next_entry_id += 1
        return entry_id


def main():
    # Example usage:
    entries = [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]
    movieRentingSystem = MovieRentingSystem(3, entries)
    print(movieRentingSystem.search(1))  # return [1,0,2]. Shops 1,0 and 2 have an unrented copy of movie 1. Their prices are [4,5,5] respectively. The cheapest shops are [1,0] (price 4 and 5 respectively). Shop 2 is also included since there are less than 5 shops.
    movieRentingSystem.rent(0, 1)         # Rent movie 1 from shop 0. Unrented movies at shop 0 are now [2,3].
    print(movieRentingSystem.report())    # return [[0,1]]. The only rented movie is [1] from shop [0].
    movieRentingSystem.drop(0, 1)         # Drop off movie 1 at shop 0. Unrented movies at shop 0 are now [1,2,3].
    print(movieRentingSystem.search(1))   # return [1,0,2]. Shops 0,1 and 2 have an unrented copy of movie 1. Their prices are [5,4,5] respectively. The cheapest shops are [1,0] (price 4 and 5 respectively). Shop 2 is also included since there are less than 5 shops.
    movieRentingSystem.rent(1, 2)         # Rent movie 2 from shop 1. Unrented movies at shop 1 are now [1].
    movieRentingSystem.rent(2, 1)         # Rent movie 1 from shop 2. Unrented movies at shop 2 are now [].
    print(movieRentingSystem.report())    # return [[2,1],[1,2]]. The rented movies are [2] from shop [1] (price=7) and [1] from shop [2] (price=5). The cheapest is [1] from shop [2] followed by [2] from shop [1].


if __name__ == "__main__":
    main()

# Your MovieRentingSystem object will be instantiated and called as such:
# obj = MovieRentingSystem(n, entries)
# param_1 = obj.search(movie)
# obj.rent(shop,movie)
# obj.drop(shop,movie)
# param_4 = obj.report()
