<script lang="ts">
  import {
    Accordion,
    Avatar,
    Progress,
    RatingGroup,
    SegmentedControl,
    Slider,
    Switch,
  } from "@skeletonlabs/skeleton-svelte";
  import {
    ChevronDownIcon,
    MailIcon,
    MoonIcon,
    SettingsIcon,
    StarIcon,
    UserIcon,
  } from "@lucide/svelte";
  import { slide } from "svelte/transition";

  let tabSet = $state(0);
  let toggleValue = $state(true);
  let sliderValue = $state([50]);
  let ratingValue = $state(3);
  let segmentedValue = $state("email");

  const faqItems = [
    {
      id: "1",
      title: "Skeleton v3 có gì mới?",
      description:
        "Skeleton v3 được viết lại hoàn toàn cho Svelte 5, hỗ trợ $state runes, không cần import component riêng lẻ như trước.",
    },
    {
      id: "2",
      title: "Tại sao không dùng import như cũ?",
      description:
        "Skeleton v3 dùng kiến trúc mới dựa trên Zag.js. Các component được import từ @skeletonlabs/skeleton-svelte thay vì @skeletonlabs/skeleton.",
    },
    {
      id: "3",
      title: "Có thể mở nhiều item cùng lúc không?",
      description:
        'Có! Chỉ cần thêm prop "multiple" vào <Accordion multiple> là có thể mở nhiều item cùng lúc.',
    },
  ];
</script>

<div class="container mx-auto p-8 space-y-8 pb-20">
  <header class="text-center">
    <h1 class="h1 font-bold">Skeleton UI Showcase</h1>
    <p class="opacity-50">Cùng xem các component xịn xò của Skeleton nào!</p>
  </header>

  <!-- Avatars -->
  <section class="card p-4 space-y-4 shadow-xl">
    <h3 class="h3 font-bold">Avatars</h3>
    <div class="flex flex-wrap gap-8 items-center justify-center">
      <Avatar>
        <Avatar.Image
          src="https://i.pravatar.cc/128?u=1"
          alt="User"
        />
        <Avatar.Fallback>JD</Avatar.Fallback>
      </Avatar>

      <Avatar>
        <Avatar.Fallback class="preset-filled-secondary-500">
          <UserIcon class="size-6" />
        </Avatar.Fallback>
      </Avatar>

      <div class="flex -space-x-3">
        {#each [1, 2, 3] as i}
          <Avatar class="border-2 border-surface-100-900">
            <Avatar.Image
              src="https://i.pravatar.cc/128?u={i}"
              alt="User {i}"
            />
            <Avatar.Fallback>U{i}</Avatar.Fallback>
          </Avatar>
        {/each}
      </div>
    </div>
  </section>

  <!-- Buttons & Badges -->
  <section class="card p-4 space-y-4 shadow-xl">
    <h3 class="h3 font-bold">Buttons & Badges</h3>
    <div class="flex flex-wrap gap-4 items-center">
      <button class="btn preset-filled-primary-500">Primary Button</button>
      <button class="btn preset-filled-secondary-500">Secondary Button</button>
      <button class="btn preset-tonal-tertiary">Ghost Tertiary</button>

      <span class="badge preset-filled-success-500">Success</span>
      <span class="badge preset-filled-warning-500">Warning</span>
      <span class="badge preset-filled-error-500">Error</span>
    </div>
  </section>

  <!-- Progress & Switch -->
  <section class="card p-4 space-y-6 shadow-xl">
    <h3 class="h3 font-bold">Progress & Selection</h3>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="space-y-4">
        <p class="font-bold">Official Progress:</p>
        <Progress
          value={64}
          max={100}
        >
          <Progress.Label class="mb-2 block">64% Complete</Progress.Label>
          <Progress.Track class="h-2 rounded-full bg-surface-200-800 overflow-hidden">
            <Progress.Range class="bg-primary-500 h-full" />
          </Progress.Track>
        </Progress>

        <Progress
          value={null}
          max={100}
        >
          <Progress.Label class="mb-2 block">Loading...</Progress.Label>
          <Progress.Track class="h-2 rounded-full bg-surface-200-800 overflow-hidden">
            <Progress.Range class="bg-tertiary-500 h-full animate-[progress_1.5s_ease-in-out_infinite] w-1/3" />
          </Progress.Track>
        </Progress>
      </div>

      <div class="space-y-4">
        <p class="font-bold">Official Switch (Toggle):</p>
        <div class="flex flex-col gap-4">
          <Switch
            checked={toggleValue}
            onCheckedChange={(e) => (toggleValue = e.checked)}
            class="flex items-center justify-between"
          >
            <Switch.Label>Notifications</Switch.Label>
            <Switch.Control class="w-12 h-6 rounded-full transition-colors duration-200 {toggleValue ? 'bg-success-500' : 'bg-surface-400'} relative">
              <Switch.Thumb class="absolute top-1 left-1 w-4 h-4 rounded-full bg-white transition-transform duration-200 {toggleValue ? 'translate-x-[22px]' : ''}" />
            </Switch.Control>
          </Switch>

          <Switch class="flex items-center justify-between">
            <Switch.Label class="flex items-center gap-2">
              <MoonIcon class="size-4" /> Dark Mode
            </Switch.Label>
            <Switch.Control class="w-12 h-6 rounded-full bg-surface-400 relative">
              <Switch.Thumb class="absolute top-1 left-1 w-4 h-4 rounded-full bg-white transition-transform" />
            </Switch.Control>
          </Switch>
        </div>
      </div>
    </div>
  </section>

  <!-- Slider & Segmented Control -->
  <section class="card p-4 space-y-6 shadow-xl">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="space-y-4">
        <h3 class="h3 font-bold">Slider</h3>
        <Slider
          value={sliderValue}
          onValueChange={(e) => (sliderValue = e.value)}
          min={0}
          max={100}
          step={1}
        >
          <Slider.Label class="mb-4 block">Volume: {sliderValue[0]}%</Slider.Label>
          <Slider.Control class="flex items-center">
            <Slider.Track class="h-2 flex-1 rounded-full bg-surface-200-800 overflow-hidden">
              <Slider.Range class="bg-primary-500 h-full" />
            </Slider.Track>
            <Slider.Thumb index={0} class="size-6 bg-surface-100-900 border-2 border-primary-500 rounded-full shadow-xl -ml-3" />
          </Slider.Control>
        </Slider>
      </div>

      <div class="space-y-4">
        <h3 class="h3 font-bold">Segmented Control</h3>
        <SegmentedControl
          value={segmentedValue}
          onValueChange={(e) => (segmentedValue = e.value)}
          class="flex bg-surface-200-800 p-1 rounded-3xl"
        >
          <SegmentedControl.Item
            value="email"
            class="flex-1"
          >
            <SegmentedControl.ItemText class="flex items-center justify-center p-2 rounded-2xl transition-colors {segmentedValue === 'email' ? 'bg-surface-100-900 shadow font-bold' : 'opacity-50'}">
              <MailIcon class="size-4 mr-2" /> Email
            </SegmentedControl.ItemText>
          </SegmentedControl.Item>
          <SegmentedControl.Item
            value="settings"
            class="flex-1"
          >
            <SegmentedControl.ItemText class="flex items-center justify-center p-2 rounded-2xl transition-colors {segmentedValue === 'settings' ? 'bg-surface-100-900 shadow font-bold' : 'opacity-50'}">
              <SettingsIcon class="size-4 mr-2" /> Settings
            </SegmentedControl.ItemText>
          </SegmentedControl.Item>
          <SegmentedControl.Item
            value="profile"
            class="flex-1"
          >
            <SegmentedControl.ItemText class="flex items-center justify-center p-2 rounded-2xl transition-colors {segmentedValue === 'profile' ? 'bg-surface-100-900 shadow font-bold' : 'opacity-50'}">
              <UserIcon class="size-4 mr-2" /> Profile
            </SegmentedControl.ItemText>
          </SegmentedControl.Item>
        </SegmentedControl>
      </div>
    </div>
  </section>

  <!-- Rating Group -->
  <section class="card p-4 space-y-4 shadow-xl">
    <h3 class="h3 font-bold">Rating Group</h3>
    <div class="flex items-center gap-4">
      <RatingGroup
        value={ratingValue}
        count={5}
        onValueChange={(e) => (ratingValue = e.value)}
      >
        <RatingGroup.Control class="flex gap-1">
          {#each Array.from({ length: 5 }) as _, i}
            <RatingGroup.Item index={i + 1}>
              {#snippet full()}<StarIcon class="size-8 fill-current text-warning-500" />{/snippet}
              {#snippet empty()}<StarIcon class="size-8 text-surface-400" />{/snippet}
            </RatingGroup.Item>
          {/each}
        </RatingGroup.Control>
      </RatingGroup>
      <span class="text-xl font-bold">{ratingValue} / 5</span>
    </div>
  </section>

  <!-- Tabs (Custom CSS based) -->
  <section class="card p-4 shadow-xl">
    <h3 class="h3 font-bold mb-4">Tabs</h3>
    <div class="flex gap-2 border-b border-surface-300-700 mb-4">
      {#each [["Home", 0], ["About", 1], ["Contact", 2]] as [label, idx]}
        <button
          class="px-4 py-2 font-medium transition-colors {tabSet === idx
            ? 'border-b-2 border-primary-500 text-primary-500'
            : 'opacity-50'}"
          onclick={() => (tabSet = idx as number)}
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

  <!-- Accordion (Skeleton v3 component) -->
  <section class="card p-4 shadow-xl space-y-2">
    <h3 class="h3 font-bold mb-4">Accordion (FAQ)</h3>

    <Accordion collapsible>
      {#each faqItems as item, i (item)}
        {#if i !== 0}
          <hr class="hr" />
        {/if}
        <Accordion.Item value={item.id}>
          <h3>
            <Accordion.ItemTrigger
              class="font-bold flex items-center justify-between gap-2 w-full py-2"
            >
              {item.title}
              <Accordion.ItemIndicator class="group">
                <ChevronDownIcon
                  class="h-5 w-5 transition group-data-[state=open]:rotate-180"
                />
              </Accordion.ItemIndicator>
            </Accordion.ItemTrigger>
          </h3>
          <Accordion.ItemContent>
            {#snippet element(attributes)}
              {#if !attributes.hidden}
                <div
                  {...attributes}
                  transition:slide={{ duration: 150 }}
                  class="pb-3 opacity-75"
                >
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
