<script>
  import Nav from "./lib/Nav.svelte";
  import Board from "./lib/Board.svelte";
</script>

<main class="app-wrapper">
  <Nav />
  <Board />
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
