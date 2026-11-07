<script lang="ts">
  import { onMount } from "svelte";
  import { ICON } from "../../ui/icons";
  import { formatUrl, stripTrackers } from "../../utils";
  import Icon from "../Icon.svelte";

  // The pill search bar: URL input, Paste, Scan, plus tracker clean-up and length hints.
  export let input = "";
  export let loading = false;
  export let formError: string | null = null;
  export let isLanding = true;
  export let onSubmit: (value: string) => void;
  export let onPaste: (() => void) | undefined = undefined;

  // The longest hint that fits the box in full; mentions whole messages when there's room.
  const HINTS = [
    "Paste a link, or a whole message with links in it",
    "Paste a link or a message",
    "Paste link or message",
    "Paste a link",
  ];
  let inputEl: HTMLInputElement;
  let placeholder = HINTS[HINTS.length - 1];

  function fitPlaceholder() {
    if (!inputEl) return;
    const cs = getComputedStyle(inputEl);
    const room = inputEl.clientWidth - parseFloat(cs.paddingLeft) - parseFloat(cs.paddingRight) - 2;
    const ctx = document.createElement("canvas").getContext("2d");
    if (!ctx) return;
    ctx.font = `${cs.fontWeight} ${cs.fontSize} ${cs.fontFamily}`;
    placeholder = HINTS.find((h) => ctx.measureText(h).width <= room) ?? HINTS[HINTS.length - 1];
  }

  onMount(() => {
    fitPlaceholder();
    // Fonts may arrive after first paint; re-measure then, and whenever the box resizes.
    document.fonts?.ready.then(fitPlaceholder);
    const ro = new ResizeObserver(fitPlaceholder);
    ro.observe(inputEl);
    return () => ro.disconnect();
  });

  let justPasted = false;
  let trackerCopied = false;
  let trackerDismissed = false;

  $: trackers = input ? stripTrackers(formatUrl(input)) : { cleaned: "", removed: [] as string[] };
  $: hasTrackers = trackers.removed.length > 0;
  $: if (input === "") trackerDismissed = false;

  async function pasteFromClipboard() {
    try {
      const text = await navigator.clipboard.readText();
      if (text) {
        input = text.trim();
        formError = null;
        onPaste?.();
        justPasted = true;
        setTimeout(() => (justPasted = false), 1500);
      }
    } catch {
      /* clipboard access denied */
    }
  }

  async function copyCleanUrl() {
    try {
      await navigator.clipboard.writeText(trackers.cleaned);
      trackerCopied = true;
      setTimeout(() => {
        trackerCopied = false;
        trackerDismissed = true;
      }, 1200);
    } catch {}
  }
</script>

<form
  class={isLanding ? "w-full max-w-2xl" : "w-full max-w-4xl mx-auto px-4"}
  on:submit|preventDefault={() => onSubmit(input)}
>
  <label class="sr-only" for="url-input">URL to analyze</label>
  <div class="scan-panel">
    <div class="mb-3 flex items-center justify-between gap-3 px-1">
      <div class="flex items-center gap-2.5">
        <span class="scan-shield" aria-hidden="true">
          <svg class="h-4 w-4" viewBox="0 0 24 24" fill="none">
            <path d="M12 3.5 19 6v5.1c0 4.2-2.8 7.4-7 9.4-4.2-2-7-5.2-7-9.4V6l7-2.5Z" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round" />
            <path d="m9 12 2 2 4-4" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </span>
        <span>
          <span class="block text-[11px] font-semibold uppercase tracking-[0.15em] text-gray-800 dark:text-gray-200">Link safety scan</span>
          <span class="mt-0.5 block text-[10px] text-gray-500 dark:text-gray-400">21 analyzers · 39 security signals</span>
        </span>
      </div>
      <span class={`hidden sm:inline-flex items-center gap-1.5 rounded-full border px-2.5 py-1 text-[9px] font-semibold uppercase tracking-[0.14em] ${loading
        ? 'border-accent-light/20 bg-accent-light/[0.05] text-accent-light dark:border-accent-dark/25 dark:text-accent-dark'
        : 'border-emerald-500/20 bg-emerald-500/[0.06] text-emerald-700 dark:text-emerald-300'}`}>
        <span class={`h-1.5 w-1.5 rounded-full ${loading ? 'animate-pulse bg-accent-light dark:bg-accent-dark' : 'bg-emerald-500'}`}></span>
        {loading ? "Scanning" : "Ready"}
      </span>
    </div>
  <div
    data-guide="search"
    class="scan-input flex items-center gap-1 rounded-[1.2rem] p-1.5 pl-3 sm:pl-4 transition-all {formError ? 'scan-input-error' : ''}"
  >
    <span class="scan-link-icon" aria-hidden="true">
      <svg class="h-4 w-4" viewBox="0 0 24 24" fill="none">
        <path d="M10 13a5 5 0 0 0 7.07 0l3-3A5 5 0 0 0 13 2.93l-1.72 1.72M14 11a5 5 0 0 0-7.07 0l-3 3A5 5 0 0 0 11 21.07l1.72-1.72" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round" />
      </svg>
    </span>
    <input
      id="url-input"
      type="text"
      class="scan-url flex-1 min-w-0 bg-transparent py-3 text-[13px] sm:text-sm font-mono text-gray-900 dark:text-gray-100 placeholder-gray-400 dark:placeholder-gray-500 focus:outline-hidden"
      bind:this={inputEl}
      {placeholder}
      bind:value={input}
      on:input={() => {
        if (formError) formError = null;
      }}
      autocomplete="url"
      inputmode="url"
      aria-invalid={formError ? "true" : undefined}
      aria-describedby={formError ? "url-error" : undefined}
      required
    />
    <button
      type="button"
      data-guide="paste"
      on:click={pasteFromClipboard}
      class="shrink-0 inline-flex items-center gap-1.5 px-3 py-2 rounded-full text-xs font-medium transition-colors {justPasted
        ? 'text-emerald-700 dark:text-emerald-400'
        : 'text-gray-500 hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800'}"
      aria-label="Paste from clipboard"
      title="Paste from clipboard"
    >
      {#if justPasted}
        <Icon path={ICON.check} strokeWidth={2.5} class="w-3.5 h-3.5" />
        <span class="hidden sm:inline">Pasted</span>
      {:else}
        <Icon path={ICON.clipboard} class="w-3.5 h-3.5" />
        <span class="hidden sm:inline">Paste</span>
      {/if}
    </button>
    <button
      type="submit"
      data-guide="scan"
      class="shrink-0 inline-flex items-center justify-center gap-2 px-5 sm:px-6 py-2.5 rounded-[0.95rem] bg-accent-light dark:bg-accent-dark text-white dark:text-gray-950 text-sm font-semibold shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md disabled:opacity-60 focus:outline-hidden focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-accent-light dark:focus-visible:ring-offset-gray-900"
      disabled={loading}
      aria-busy={loading}
      aria-label={loading ? "Scanning URL, please wait" : "Scan"}
    >
      {#if loading}
        <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24" aria-hidden="true">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"
          ></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
          ></path>
        </svg>
        Scanning
      {:else}
        Scan
        <Icon path={ICON.arrowRight} />
      {/if}
    </button>
  </div>

  <p class="scan-privacy mt-2.5 flex items-center justify-center gap-1.5 text-[10px] text-gray-500 dark:text-gray-400">
    <svg class="h-3 w-3 text-accent-light dark:text-accent-dark" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <path d="M12 3.5 19 6v5.1c0 4.2-2.8 7.4-7 9.4-4.2-2-7-5.2-7-9.4V6l7-2.5Z" stroke="currentColor" stroke-width="1.8" stroke-linejoin="round" />
      <path d="m9 12 2 2 4-4" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
    </svg>
    Privacy-first · No account required
  </p>
  </div>

  {#if formError}
    <p
      id="url-error"
      class="text-red-600 dark:text-red-400 text-xs mt-2 pl-5 text-left"
      role="alert"
    >
      {formError}
    </p>
  {:else if input.length > 1848 && !/\s/.test(input.trim())}
    <p
      class="text-[11px] mt-1.5 pr-5 text-right {input.length >= 2048
        ? 'text-red-500'
        : 'text-gray-500'}"
    >
      {2048 - input.length} chars remaining
    </p>
  {/if}

  {#if hasTrackers && !trackerDismissed}
    <div class="mt-2 px-1 text-xs space-y-1.5 text-center">
      <p class="text-gray-500">
        <span class="text-gray-600 dark:text-gray-400 font-medium"
          >{trackers.removed.length} tracker{trackers.removed.length > 1 ? "s" : ""} found —</span
        >
        <span class="break-all"> {trackers.removed.join(", ")}</span>
      </p>
      <p class="text-gray-500">Click to copy the clean link</p>
      <button
        type="button"
        on:click={copyCleanUrl}
        class="w-full font-mono text-xs break-all transition-colors cursor-copy {trackerCopied
          ? 'text-emerald-600 dark:text-emerald-400'
          : 'text-gray-500 hover:text-gray-800 dark:hover:text-gray-300'}"
      >
        {trackerCopied ? "Copied!" : trackers.cleaned}
      </button>
    </div>
  {/if}
</form>

<style>
  .scan-panel {
    padding: 0.75rem;
    border: 1px solid rgb(211 202 185 / 0.72);
    border-radius: 1.6rem;
    background: rgb(255 255 255 / 0.78);
    box-shadow: 0 22px 58px rgb(30 41 59 / 0.07);
    backdrop-filter: blur(14px);
  }
  .scan-input {
    border: 1px solid rgb(211 202 185 / 0.82);
    background: rgb(250 247 242 / 0.68);
  }
  .scan-input:focus-within {
    border-color: rgb(52 97 201 / 0.56);
    box-shadow: 0 0 0 3px rgb(52 97 201 / 0.1);
  }
  .scan-input-error,
  .scan-input-error:focus-within {
    border-color: rgb(239 68 68 / 0.58);
    box-shadow: 0 0 0 3px rgb(239 68 68 / 0.08);
  }
  .scan-shield {
    display: grid;
    width: 2.25rem;
    height: 2.25rem;
    place-items: center;
    border: 1px solid rgb(52 97 201 / 0.14);
    border-radius: 0.85rem;
    background: rgb(52 97 201 / 0.08);
    color: #3461c9;
  }
  .scan-link-icon {
    display: grid;
    width: 2rem;
    height: 2rem;
    flex: 0 0 auto;
    place-items: center;
    color: #7c7366;
  }
  .scan-url::placeholder {
    font-family: 'Geist', system-ui, sans-serif;
  }
  :global(.dark) .scan-panel {
    border-color: rgb(70 63 55 / 0.82);
    background: rgb(26 23 20 / 0.82);
    box-shadow: 0 22px 58px rgb(0 0 0 / 0.24);
  }
  :global(.dark) .scan-input {
    border-color: rgb(70 63 55 / 0.9);
    background: rgb(15 13 11 / 0.55);
  }
  :global(.dark) .scan-input:focus-within {
    border-color: rgb(142 176 255 / 0.58);
    box-shadow: 0 0 0 3px rgb(142 176 255 / 0.12);
  }
  :global(.dark) .scan-input-error,
  :global(.dark) .scan-input-error:focus-within {
    border-color: rgb(248 113 113 / 0.58);
    box-shadow: 0 0 0 3px rgb(248 113 113 / 0.1);
  }
  :global(.dark) .scan-shield {
    border-color: rgb(142 176 255 / 0.2);
    background: rgb(142 176 255 / 0.1);
    color: #a9c2ff;
  }
  :global(.dark) .scan-link-icon {
    color: #a69c8c;
  }
  @media (max-width: 400px) {
    .scan-panel {
      padding: 0.6rem;
    }
    .scan-url {
      font-size: 0.75rem;
    }
  }
</style>
