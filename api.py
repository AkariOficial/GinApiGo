from requests import get

def main():
    r = get("http://localhost:6765/ping").json()
    print(r)

main()
