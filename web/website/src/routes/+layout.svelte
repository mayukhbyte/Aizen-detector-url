<script lang="ts">
  import ThemeToggle from "$lib/components/ThemeToggle.svelte";
  import { page } from "$app/stores";
  import Guide from "$lib/components/guide/Guide.svelte";
  import { vettyMemory } from "$lib/components/guide/store";
  import { theme } from "$lib/theme";
  import { REPO } from "$lib/site";
  import { onMount } from "svelte";
  import "../app.css";

  const year = new Date().getFullYear();

  const FOOTER_LINKS: { label: string; href: string; external?: boolean }[] = [
    { label: "How it works", href: "/how-it-works" },
    { label: "About", href: "/about" },
    { label: "Privacy", href: "/privacy" },
  ];

  $: isHome = $page.url.pathname === "/";
  $: isAdmin = $page.url.pathname.startsWith("/admin");

  const schemaWebSite = {
    "@context": "https://schema.org",
    "@type": "WebSite",
    name: "aizen.detector",
    alternateName: ["aizen.detector", "Aizen Detector"],
    url: "https://aizen.detector",
    description:
      "Free real-time URL scanner. Paste any link and get an instant phishing and safety verdict, no signup needed.",
  };

  const schemaOrganization = {
    "@context": "https://schema.org",
    "@type": "Organization",
    name: "aizen.detector",
    alternateName: ["aizen.detector", "Aizen Detector"],
    url: "https://aizen.detector",
  };

  onMount(() => {
    theme.init();
  });
</script>

<svelte:head>
  {@html `<script type="application/ld+json">${JSON.stringify(schemaWebSite)}</script>`}
  {@html `<script type="application/ld+json">${JSON.stringify(schemaOrganization)}</script>`}
</svelte:head>

<div class="relative min-h-screen flex flex-col">
  <nav class="relative z-20 w-full max-w-6xl mx-auto px-5 sm:px-6 py-5 sm:py-6 flex items-center justify-between gap-4">
    <!-- Compact brand anchor; the larger hero wordmark reinforces the product name. -->
    <a
      href="/"
      on:click={() => (location.href = "/")}
      aria-current={isHome ? "page" : undefined}
      aria-label="Aizen Detector home"
      class="group inline-flex items-center gap-2.5 font-serif text-[1.65rem] leading-none tracking-[-0.02em]"
    >
      <img
        src="/aizen-detector-normal.png"
        alt="Aizen Detector Logo"
        class="w-7 h-7 rounded-lg object-contain shadow-sm border border-cyan-500/20"
      />
      <span>
        <span
          class="text-gray-400 dark:text-gray-500 group-hover:text-gray-500 dark:group-hover:text-gray-400 transition-colors"
          >aizen</span
        ><span class="-mx-[0.025em] text-accent-light dark:text-accent-dark">.</span><span
          class="text-gray-900 dark:text-gray-100">detector</span
        >
      </span>
    </a>
    <div
      class="hidden sm:flex items-center gap-1 rounded-full border border-gray-200/80 bg-white/50 p-1 text-sm shadow-sm shadow-gray-900/[0.02] dark:border-gray-800 dark:bg-gray-900/50"
      role="group"
      aria-label="Main navigation"
    >
      {#each FOOTER_LINKS.slice(0, 2) as link}
        <a
          href={link.href}
          aria-current={$page.url.pathname === link.href ? "page" : undefined}
          class="rounded-full px-3.5 py-2 transition-colors {$page.url.pathname === link.href
            ? 'bg-gray-100 text-gray-900 dark:bg-gray-800 dark:text-gray-100'
            : 'text-gray-500 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-100'}"
        >
          {link.label}
        </a>
      {/each}
    </div>
    <div class="flex items-center gap-2">
      <ThemeToggle />
    </div>
  </nav>
  <main class="flex-1">
    <slot />
  </main>

  {#if !isAdmin}
    <Guide />
  {/if}

  <footer class="text-gray-500 dark:text-gray-400 py-8">
    <div
      class="max-w-5xl mx-auto px-6 pt-8 border-t border-gray-200 dark:border-gray-800 flex flex-col md:flex-row md:justify-between items-center gap-4 text-sm"
    >
      <!-- Left: site links, then (if hidden) Vetty -->
      <nav
        aria-label="Footer"
        data-guide="learn-more"
        class="flex flex-wrap items-center justify-center md:justify-start gap-x-5 sm:gap-x-4 gap-y-2"
      >
        {#each FOOTER_LINKS as link, i}
          {#if i > 0}<span
              class="hidden sm:inline text-gray-300 dark:text-gray-700"
              aria-hidden="true">·</span
            >{/if}
          <a
            href={link.href}
            target={link.external ? "_blank" : undefined}
            rel={link.external ? "noopener noreferrer" : undefined}
            aria-current={$page.url.pathname === link.href ? "page" : undefined}
            class="transition-colors hover:text-gray-900 dark:hover:text-gray-100 {$page.url
              .pathname === link.href
              ? 'text-gray-900 dark:text-gray-100'
              : 'text-gray-500 dark:text-gray-400'}"
          >
            {link.label}
          </a>
        {/each}
        {#if $vettyMemory.hidden}
          <span class="hidden sm:inline text-gray-300 dark:text-gray-700" aria-hidden="true">·</span
          >
          <button
            type="button"
            on:click={() => vettyMemory.set({ hidden: false })}
            class="text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors"
          >
            Show Friday 📎
          </button>
        {/if}
      </nav>

      <!-- Right: Copyright -->
      <p class="text-gray-500 dark:text-gray-400 text-center md:text-right">
        Aizen Detector © {year}
      </p>
    </div>
  </footer>
</div>
