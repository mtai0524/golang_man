<script>
  import { TabGroup, Tab, TabAnchor } from '@skeletonlabs/skeleton';
  import { ProgressBar } from '@skeletonlabs/skeleton';
  import { SlideToggle } from '@skeletonlabs/skeleton';

  let tabSet = 0;
  let toggleValue = true;
</script>

<div class="container mx-auto p-8 space-y-8">
  <header class="text-center">
    <h1 class="h1">Skeleton UI Showcase</h1>
    <p class="opacity-50">Cùng xem các component xịn xò của Skeleton nào!</p>
  </header>

  <!-- Section 1: Buttons and Badges -->
  <section class="card p-4 space-y-4 shadow-xl">
    <h3 class="h3">Buttons & Badges</h3>
    <div class="flex flex-wrap gap-4 items-center">
      <button class="btn variant-filled-primary">Primary Button</button>
      <button class="btn variant-filled-secondary">Secondary Button</button>
      <button class="btn variant-ghost-tertiary">Ghost Tertiary</button>
      <button class="btn btn-icon variant-filled" aria-label="Heart"><i class="fa-solid fa-heart"></i></button>
      
      <span class="badge variant-filled-success">Success</span>
      <span class="badge variant-filled-warning">Warning</span>
      <span class="badge variant-filled-error">Error</span>
    </div>
  </section>

  <!-- Section 2: Cards and Progress -->
  <section class="grid grid-cols-1 md:grid-cols-2 gap-4">
    <!-- Card Data -->
    <a href="/" class="card card-hover p-4 overflow-hidden block">
      <header>
        <img src="https://images.unsplash.com/photo-1614850523459-c2f4c699c52e?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80" class="bg-black/50 w-full aspect-[2/1] object-cover" alt="Post" />
      </header>
      <div class="p-4 space-y-4">
        <h3 class="h3" data-toc-ignore>Svelte & Wails</h3>
        <article>
          <p>This is an example of a Skeleton Card component. It looks gorgeous out of the box with minimal effort!</p>
        </article>
      </div>
      <hr class="opacity-50" />
      <footer class="p-4 flex justify-between items-center">
        <div class="flex items-center gap-2">
          <div class="avatar w-8 h-8 rounded-full bg-surface-300-600-token"></div>
          <span>Avatar</span>
        </div>
        <button class="btn variant-ringed-primary">Read More</button>
      </footer>
    </a>

    <!-- Interactive Elements -->
    <div class="card p-4 space-y-6 shadow-xl flex flex-col justify-center">
      <h3 class="h3">Interactive Elements</h3>
      
      <div>
        <p class="mb-2">Progress Bar:</p>
        <ProgressBar value={64} max={100} meter="bg-primary-500" track="bg-surface-200-700-token" />
      </div>

      <div>
        <p class="mb-2">Indeterminate Progress:</p>
        <ProgressBar meter="bg-tertiary-500" track="bg-surface-200-700-token" />
      </div>

      <div class="flex items-center justify-between">
        <p>Slide Toggle:</p>
        <SlideToggle name="slide" bind:checked={toggleValue} active="bg-success-500" />
      </div>
    </div>
  </section>

  <!-- Section 3: Tabs -->
  <section class="card p-4 shadow-xl">
    <TabGroup>
      <Tab bind:group={tabSet} name="tab1" value={0}>Home</Tab>
      <Tab bind:group={tabSet} name="tab2" value={1}>About</Tab>
      <Tab bind:group={tabSet} name="tab3" value={2}>Contact</Tab>
      <!-- Tab Panels --->
      <svelte:fragment slot="panel">
        {#if tabSet === 0}
          <div class="p-4 text-center text-lg">
            Chào mừng bạn đến với Tab 1. Đây là nội dung của trang chủ.
          </div>
        {:else if tabSet === 1}
          <div class="p-4 text-center text-lg">
            Tab 2: Giới thiệu về ứng dụng Desktop tuyệt vời này.
          </div>
        {:else if tabSet === 2}
          <div class="p-4 text-center text-lg">
            Tab 3: Liên hệ với chúng tôi nếu bạn có bất kỳ câu hỏi nào.
          </div>
        {/if}
      </svelte:fragment>
    </TabGroup>
  </section>
</div>
