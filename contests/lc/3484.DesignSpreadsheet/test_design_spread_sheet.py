from design_spread_sheet import Spreadsheet

def test_example_flow():
    ss = Spreadsheet(3)
    assert ss.getValue("=5+7") == 12
    ss.setCell("A1", 10)
    assert ss.getValue("=A1+6") == 16
    ss.setCell("B2", 15)
    assert ss.getValue("=A1+B2") == 25
    ss.resetCell("A1")
    assert ss.getValue("=A1+B2") == 15  # A1 теперь 0

def test_unset_defaults_to_zero_and_zero_is_stored():
    ss = Spreadsheet(2)
    assert ss.getValue("=X9+3") == 3  # неизвестная ячейка -> 0
    ss.setCell("C1", 0)
    assert ss.getValue("=C1+0") == 0
    ss.setCell("C1", 7)
    ss.resetCell("C1")
    assert ss.getValue("=C1+1") == 1

def test_numbers_and_cells_mixed():
    ss = Spreadsheet(5)
    ss.setCell("Z5", 100)
    assert ss.getValue("=Z5+900") == 1000