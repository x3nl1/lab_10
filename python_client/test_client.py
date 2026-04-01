from client import calculate

def test_calculate():
    res = calculate([1, 2, 3])
    assert res["sum"] == 14