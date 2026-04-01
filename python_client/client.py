import requests

BASE = "http://localhost:8080"

def calculate(values):
    return requests.post(f"{BASE}/calculate", json={"values": values}).json()