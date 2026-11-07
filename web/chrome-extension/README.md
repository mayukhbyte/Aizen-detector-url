# aizen.detector for Chrome

Real-time browser extension that checks links and pages with **aizen.detector** directly in Chrome.

- **Automated Protection**: Warns before visiting potentially malicious or deceptive web pages.
- **Password Shield**: Alerts before typing credentials into untrusted domains.
- **Context Menu Integration**: Right-click any URL or highlighted text to immediately check it with aizen.detector.
- **Real-Time Visual Indicators**: Color-coded badges indicating domain safety directly on your toolbar.

---

## 🛠️ Build & Install

```bash
cd web/chrome-extension
npm install
npm run build
```

1. Navigate to `chrome://extensions` in Google Chrome.
2. Enable **Developer mode** in the top right.
3. Click **Load unpacked** and select the `web/chrome-extension/dist` directory.
