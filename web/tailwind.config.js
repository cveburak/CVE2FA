export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'system-ui', '-apple-system', 'Segoe UI', 'sans-serif'],
        mono: ['"JetBrains Mono"', 'ui-monospace', 'SFMono-Regular', 'Menlo', 'monospace'],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        cvetheme: {
          "primary": "#0f766e",
          "primary-focus": "#0d6d66",
          "primary-content": "#ffffff",
          "secondary": "#1e293b",
          "secondary-focus": "#0f172a",
          "secondary-content": "#ffffff",
          "accent": "#0891b2",
          "accent-focus": "#0e7490",
          "accent-content": "#ffffff",
          "neutral": "#475569",
          "neutral-focus": "#334155",
          "neutral-content": "#f8fafc",
          "base-100": "#ffffff",
          "base-200": "#f1f5f9",
          "base-300": "#e2e8f0",
          "base-content": "#1e293b",
          "info": "#0284c7",
          "success": "#16a34a",
          "warning": "#d97706",
          "error": "#dc2626",
          "--rounded-box": "1rem",
          "--rounded-btn": "0.5rem",
          "--rounded-badge": "1.9rem",
          "--animation-btn": "0.2s",
          "--animation-input": "0.2s",
          "--btn-focus-scale": "0.98",
          "--btn-text-case": "none",
          "--border-btn": "1px",
          "--tab-border": "1px",
          "--tab-radius": "0.5rem",
        },
      },
    ],
    darkTheme: "cvetheme",
  },
}