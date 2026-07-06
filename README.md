<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://calscall.online/static/calscall-logo.jpg">
    <img src="https://calscall.online/static/calscall-logo.jpg" alt="CalsCall Logo" width="200">
  </picture>
</p>

<h1 align="center">CalsCall Nutrition API</h1>

<p align="center">
  <b>One barcode. Instant nutrition data.</b><br>
  Scan any UPC / EAN / GTIN barcode → get calories, protein, carbs, fat, allergens & more in under 5ms.
</p>

<p align="center">
  <a href="https://calscall.online/admin/"><img src="https://img.shields.io/badge/status-live-brightgreen?style=flat-square" alt="Status: Live"></a>
  <a href="#quick-start"><img src="https://img.shields.io/badge/demo-try%20it%20now-blue?style=flat-square" alt="Try it now"></a>
  <a href="https://rapidapi.com/"><img src="https://img.shields.io/badge/marketplace-rapidapi-orange?style=flat-square" alt="RapidAPI Marketplace"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-lightgrey?style=flat-square" alt="MIT License"></a>
</p>

---

## 🚀 Quick Start

```bash
# One curl command — just swap in your API key
curl -H "X-API-Key: your-key-here" \
  "https://calscall.online/v1/nutrition?barcode=5053990194769"
```

**Response (12ms):**
```json
{
  "found": true,
  "source": "cache",
  "p": {
    "b": "5053990194769",
    "n": "Coca-Cola 330ml",
    "c": 42.0,
    "h": 10.6,
    "p": 0.0,
    "f": 0.0,
    "s": 330.0,
    "i": "https://calscall.online/static/example-product.jpg",
    "ct": "Beverages,Soft Drinks",
    "ns": "D",
    "nv": 4,
    "al": ""
  },
  "sv": {
    "c": 138.6,
    "h": 35.0,
    "p": 0.0,
    "f": 0.0
  }
}
```

---

## 🔑 Get an API Key

| Plan | Requests | Price | Best For |
|------|----------|-------|----------|
| **Free** | 100/day | $0 | Evaluation & testing |
| **Starter** | 5,000/month | $9.99 | Personal projects |
| **Pro** | 50,000/month | $49.99 | Production apps |
| **Enterprise** | Custom | Contact us | High-volume |

**[→ Get your key on RapidAPI](https://rapidapi.com/)** _(coming soon to AWS Marketplace & Gumroad)_

---

## 📖 API Reference

### `GET /v1/nutrition`

Look up a product by its barcode.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `barcode` | string | ✅ | 8–13 digit UPC / EAN / GTIN |

| Header | Required | Description |
|--------|----------|-------------|
| `X-API-Key` | ✅ | Your API key |

### Response Fields

| Field | Description |
|-------|-------------|
| `found` | Whether a product was found |
| `source` | Data origin: `cache` or `external` |
| `p.b` | Barcode |
| `p.n` | Product name |
| `p.c` | Calories per 100g (kcal) |
| `p.h` | Carbohydrates per 100g (g) |
| `p.p` | Protein per 100g (g) |
| `p.f` | Fat per 100g (g) |
| `p.s` | Serving size (g) — `0` if unknown |
| `p.i` | Product image URL |
| `p.ct` | Categories (comma-separated) |
| `p.ns` | Nutri-Score grade (A–E) |
| `p.nv` | NOVA food processing group (1–4) |
| `p.al` | Allergens (comma-separated) |
| `sv` | Per-serving values (present when `p.s` > 0) |

### Error Responses

| Status | Meaning |
|--------|---------|
| `400` | Invalid barcode (must be 8–13 digits) |
| `429` | Rate limit exceeded |
| `401` | Missing or invalid API key |

---

## 💻 Code Examples

### Python

```python
import requests

API_KEY = "your-key-here"
BARCODE = "5053990194769"

headers = {"X-API-Key": API_KEY}
resp = requests.get(
    f"https://calscall.online/v1/nutrition",
    params={"barcode": BARCODE},
    headers=headers
)

data = resp.json()
if data["found"]:
    product = data["p"]
    print(f"{product['n']} — {product['c']} kcal/100g")
    if "sv" in data:
        print(f"Per serving: {data['sv']['c']} kcal")
```

### JavaScript (Node.js)

```javascript
const API_KEY = "your-key-here";
const BARCODE = "5053990194769";

const resp = await fetch(
  `https://calscall.online/v1/nutrition?barcode=${BARCODE}`,
  { headers: { "X-API-Key": API_KEY } }
);

const data = await resp.json();
if (data.found) {
  console.log(`${data.p.n} — ${data.p.c} kcal/100g`);
  if (data.sv) {
    console.log(`Per serving: ${data.sv.c} kcal`);
  }
}
```

### cURL

```bash
curl -s -H "X-API-Key: your-key-here" \
  "https://calscall.online/v1/nutrition?barcode=5053990194769" | jq .
```

### Go

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://calscall.online/v1/nutrition?barcode=5053990194769", nil)
	req.Header.Set("X-API-Key", "your-key-here")
	
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	
	if data["found"].(bool) {
		p := data["p"].(map[string]interface{})
		fmt.Printf("%v — %v kcal/100g\n", p["n"], p["c"])
	}
}
```

---

## 🏗 Architecture

```
   Barcode
     │
     ▼
┌────────────┐   ┌────────────────────┐
│  FastAPI   │──▶│  PostgreSQL Cache  │
│  (v1/*)    │   │  (<5ms hits)       │
└────────────┘   └─────────┬──────────┘
                           │ (cache miss)
                           ▼
                  ┌────────────────┐
                  │  Data Sources  │
                  │  (multi-source │
                  │   fallback)    │
                  └────────────────┘
```

- **First lookup**: checks cache → misses → fetches from upstream sources → caches the result
- **Repeat lookups**: served from PostgreSQL in under 5ms
- **Multi-source fallback**: cascading data sources for maximum coverage

---

## 🧪 Try It

Pick any product with a barcode and try it yourself:

```bash
# A few to test with
curl -H "X-API-Key: your-key" "https://calscall.online/v1/nutrition?barcode=5053990194769"  # Coca-Cola
curl -H "X-API-Key: your-key" "https://calscall.online/v1/nutrition?barcode=5449000000996"  # Another soda
curl -H "X-API-Key: your-key" "https://calscall.online/v1/nutrition?barcode=5000159017002"  # Walkers Crisps
```

---

## 📊 Status & Monitoring

- **API Status:** [`https://calscall.online/`](https://calscall.online/) — the base URL is the health check
- **Admin Dashboard:** [`https://calscall.online/admin/`](https://calscall.online/admin/) — public stats (uptime, request volume, cache hit rate)

---

## 🤝 Contributing

This repo hosts the public documentation, SDK clients, and code examples for the CalsCall API.

- 🐛 **Found a bug in the docs?** Open an issue
- 💡 **Want a client library for another language?** Request it
- 📦 **Built something cool with the API?** We'd love to hear about it

The actual API server code is proprietary — but the SDK wrappers and examples here are open source (MIT).

---

## 📜 License

Documentation and code examples in this repo are MIT licensed. The CalsCall API service itself is commercial — see our [Terms of Use](TERMS.md) for service usage.

---

<p align="center">
  <a href="https://calscall.online">🌐 calscall.online</a> &nbsp;·&nbsp;
  <a href="https://rapidapi.com/">⚡ RapidAPI</a> &nbsp;·&nbsp;
  <a href="https://calscall.online/admin/">📊 Dashboard</a>
</p>
