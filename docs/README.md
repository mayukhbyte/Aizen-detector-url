# aizen.detector Documentation

Welcome to the technical documentation for **aizen.detector**, a modern, real-time phishing detection and URL analysis suite created by **Mayukh** ([@mayukhbyte](https://github.com/mayukhbyte)).

---

## 🚀 Live Production Links

| Service | Platform | URL |
|---|---|---|
| **Frontend Application** | Vercel (SvelteKit) | [https://aizen-detector.vercel.app](https://aizen-detector.vercel.app) |
| **Backend REST API** | Vercel Serverless (Go) | [https://aizen-detector-api.vercel.app](https://aizen-detector-api.vercel.app) |
| **API Health Check** | Vercel Serverless | [https://aizen-detector-api.vercel.app/health](https://aizen-detector-api.vercel.app/health) |
| **Interactive Swagger API** | Vercel Serverless | [https://aizen-detector-api.vercel.app/swagger/index.html](https://aizen-detector-api.vercel.app/swagger/index.html) |
| **GitHub Repository** | GitHub (Private) | [https://github.com/mayukhbyte/aizen-detector](https://github.com/mayukhbyte/aizen-detector) |

---

## 🏛 System Architecture

aizen.detector is built as a serverless-first, modular architecture:

1. **Web Frontend (`web/website`)**:
   - Built with **SvelteKit 5** and **TailwindCSS**.
   - Features **Friday**, an interactive paperclip assistant that breaks down risk factors into plain language.
   - Live URL preview screenshot cards rendered on-demand.
   - Hosted and deployed on Vercel CDN.

2. **Backend Engine (`server/`)**:
   - High-throughput **Go REST API** running on Vercel Serverless Functions.
   - Executes **21 concurrent analyzers** across 7 security categories using Go routines.
   - Evaluates **39 distinct security signals** and computes an explainable trust/risk verdict.
   - Built-in cloud screenshot fallback rendering for environments without local headless Chrome.

3. **Chrome Extension (`web/chrome-extension`)**:
   - Client-side browser protection alerting users before entering credentials on suspicious domains.

---

## ⚙️ Detection Engine & Scoring

The engine calculates an explainable safety verdict:

$$\text{Final Score} = \text{clamp}\left(50 + (\text{Trust Score} - \text{Risk Score}) \times 0.5\right)$$

- **Safe**: $\ge 65$
- **Suspicious**: $30 - 64$
- **Risky**: $< 30$

### 7 Security Categories
1. **URL Structure**: Raw IP hosts, Punycode/IDN spoofing, URL shorteners, excessive length/depth, phishing keywords, homoglyph character analysis.
2. **Network & HTTP**: Hop counts, cross-domain redirection, HSTS protocol enforcement, status codes.
3. **DNS Validation**: NS records, MX records, IP resolution.
4. **TLS/SSL Encryption**: Chain validity, certificate expiration, known-bad fingerprints, CT logging verification.
5. **Domain Intelligence**: Popularity rank, TLD trust tier, RDAP/WHOIS domain age, Shannon entropy calculation.
6. **Typosquatting & Combo-squatting**: Comparison against top-ranking global domains and protected brand names.
7. **Content Inspection**: Credential phishing forms, credit card collection fields, hidden iframes, deceptive branding.

---

## 🔌 API Reference

### Analyze URL
```http
GET /api/v1/analyze?url={target_url}
```

#### Example Request:
```bash
curl "https://aizen-detector-api.vercel.app/api/v1/analyze?url=https://google.com"
```

#### Response:
```json
{
  "url": "https://google.com",
  "domain": "google.com",
  "result": {
    "trust_score": 100,
    "risk_score": 0,
    "final_score": 100,
    "verdict": "Safe",
    "reasons": {
      "good_reasons": ["Long-standing domain history", "Valid TLS/SSL certificate with secure cipher"],
      "neutral_reasons": ["Standard top-level domain"],
      "bad_reasons": []
    }
  }
}
```

### Screenshot Preview
```http
GET /api/v1/screenshot?url={target_url}
```
Returns a binary PNG image of the rendered webpage.

---

## 🔒 Security & Firewall Configurations

For standalone or self-hosted deployments:
- Only standard web traffic ports (`80` / `443`) should be exposed externally.
- In production, set `CORS_ALLOWED_ORIGINS` to `https://aizen-detector.vercel.app`.
- Internal services communicate over isolated virtual networks.
- Admin endpoints require signed Bearer tokens generated via `ADMIN_JWT_SECRET`.

---

## 👤 Author & Maintainer

**Mayukh** ([@mayukhbyte](https://github.com/mayukhbyte))  
*Contact: [hi@aizen.detector](mailto:hi@aizen.detector)*  
*License: MIT License*
