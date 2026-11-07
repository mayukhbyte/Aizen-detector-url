<div align="center">

<img src="web/website/static/aizen-detector-normal.png" alt="Aizen Detector Logo" width="120" height="120" />

# aizen.detector

**sketchy link? just aizen detect it**

**Live Website: [https://aizen-detector.vercel.app](https://aizen-detector.vercel.app)** · **Live API: [https://aizen-detector-api.vercel.app](https://aizen-detector-api.vercel.app)**

Real-time phishing detection and URL safety analysis engine. Paste any link to instantly receive a comprehensive trust score, fully explainable verdict, and deep security report. Built with explainable heuristics, complete transparency, and zero black-box obscurity.

**21 concurrent analyzers · 39 security signals · DNS · TLS · RDAP/WHOIS · typosquatting · redirect chains · short-link resolution · content analysis · threat intel · REST API · Chrome extension**

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![SvelteKit](https://img.shields.io/badge/SvelteKit-5-orange?logo=svelte)](https://svelte.dev)
[![Vercel](https://img.shields.io/badge/Vercel-Deployed-black?logo=vercel)](https://aizen-detector.vercel.app)
[![Author](https://img.shields.io/badge/Author-mayukh-blue)](https://github.com/mayukhbyte)

[⚡ Live Demo](https://aizen-detector.vercel.app) · [⚙️ Detection Engine](#detection-engine) · [🏛 Architecture](#architecture) · [🚀 API Reference](#api-example)

<sub>Created & maintained by Mayukh (@mayukhbyte)</sub>

</div>

---

## 🎯 At a Glance

- **Instant Live Crawl**: Real-time evaluation of arbitrary URLs, including deep redirection hop resolution.
- **Explainable Scoring**: Every verdict breaks down into discrete good, bad, and neutral signals.
- **Privacy-First**: No tracking, no stored browsing history, no third-party cookies.
- **Microservice & Serverless Ready**: Full support for Vercel Serverless (Go runtime) and SvelteKit frontend.
- **Friday the Assistant**: Interactive in-app guide assisting users with security insights and warnings.

---

## 🚀 Live Deployments

- **Frontend Application**: [https://aizen-detector.vercel.app](https://aizen-detector.vercel.app)
- **Backend API**: [https://aizen-detector-api.vercel.app](https://aizen-detector-api.vercel.app)
- **API Health Check**: [https://aizen-detector-api.vercel.app/health](https://aizen-detector-api.vercel.app/health)
- **Interactive Documentation**: [https://aizen-detector-api.vercel.app/swagger/index.html](https://aizen-detector-api.vercel.app/swagger/index.html)

---

## 💻 Quick Start (Local Development)

### 1. Clone the repository
```bash
git clone https://github.com/mayukhbyte/aizen-detector.git
cd aizen-detector
```

### 2. Run the Backend (Go)
```bash
cd server
go run ./cmd/aizen-detector
```
*Backend runs on `http://localhost:8080`*

### 3. Run the Frontend (SvelteKit)
```bash
cd ../web/website
npm install
npm run dev
```
*Frontend runs on `http://localhost:5173`*

---

## 📎 Meet Friday

**Friday** is the AI paperclip assistant residing in the corner of `aizen.detector`. Friday guides users through unfamiliar technical reports, flags critical vulnerabilities, and explains risk scores in plain language.

---

## 🔌 API Example

Analyze any URL via standard HTTP GET:

```bash
curl "https://aizen-detector-api.vercel.app/api/v1/analyze?url=https://example.com"
```

### Sample Response:
```json
{
  "url": "https://example.com",
  "domain": "example.com",
  "result": {
    "trust_score": 100,
    "risk_score": 0,
    "final_score": 100,
    "verdict": "Safe",
    "reasons": {
      "good_reasons": [
        "Long-standing domain history",
        "Valid TLS/SSL certificate with secure cipher"
      ],
      "neutral_reasons": [
        "Standard top-level domain"
      ],
      "bad_reasons": []
    }
  }
}
```

---

## ⚙️ Detection Engine

**21 concurrent goroutines** evaluate **7 distinct security categories**, generating **39 individual signals**:

$$\text{Final Score} = \text{clamp}\left(50 + (\text{Trust Score} - \text{Risk Score}) \times 0.5\right)$$

- **Safe**: $\ge 65$
- **Suspicious**: $30 - 64$
- **Risky**: $< 30$

### Core Signal Categories
1. **URL Structure**: Raw IP hosts, Punycode/IDN spoofing, URL shorteners, excessive length/depth, phishing keywords, homoglyph character analysis.
2. **Network & HTTP**: Hop counts, cross-domain redirection, HSTS protocol enforcement, status codes.
3. **DNS Validation**: NS records, MX records, IP resolution.
4. **TLS/SSL Encryption**: Chain validity, certificate expiration, known-bad fingerprints, CT logging verification.
5. **Domain Intelligence**: Popularity rank, TLD trust tier, RDAP/WHOIS domain age, Shannon entropy calculation.
6. **Typosquatting & Combo-squatting**: Comparison against top-ranking global domains and protected brand names.
7. **Content Inspection**: Credential phishing forms, credit card collection fields, hidden iframes, deceptive branding.

---

## 🏛 Architecture

```text
aizen-detector/
├── server/
│   ├── api/index.go             # Vercel Go Serverless Function bridge
│   ├── cmd/aizen-detector/      # Standalone server entrypoint
│   ├── pkg/
│   │   ├── analyzer/            # Goroutine runner & score aggregation
│   │   ├── handler/             # REST endpoints (Gin framework)
│   │   ├── service/             # 21 concurrent check implementations
│   │   └── docs/                # OpenAPI / Swagger specification
│   └── vercel.json              # Backend serverless configuration
├── web/
│   ├── website/                 # SvelteKit 5 web interface
│   └── chrome-extension/        # Browser extension integration
└── assets/                      # Brand assets and logos
```

---

## 👤 Author

**Mayukh** ([@mayukhbyte](https://github.com/mayukhbyte))  
*Contact: [hi@aizen.detector](mailto:hi@aizen.detector)*
