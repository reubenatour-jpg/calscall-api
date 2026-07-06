/**
 * CalsCall Nutrition API — Node.js example
 *
 * Usage:
 *   node scan_barcode.mjs 5053990194769
 */

const API_KEY = "your-key-here";
const BASE_URL = "https://calscall.online";

async function lookupBarcode(barcode) {
  const resp = await fetch(
    `${BASE_URL}/v1/nutrition?barcode=${encodeURIComponent(barcode)}`,
    { headers: { "X-API-Key": API_KEY } }
  );
  if (!resp.ok) throw new Error(`HTTP ${resp.status}: ${await resp.text()}`);
  return resp.json();
}

const barcode = process.argv[2] || "5053990194769";

lookupBarcode(barcode)
  .then((data) => {
    if (data.found) {
      const p = data.p;
      console.log(`✅ ${p.n}`);
      console.log(`   Calories:  ${p.c} kcal/100g`);
      console.log(`   Protein:   ${p.p}g/100g`);
      console.log(`   Carbs:     ${p.h}g/100g`);
      console.log(`   Fat:       ${p.f}g/100g`);
      if (data.sv) console.log(`   Per serve: ${data.sv.c} kcal`);
      console.log(`   Source:    ${data.source}`);
      if (p.ns) console.log(`   NutriScore: ${p.ns}`);
      if (p.al) console.log(`   Allergens:  ${p.al}`);
    } else {
      console.log(`❌ Product not found for barcode ${barcode}`);
    }
  })
  .catch(console.error);
