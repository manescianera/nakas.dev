<script>
	let message = "meanwhile.."
	let href = "https://github.com/manescianera"
</script>

<a {href}>{message}</a>

<style>
a {
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);

	opacity: 1;
	text-decoration: none;
	font-size: calc(0.75em + 1vw);
	color: #bcc5d8;
	border-radius: 3px;
	padding: 1rem;
	animation: appear 2000ms;
}

a:hover {
	color: #bf616a;
	background: #232835;
	text-decoration: underline;
}

@keyframes appear {
	from {
		opacity: 0;
	}
	to {
		opacity: 1;
	}
}
</style>
