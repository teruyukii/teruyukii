<script lang="ts">
	import '$lib/styles/global.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/generics/Navbar.svelte';
	let { children } = $props();
	import { env } from '$env/dynamic/public';
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
{#if env.PUBLIC_MAINTENANCE_PAGE === 'true'}
	<div class="flex min-h-screen flex-col items-center justify-center p-4 text-center">
		<img src={favicon} alt="Favicon" class="mb-4 h-16 w-16" />
		<h1>Maintenance Mode</h1>
		<p>We are currently performing maintenance on our site. Please check back soon.</p>
	</div>
{:else}
	<Navbar />
	{@render children()}
{/if}
