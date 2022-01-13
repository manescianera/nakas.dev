<script>
	import { Router, Route } from "svelte-routing";
	import Home from "./pages/Home.svelte"
	export let url = "";
</script>

<main>
	<div id="content">
		<Router {url}>
			<Route basepath component={Home} />
		</Router>
	</div>
</main>
