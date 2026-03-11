<script>
  import { flip } from "svelte/animate";
  import { dndzone } from "svelte-dnd-action";

  const flipDurationMs = 200;

  // Initial Data Structure for Trello Board
  let columns = [
    {
      id: "c1",
      name: "To Do",
      items: [
        { id: "i1", name: "Learn Wails" },
        { id: "i2", name: "Implement Svelte DND" },
        { id: "i3", name: "Create Trello UI" },
      ],
    },
    {
      id: "c2",
      name: "In Progress",
      items: [{ id: "i4", name: "Build Desktop App" }],
    },
    {
      id: "c3",
      name: "Review",
      items: [],
    },
    {
      id: "c4",
      name: "Done",
      items: [
        { id: "i5", name: "Setup Go environment" },
        { id: "i6", name: "Install Wails CLI" },
      ],
    },
  ];

  function handleDndConsiderColumns(e) {
    columns = e.detail.items;
  }

  function handleDndFinalizeColumns(e) {
    columns = e.detail.items;
  }

  function handleDndConsiderCards(columnId, e) {
    const colIdx = columns.findIndex((c) => c.id === columnId);
    columns[colIdx].items = e.detail.items;
    columns = [...columns];
  }

  function handleDndFinalizeCards(columnId, e) {
    const colIdx = columns.findIndex((c) => c.id === columnId);
    columns[colIdx].items = e.detail.items;
    columns = [...columns];
  }
</script>

<div class="board-container">
  <section
    class="board"
    use:dndzone={{ items: columns, flipDurationMs, type: "columns" }}
    on:consider={handleDndConsiderColumns}
    on:finalize={handleDndFinalizeColumns}
  >
    {#each columns as column (column.id)}
      <div class="column-wrapper" animate:flip={{ duration: flipDurationMs }}>
        <div class="column">
          <div class="column-header">
            <h4>{column.name}</h4>
            <span class="count">{column.items.length}</span>
          </div>

          <div
            class="cards-container"
            use:dndzone={{ items: column.items, flipDurationMs, type: "cards" }}
            on:consider={(e) => handleDndConsiderCards(column.id, e)}
            on:finalize={(e) => handleDndFinalizeCards(column.id, e)}
          >
            {#each column.items as item (item.id)}
              <div class="card" animate:flip={{ duration: flipDurationMs }}>
                {item.name}
              </div>
            {/each}
          </div>

          <div class="add-card-btn">
            <span>+ Add a card</span>
          </div>
        </div>
      </div>
    {/each}
    <div class="add-column-btn">+ Add another list</div>
  </section>
</div>

<style>
  .board-container {
    height: calc(100vh - 50px); /* Subtract Nav height */
    overflow-x: auto;
    overflow-y: hidden;
    background-color: #0079bf; /* Trello background */
    padding: 20px;
    box-sizing: border-box;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  }

  .board {
    display: flex;
    align-items: flex-start;
    height: 100%;
    gap: 15px;
  }

  .column-wrapper {
    height: 100%;
  }

  .column {
    background-color: #ebecf0;
    border-radius: 6px;
    width: 272px;
    max-height: 100%;
    display: flex;
    flex-direction: column;
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
    flex-shrink: 0;
  }

  .column-header {
    padding: 10px 15px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: grab;
  }

  .column-header h4 {
    margin: 0;
    font-size: 1rem;
    font-weight: 600;
    color: #172b4d;
  }

  .count {
    background: #dfe1e6;
    border-radius: 12px;
    padding: 2px 8px;
    font-size: 0.8rem;
    color: #172b4d;
  }

  .cards-container {
    padding: 0 10px;
    flex: 1;
    overflow-y: auto;
    min-height: 20px;
  }

  .card {
    background-color: #ffffff;
    border-radius: 3px;
    box-shadow: 0 1px 2px rgba(9, 30, 66, 0.25);
    padding: 10px;
    margin-bottom: 8px;
    cursor: grab;
    font-size: 0.95rem;
    color: #172b4d;
    word-wrap: break-word;
  }

  .card:hover {
    background-color: #f4f5f7;
  }

  .add-card-btn {
    padding: 10px 15px;
    color: #5e6c84;
    cursor: pointer;
    border-radius: 0 0 6px 6px;
    font-size: 0.95rem;
  }

  .add-card-btn:hover {
    background-color: rgba(9, 30, 66, 0.08);
    color: #172b4d;
  }

  .add-column-btn {
    background-color: rgba(255, 255, 255, 0.24);
    border-radius: 6px;
    width: 272px;
    padding: 12px 15px;
    color: white;
    cursor: pointer;
    flex-shrink: 0;
    font-size: 0.95rem;
  }

  .add-column-btn:hover {
    background-color: rgba(255, 255, 255, 0.32);
  }

  /* Scrollbar styling for cards container */
  .cards-container::-webkit-scrollbar {
    width: 8px;
  }

  .cards-container::-webkit-scrollbar-track {
    background: rgba(9, 30, 66, 0.08);
    border-radius: 4px;
  }

  .cards-container::-webkit-scrollbar-thumb {
    background: #bfc2ce;
    border-radius: 4px;
  }

  .cards-container::-webkit-scrollbar-thumb:hover {
    background: #a5a8b5;
  }
</style>
