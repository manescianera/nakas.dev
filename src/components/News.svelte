<script>
  import { onMount } from "svelte"

  const date = Date.parse(new Date().toISOString()) / 1000
  const title = "top"
  const api = "https://hacker-news.firebaseio.com/v0"

  const getAgo = (time) => {
    let diffInMilliSeconds = Math.abs(date - time)

    const days = Math.floor(diffInMilliSeconds / 86400)
    diffInMilliSeconds -= days * 86400

    const hours = Math.floor(diffInMilliSeconds / 3600) % 24
    diffInMilliSeconds -= hours * 3600

    const minutes = Math.floor(diffInMilliSeconds / 60) % 60
    diffInMilliSeconds -= minutes * 60

    let diff = ""
    if (days) {
      diff += days === 1 ? `${days} day ` : `${days} days `
    }

    if (hours) {
      diff += hours === 1 ? `${hours} hour ` : `${hours} hours `
    }

    if (!days && !hours && minutes) {
      diff += minutes === 1 ? `${minutes} minute` : `${minutes} minutes`
    }

    return diff
  }

  async function fetchIds() {
    return fetch(`${api}/topstories.json`).then((x) => x.json())
  }

  async function fetchItem(id) {
    return fetch(`${api}/item/${id}.json`).then((x) => x.json())
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

    promises.forEach((p) => {
      p.then((res) => {
        items.push(res)
      })
    })

    loaded = true
  }

  onMount(fetchData())
</script>

<div id="items">
  <h1 class="title">{title}</h1>
  {#if loaded}
    {#each items as item}
      <div class="item">
        <a class="item-title" href={item.url} target="_blank">{item.title}</a>
        <div class="item-footer">
          {item.score} points by {item.by}
          {getAgo(item.time)} ago
        </div>
      </div>
    {/each}
  {:else}
    <p>Loading...</p>
  {/if}
</div>

<style>
  #items {
    padding: var(--global-padding);
    margin: auto;
    width: 80vw;
  }

  .title {
    color: var(--global-color-white);
    padding: calc(var(--global-padding) / 2);
  }

  .item {
    padding: calc(var(--global-padding) / 2);
    border-radius: var(--global-border-radius);
  }

  .item:hover {
    background: var(--global-color-bg-secondary);
  }

  .item-title {
    color: var(--global-color-white);
    text-decoration: none;
  }

  .item-title:hover {
    color: var(--global-color-red);
    text-decoration: underline;
  }

  .item-title:visited {
    color: var(--global-color-dimmed);
  }

  .item-title:visited:hover {
    color: var(--global-color-red);
  }

  .item-footer {
    color: var(--global-color-dimmed);
    font-size: 0.8rem;
  }
</style>
