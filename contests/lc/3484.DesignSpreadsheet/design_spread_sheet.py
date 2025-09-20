# lc_3484_spreadsheet.py
from typing import Dict

class Spreadsheet:
    def __init__(self, rows: int):
        # храним только установленные клетки (sparse)
        self.rows = rows
        self.val: Dict[str, int] = {}

    def setCell(self, cell: str, value: int) -> None:
        # по условию значения 0..1e5; 0 тоже допустим и означает «храним 0»
        self.val[cell] = value

    def resetCell(self, cell: str) -> None:
        # reset именно в 0 (не «удалить»), поэтому будем хранить 0
        self.val[cell] = 0

    def getValue(self, formula: str) -> int:
        # формат: "=TOKEN+TOKEN"
        assert formula and formula[0] == '='
        body = formula[1:]
        i = body.find('+')
        left = body[:i]
        right = body[i + 1 :]
        return self._get_token(left) + self._get_token(right)

    def _get_token(self, tok: str) -> int:
        # токен — либо целое число, либо ссылка вида [A-Z][1..]
        if tok and tok[0].isdigit():
            return int(tok)
        # ссылка на ячейку
        return self.val.get(tok, 0)
