<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<nav>
	<a href="/" class="logo">My App</a>
	<a href="/about" class="nav-link">About</a>
	<a href="/contact" class="nav-link">Contact</a>
</nav>
{@render children()}
