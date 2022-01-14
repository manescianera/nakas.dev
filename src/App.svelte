<script>
	import router from 'page'
	import Home from './components/Home.svelte'
	import News from './components/News.svelte'

	let page

	router('/', () => page = Home)
	router('/news', () => page = News)

	router.start()
</script>


<main>
	<div id="content">
		<svelte:component this={page} />
	</div>
</main>
