from client import get_token, protected

def test_protected():
    token = get_token()
    status = protected(token)
    assert status == 200