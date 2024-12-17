import requests

def test_availability():
    r = requests.get("127.0.0.1:80")
    assert r.status_code == 200