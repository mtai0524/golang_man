<script>
  import Nav from "./lib/Nav.svelte";
  import Board from "./lib/Board.svelte";

  let showShowcase = false;

  function toggleView() {
    showShowcase = !showShowcase;
  }
</script>

<main class="app-wrapper">
  <Nav />

  <div class="view-toggle p-4 flex justify-end bg-surface-100-800-token">
    <button class="btn variant-filled-primary" on:click={toggleView}>
      {showShowcase ? "View Trello Board" : "View Skeleton Components"}
    </button>
  </div>

  <div class="content-area flex-1 overflow-auto">
    {#if showShowcase}
      <SkeletonShowcase />
    {:else}
      <Board />
    {/if}
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    overflow: hidden; /* Prevent scrolling on body, the board will scroll horizontally */
  }

  .app-wrapper {
    height: 100vh;
    display: flex;
    flex-direction: column;
  }
</style>
