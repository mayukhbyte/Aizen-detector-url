<div
  class="backdrop pointer-events-none absolute inset-x-0 -top-24 h-[44rem] overflow-hidden"
  aria-hidden="true"
>
  <svg class="signal-art absolute inset-x-0 top-0 mx-auto" viewBox="0 0 1080 620" fill="none">
    <g class="orbit-lines">
      <circle cx="540" cy="206" r="105" />
      <circle cx="540" cy="206" r="158" stroke-dasharray="2 8" />
      <circle cx="540" cy="206" r="214" stroke-dasharray="1 10" />
      <path d="M540 206 760 111" />
      <path d="M540 206 322 292" />
      <path d="M346 133h388" stroke-dasharray="3 9" />
      <path d="M373 346h334" stroke-dasharray="3 9" />
    </g>
    <g class="signal-nodes">
      <circle cx="760" cy="111" r="5" />
      <circle cx="322" cy="292" r="4" />
      <circle cx="346" cy="133" r="3" />
      <circle cx="734" cy="346" r="3" />
      <circle cx="540" cy="48" r="3" />
      <circle cx="540" cy="364" r="3" />
    </g>
    <g class="shield-mark">
      <path d="M540 147 578 163v28c0 25-16 43-38 55-22-12-38-30-38-55v-28l38-16Z" />
      <path d="m523 204 12 12 24-26" />
    </g>
    <path class="scan-sweep" d="M540 206 677 147" />
  </svg>
  <div class="backdrop-glow absolute inset-0"></div>
</div>

<style>
  .backdrop {
    background:
      radial-gradient(ellipse 46rem 29rem at 50% 7%, rgb(214 226 255 / 0.56), transparent 72%),
      radial-gradient(ellipse 40rem 20rem at 50% 0%, rgb(255 255 255 / 0.74), transparent 75%);
  }
  .signal-art {
    width: min(100%, 1080px);
    height: 100%;
    color: #3461c9;
    opacity: 0.32;
    mask-image: linear-gradient(to bottom, #000 0%, rgb(0 0 0 / 0.78) 42%, transparent 96%);
  }
  .orbit-lines circle,
  .orbit-lines path {
    stroke: currentColor;
    stroke-width: 1;
    opacity: 0.19;
  }
  .signal-nodes circle {
    fill: currentColor;
    opacity: 0.55;
  }
  .shield-mark path:first-child {
    fill: rgb(255 255 255 / 0.56);
    stroke: currentColor;
    stroke-width: 1.3;
  }
  .shield-mark path:last-child {
    stroke: currentColor;
    stroke-width: 2.1;
    stroke-linecap: round;
    stroke-linejoin: round;
  }
  .scan-sweep {
    stroke: currentColor;
    stroke-width: 1.5;
    stroke-linecap: round;
    opacity: 0.48;
  }
  .backdrop-glow {
    background: linear-gradient(to bottom, transparent 46%, var(--color-gray-50) 100%);
  }
  :global(.dark) .backdrop {
    background:
      radial-gradient(ellipse 46rem 29rem at 50% 7%, rgb(47 75 130 / 0.22), transparent 72%),
      radial-gradient(ellipse 40rem 20rem at 50% 0%, rgb(255 255 255 / 0.025), transparent 75%);
  }
  :global(.dark) .signal-art {
    color: #8eb0ff;
    opacity: 0.2;
  }
  :global(.dark) .shield-mark path:first-child {
    fill: rgb(15 13 11 / 0.2);
  }
  :global(.dark) .backdrop-glow {
    background: linear-gradient(to bottom, transparent 46%, var(--color-gray-950) 100%);
  }
  @media (max-width: 640px) {
    .signal-art {
      width: 150%;
      max-width: none;
      left: 50%;
      transform: translateX(-50%);
      opacity: 0.18;
    }
  }
</style>
