<script>
  import { onMount } from "svelte"

  const api = "https://hacker-news.firebaseio.com/v0"

  async function fetchIds() {
    return fetch(`${api}/topstories.json`).then(x => x.json())
  }

  async function fetchItem(id) {
    return fetch(`${api}/item/${id}.json`).then(x => x.json())
  }

  const total = 30
  let items = []
  let loaded = false

  async function fetchData() {
    const ids = await fetchIds()
    let promises = []

    for (let i = 0; i < total; i++) {
      promises.push(fetchItem(ids[i]))
    }

    await Promise.all(promises)

    promises.forEach(p => {
      p.then(res => {
        console.log(res)
        items.push(res)
      })
    })

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
