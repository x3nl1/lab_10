import requests

BASE = "http://localhost:8080"

def calculate(values):
    return requests.post(f"{BASE}/calculate", json={"values": values}).json()

def get_token():
    return requests.get(f"{BASE}/token").json()["token"]

def protected(token):
    return requests.get(
        f"{BASE}/protected",
        headers={"Authorization": token}
    ).status_code