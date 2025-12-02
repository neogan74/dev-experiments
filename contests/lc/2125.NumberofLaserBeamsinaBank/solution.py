from typing import List

class Solution:
    def numberOfBeams(self, bank: List[str]) -> int:
        total_beams = 0
        prev_count = 0  # Количество устройств в последней непустой строке

        for row in bank:
            count = row.count('1')  # Считаем количество '1' в текущей строке
            if count == 0:
                continue  # Пропускаем строки без устройств
            total_beams += prev_count * count  # Количество лучей между строками
            prev_count = count  # Обновляем предыдущее количество

        return total_beams