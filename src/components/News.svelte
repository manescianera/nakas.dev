<script>
  import { onMount } from 'svelte'

  const date = Date.parse(new Date().toISOString()) / 1000
  const api = 'https://hacker-news.firebaseio.com/v0'
  const itemsTotal = 30
  const freq = 30

  let promises = []
  let loaded = false

  let news = {
    header: undefined,
    items: [],
    random: {
      titles: Array(itemsTotal).fill(undefined),
      footers: Array(itemsTotal).fill(undefined),
    },
  }

  const getAgo = (time) => {
    let diffInMilliSeconds = Math.abs(date - time)

    const days = Math.floor(diffInMilliSeconds / 86400)
    diffInMilliSeconds -= days * 86400

    const hours = Math.floor(diffInMilliSeconds / 3600) % 24
    diffInMilliSeconds -= hours * 3600

    const minutes = Math.floor(diffInMilliSeconds / 60) % 60
    diffInMilliSeconds -= minutes * 60

    let diff = ''
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

  function randomStr(length) {
    let result = ''
    const chars = 'abcdefghijklmnopqrstuvwxyz'

    for (let i = 0; i < length; i++) {
      result += chars.charAt(Math.floor(Math.random() * chars.length))
    }

    return result
  }

  function randomInt(min, max) {
    return Math.floor(Math.random() * (max - min) + min)
  }

  async function randomizeHeader() {
    news.header = randomStr(3)
  }

  function randomizeItemTitle() {
    let title = []
    const uppercase = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

    while (title.length < randomInt(3, 12)) {
      let word = randomStr(randomInt(2, 10))
      word.length > 3
        ? (word = uppercase.charAt(Math.floor(Math.random() * uppercase.length)) + word)
        : word
      title.push(word)
    }

    return title.join(' ')
  }

  async function randomizeItemTitles() {
    for (let i = 0; i < news.random.titles.length; i++) {
      news.random.titles[i] = randomizeItemTitle()
    }
  }

  function randomizeItemFooter() {
    return [
      randomInt(1, 500),
      randomStr(6),
      randomStr(2),
      randomStr(8),
      randomInt(1, 59),
      randomStr(5),
      randomStr(3),
    ].join(' ')
  }

  async function randomizeItemFooters() {
    for (let i = 0; i < news.random.footers.length; i++) {
      news.random.footers[i] = randomizeItemFooter()
    }
  }

  async function whileLoading() {
    randomizeHeader()
    randomizeItemTitles()
    randomizeItemFooters()
  }

  async function fetchIds() {
    return fetch(`${api}/topstories.json`).then((x) => x.json())
  }

  async function fetchItem(id) {
    return fetch(`${api}/item/${id}.json`).then((x) => x.json())
  }

  async function fetchData() {
    const ids = await fetchIds()

    for (let i = 0; i < itemsTotal; i++) {
      promises.push(fetchItem(ids[i]))
    }

    await Promise.all(promises)

    promises.forEach((p) => {
      p.then((res) => {
        news.items.push(res)
      })
    })

    loaded = true
  }

  async function initNews() {
    fetchData()
    while (!loaded) {
      whileLoading()
      await new Promise((resolve) => setTimeout(resolve, freq))
    }
    news.header = 'top'
  }

  onMount(initNews())
</script>

<div id="items">
  <h1 class="title">{news.header}</h1>
  {#each Array(itemsTotal) as _, i}
    <div class="item">
      <a
        class="item-title"
        href={loaded ? news.items[i].url : ''}
        target="_blank"
        style={loaded ? '' : 'color: var(--global-color-white);'}
        >{loaded ? news.items[i].title : news.random.titles[i]}</a
      >
      <div class="item-footer">
        {loaded
          ? `${news.items[i].score} points by ${news.items[i].by} ${getAgo(news.items[i].time)} ago`
          : news.random.footers[i]}
      </div>
    </div>
  {/each}
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
