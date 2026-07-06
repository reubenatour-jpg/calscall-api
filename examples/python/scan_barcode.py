"""
CalsCall Nutrition API — Python example

Install:
    pip install requests

Usage:
    python scan_barcode.py 5053990194769
"""

import sys
import requests

API_KEY = "your-key-here"
BASE_URL = "https://calscall.online"


def lookup_barcode(barcode: str) -> dict:
    headers = {"X-API-Key": API_KEY}
    resp = requests.get(
        f"{BASE_URL}/v1/nutrition",
        params={"barcode": barcode},
        headers=headers,
        timeout=10,
    )
    resp.raise_for_status()
    return resp.json()


if __name__ == "__main__":
    barcode = sys.argv[1] if len(sys.argv) > 1 else "5053990194769"
    data = lookup_barcode(barcode)

    if data["found"]:
        p = data["p"]
        print(f"✅ {p['n']}")
        print(f"   Calories:  {p['c']} kcal/100g")
        print(f"   Protein:   {p['p']}g/100g")
        print(f"   Carbs:     {p['h']}g/100g")
        print(f"   Fat:       {p['f']}g/100g")
        if data.get("sv"):
            print(f"   Per serve: {data['sv']['c']} kcal")
        print(f"   Source:    {data['source']}")
        if p["ns"]:
            print(f"   NutriScore: {p['ns']}")
        if p["al"]:
            print(f"   Allergens:  {p['al']}")
    else:
        print(f"❌ Product not found for barcode {barcode}")
