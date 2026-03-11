<script lang="ts">
  import { Accordion } from '@skeletonlabs/skeleton-svelte';
  import { ChevronDownIcon } from '@lucide/svelte';
  import { slide } from 'svelte/transition';

  let tabSet = $state(0);
  let toggleValue = $state(true);

  const faqItems = [
    {
      id: '1',
      title: 'Skeleton v3 có gì mới?',
      description: 'Skeleton v3 được viết lại hoàn toàn cho Svelte 5, hỗ trợ $state runes, không cần import component riêng lẻ như trước.',
    },
    {
      id: '2',
      title: 'Tại sao không dùng import như cũ?',
      description: 'Skeleton v3 dùng kiến trúc mới dựa trên Zag.js. Các component được import từ @skeletonlabs/skeleton-svelte thay vì @skeletonlabs/skeleton.',
    },
    {
      id: '3',
      title: 'Có thể mở nhiều item cùng lúc không?',
      description: 'Có! Chỉ cần thêm prop "multiple" vào <Accordion multiple> là có thể mở nhiều item cùng lúc.',
    },
  ];
</script>

<div class="container mx-auto p-8 space-y-8">
  <header class="text-center">
    <h1 class="h1">Skeleton UI Showcase</h1>
    <p class="opacity-50">Cùng xem các component xịn xò của Skeleton nào!</p>
  </header>

  <!-- Buttons & Badges -->
  <section class="card p-4 space-y-4 shadow-xl">
    <h3 class="h3">Buttons & Badges</h3>
    <div class="flex flex-wrap gap-4 items-center">
      <button class="btn preset-filled-primary-500">Primary Button</button>
      <button class="btn preset-filled-secondary-500">Secondary Button</button>
      <button class="btn preset-tonal-tertiary">Ghost Tertiary</button>

      <span class="badge preset-filled-success-500">Success</span>
      <span class="badge preset-filled-warning-500">Warning</span>
      <span class="badge preset-filled-error-500">Error</span>
    </div>
  </section>

  <!-- Progress + Toggle -->
  <section class="card p-4 space-y-6 shadow-xl">
    <h3 class="h3">Progress Bars</h3>

    <div>
      <p class="mb-2">Progress Bar (64%):</p>
      <div class="progress h-2 rounded-full bg-surface-200-800">
        <div class="h-full rounded-full bg-primary-500" style="width: 64%"></div>
      </div>
    </div>

    <div>
      <p class="mb-2">Indeterminate Progress:</p>
      <div class="progress-indeterminate h-2 rounded-full bg-surface-200-800 overflow-hidden">
        <div class="h-full w-1/3 rounded-full bg-tertiary-500 animate-[progress_1.5s_ease-in-out_infinite]"></div>
      </div>
    </div>

    <div class="flex items-center justify-between">
      <p>Slide Toggle:</p>
     <button
  role="switch"
  aria-label="Slide toggle"
  aria-checked={toggleValue}
  onclick={() => (toggleValue = !toggleValue)}
  class="w-12 h-6 rounded-full transition-colors duration-200 {toggleValue ? 'bg-success-500' : 'bg-surface-400'} relative"
>
  <span class="absolute top-1 left-1 w-4 h-4 rounded-full bg-white transition-transform duration-200 {toggleValue ? 'translate-x-6' : ''}"></span>
</button>
    </div>
  </section>

  <!-- Tabs -->
  <section class="card p-4 shadow-xl">
    <div class="flex gap-2 border-b border-surface-300-700 mb-4">
      {#each [["Home", 0], ["About", 1], ["Contact", 2]] as [label, idx]}
        <button
          class="px-4 py-2 font-medium transition-colors {tabSet === idx ? 'border-b-2 border-primary-500 text-primary-500' : 'opacity-50'}"
          onclick={() => (tabSet = idx)}
        >
          {label}
        </button>
      {/each}
    </div>

    {#if tabSet === 0}
      <p class="p-4 text-center text-lg">Chào mừng bạn đến với Tab 1.</p>
    {:else if tabSet === 1}
      <p class="p-4 text-center text-lg">Tab 2: Giới thiệu về ứng dụng.</p>
    {:else}
      <p class="p-4 text-center text-lg">Tab 3: Liên hệ với chúng tôi.</p>
    {/if}
  </section>

  <!-- ✅ Accordion (Skeleton v3 component) -->
  <section class="card p-4 shadow-xl space-y-2">
    <h3 class="h3 mb-4">Accordion (FAQ)</h3>

    <Accordion collapsible>
      {#each faqItems as item, i (item)}
        {#if i !== 0}
          <hr class="hr" />
        {/if}
        <Accordion.Item value={item.id}>
          <h3>
            <Accordion.ItemTrigger class="font-bold flex items-center justify-between gap-2 w-full py-2">
              {item.title}
              <Accordion.ItemIndicator class="group">
                <ChevronDownIcon class="h-5 w-5 transition group-data-[state=open]:rotate-180" />
              </Accordion.ItemIndicator>
            </Accordion.ItemTrigger>
          </h3>
          <Accordion.ItemContent>
            {#snippet element(attributes)}
              {#if !attributes.hidden}
                <div {...attributes} transition:slide={{ duration: 150 }} class="pb-3 opacity-75">
                  {item.description}
                </div>
              {/if}
            {/snippet}
          </Accordion.ItemContent>
        </Accordion.Item>
      {/each}
    </Accordion>
  </section>
</div>