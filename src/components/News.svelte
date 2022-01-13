<script>
  import { onMount } from "svelte"

  const api = "https://hacker-news.firebaseio.com/v0"

  async function fetchIds() {
    return fetch(`${api}/topstories.json`).then(x => x.json())
  }

  async function fetchItem(id) {
    return fetch(`${api}/item/${id}.json`).then(x => x.json())
  }

  let items = []
  let loaded = false

  async function fetchData() {
    const data = await fetchIds()

    for (const [i, e] of data.entries()) {
      if (i < 30) {
        const item = await fetchItem(e)
        items.push(item)
      }
    }
    loaded = true
  }

  onMount(fetchData())
</script>

{#if loaded}
  {#each items as item }
    <p>{item.by}</p>
  {/each}
{:else}
  <p>Loading...</p>
{/if}

<style>
  p {
    color: white;
  }
</style>
