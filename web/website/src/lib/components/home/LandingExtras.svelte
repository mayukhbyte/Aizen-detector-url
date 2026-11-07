<script lang="ts">
  import { onMount, tick } from "svelte";
  import { EXAMPLES } from "../../data/examples";
  // Example links to try, and the short list of promises under the search bar.
  export let onTry: (url: string) => void;

  const POINTS = ["Free and open source", "No tracking", "Explains every result"];

  // Keep the chips to two rows by measuring, not by breakpoint: drop the legit/fake
  // labels (level 1), then "Try" too (level 2), until they fit.
  let level = 0;
  let row: HTMLDivElement;

  const rows = () =>
    new Set(
      // By vertical centre: "Try" is shorter than the chips beside it.
      [...row.children].map((el) => {
        const e = el as HTMLElement;
        return Math.round((e.offsetTop + e.offsetHeight / 2) / 8);
      })
    ).size;

  async function fit() {
    if (!row) return;
    for (level = 0; level < 2; level++) {
      await tick();
      if (rows() <= 2) return;
    }
  }

  onMount(() => {
    fit();
    document.fonts?.ready.then(fit);
    let width = row.clientWidth;
    const ro = new ResizeObserver(() => {
      if (row.clientWidth !== width) {
        width = row.clientWidth;
        fit();
      }
    });
    ro.observe(row);
    return () => ro.disconnect();
  });
</script>

<div class="mt-8 mb-2 flex flex-wrap items-center justify-center gap-x-2 gap-y-1 text-[10px]">
  <span class="font-semibold uppercase tracking-[0.18em] text-gray-700 dark:text-gray-300">Sample scans</span>
  <span class="hidden min-[420px]:inline text-gray-400 dark:text-gray-500">·</span>
  <span class="text-gray-500 dark:text-gray-400">Choose one to preview a report</span>
</div>

<div
  bind:this={row}
  data-guide="examples"
  class="flex flex-wrap justify-center items-center gap-2"
>
  {#if level < 2}
    <span class="text-xs text-gray-500 mr-1">Try</span>
  {/if}
  {#each EXAMPLES as example}
    <button
      type="button"
      data-example={example.url}
      on:click={() => onTry(example.url)}
      title={example.hint === "Legit"
        ? "The real site"
        : "A lookalike, spelled with letters from another alphabet"}
      class="inline-flex items-center py-1.5 {level < 2
        ? 'gap-2 px-3.5'
        : 'gap-1.5 px-2.5'} rounded-full border border-gray-300/80 bg-white/55 text-gray-700 shadow-sm shadow-gray-900/[0.02] transition-all hover:-translate-y-0.5 hover:border-accent-light/40 hover:bg-white hover:shadow-md dark:border-gray-800 dark:bg-gray-900/55 dark:text-gray-300 dark:hover:border-accent-dark/35 dark:hover:bg-gray-900"
    >
      <span
        class="w-1.5 h-1.5 rounded-full {example.hint === 'Legit'
          ? 'bg-emerald-500'
          : 'bg-red-500'}"
      ></span>
      <span class="font-mono">{example.label}</span>
      {#if level === 0}
        <span class="text-[11px] text-gray-400 dark:text-gray-500"
          >{example.hint.toLowerCase()}</span
        >
      {/if}
    </button>
  {/each}
</div>

<div class="mt-10 grid w-full max-w-3xl grid-cols-1 gap-2 sm:grid-cols-3 sm:gap-3">
  {#each POINTS as point}
    <div
      class="flex items-center justify-center gap-2 rounded-full border border-gray-200/80 bg-white/60 px-3 py-2.5 text-xs text-gray-700 shadow-sm shadow-gray-900/[0.02] dark:border-gray-800 dark:bg-gray-900/55 dark:text-gray-300"
    >
      <span class="grid h-5 w-5 shrink-0 place-items-center rounded-full bg-accent-light/10 text-accent-light dark:bg-accent-dark/10 dark:text-accent-dark">
        <svg class="h-3 w-3" viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="m5 12 4 4L19 6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
      </span>
      <span>{point}</span>
    </div>
  {/each}
</div>
